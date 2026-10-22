// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"text/template"
	"unicode"
)

func main() {
	inDir := flag.String("in", "extractor/output", "extractor JSON directory")
	rulesDir := flag.String("rules", "rules", "output directory for generated rules")
	docsDir := flag.String("docs", "docs/rules", "output directory for generated docs")
	packages := flag.String("packages", "", "comma-separated packages (empty = all JSON files)")
	flag.Parse()

	want := map[string]bool{}
	if strings.TrimSpace(*packages) != "" && *packages != "all" {
		for _, p := range strings.Split(*packages, ",") {
			p = strings.TrimSpace(p)
			if p != "" {
				want[p] = true
			}
		}
	}

	entries, err := os.ReadDir(*inDir)
	if err != nil {
		fatalf("read %s: %v", *inDir, err)
	}

	var names []string
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".json") || e.Name() == "index.json" || e.Name() == "overlay-index.json" || strings.HasSuffix(e.Name(), ".overlay.json") {
			continue
		}
		pkg := strings.TrimSuffix(e.Name(), ".json")
		if len(want) > 0 && !want[pkg] {
			continue
		}
		names = append(names, e.Name())
	}
	sort.Strings(names)

	var all []ResourceJSON
	for _, name := range names {
		raw, err := os.ReadFile(filepath.Join(*inDir, name))
		if err != nil {
			fatalf("read %s: %v", name, err)
		}
		var pj PackageJSON
		if err := json.Unmarshal(raw, &pj); err != nil {
			fatalf("parse %s: %v", name, err)
		}
		all = append(all, pj.Resources...)
	}
	sort.Slice(all, func(i, j int) bool {
		if all[i].Name != all[j].Name {
			return all[i].Name < all[j].Name
		}
		if all[i].Kind != all[j].Kind {
			return all[i].Kind < all[j].Kind
		}
		return all[i].Package < all[j].Package
	})

	if err := os.MkdirAll(*rulesDir, 0o755); err != nil {
		fatalf("mkdir rules: %v", err)
	}
	if err := os.MkdirAll(*docsDir, 0o755); err != nil {
		fatalf("mkdir docs: %v", err)
	}

	overlays := loadOverlays(*inDir)

	var generated []string
	var typeNames []string
	var attrRuleCount int
	for _, res := range all {
		id := identOf(res)
		spec := buildSpec(res)
		if ov, ok := overlays[overlayKey{Name: res.Name, Kind: res.Kind}]; ok {
			applyOverlay(&spec, ov)
		} else if ov, ok := overlays[overlayKey{Name: res.Name, Kind: ""}]; ok {
			applyOverlay(&spec, ov)
		}
		checkNames(res, spec)
		attrRules := collectAttrRules(id, spec)
		src, err := renderRule(res, spec, id, attrRules)
		if err != nil {
			fatalf("render %s: %v", id.FileBase, err)
		}
		formatted, err := format.Source(src)
		if err != nil {
			fatalf("gofmt %s: %v\n%s", id.FileBase, err, src)
		}
		rulePath := filepath.Join(*rulesDir, id.FileBase+"_gen.go")
		if err := os.WriteFile(rulePath, formatted, 0o644); err != nil {
			fatalf("write %s: %v", rulePath, err)
		}
		testSrc, err := renderTest(res, spec, id, attrRules)
		if err != nil {
			fatalf("render test %s: %v", id.FileBase, err)
		}
		testFmt, err := format.Source(testSrc)
		if err != nil {
			fatalf("gofmt test %s: %v\n%s", id.FileBase, err, testSrc)
		}
		testPath := filepath.Join(*rulesDir, id.FileBase+"_gen_test.go")
		if err := os.WriteFile(testPath, testFmt, 0o644); err != nil {
			fatalf("write %s: %v", testPath, err)
		}
		doc := renderDoc(res, spec, id, attrRules)
		docPath := filepath.Join(*docsDir, id.FileBase+".md")
		if err := os.WriteFile(docPath, []byte(doc), 0o644); err != nil {
			fatalf("write %s: %v", docPath, err)
		}
		generated = append(generated, id.FileBase)
		typeNames = append(typeNames, id.TypeName)
		for _, ar := range attrRules {
			typeNames = append(typeNames, ar.TypeName)
		}
		attrRuleCount += len(attrRules)
		fmt.Printf("generated %s (%d attribute rules)\n", id.FileBase, len(attrRules))
	}

	if len(want) == 0 {
		providerSrc, err := renderProvider(typeNames)
		if err != nil {
			fatalf("provider: %v", err)
		}
		providerFmt, err := format.Source(providerSrc)
		if err != nil {
			fatalf("gofmt provider: %v\n%s", err, providerSrc)
		}
		if err := os.WriteFile(filepath.Join(*rulesDir, "provider.go"), providerFmt, 0o644); err != nil {
			fatalf("write provider: %v", err)
		}

		if err := writeDocsIndex(*docsDir, all); err != nil {
			fatalf("docs index: %v", err)
		}

		if err := removeOrphans(*rulesDir, *docsDir, generated); err != nil {
			fatalf("orphans: %v", err)
		}
	}

	fmt.Printf("done: %d type rules, %d attribute rules (%d total)\n", len(generated), attrRuleCount, len(typeNames))
}

func fatalf(format string, args ...any) {
	fmt.Fprintf(os.Stderr, "generator: "+format+"\n", args...)
	os.Exit(1)
}

type PackageJSON struct {
	Package   string         `json:"package"`
	Resources []ResourceJSON `json:"resources"`
	Skipped   []struct {
		Name   string `json:"name"`
		Reason string `json:"reason"`
	} `json:"skipped"`
}

type ResourceJSON struct {
	Name        string               `json:"name"`
	Function    string               `json:"function"`
	Package     string               `json:"package"`
	Kind        string               `json:"kind"`
	Description string               `json:"description"`
	Attributes  map[string]*AttrJSON `json:"attributes"`
}

