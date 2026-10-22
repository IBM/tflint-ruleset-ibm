// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type attrRule struct {
	TypeName  string
	RuleName  string
	Path      string
	PathParts []string
	Attr      specAttr
}

func collectAttrRules(id ident, spec resourceSpec) []attrRule {
	var out []attrRule
	seen := map[string]bool{}
	var walk func(attrs []specAttr, blocks []specBlock, parts []string)
	walk = func(attrs []specAttr, blocks []specBlock, parts []string) {
		for _, a := range attrs {
			if a.ComputedOnly {
				continue
			}
			p := append(append([]string{}, parts...), a.Name)
			leaf := a
			leaf.Required = false
			out = append(out, makeAttrRule(id, leaf, p, seen))
		}
		for _, b := range blocks {
			walk(b.Attrs, b.Blocks, append(append([]string{}, parts...), b.Type))
		}
	}
	walk(spec.Attrs, spec.Blocks, nil)
	return out
}

func makeAttrRule(id ident, a specAttr, parts []string, seen map[string]bool) attrRule {
	ruleSuf, typeSuf := pathSuffix(parts)
	typeName := id.TypeName + "Invalid" + typeSuf
	ruleName := id.RuleName + "_invalid_" + ruleSuf
	if seen[typeName] {
		n := 2
		for seen[typeName+strconv.Itoa(n)] {
			n++
		}
		typeName = typeName + strconv.Itoa(n)
		ruleName = ruleName + "_" + strconv.Itoa(n)
	}
	seen[typeName] = true
	return attrRule{
		TypeName:  typeName,
		RuleName:  ruleName,
		Path:      strings.Join(parts, "."),
		PathParts: parts,
		Attr:      a,
	}
}

func pathSuffix(parts []string) (ruleSuffix, typeSuffix string) {
	ruleSuffix = strings.Join(parts, "__")
	pascals := make([]string, len(parts))
	for i, p := range parts {
		pascals[i] = toPascal(p)
	}
	typeSuffix = strings.Join(pascals, "In")
	return
}

func renderAttrRules(id ident, attrRules []attrRule) string {
	if len(attrRules) == 0 {
		return ""
	}
	var b strings.Builder
	for _, ar := range attrRules {
		b.WriteString("\n// New" + ar.TypeName + "Rule validates the " + strconv.Quote(ar.Path) + " argument of " + id.Name + " " + id.Noun + ".\n")
		b.WriteString("func New" + ar.TypeName + "Rule() tflint.Rule {\n")
		b.WriteString("\treturn NewArgumentRule(" + strconv.Quote(ar.RuleName) + ", project.ReferenceLink(" + strconv.Quote(id.FileBase) + "), ResourceCheck{\n")
		b.WriteString("\t\tResource: " + strconv.Quote(id.Name) + ",\n")
		if len(ar.PathParts) <= 1 {
			b.WriteString("\t\tAttrs: []AttrCheck{{ " + attrFields(ar.Attr) + " }},\n")
		} else {
			blocks := nestBlocks(ar.PathParts[:len(ar.PathParts)-1], ar.Attr)
			b.WriteString("\t\tBlocks: []BlockCheck{")
			for i, bl := range blocks {
				if i > 0 {
					b.WriteString(", ")
				}
				b.WriteString(blockLit(bl))
			}
			b.WriteString("},\n")
		}
		b.WriteString("\t}, " + id.CheckFn + ")\n}\n")
	}
	return b.String()
}

func nestBlocks(blockPath []string, leaf specAttr) []specBlock {
	if len(blockPath) == 0 {
		return nil
	}
	inner := specBlock{Type: blockPath[len(blockPath)-1], Attrs: []specAttr{leaf}}
	for i := len(blockPath) - 2; i >= 0; i-- {
		inner = specBlock{Type: blockPath[i], Blocks: []specBlock{inner}}
	}
	return []specBlock{inner}
}

