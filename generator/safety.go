// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package main

import (
	"regexp"
	"sort"
	"strings"

	"github.com/IBM/tflint-ruleset-ibm/internal/redact"
)

// Names from the provider source end up in Go identifiers, comments, file
// paths and HCL fixtures. Anything outside this shape is refused so a
// tampered input cannot inject code or escape the output directories.
var namePattern = regexp.MustCompile(`^[a-z_][a-z0-9_]*$`)

func validName(name string) bool {
	return namePattern.MatchString(name)
}

// checkNames refuses a resource whose type, argument or block names are
// not plain Terraform identifiers.
func checkNames(res ResourceJSON, spec resourceSpec) {
	if !validName(res.Name) {
		fatalf("refusing unsafe resource name %q", res.Name)
	}
	var walk func(attrs []specAttr, blocks []specBlock)
	walk = func(attrs []specAttr, blocks []specBlock) {
		for _, a := range attrs {
			if !validName(a.Name) {
				fatalf("refusing unsafe argument name %q in %s", a.Name, res.Name)
			}
		}
		for _, b := range blocks {
			if !validName(b.Type) {
				fatalf("refusing unsafe block name %q in %s", b.Type, res.Name)
			}
			walk(b.Attrs, b.Blocks)
		}
	}
	walk(spec.Attrs, spec.Blocks)
}

// fixture guards HCL that is written into a Go raw string literal.
func fixture(hcl string) string {
	if strings.Contains(hcl, "`") {
		fatalf("refusing test fixture containing a backtick:\n%s", hcl)
	}
	return hcl
}

// shown mirrors rules.shown so generated expectations match runtime output.
func shown(a specAttr, val string) string {
	return redact.Display(redact.IsSensitive(a.Name, a.Sensitive), val)
}

// mdText neutralizes raw HTML in upstream descriptions written to Markdown.
func mdText(s string) string {
	return strings.NewReplacer("<", "&lt;", ">", "&gt;").Replace(s)
}

// sensitiveNames lists every argument in spec that holds a secret.
func sensitiveNames(attrs []specAttr, blocks []specBlock, out map[string]bool) map[string]bool {
	for _, a := range attrs {
		if !a.ComputedOnly && redact.IsSensitive(a.Name, a.Sensitive) {
			out[a.Name] = true
		}
	}
	for _, b := range blocks {
		sensitiveNames(b.Attrs, b.Blocks, out)
	}
	return out
}

var hclAssign = regexp.MustCompile(`^(\s*)([a-z_][a-z0-9_]*) = .*$`)

// docExample rewrites literal secrets in a documentation example to
// references to sensitive variables, so the docs never model hardcoding.
func docExample(hcl string, spec resourceSpec) string {
	secrets := sensitiveNames(spec.Attrs, spec.Blocks, map[string]bool{})
	if len(secrets) == 0 {
		return hcl
	}
	lines := strings.Split(hcl, "\n")
	used := map[string]bool{}
	for i, line := range lines {
		m := hclAssign.FindStringSubmatch(line)
		if m == nil || !secrets[m[2]] {
			continue
		}
		lines[i] = m[1] + m[2] + " = var." + m[2]
		used[m[2]] = true
	}
	if len(used) == 0 {
		return hcl
	}
	names := make([]string, 0, len(used))
	for n := range used {
		names = append(names, n)
	}
	sort.Strings(names)
	var b strings.Builder
	for _, n := range names {
		b.WriteString("variable \"" + n + "\" {\n  sensitive = true\n}\n\n")
	}
	b.WriteString(strings.Join(lines, "\n"))
	return b.String()
}