type AttrJSON struct {
	Name        string        `json:"name"`
	Type        string        `json:"type"`
	Required    bool          `json:"required"`
	Optional    bool          `json:"optional"`
	Computed    bool          `json:"computed"`
	ForceNew    bool          `json:"force_new"`
	Sensitive   bool          `json:"sensitive"`
	Description string        `json:"description"`
	Default     string        `json:"default"`
	Validate    *ValidateJSON `json:"validate"`
	Elem        *ElemJSON     `json:"elem"`
	MinItems    *int          `json:"min_items"`
	MaxItems    *int          `json:"max_items"`
}

type ElemJSON struct {
	Kind       string               `json:"kind"`
	Type       string               `json:"type"`
	Validate   *ValidateJSON        `json:"validate"`
	Attributes map[string]*AttrJSON `json:"attributes"`
}

type ValidateJSON struct {
	Kind     string   `json:"kind"`
	FuncName string   `json:"func_name"`
	Enum     []string `json:"enum"`
	IntEnum  []int    `json:"int_enum"`
	Min      *int     `json:"min"`
	Max      *int     `json:"max"`
	MinLen   *int     `json:"min_len"`
	MaxLen   *int     `json:"max_len"`
	Regexp   string   `json:"regexp"`
	NoZero   bool     `json:"no_zero"`
}

type specAttr struct {
	Name           string
	Type           string
	Required       bool
	Enum           []string
	IntEnum        []int
	Min            *int
	Max            *int
	MinLen         *int
	MaxLen         *int
	Regexp         string
	Kind           string
	NoZero         bool
	List           bool
	ComputedOnly   bool
	Sensitive      bool
	APIDescription string
}

type specBlock struct {
	Type     string
	Required bool
	Attrs    []specAttr
	Blocks   []specBlock
}

type resourceSpec struct {
	Resource  string
	Attrs     []specAttr
	Blocks    []specBlock
	APISchema string
	APISpec   string
}

func configurable(a *AttrJSON) bool {
	if a.Required && !a.Computed {
		return true
	}
	if a.Optional {
		return true
	}
	return false
}

func isNestedBlock(a *AttrJSON) bool {
	if a.Elem == nil || a.Elem.Kind != "resource" {
		return false
	}
	return a.Type == "list" || a.Type == "set"
}

func buildSpec(res ResourceJSON) resourceSpec {
	spec := resourceSpec{Resource: res.Name}
	names := sortedKeys(res.Attributes)
	for _, name := range names {
		a := res.Attributes[name]
		if a == nil {
			continue
		}
		if a.Computed && !a.Optional && !a.Required {
			spec.Attrs = append(spec.Attrs, specAttr{
				Name:         a.Name,
				Type:         a.Type,
				ComputedOnly: true,
				List:         a.Type == "list" || a.Type == "set",
			})
			continue
		}
		if !configurable(a) {
			continue
		}
		if isNestedBlock(a) {
			spec.Blocks = append(spec.Blocks, buildBlock(name, a))
			continue
		}
		spec.Attrs = append(spec.Attrs, buildAttr(a))
	}
	return spec
}

func buildBlock(name string, a *AttrJSON) specBlock {
	b := specBlock{Type: name, Required: a.Required && !a.Computed}
	if a.Elem != nil {
		nestedNames := sortedKeys(a.Elem.Attributes)
		for _, n := range nestedNames {
			na := a.Elem.Attributes[n]
			if na == nil || !configurable(na) {
				continue
			}
			if isNestedBlock(na) {
				b.Blocks = append(b.Blocks, buildBlock(n, na))
				continue
			}
			b.Attrs = append(b.Attrs, buildAttr(na))
		}
	}
	return b
}

func buildAttr(a *AttrJSON) specAttr {
	s := specAttr{
		Name:      a.Name,
		Type:      a.Type,
		Required:  a.Required && !a.Computed,
		List:      a.Type == "list" || a.Type == "set",
		Sensitive: a.Sensitive,
	}
	if a.Validate != nil {
		applyValidate(&s, a.Validate)
	}
	if a.Elem != nil && a.Elem.Kind == "schema" && a.Elem.Validate != nil {
		applyValidate(&s, a.Elem.Validate)
		s.List = true
	}
	if s.Kind == "" && s.List && a.Type != "map" {
		// keep list flag for EvaluateExpr []string when we have string constraints
		if s.Regexp == "" && len(s.Enum) == 0 && !s.NoZero {
			s.List = false
		}
	}
	normalizeConstraints(&s)
	return s
}

func normalizeConstraints(s *specAttr) {
	if s.Regexp != "" && trivialRegexp(s.Regexp) {
		s.Regexp = ""
		if s.Kind == "regexp" {
			s.Kind = ""
		}
		s.NoZero = true
	}
	if s.Regexp != "" {
		dummy := stringMatching(s.Regexp, s.MinLen, s.MaxLen)
		re, err := regexp.Compile(s.Regexp)
		if err != nil || !re.MatchString(dummy) {
			s.Regexp = ""
			if s.Kind == "regexp" {
				s.Kind = ""
			}
		}
	}
	switch s.Kind {
	case "cidr", "ipv4", "crn", "zone", "uuid", "email", "uri", "hostname", "json", "ip_or_cidr", "port":
		s.Regexp = ""
	}
}

func trivialRegexp(p string) bool {
	switch p {
	case "^.+$", "^.*$", ".*", ".+", "^(.*)$", "(.*?)", "^(.*?)$", "^[\\s\\S]+$", "^[\\s\\S]*$":
		return true
	}
	re, err := regexp.Compile(p)
	if err != nil {
		return true
	}
	if re.MatchString("") && re.MatchString("!!!invalid!!!") {
		return true
	}
	return false
}

func applyValidate(s *specAttr, v *ValidateJSON) {
	if v == nil {
		return
	}
	s.NoZero = s.NoZero || v.NoZero
	if len(v.Enum) > 0 {
		s.Enum = v.Enum
	}
	if len(v.IntEnum) > 0 {
		s.IntEnum = v.IntEnum
	}
	if v.Min != nil {
		s.Min = v.Min
	}
	if v.Max != nil {
		s.Max = v.Max
	}
	if v.MinLen != nil {
		s.MinLen = v.MinLen
	}
	if v.MaxLen != nil {
		s.MaxLen = v.MaxLen
	}
	if v.Regexp != "" {
		s.Regexp = v.Regexp
	}
	switch v.Kind {
	case "invoke", "func", "cloud_data":
		// ignore as a check kind unless we already have constraints
	default:
		if s.Kind == "" || s.Kind == "no_zero" {
			s.Kind = v.Kind
		}
	}
}