func renderAttrTests(id ident, spec resourceSpec, attrRules []attrRule) string {
	if len(attrRules) == 0 {
		return ""
	}
	var b strings.Builder
	b.WriteString("\nfunc Test_" + id.TypeName + "AttrRules(t *testing.T) {\n")
	b.WriteString("\tvalid := `\n" + fixture(passingHCL(id.Block, id.Name, spec)) + "`\n")
	b.WriteString("\tfor _, rule := range []tflint.Rule{\n")
	for _, ar := range attrRules {
		b.WriteString("\t\tNew" + ar.TypeName + "Rule(),\n")
	}
	b.WriteString("\t} {\n")
	b.WriteString("\t\tt.Run(rule.Name()+\"_valid\", func(t *testing.T) {\n")
	b.WriteString("\t\t\trunner := helper.TestRunner(t, map[string]string{\"resource.tf\": valid})\n")
	b.WriteString("\t\t\tif err := rule.Check(runner); err != nil {\n")
	b.WriteString("\t\t\t\tt.Fatalf(\"Unexpected error occurred: %s\", err)\n")
	b.WriteString("\t\t\t}\n")
	b.WriteString("\t\t\thelper.AssertIssuesWithoutRange(t, helper.Issues{}, runner.Issues)\n")
	b.WriteString("\t\t})\n")
	b.WriteString("\t}\n")
	for _, ar := range attrRules {
		hcl, msg, ok := failingAttrHCL(id.Block, id.Name, spec, ar)
		if !ok {
			continue
		}
		b.WriteString("\tt.Run(" + strconv.Quote("invalid "+ar.Path) + ", func(t *testing.T) {\n")
		b.WriteString("\t\trule := New" + ar.TypeName + "Rule()\n")
		b.WriteString("\t\trunner := helper.TestRunner(t, map[string]string{\"resource.tf\": `\n" + fixture(hcl) + "`})\n")
		b.WriteString("\t\tif err := rule.Check(runner); err != nil {\n")
		b.WriteString("\t\t\tt.Fatalf(\"Unexpected error occurred: %s\", err)\n")
		b.WriteString("\t\t}\n")
		b.WriteString("\t\thelper.AssertIssuesWithoutRange(t, helper.Issues{{\n")
		b.WriteString("\t\t\tRule:    rule,\n")
		b.WriteString("\t\t\tMessage: " + strconv.Quote(msg) + ",\n")
		b.WriteString("\t\t}}, runner.Issues)\n")
		b.WriteString("\t})\n")
	}
	b.WriteString("}\n")
	return b.String()
}

func failingAttrHCL(block, resource string, spec resourceSpec, ar attrRule) (string, string, bool) {
	val, msg, ok := failingAttrValue(ar.Attr)
	if !ok {
		return "", "", false
	}
	hcl := hclHeader(block, resource) + bodyHCLForcePath(spec.Attrs, spec.Blocks, "  ", ar.PathParts, val) + "}\n"
	return hcl, msg, true
}

func failingAttrValue(a specAttr) (string, string, bool) {
	quoteList := func(v string) string {
		return listWrap(a, v)
	}
	if len(a.IntEnum) > 0 || a.Kind == "int_enum" || a.Kind == "int_range" || a.Kind == "int_min" || a.Kind == "int_max" {
		if len(a.IntEnum) > 0 {
			msg := fmt.Sprintf("%d is not a valid value for %q", 999999, a.Name)
			return quoteList("999999"), msg, true
		}
		if a.Min != nil {
			bad := *a.Min - 1
			msg := fmt.Sprintf("%q must be at least %d, got %d", a.Name, *a.Min, bad)
			return quoteList(strconv.Itoa(bad)), msg, true
		}
		if a.Max != nil {
			bad := *a.Max + 1
			msg := fmt.Sprintf("%q must be at most %d, got %d", a.Name, *a.Max, bad)
			return quoteList(strconv.Itoa(bad)), msg, true
		}
	}
	if len(a.Enum) > 0 {
		val := "not-a-valid-enum-value"
		msg := fmt.Sprintf("%s is not a valid value for %q. Valid options: %s", shown(a, val), a.Name, strings.Join(a.Enum, ", "))
		return quoteList(strconv.Quote(val)), msg, true
	}
	if len(a.IntEnum) > 0 {
		msg := fmt.Sprintf("%d is not a valid value for %q", 999999, a.Name)
		return quoteList("999999"), msg, true
	}
	switch a.Kind {
	case "cidr":
		return quoteList(strconv.Quote("not-a-cidr")), fmt.Sprintf("%s is not a valid CIDR block", shown(a, "not-a-cidr")), true
	case "ipv4":
		return quoteList(strconv.Quote("not-an-ip")), fmt.Sprintf("%s is not a valid IPv4 address", shown(a, "not-an-ip")), true
	case "ip_or_cidr":
		bad := padToLen("nope", a.MinLen, a.MaxLen)
		return quoteList(strconv.Quote(bad)), fmt.Sprintf("%s is not a valid IP address or CIDR block", shown(a, bad)), true
	case "crn":
		bad := padToLen("not-a-crn", a.MinLen, a.MaxLen)
		return quoteList(strconv.Quote(bad)), fmt.Sprintf("%s is not a valid CRN", shown(a, bad)), true
	case "zone":
		return quoteList(strconv.Quote("not-a-zone")), fmt.Sprintf("%s is not a valid availability zone name", shown(a, "not-a-zone")), true
	case "port":
		return quoteList("0"), fmt.Sprintf("%q must be a valid port in the range 1-65535, got %d", a.Name, 0), true
	case "json":
		return quoteList(strconv.Quote("not-json")), fmt.Sprintf("%q is not valid JSON", a.Name), true
	case "uuid":
		bad := padToLen("not-a-uuid", a.MinLen, a.MaxLen)
		return quoteList(strconv.Quote(bad)), fmt.Sprintf("%s is not a valid UUID", shown(a, bad)), true
	case "email":
		return quoteList(strconv.Quote("not-an-email")), fmt.Sprintf("%s is not a valid email address", shown(a, "not-an-email")), true
	case "uri":
		bad := padToLen("not a uri", a.MinLen, a.MaxLen)
		return quoteList(strconv.Quote(bad)), fmt.Sprintf("%s is not a valid URI", shown(a, bad)), true
	case "hostname":
		return quoteList(strconv.Quote("-bad-host")), fmt.Sprintf("%s is not a valid hostname", shown(a, "-bad-host")), true
	}
	if a.Regexp != "" {
		dummy := regexpFailDummy(a)
		if dummy != "" {
			msg := fmt.Sprintf("%s does not match expected format for %q", shown(a, dummy), a.Name)
			return quoteList(strconv.Quote(dummy)), msg, true
		}
	}
	if a.Min != nil {
		bad := *a.Min - 1
		msg := fmt.Sprintf("%q must be at least %d, got %d", a.Name, *a.Min, bad)
		return quoteList(strconv.Itoa(bad)), msg, true
	}
	if a.Max != nil {
		bad := *a.Max + 1
		msg := fmt.Sprintf("%q must be at most %d, got %d", a.Name, *a.Max, bad)
		return quoteList(strconv.Itoa(bad)), msg, true
	}
	if a.NoZero {
		msg := fmt.Sprintf("%q must not be empty", a.Name)
		return quoteList(strconv.Quote("")), msg, true
	}
	if a.MinLen != nil && *a.MinLen > 0 {
		msg := fmt.Sprintf("%q must be at least %d characters", a.Name, *a.MinLen)
		return quoteList(strconv.Quote("")), msg, true
	}
	return "", "", false
}

