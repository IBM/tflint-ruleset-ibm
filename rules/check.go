// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package rules

import (
	"encoding/json"
	"fmt"
	"net"
	"regexp"
	"strings"
	"sync"

	"github.com/IBM/tflint-ruleset-ibm/internal/redact"
	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AttrCheck is a generated per-attribute validation spec.
type AttrCheck struct {
	Name         string
	Required     bool
	Enum         []string
	IntEnum      []int
	Min          *int
	Max          *int
	MinLen       *int
	MaxLen       *int
	Regexp       string
	Kind         string
	NoZero       bool
	List         bool
	ComputedOnly bool
	// Sensitive marks attributes whose values must never appear in messages.
	Sensitive bool
}

// BlockCheck is a generated nested-block validation spec.
type BlockCheck struct {
	Type     string
	Required bool
	Attrs    []AttrCheck
	Blocks   []BlockCheck
}

// ResourceCheck is the full generated spec for one resource type.
type ResourceCheck struct {
	Resource string
	Attrs    []AttrCheck
	Blocks   []BlockCheck
}

// RunResourceChecks loads resource content with schema and applies checks.
func RunResourceChecks(runner tflint.Runner, rule tflint.Rule, schema *hclext.BodySchema, spec ResourceCheck) error {
	resources, err := runner.GetResourceContent(spec.Resource, schema, nil)
	if err != nil {
		return err
	}
	for _, resource := range resources.Blocks {
		if err := applyBodyChecks(runner, rule, resource.Body, resource.DefRange, spec.Attrs, spec.Blocks, spec.Resource); err != nil {
			return err
		}
	}
	return nil
}

// RunDataSourceChecks loads data source content with schema and applies checks.
func RunDataSourceChecks(runner tflint.Runner, rule tflint.Rule, schema *hclext.BodySchema, spec ResourceCheck) error {
	content, err := runner.GetModuleContent(&hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{
				Type:       "data",
				LabelNames: []string{"type", "name"},
				Body:       schema,
			},
		},
	}, nil)
	if err != nil {
		return err
	}
	for _, block := range content.Blocks {
		if len(block.Labels) == 0 || block.Labels[0] != spec.Resource {
			continue
		}
		if err := applyBodyChecks(runner, rule, block.Body, block.DefRange, spec.Attrs, spec.Blocks, spec.Resource); err != nil {
			return err
		}
	}
	return nil
}

func applyBodyChecks(runner tflint.Runner, rule tflint.Rule, body *hclext.BodyContent, defRange hcl.Range, attrs []AttrCheck, blocks []BlockCheck, resource string) error {
	for _, ac := range attrs {
		attr, exists := body.Attributes[ac.Name]
		if ac.ComputedOnly {
			if exists {
				rng := defRange
				if attr.Expr != nil {
					rng = attr.Expr.Range()
				}
				if err := runner.EmitIssue(rule, fmt.Sprintf("%q is computed and cannot be set on %s", ac.Name, resource), rng); err != nil {
					return err
				}
			}
			continue
		}
		if !exists {
			if ac.Required {
				if err := runner.EmitIssue(rule, fmt.Sprintf("%q is a required attribute of %s", ac.Name, resource), defRange); err != nil {
					return err
				}
			}
			continue
		}
		if !needsValueCheck(ac) {
			continue
		}
		if err := applyAttrCheck(runner, rule, attr, ac); err != nil {
			return err
		}
	}
	for _, bc := range blocks {
		found := 0
		for _, b := range body.Blocks {
			if b.Type != bc.Type {
				continue
			}
			found++
			if err := applyBodyChecks(runner, rule, b.Body, b.DefRange, bc.Attrs, bc.Blocks, resource+"."+bc.Type); err != nil {
				return err
			}
		}
		if bc.Required && found == 0 {
			if err := runner.EmitIssue(rule, fmt.Sprintf("%q is a required block of %s", bc.Type, resource), defRange); err != nil {
				return err
			}
		}
	}
	return nil
}

func applyAttrCheck(runner tflint.Runner, rule tflint.Rule, attr *hclext.Attribute, ac AttrCheck) error {
	if ac.List || len(ac.Enum) > 0 && ac.Kind == "enum_list" {
		return runner.EvaluateExpr(attr.Expr, func(vals []string) error {
			for _, val := range vals {
				if err := emitStringIssues(runner, rule, attr, ac, val); err != nil {
					return err
				}
			}
			return nil
		}, nil)
	}
	if ac.Kind == "int_range" || ac.Kind == "int_min" || ac.Kind == "int_max" || ac.Kind == "port" || ac.Kind == "int_enum" || len(ac.IntEnum) > 0 || ((ac.Min != nil || ac.Max != nil) && ac.Regexp == "" && len(ac.Enum) == 0) {
		return runner.EvaluateExpr(attr.Expr, func(val int) error {
			return emitIntIssues(runner, rule, attr, ac, val)
		}, nil)
	}
	return runner.EvaluateExpr(attr.Expr, func(val string) error {
		return emitStringIssues(runner, rule, attr, ac, val)
	}, nil)
}