func sortedKeys[T any](m map[string]T) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func toPascal(name string) string {
	parts := strings.Split(name, "_")
	var b strings.Builder
	for _, p := range parts {
		if p == "" {
			continue
		}
		runes := []rune(p)
		runes[0] = unicode.ToUpper(runes[0])
		b.WriteString(string(runes))
	}
	return b.String()
}

func goLitStringSlice(vals []string) string {
	if len(vals) == 0 {
		return "nil"
	}
	var b strings.Builder
	b.WriteString("[]string{")
	for i, v := range vals {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteString(strconv.Quote(v))
	}
	b.WriteString("}")
	return b.String()
}

func goLitIntSlice(vals []int) string {
	if len(vals) == 0 {
		return "nil"
	}
	var b strings.Builder
	b.WriteString("[]int{")
	for i, v := range vals {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteString(strconv.Itoa(v))
	}
	b.WriteString("}")
	return b.String()
}

func goLitIntPtr(v *int) string {
	if v == nil {
		return "nil"
	}
	return fmt.Sprintf("intPtr(%d)", *v)
}

var ruleTemplate = template.Must(template.New("rule").Funcs(template.FuncMap{
	"pascal":     toPascal,
	"quote":      strconv.Quote,
	"strings":    goLitStringSlice,
	"ints":       goLitIntSlice,
	"intptr":     goLitIntPtr,
	"attrfields": attrFields,
	"blocklit":   blockLit,
}).Parse(`// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
// Code generated by generator; DO NOT EDIT.

package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/IBM/tflint-ruleset-ibm/project"
)

// {{ .TypeName }}Rule validates the schema of {{ .Name }} {{ .Noun }}.
type {{ .TypeName }}Rule struct {
	tflint.DefaultRule
	resourceType string
}

// New{{ .TypeName }}Rule returns a new rule.
func New{{ .TypeName }}Rule() *{{ .TypeName }}Rule {
	return &{{ .TypeName }}Rule{resourceType: {{ quote .Name }}}
}

func (r *{{ .TypeName }}Rule) Name() string { return {{ quote .RuleName }} }

func (r *{{ .TypeName }}Rule) Enabled() bool { return true }

func (r *{{ .TypeName }}Rule) Severity() tflint.Severity { return tflint.ERROR }

func (r *{{ .TypeName }}Rule) Link() string { return project.ReferenceLink({{ quote .FileBase }}) }

func (r *{{ .TypeName }}Rule) Check(runner tflint.Runner) error {
	spec := {{ .TypeName }}Spec()
	return {{ .CheckFn }}(runner, r, BodySchemaFromSpec(spec), spec)
}

func {{ .TypeName }}Spec() ResourceCheck {
	return ResourceCheck{
		Resource: {{ quote .Name }},
		Attrs: []AttrCheck{
{{- range .Spec.Attrs }}
			{ {{ attrfields . }} },
{{- end }}
		},
		Blocks: []BlockCheck{
{{- range .Spec.Blocks }}
			{{ blocklit . }},
{{- end }}
		},
	}
}
`))

func attrFields(a specAttr) string {
	var parts []string
	parts = append(parts, "Name: "+strconv.Quote(a.Name))
	if a.Required {
		parts = append(parts, "Required: true")
	}
	if len(a.Enum) > 0 {
		parts = append(parts, "Enum: "+goLitStringSlice(a.Enum))
	}
	if len(a.IntEnum) > 0 {
		parts = append(parts, "IntEnum: "+goLitIntSlice(a.IntEnum))
	}
	if a.Min != nil {
		parts = append(parts, "Min: "+goLitIntPtr(a.Min))
	}
	if a.Max != nil {
		parts = append(parts, "Max: "+goLitIntPtr(a.Max))
	}
	if a.MinLen != nil {
		parts = append(parts, "MinLen: "+goLitIntPtr(a.MinLen))
	}
	if a.MaxLen != nil {
		parts = append(parts, "MaxLen: "+goLitIntPtr(a.MaxLen))
	}
	if a.Regexp != "" {
		parts = append(parts, "Regexp: "+strconv.Quote(a.Regexp))
	}
	if a.Kind != "" && a.Kind != "no_zero" {
		parts = append(parts, "Kind: "+strconv.Quote(a.Kind))
	}
	if a.NoZero {
		parts = append(parts, "NoZero: true")
	}
	if a.List {
		parts = append(parts, "List: true")
	}
	if a.ComputedOnly {
		parts = append(parts, "ComputedOnly: true")
	}
	if a.Sensitive {
		parts = append(parts, "Sensitive: true")
	}
	return strings.Join(parts, ", ")
}

func blockLit(b specBlock) string {
	var sb strings.Builder
	sb.WriteString("{ Type: " + strconv.Quote(b.Type))
	if b.Required {
		sb.WriteString(", Required: true")
	}
	sb.WriteString(", Attrs: []AttrCheck{")
	for i, a := range b.Attrs {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString("{ " + attrFields(a) + " }")
	}
	sb.WriteString("}, Blocks: []BlockCheck{")
	for i, nb := range b.Blocks {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(blockLit(nb))
	}
	sb.WriteString("} }")
	return sb.String()
}

type ident struct {
	Name     string
	RuleName string
	TypeName string
	FileBase string
	Block    string
	Noun     string
	CheckFn  string
}

func identOf(res ResourceJSON) ident {
	if res.Kind == "data" {
		return ident{
			Name:     res.Name,
			RuleName: res.Name + "_data",
			TypeName: toPascal(res.Name) + "Data",
			FileBase: "data_" + res.Name,
			Block:    "data",
			Noun:     "data sources",
			CheckFn:  "RunDataSourceChecks",
		}
	}
	return ident{
		Name:     res.Name,
		RuleName: res.Name,
		TypeName: toPascal(res.Name),
		FileBase: res.Name,
		Block:    "resource",
		Noun:     "resources",
		CheckFn:  "RunResourceChecks",
	}
}