func listWrap(a specAttr, v string) string {
	if a.List && !strings.HasPrefix(strings.TrimSpace(v), "[") {
		return "[" + v + "]"
	}
	return v
}

func bodyHCLForcePath(attrs []specAttr, blocks []specBlock, indent string, path []string, value string) string {
	var b strings.Builder
	if len(path) == 0 {
		return b.String()
	}
	if len(path) == 1 {
		for _, a := range attrs {
			if a.Name == path[0] {
				b.WriteString(indent + a.Name + " = " + value + "\n")
				continue
			}
			if includeInFixture(a) {
				b.WriteString(indent + a.Name + " = " + dummyValue(a) + "\n")
			}
		}
		for _, bl := range blocks {
			if includeBlock(bl) {
				b.WriteString(indent + bl.Type + " {\n")
				b.WriteString(bodyHCL(bl.Attrs, bl.Blocks, indent+"  ", false))
				b.WriteString(indent + "}\n")
			}
		}
		return b.String()
	}
	for _, a := range attrs {
		if includeInFixture(a) {
			b.WriteString(indent + a.Name + " = " + dummyValue(a) + "\n")
		}
	}
	wrote := false
	for _, bl := range blocks {
		if bl.Type == path[0] {
			b.WriteString(indent + bl.Type + " {\n")
			b.WriteString(bodyHCLForcePath(bl.Attrs, bl.Blocks, indent+"  ", path[1:], value))
			b.WriteString(indent + "}\n")
			wrote = true
			continue
		}
		if includeBlock(bl) {
			b.WriteString(indent + bl.Type + " {\n")
			b.WriteString(bodyHCL(bl.Attrs, bl.Blocks, indent+"  ", false))
			b.WriteString(indent + "}\n")
		}
	}
	if !wrote {
		b.WriteString(indent + path[0] + " {\n")
		b.WriteString(bodyHCLForcePath(nil, nil, indent+"  ", path[1:], value))
		b.WriteString(indent + "}\n")
	}
	return b.String()
}

func regexpFailDummy(a specAttr) string {
	re, err := regexp.Compile(a.Regexp)
	if err != nil {
		return ""
	}
	candidates := []string{"!!!invalid!!!", "###", "not-valid", "!!!!!", "xx", "x", " "}
	for _, dummy := range candidates {
		if a.MinLen != nil && len(dummy) < *a.MinLen {
			dummy = dummy + strings.Repeat("!", *a.MinLen-len(dummy))
		}
		if a.MaxLen != nil && *a.MaxLen > 0 && len(dummy) > *a.MaxLen {
			dummy = dummy[:*a.MaxLen]
		}
		if dummy != "" && !re.MatchString(dummy) {
			if a.MinLen != nil && len(dummy) < *a.MinLen {
				continue
			}
			if a.MaxLen != nil && len(dummy) > *a.MaxLen {
				continue
			}
			return dummy
		}
	}
	return ""
}