func emitStringIssues(runner tflint.Runner, rule tflint.Rule, attr *hclext.Attribute, ac AttrCheck, val string) error {
	rng := attr.Expr.Range()
	if ac.NoZero && strings.TrimSpace(val) == "" {
		return runner.EmitIssue(rule, fmt.Sprintf("%q must not be empty", ac.Name), rng)
	}
	if len(ac.Enum) > 0 && !inStringEnum(val, ac.Enum) {
		return runner.EmitIssue(rule, fmt.Sprintf("%s is not a valid value for %q. Valid options: %s", shown(ac, val), ac.Name, strings.Join(ac.Enum, ", ")), rng)
	}
	if ac.MinLen != nil && len(val) < *ac.MinLen {
		return runner.EmitIssue(rule, fmt.Sprintf("%q must be at least %d characters", ac.Name, *ac.MinLen), rng)
	}
	if ac.MaxLen != nil && len(val) > *ac.MaxLen {
		return runner.EmitIssue(rule, fmt.Sprintf("%q must be at most %d characters", ac.Name, *ac.MaxLen), rng)
	}
	if ac.Regexp != "" {
		re, err := compiledRegexp(ac.Regexp)
		if err == nil && !re.MatchString(val) {
			if err := runner.EmitIssue(rule, fmt.Sprintf("%s does not match expected format for %q", shown(ac, val), ac.Name), rng); err != nil {
				return err
			}
		}
	}
	switch ac.Kind {
	case "cidr":
		if _, _, err := net.ParseCIDR(val); err != nil {
			if err := runner.EmitIssue(rule, fmt.Sprintf("%s is not a valid CIDR block", shown(ac, val)), rng); err != nil {
				return err
			}
		}
	case "ipv4":
		ip := net.ParseIP(val)
		if ip == nil || ip.To4() == nil {
			if err := runner.EmitIssue(rule, fmt.Sprintf("%s is not a valid IPv4 address", shown(ac, val)), rng); err != nil {
				return err
			}
		}
	case "ip_or_cidr":
		if net.ParseIP(val) == nil {
			if _, _, err := net.ParseCIDR(val); err != nil {
				if err := runner.EmitIssue(rule, fmt.Sprintf("%s is not a valid IP address or CIDR block", shown(ac, val)), rng); err != nil {
					return err
				}
			}
		}
	case "crn":
		if !validCRN(val) {
			if err := runner.EmitIssue(rule, fmt.Sprintf("%s is not a valid CRN", shown(ac, val)), rng); err != nil {
				return err
			}
		}
	case "zone":
		if !validZone(val) {
			if err := runner.EmitIssue(rule, fmt.Sprintf("%s is not a valid availability zone name", shown(ac, val)), rng); err != nil {
				return err
			}
		}
	case "json":
		if !json.Valid([]byte(val)) {
			if err := runner.EmitIssue(rule, fmt.Sprintf("%q is not valid JSON", ac.Name), rng); err != nil {
				return err
			}
		}
	case "uuid":
		if !validUUID(val) {
			if err := runner.EmitIssue(rule, fmt.Sprintf("%s is not a valid UUID", shown(ac, val)), rng); err != nil {
				return err
			}
		}
	case "email":
		if !validEmail(val) {
			if err := runner.EmitIssue(rule, fmt.Sprintf("%s is not a valid email address", shown(ac, val)), rng); err != nil {
				return err
			}
		}
	case "uri":
		if !validURI(val) {
			if err := runner.EmitIssue(rule, fmt.Sprintf("%s is not a valid URI", shown(ac, val)), rng); err != nil {
				return err
			}
		}
	case "hostname":
		if !validHostname(val) {
			if err := runner.EmitIssue(rule, fmt.Sprintf("%s is not a valid hostname", shown(ac, val)), rng); err != nil {
				return err
			}
		}
	}
	return nil
}

func emitIntIssues(runner tflint.Runner, rule tflint.Rule, attr *hclext.Attribute, ac AttrCheck, val int) error {
	rng := attr.Expr.Range()
	if ac.Kind == "port" {
		if val < 1 || val > 65535 {
			return runner.EmitIssue(rule, fmt.Sprintf("%q must be a valid port in the range 1-65535, got %d", ac.Name, val), rng)
		}
	}
	if len(ac.IntEnum) > 0 && !inIntEnum(val, ac.IntEnum) {
		return runner.EmitIssue(rule, fmt.Sprintf("%d is not a valid value for %q", val, ac.Name), rng)
	}
	if ac.Min != nil && val < *ac.Min {
		if err := runner.EmitIssue(rule, fmt.Sprintf("%q must be at least %d, got %d", ac.Name, *ac.Min, val), rng); err != nil {
			return err
		}
	}
	if ac.Max != nil && val > *ac.Max {
		if err := runner.EmitIssue(rule, fmt.Sprintf("%q must be at most %d, got %d", ac.Name, *ac.Max, val), rng); err != nil {
			return err
		}
	}
	return nil
}