type ruleView struct {
	Name     string
	RuleName string
	TypeName string
	FileBase string
	Noun     string
	CheckFn  string
	Spec     resourceSpec
}

func renderRule(res ResourceJSON, spec resourceSpec, id ident, attrRules []attrRule) ([]byte, error) {
	var buf bytes.Buffer
	err := ruleTemplate.Execute(&buf, ruleView{
		Name:     id.Name,
		RuleName: id.RuleName,
		TypeName: id.TypeName,
		FileBase: id.FileBase,
		Noun:     id.Noun,
		CheckFn:  id.CheckFn,
		Spec:     spec,
	})
	if err != nil {
		return nil, err
	}
	buf.WriteString(renderAttrRules(id, attrRules))
	return buf.Bytes(), nil
}

func renderProvider(typeNames []string) ([]byte, error) {
	var b strings.Builder
	b.WriteString("// Copyright IBM Corp. 2026 All Rights Reserved.\n")
	b.WriteString("// Licensed under the Mozilla Public License v2.0\n")
	b.WriteString("// Code generated by generator; DO NOT EDIT.\n\npackage rules\n\n")
	b.WriteString("import \"github.com/terraform-linters/tflint-plugin-sdk/tflint\"\n\n")
	b.WriteString("// Rules is the complete generated ruleset.\n")
	b.WriteString("var Rules = []tflint.Rule{\n")
	for _, n := range typeNames {
		b.WriteString("\tNew" + n + "Rule(),\n")
	}
	b.WriteString("}\n")
	return []byte(b.String()), nil
}

func renderTest(res ResourceJSON, spec resourceSpec, id ident, attrRules []attrRule) ([]byte, error) {
	passHCL := passingHCL(id.Block, res.Name, spec)
	failHCL, failMsg, hasFail := failingHCL(id.Block, res.Name, spec)
	var b strings.Builder
	b.WriteString("// Copyright IBM Corp. 2026 All Rights Reserved.\n")
	b.WriteString("// Licensed under the Mozilla Public License v2.0\n")
	b.WriteString("// Code generated by generator; DO NOT EDIT.\n\npackage rules\n\n")
	b.WriteString("import (\n\t\"testing\"\n\n\t\"github.com/terraform-linters/tflint-plugin-sdk/helper\"\n")
	if len(attrRules) > 0 {
		b.WriteString("\t\"github.com/terraform-linters/tflint-plugin-sdk/tflint\"\n")
	}
	b.WriteString(")\n\n")
	b.WriteString("func Test_" + id.TypeName + "(t *testing.T) {\n")
	b.WriteString("\tcases := []struct {\n\t\tName     string\n\t\tContent  string\n\t\tExpected helper.Issues\n\t}{\n")
	b.WriteString("\t\t{\n\t\t\tName: \"valid configuration\",\n")
	b.WriteString("\t\t\tContent: `\n" + fixture(passHCL) + "`,\n")
	b.WriteString("\t\t\tExpected: helper.Issues{},\n\t\t},\n")
	if hasFail {
		b.WriteString("\t\t{\n\t\t\tName: \"invalid configuration\",\n")
		b.WriteString("\t\t\tContent: `\n" + fixture(failHCL) + "`,\n")
		b.WriteString("\t\t\tExpected: helper.Issues{\n\t\t\t\t{\n")
		b.WriteString("\t\t\t\t\tRule:    New" + id.TypeName + "Rule(),\n")
		b.WriteString("\t\t\t\t\tMessage: " + strconv.Quote(failMsg) + ",\n")
		b.WriteString("\t\t\t\t},\n\t\t\t},\n\t\t},\n")
	}
	b.WriteString("\t}\n\n")
	b.WriteString("\trule := New" + id.TypeName + "Rule()\n")
	b.WriteString("\tfor _, tc := range cases {\n")
	b.WriteString("\t\tt.Run(tc.Name, func(t *testing.T) {\n")
	b.WriteString("\t\t\trunner := helper.TestRunner(t, map[string]string{\"resource.tf\": tc.Content})\n")
	b.WriteString("\t\t\tif err := rule.Check(runner); err != nil {\n")
	b.WriteString("\t\t\t\tt.Fatalf(\"Unexpected error occurred: %s\", err)\n")
	b.WriteString("\t\t\t}\n")
	b.WriteString("\t\t\thelper.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)\n")
	b.WriteString("\t\t})\n\t}\n}\n")
	b.WriteString(renderAttrTests(id, spec, attrRules))
	return []byte(b.String()), nil
}

func hclHeader(block, resource string) string {
	return block + " \"" + resource + "\" \"test\" {\n"
}

func passingHCL(block, resource string, spec resourceSpec) string {
	return hclHeader(block, resource) + bodyHCL(spec.Attrs, spec.Blocks, "  ", false) + "}\n"
}

func failingHCL(block, resource string, spec resourceSpec) (string, string, bool) {
	for _, a := range spec.Attrs {
		if a.Required && !a.ComputedOnly {
			hcl := hclHeader(block, resource) + bodyHCL(spec.Attrs, spec.Blocks, "  ", true) + "}\n"
			return hcl, fmt.Sprintf("%q is a required attribute of %s", a.Name, resource), true
		}
	}
	for _, b := range spec.Blocks {
		if b.Required {
			hcl := hclHeader(block, resource) + bodyHCL(spec.Attrs, spec.Blocks, "  ", true) + "}\n"
			return hcl, fmt.Sprintf("%q is a required block of %s", b.Type, resource), true
		}
	}
	for _, a := range spec.Attrs {
		if len(a.Enum) > 0 {
			val := "not-a-valid-enum-value"
			hcl := hclHeader(block, resource) + bodyHCLWithOverride(spec.Attrs, spec.Blocks, "  ", a.Name, strconv.Quote(val)) + "}\n"
			msg := fmt.Sprintf("%s is not a valid value for %q. Valid options: %s", shown(a, val), a.Name, strings.Join(a.Enum, ", "))
			return hcl, msg, true
		}
		if len(a.IntEnum) > 0 {
			hcl := hclHeader(block, resource) + bodyHCLWithOverride(spec.Attrs, spec.Blocks, "  ", a.Name, "999999") + "}\n"
			msg := fmt.Sprintf("%d is not a valid value for %q", 999999, a.Name)
			return hcl, msg, true
		}
		switch a.Kind {
		case "cidr":
			hcl := hclHeader(block, resource) + bodyHCLWithOverride(spec.Attrs, spec.Blocks, "  ", a.Name, strconv.Quote("not-a-cidr")) + "}\n"
			return hcl, fmt.Sprintf("%s is not a valid CIDR block", shown(a, "not-a-cidr")), true
		case "ipv4":
			hcl := hclHeader(block, resource) + bodyHCLWithOverride(spec.Attrs, spec.Blocks, "  ", a.Name, strconv.Quote("not-an-ip")) + "}\n"
			return hcl, fmt.Sprintf("%s is not a valid IPv4 address", shown(a, "not-an-ip")), true
		case "ip_or_cidr":
			hcl := hclHeader(block, resource) + bodyHCLWithOverride(spec.Attrs, spec.Blocks, "  ", a.Name, strconv.Quote("nope")) + "}\n"
			return hcl, fmt.Sprintf("%s is not a valid IP address or CIDR block", shown(a, "nope")), true
		case "crn":
			hcl := hclHeader(block, resource) + bodyHCLWithOverride(spec.Attrs, spec.Blocks, "  ", a.Name, strconv.Quote("not-a-crn")) + "}\n"
			return hcl, fmt.Sprintf("%s is not a valid CRN", shown(a, "not-a-crn")), true
		case "zone":
			hcl := hclHeader(block, resource) + bodyHCLWithOverride(spec.Attrs, spec.Blocks, "  ", a.Name, strconv.Quote("not-a-zone")) + "}\n"
			return hcl, fmt.Sprintf("%s is not a valid availability zone name", shown(a, "not-a-zone")), true
		case "port":
			hcl := hclHeader(block, resource) + bodyHCLWithOverride(spec.Attrs, spec.Blocks, "  ", a.Name, "0") + "}\n"
			return hcl, fmt.Sprintf("%q must be a valid port in the range 1-65535, got %d", a.Name, 0), true
		case "json":
			hcl := hclHeader(block, resource) + bodyHCLWithOverride(spec.Attrs, spec.Blocks, "  ", a.Name, strconv.Quote("not-json")) + "}\n"
			return hcl, fmt.Sprintf("%q is not valid JSON", a.Name), true
		case "uuid":
			bad := padToLen("not-a-uuid", a.MinLen, a.MaxLen)
			hcl := hclHeader(block, resource) + bodyHCLWithOverride(spec.Attrs, spec.Blocks, "  ", a.Name, strconv.Quote(bad)) + "}\n"
			return hcl, fmt.Sprintf("%s is not a valid UUID", shown(a, bad)), true
		case "email":
			hcl := hclHeader(block, resource) + bodyHCLWithOverride(spec.Attrs, spec.Blocks, "  ", a.Name, strconv.Quote("not-an-email")) + "}\n"
			return hcl, fmt.Sprintf("%s is not a valid email address", shown(a, "not-an-email")), true
		case "uri":
			bad := padToLen("not a uri", a.MinLen, a.MaxLen)
			hcl := hclHeader(block, resource) + bodyHCLWithOverride(spec.Attrs, spec.Blocks, "  ", a.Name, strconv.Quote(bad)) + "}\n"
			return hcl, fmt.Sprintf("%s is not a valid URI", shown(a, bad)), true
		case "hostname":
			hcl := hclHeader(block, resource) + bodyHCLWithOverride(spec.Attrs, spec.Blocks, "  ", a.Name, strconv.Quote("-bad-host")) + "}\n"
			return hcl, fmt.Sprintf("%s is not a valid hostname", shown(a, "-bad-host")), true
		}
		if a.Regexp != "" {
			dummy := regexpFailDummy(a)
			if dummy != "" {
				hcl := hclHeader(block, resource) + bodyHCLWithOverride(spec.Attrs, spec.Blocks, "  ", a.Name, strconv.Quote(dummy)) + "}\n"
				return hcl, fmt.Sprintf("%s does not match expected format for %q", shown(a, dummy), a.Name), true
			}
		}
		if a.Min != nil {
			bad := *a.Min - 1
			hcl := hclHeader(block, resource) + bodyHCLWithOverride(spec.Attrs, spec.Blocks, "  ", a.Name, strconv.Itoa(bad)) + "}\n"
			return hcl, fmt.Sprintf("%q must be at least %d, got %d", a.Name, *a.Min, bad), true
		}
		if a.Max != nil {
			bad := *a.Max + 1
			hcl := hclHeader(block, resource) + bodyHCLWithOverride(spec.Attrs, spec.Blocks, "  ", a.Name, strconv.Itoa(bad)) + "}\n"
			return hcl, fmt.Sprintf("%q must be at most %d, got %d", a.Name, *a.Max, bad), true
		}
		if a.NoZero {
			hcl := hclHeader(block, resource) + bodyHCLWithOverride(spec.Attrs, spec.Blocks, "  ", a.Name, strconv.Quote("")) + "}\n"
			return hcl, fmt.Sprintf("%q must not be empty", a.Name), true
		}
		if a.MinLen != nil && *a.MinLen > 0 {
			hcl := hclHeader(block, resource) + bodyHCLWithOverride(spec.Attrs, spec.Blocks, "  ", a.Name, strconv.Quote("")) + "}\n"
			return hcl, fmt.Sprintf("%q must be at least %d characters", a.Name, *a.MinLen), true
		}
	}
	for _, a := range spec.Attrs {
		if !a.ComputedOnly {
			continue
		}
		val := strconv.Quote("computed")
		if a.List {
			val = "[]"
		}
		hcl := hclHeader(block, resource) + bodyHCLWithOverride(spec.Attrs, spec.Blocks, "  ", a.Name, val) + "}\n"
		return hcl, fmt.Sprintf("%q is computed and cannot be set on %s", a.Name, resource), true
	}
	return passingHCL(block, resource, spec), "", false
}

func includeInFixture(a specAttr) bool {
	return a.Required && !a.ComputedOnly
}