// shown renders a value for a message without leaking secrets.
func shown(ac AttrCheck, val string) string {
	return redact.Display(redact.IsSensitive(ac.Name, ac.Sensitive), val)
}

var regexpCache sync.Map

// compiledRegexp compiles each generated pattern once per process.
func compiledRegexp(pattern string) (*regexp.Regexp, error) {
	if re, ok := regexpCache.Load(pattern); ok {
		return re.(*regexp.Regexp), nil
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	regexpCache.Store(pattern, re)
	return re, nil
}

func inStringEnum(val string, enum []string) bool {
	for _, e := range enum {
		if val == e {
			return true
		}
	}
	return false
}

func inIntEnum(val int, enum []int) bool {
	for _, e := range enum {
		if val == e {
			return true
		}
	}
	return false
}

var crnPattern = regexp.MustCompile(`^crn:v[0-9]+:[^:]+:[^:]+:[^:]+:[^:]*:.+$`)

func validCRN(val string) bool {
	return crnPattern.MatchString(val)
}

var zonePattern = regexp.MustCompile(`^[a-z]{2}-[a-z0-9]+(?:-[a-z0-9]+)*-[0-9]+$`)

func validZone(val string) bool {
	return zonePattern.MatchString(val)
}

var uuidPattern = regexp.MustCompile(`(?i)^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)

func validUUID(val string) bool {
	return uuidPattern.MatchString(val)
}

var emailPattern = regexp.MustCompile(`^[^@\s]+@[^@\s]+\.[^@\s]+$`)

func validEmail(val string) bool {
	return emailPattern.MatchString(val)
}

func validURI(val string) bool {
	return strings.Contains(val, "://") || strings.HasPrefix(val, "/")
}

var hostnamePattern = regexp.MustCompile(`^[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`)

func validHostname(val string) bool {
	if val == "" || len(val) > 253 || strings.HasPrefix(val, "-") || strings.HasSuffix(val, "-") {
		return false
	}
	return hostnamePattern.MatchString(val)
}

// BodySchemaFromSpec builds the hclext schema from a ResourceCheck.
func BodySchemaFromSpec(spec ResourceCheck) *hclext.BodySchema {
	return bodySchema(spec.Attrs, spec.Blocks)
}

func bodySchema(attrs []AttrCheck, blocks []BlockCheck) *hclext.BodySchema {
	s := &hclext.BodySchema{}
	for _, a := range attrs {
		s.Attributes = append(s.Attributes, hclext.AttributeSchema{Name: a.Name, Required: false})
	}
	for _, b := range blocks {
		s.Blocks = append(s.Blocks, hclext.BlockSchema{
			Type: b.Type,
			Body: bodySchema(b.Attrs, b.Blocks),
		})
	}
	return s
}

func needsValueCheck(ac AttrCheck) bool {
	return len(ac.Enum) > 0 || len(ac.IntEnum) > 0 || ac.Min != nil || ac.Max != nil ||
		ac.MinLen != nil || ac.MaxLen != nil || ac.Regexp != "" || ac.NoZero || ac.Kind != ""
}

func intPtr(n int) *int { return &n }

// ArgumentRule is a generated per-argument rule. One type serves every
// writable attribute so the compiler is not asked to digest thousands of
// unique structs.
type ArgumentRule struct {
	tflint.DefaultRule
	ruleName string
	link     string
	spec     ResourceCheck
	run      func(tflint.Runner, tflint.Rule, *hclext.BodySchema, ResourceCheck) error
}

// NewArgumentRule constructs an attribute-level schema rule.
func NewArgumentRule(name, link string, spec ResourceCheck, run func(tflint.Runner, tflint.Rule, *hclext.BodySchema, ResourceCheck) error) *ArgumentRule {
	return &ArgumentRule{ruleName: name, link: link, spec: spec, run: run}
}

func (r *ArgumentRule) Name() string { return r.ruleName }

func (r *ArgumentRule) Enabled() bool { return true }

func (r *ArgumentRule) Severity() tflint.Severity { return tflint.ERROR }

func (r *ArgumentRule) Link() string { return r.link }

func (r *ArgumentRule) Check(runner tflint.Runner) error {
	if r.run == nil {
		return nil
	}
	return r.run(runner, r, BodySchemaFromSpec(r.spec), r.spec)
}