func hasConstraint(a specAttr) bool {
	return len(a.Enum) > 0 || len(a.IntEnum) > 0 || a.Min != nil || a.Max != nil ||
		a.MinLen != nil || a.MaxLen != nil || a.Regexp != "" || a.NoZero ||
		(a.Kind != "" && a.Kind != "no_zero")
}

func includeBlock(b specBlock) bool {
	return b.Required
}

func bodyHCL(attrs []specAttr, blocks []specBlock, indent string, skipFirstRequired bool) string {
	var b strings.Builder
	skipped := false
	for _, a := range attrs {
		if !includeInFixture(a) {
			continue
		}
		if skipFirstRequired && a.Required && !skipped {
			skipped = true
			continue
		}
		b.WriteString(indent + a.Name + " = " + dummyValue(a) + "\n")
	}
	for _, bl := range blocks {
		if !includeBlock(bl) {
			continue
		}
		if skipFirstRequired && bl.Required && !skipped {
			skipped = true
			continue
		}
		b.WriteString(indent + bl.Type + " {\n")
		b.WriteString(bodyHCL(bl.Attrs, bl.Blocks, indent+"  ", false))
		b.WriteString(indent + "}\n")
	}
	return b.String()
}

func bodyHCLWithOverride(attrs []specAttr, blocks []specBlock, indent, name, value string) string {
	var b strings.Builder
	for _, a := range attrs {
		if !includeInFixture(a) && a.Name != name {
			continue
		}
		if a.Name == name {
			v := value
			if a.List && !strings.HasPrefix(strings.TrimSpace(v), "[") {
				v = "[" + v + "]"
			}
			b.WriteString(indent + a.Name + " = " + v + "\n")
			continue
		}
		b.WriteString(indent + a.Name + " = " + dummyValue(a) + "\n")
	}
	for _, bl := range blocks {
		if !includeBlock(bl) {
			continue
		}
		b.WriteString(indent + bl.Type + " {\n")
		b.WriteString(bodyHCLWithOverride(bl.Attrs, bl.Blocks, indent+"  ", name, value))
		b.WriteString(indent + "}\n")
	}
	return b.String()
}

func dummyValue(a specAttr) string {
	scalar := dummyScalar(a)
	if a.List {
		return "[" + scalar + "]"
	}
	return scalar
}

func dummyScalar(a specAttr) string {
	n := strings.ToLower(a.Name)
	if a.Type == "bool" {
		return "true"
	}
	if len(a.Enum) > 0 {
		return strconv.Quote(a.Enum[0])
	}
	if len(a.IntEnum) > 0 {
		return strconv.Itoa(a.IntEnum[0])
	}
	switch a.Kind {
	case "cidr":
		return strconv.Quote("10.0.0.0/24")
	case "ipv4", "ip_or_cidr":
		return strconv.Quote("10.0.0.1")
	case "crn":
		return strconv.Quote(padToLen("crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::", a.MinLen, a.MaxLen))
	case "zone":
		return strconv.Quote("us-south-1")
	case "port":
		return "80"
	case "json":
		return strconv.Quote(`{"ok":true}`)
	case "uuid":
		return strconv.Quote("01234567-89ab-cdef-0123-456789abcdef")
	case "email":
		return strconv.Quote("user@example.com")
	case "uri":
		return strconv.Quote(padToLen("https://example.com", a.MinLen, a.MaxLen))
	case "hostname":
		return strconv.Quote("example.com")
	}
	if a.Type == "int" || a.Type == "float" || a.Kind == "int_range" || a.Kind == "int_min" || a.Kind == "int_max" || a.Kind == "int_enum" || ((a.Min != nil || a.Max != nil) && a.Regexp == "" && len(a.Enum) == 0 && a.Kind == "") {
		if a.Min != nil {
			return strconv.Itoa(*a.Min)
		}
		return "1"
	}
	if strings.Contains(n, "crn") || strings.HasSuffix(n, "_guid") || n == "resource_instance_id" || n == "service_instance" {
		if a.Regexp == "" && a.Kind != "crn" {
			return strconv.Quote("crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::")
		}
	}
	if strings.Contains(n, "uuid") || strings.HasSuffix(n, "_id") || strings.HasSuffix(n, "_guid") {
		if a.Regexp == "" {
			return strconv.Quote("01234567-89ab-cdef-0123-456789abcdef")
		}
	}
	if strings.Contains(n, "cidr") {
		return strconv.Quote("10.0.0.0/24")
	}
	if n == "address" || strings.Contains(n, "ipv4") || strings.HasSuffix(n, "_ip") {
		return strconv.Quote("10.0.0.1")
	}
	if strings.Contains(n, "cron") {
		return strconv.Quote("0 0 * * *")
	}
	if strings.Contains(n, "api_version") {
		return strconv.Quote("2024-01-01")
	}
	if a.Regexp != "" {
		return strconv.Quote(stringMatching(a.Regexp, a.MinLen, a.MaxLen))
	}
	if a.MinLen != nil || a.MaxLen != nil {
		n := 1
		if a.MinLen != nil && *a.MinLen > n {
			n = *a.MinLen
		}
		if n < 7 {
			n = 7
		}
		if a.MaxLen != nil && *a.MaxLen > 0 && n > *a.MaxLen {
			n = *a.MaxLen
		}
		if n < 1 {
			n = 1
		}
		return strconv.Quote(strings.Repeat("a", n))
	}
	if a.Type == "int" || a.Type == "float" || a.Kind == "int_range" || a.Kind == "int_min" || a.Kind == "int_max" || a.Kind == "int_enum" {
		if a.Min != nil {
			return strconv.Itoa(*a.Min)
		}
		return "1"
	}
	if a.Min != nil || a.Max != nil {
		if a.Min != nil {
			return strconv.Itoa(*a.Min)
		}
		return "1"
	}
	return strconv.Quote("example")
}

func stringMatching(pattern string, minLen, maxLen *int) string {
	candidates := []string{
		"example-name", "example", "test", "foo1", "my-resource",
		"env:prod", "foo:bar", "key:value",
		"10.0.0.0/24", "10.0.0.1",
		"0 0 * * *", "* * * * *",
		"us-south-1", "http://example.com", "https://github.com/org/repo.git",
		"https://example.com", "user@example.com",
		"2024-01-01", "v1", "1", "10", "12", "99", "0", "01",
		"1s", "10s", "1h", "24h", "1d", "60s", "1m",
		"01234567-89ab-cdef-0123-456789abcdef",
		"01234567-89ab-4cde-8123-456789abcdef",
		"0123456789abcdef0123456789abcdef",
		"crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::",
		"crn:v1:bluemix:public:is:us-south:a/account:vpc:identifier",
		"crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef",
		"crn:v1:bluemix:public:hs-crypto:us-south:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef:key:identifier",
		"terraform_v1.5", "terraform_v1.5.0",
		"python-3.11", "nodejs-20",
		"1.0.0_1", "1",
		"AccessGroupId-01234567-89ab-cdef-0123-456789abcdef",
		"nqn.2014-08.org.nvmexpress:uuid",
		"PREFIX", "FOO_BAR",
		"aaaaaaaa",
		"ssh://git@github.com/org/repo.git",
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa:o:01234567-89ab-cdef-0123-456789abcdef",
		"-----BEGIN " + "PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC0example\n-----END " + "PRIVATE KEY-----\n", // fake key, split for secret scanners
		"-----BEGIN RSA " + "PRIVATE KEY-----\nMIIEpAIBAAKCAQEA0examplekeymaterialhere\n-----END RSA " + "PRIVATE KEY-----\n",
		"-----BEGIN CERTIFICATE REQUEST-----\nMIIBGDCCATACAQAw\n-----END CERTIFICATE REQUEST-----\n",
		"-----BEGIN CERTIFICATE-----\nMIIBGDCCATACAQAw\n-----END CERTIFICATE-----\n",
		"smtp_custom.notification", "slack.notification", "webhook.notification",
		"pagerduty.notification", "event_streams.notification",
		"ibmcejob.notification", "ibmceapp.notification",
		"app_configuration.notification", "smtp_custom.invitation",
		"Example", "Jane Doe", "John Smith", "Example Name",
		"P1Y", "P30D", "PT1H", "2024-01-01T00:00:00Z", "2024-01-01T00:00:00.000Z",
		"-----BEGIN CERTIFICATE-----MIIBGDCCATACAQAwFCAQAwexampleCERTIFICATE-----END CERTIFICATE-----",
	}
	if lit := leadingLiteral(pattern); lit != "" {
		uuid := "01234567-89ab-cdef-0123-456789abcdef"
		candidates = append([]string{lit, lit + uuid, lit + "-" + uuid, lit + "1", lit + "example"}, candidates...)
	}
	for _, alt := range regexpAlternatives(pattern) {
		candidates = append([]string{alt}, candidates...)
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		return "example"
	}
	fits := func(c string) bool {
		if minLen != nil && len(c) < *minLen {
			return false
		}
		if maxLen != nil && len(c) > *maxLen {
			return false
		}
		return re.MatchString(c)
	}
	for _, c := range candidates {
		if fits(c) {
			return c
		}
		if minLen != nil && len(c) < *minLen {
			padded := c + strings.Repeat("x", *minLen-len(c))
			if fits(padded) {
				return padded
			}
		}
	}
	alpha := "abcdefghijklmnopqrstuvwxyz0123456789"
	maxN := 16
	if maxLen != nil && *maxLen < maxN {
		maxN = *maxLen
	}
	start := 1
	if minLen != nil && *minLen > 0 {
		start = *minLen
	}
	if start > 64 {
		start = 64
	}
	if maxN < start {
		maxN = start
	}
	for n := start; n <= maxN; n++ {
		var sb strings.Builder
		for i := 0; i < n; i++ {
			sb.WriteByte(alpha[i%len(alpha)])
		}
		c := sb.String()
		if re.MatchString(c) {
			return c
		}
		digits := strings.Repeat("1", n)
		if re.MatchString(digits) {
			return digits
		}
	}
	if minLen != nil && *minLen > 0 {
		base := strings.Repeat("a", *minLen)
		if re.MatchString(base) {
			return base
		}
		digits := strings.Repeat("1", *minLen)
		if re.MatchString(digits) {
			return digits
		}
	}
	return "example"
}

func leadingLiteral(pattern string) string {
	p := strings.TrimPrefix(pattern, "^")
	var b strings.Builder
	for i := 0; i < len(p); i++ {
		c := p[i]
		switch c {
		case '\\':
			if i+1 < len(p) {
				b.WriteByte(p[i+1])
				i++
			}
		case '[', '(', '.', '*', '+', '?', '{', '|', '$':
			return b.String()
		default:
			b.WriteByte(c)
		}
	}
	return b.String()
}

func regexpAlternatives(pattern string) []string {
	p := strings.TrimPrefix(pattern, "^")
	p = strings.TrimSuffix(p, "$")
	if strings.HasPrefix(p, "(") && strings.HasSuffix(p, ")") {
		p = p[1 : len(p)-1]
	}
	if !strings.Contains(p, "|") || strings.ContainsAny(p, "[]{}*+?\\") {
		return nil
	}
	return strings.Split(p, "|")
}

func padToLen(s string, minLen, maxLen *int) string {
	if minLen != nil && len(s) < *minLen {
		s += strings.Repeat("x", *minLen-len(s))
	}
	if maxLen != nil && *maxLen > 0 && len(s) > *maxLen {
		s = s[:*maxLen]
	}
	return s
}

func renderDoc(res ResourceJSON, spec resourceSpec, id ident, attrRules []attrRule) string {
	var b strings.Builder
	b.WriteString("# " + id.RuleName + "\n\n")
	b.WriteString(mdText(res.Description) + "\n\n")
	b.WriteString("This rule validates `" + res.Name + "` " + id.Noun + " against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.\n\n")
	if spec.APISchema != "" {
		b.WriteString("OpenAPI schema: `" + spec.APISchema + "`")
		if spec.APISpec != "" {
			b.WriteString(" (`" + spec.APISpec + "`)")
		}
		b.WriteString(".\n\n")
	}
	b.WriteString("## Configuration\n\n```hcl\nrule \"" + id.RuleName + "\" {\n  enabled = true\n}\n```\n\n")
	b.WriteString("## Checks\n\n")
	if len(spec.Attrs) == 0 && len(spec.Blocks) == 0 {
		b.WriteString("This " + strings.TrimSuffix(id.Noun, "s") + " exposes no user-configurable arguments beyond provider defaults.\n")
		return b.String()
	}
	b.WriteString("| Argument | Required | Constraints |\n| --- | --- | --- |\n")
	writeDocAttrs(&b, spec.Attrs, "")
	writeDocBlocks(&b, spec.Blocks, "")
	if len(attrRules) > 0 {
		b.WriteString("\n## Attribute rules\n\nEach writable argument has its own rule so checks can be enabled or disabled independently.\n\n")
		b.WriteString("| Rule | Argument |\n| --- | --- |\n")
		for _, ar := range attrRules {
			b.WriteString("| `" + ar.RuleName + "` | `" + ar.Path + "` |\n")
		}
	}
	b.WriteString("\n## Example\n\n### Valid\n\n```hcl\n")
	b.WriteString(docExample(passingHCL(id.Block, res.Name, spec), spec))
	b.WriteString("```\n")
	return b.String()
}

func writeDocAttrs(b *strings.Builder, attrs []specAttr, prefix string) {
	for _, a := range attrs {
		req := "no"
		if a.Required {
			req = "yes"
		}
		b.WriteString("| `" + prefix + a.Name + "` | " + req + " | " + constraintText(a) + " |\n")
	}
}

func writeDocBlocks(b *strings.Builder, blocks []specBlock, prefix string) {
	for _, bl := range blocks {
		req := "no"
		if bl.Required {
			req = "yes"
		}
		b.WriteString("| `" + prefix + bl.Type + "` (block) | " + req + " | nested block |\n")
		writeDocAttrs(b, bl.Attrs, prefix+bl.Type+".")
		writeDocBlocks(b, bl.Blocks, prefix+bl.Type+".")
	}
}

func constraintText(a specAttr) string {
	var parts []string
	if len(a.Enum) > 0 {
		parts = append(parts, "one of: `"+strings.Join(a.Enum, "`, `")+"`")
	}
	if a.Regexp != "" {
		parts = append(parts, "pattern `"+a.Regexp+"`")
	}
	if a.Min != nil && a.Max != nil {
		parts = append(parts, fmt.Sprintf("range %d–%d", *a.Min, *a.Max))
	} else if a.Min != nil {
		parts = append(parts, fmt.Sprintf("min %d", *a.Min))
	} else if a.Max != nil {
		parts = append(parts, fmt.Sprintf("max %d", *a.Max))
	}
	if a.Kind == "cidr" || a.Kind == "ipv4" || a.Kind == "crn" || a.Kind == "zone" || a.Kind == "port" || a.Kind == "json" || a.Kind == "ip_or_cidr" || a.Kind == "uuid" || a.Kind == "email" || a.Kind == "uri" || a.Kind == "hostname" {
		parts = append(parts, a.Kind+" format")
	}
	if a.APIDescription != "" && len(parts) == 0 {
		parts = append(parts, "API: "+mdText(a.APIDescription))
	}
	if len(parts) == 0 {
		return "—"
	}
	return strings.Join(parts, "; ")
}

func writeDocsIndex(dir string, all []ResourceJSON) error {
	byPkg := map[string][]ResourceJSON{}
	for _, r := range all {
		byPkg[r.Package] = append(byPkg[r.Package], r)
	}
	pkgs := sortedKeys(byPkg)
	var b strings.Builder
	b.WriteString("# Rules\n\n")
	b.WriteString("Schema-validation rules generated from the IBM Cloud Terraform provider and IBM Cloud OpenAPI documents. Each type has a schema rule plus one rule per writable argument.\n\n")
	b.WriteString("These rules check schema validity only. They do not assess security posture (for example open security groups, broad IAM policies or missing encryption), so run a policy scanner alongside them.\n\n")
	for _, pkg := range pkgs {
		b.WriteString("## " + pkg + "\n\n")
		b.WriteString("| Rule | Kind | Description |\n| --- | --- | --- |\n")
		for _, r := range byPkg[pkg] {
			id := identOf(r)
			kind := r.Kind
			if kind == "" {
				kind = "resource"
			}
			desc := strings.ReplaceAll(mdText(firstSentence(r.Description)), "|", "\\|")
			b.WriteString("| [" + id.RuleName + "](" + id.FileBase + ".md) | " + kind + " | " + desc + " |\n")
		}
		b.WriteString("\n")
	}
	return os.WriteFile(filepath.Join(dir, "README.md"), []byte(b.String()), 0o644)
}

func firstSentence(desc string) string {
	desc = strings.TrimSpace(desc)
	if i := strings.IndexByte(desc, '.'); i > 0 && i < 200 {
		return desc[:i+1]
	}
	return desc
}

func removeOrphans(rulesDir, docsDir string, generated []string) error {
	keep := map[string]bool{}
	for _, base := range generated {
		keep[base+"_gen.go"] = true
		keep[base+"_gen_test.go"] = true
		keep[base+".md"] = true
	}
	entries, err := os.ReadDir(rulesDir)
	if err != nil {
		return err
	}
	for _, e := range entries {
		name := e.Name()
		if e.IsDir() {
			continue
		}
		if strings.HasSuffix(name, "_gen.go") || strings.HasSuffix(name, "_gen_test.go") {
			if !keep[name] {
				if err := os.Remove(filepath.Join(rulesDir, name)); err != nil {
					return err
				}
				fmt.Printf("removed stale %s\n", name)
			}
		}
	}
	docs, err := os.ReadDir(docsDir)
	if err != nil {
		return err
	}
	for _, e := range docs {
		name := e.Name()
		if e.IsDir() || name == "README.md" || !strings.HasSuffix(name, ".md") {
			continue
		}
		if !keep[name] {
			if err := os.Remove(filepath.Join(docsDir, name)); err != nil {
				return err
			}
			fmt.Printf("removed stale %s\n", name)
		}
	}
	return nil
}
