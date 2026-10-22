// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

type extractPkg struct {
	Package   string `json:"package"`
	Resources []struct {
		Name       string                 `json:"name"`
		Kind       string                 `json:"kind"`
		Package    string                 `json:"package"`
		Function   string                 `json:"function"`
		Attributes map[string]interface{} `json:"attributes"`
	} `json:"resources"`
	Skipped []struct {
		Name   string `json:"name"`
		Reason string `json:"reason"`
	} `json:"skipped"`
}

type row struct {
	Package     string `json:"package"`
	Name        string `json:"name"`
	Kind        string `json:"kind"`
	Extracted   bool   `json:"extracted"`
	Rule        bool   `json:"rule"`
	Test        bool   `json:"test"`
	PassFixture bool   `json:"pass_fixture"`
	FailFixture bool   `json:"fail_fixture"`
	FailNA      bool   `json:"fail_na"`
	Registered  bool   `json:"registered"`
	Docs        bool   `json:"docs"`
	DocsIndex   bool   `json:"docs_index"`
	CheckFn     bool   `json:"check_fn"`
	Status      string `json:"status"`
	Missing     string `json:"missing,omitempty"`
	AttrCount   int    `json:"attr_count"`
	Constructor string `json:"constructor"`
}

func main() {
	provider := flag.String("provider", "", "path to IBM Cloud Terraform provider source root")
	root := flag.String("root", ".", "tflint-ruleset-ibm repository root")
	flag.Parse()
	if *provider == "" {
		fmt.Fprintln(os.Stderr, "coverage: -provider is required")
		os.Exit(2)
	}
	absRoot, _ := filepath.Abs(*root)
	absProv, _ := filepath.Abs(*provider)

	types, err := parseProviderMaps(filepath.Join(absProv, "ibm", "provider", "provider.go"))
	if err != nil {
		fatalf("parse provider maps: %v", err)
	}
	if len(types) == 0 {
		fatalf("no resources or data sources found in provider maps")
	}

	extracted := loadExtracted(filepath.Join(absRoot, "extractor", "output"))
	regSrc := readFile(filepath.Join(absRoot, "rules", "provider.go"))
	docsIndex := readFile(filepath.Join(absRoot, "docs", "rules", "README.md"))
	readme := readFile(filepath.Join(absRoot, "README.md"))

	var rows []row
	var gaps []row
	nRes, nData := 0, 0
	emptySchema := 0
	for _, t := range types {
		if t.Kind == "data" {
			nData++
		} else {
			nRes++
		}
		r := auditOne(absRoot, t, extracted, regSrc, docsIndex)
		rows = append(rows, r)
		if r.Status == "gap" {
			gaps = append(gaps, r)
		}
		if r.FailNA {
			emptySchema++
		}
	}

	orphans := findOrphans(absRoot, types)
	dups := findDuplicates(types, regSrc)

	resFiles, dataFiles, _ := listSourceFiles(filepath.Join(absProv, "ibm", "service"))
	pkgs, _ := listServicePackages(filepath.Join(absProv, "ibm", "service"))

	nReg := 0
	for _, line := range strings.Split(regSrc, "\n") {
		s := strings.TrimSpace(line)
		if strings.HasPrefix(s, "New") && strings.Contains(s, "Rule()") {
			nReg++
		}
	}
	nAttr := nReg - len(types)
	if nAttr < 0 {
		nAttr = 0
	}
	readmeOK := readmeHasCounts(readme, nRes, nData, len(types), 0)

	writeCoverageMD(filepath.Join(absRoot, "COVERAGE.md"), rows, nRes, nData, gaps, orphans, resFiles, dataFiles, pkgs, emptySchema, nAttr)
	patchReadme(filepath.Join(absRoot, "README.md"), readme, nRes, nData, len(types), len(pkgs), nAttr)
	writeReportJSON(filepath.Join(absRoot, "coverage", "report.json"), rows, nRes, nData, gaps, orphans, resFiles, dataFiles)

	fmt.Printf("coverage: %d resources, %d data sources, %d rules, %d gaps, %d orphans, %d empty-schema\n",
		nRes, nData, len(types), len(gaps), len(orphans), emptySchema)

	failed := false
	if len(gaps) > 0 {
		fmt.Fprintf(os.Stderr, "coverage: %d types missing artifacts\n", len(gaps))
		for _, g := range gaps {
			fmt.Fprintf(os.Stderr, "  %s %s (%s): %s\n", g.Kind, g.Name, g.Package, g.Missing)
		}
		failed = true
	}
	if len(orphans) > 0 {
		fmt.Fprintf(os.Stderr, "coverage: %d stale generated files\n", len(orphans))
		for _, o := range orphans {
			fmt.Fprintf(os.Stderr, "  stale %s\n", o)
		}
		failed = true
	}
	if len(dups) > 0 {
		fmt.Fprintf(os.Stderr, "coverage: duplicates: %s\n", strings.Join(dups, ", "))
		failed = true
	}
	if !readmeOK {
		// patchReadme rewrites; re-read
		readme = readFile(filepath.Join(absRoot, "README.md"))
		if !readmeHasCounts(readme, nRes, nData, len(types), 0) {
			fmt.Fprintln(os.Stderr, "coverage: README coverage table does not match provider totals")
			failed = true
		}
	}
	if failed {
		os.Exit(1)
	}
	fmt.Println("coverage: ok")
}

func auditOne(root string, t providerType, extracted map[string]extractHit, regSrc, docsIndex string) row {
	key := t.Kind + "\x00" + t.Name
	base := fileBase(t.Name, t.Kind)
	rn := ruleName(t.Name, t.Kind)
	tn := typeName(t.Name, t.Kind)
	rulePath := filepath.Join(root, "rules", base+"_gen.go")
	testPath := filepath.Join(root, "rules", base+"_gen_test.go")
	docPath := filepath.Join(root, "docs", "rules", base+".md")
	ruleSrc := readFile(rulePath)
	testSrc := readFile(testPath)

	r := row{
		Package:     t.Package,
		Name:        t.Name,
		Kind:        t.Kind,
		Constructor: t.Function,
	}
	hit, ok := extracted[key]
	r.Extracted = ok
	if ok {
		r.AttrCount = hit.attrs
		if hit.pkg != "" && hit.pkg != t.Package {
			r.Missing = joinMiss(r.Missing, "package mismatch extract="+hit.pkg+" map="+t.Package)
		}
	}
	r.Rule = fileExists(rulePath)
	r.Test = fileExists(testPath)
	r.Docs = fileExists(docPath)
	r.DocsIndex = strings.Contains(docsIndex, rn)
	r.Registered = strings.Contains(regSrc, "New"+tn+"Rule()")
	r.PassFixture = strings.Contains(testSrc, `Name: "valid configuration"`)
	r.FailFixture = strings.Contains(testSrc, `Name: "invalid configuration"`)
	if r.Extracted && !hit.checkable {
		r.FailNA = true
	}
	checkFn := "RunResourceChecks"
	if t.Kind == "data" {
		checkFn = "RunDataSourceChecks"
	}
	r.CheckFn = strings.Contains(ruleSrc, checkFn) && strings.Contains(ruleSrc, "func (r *"+tn+"Rule) Check(")
	if strings.Contains(strings.ToLower(ruleSrc), "todo") || strings.Contains(ruleSrc, "panic(\"unimplemented\")") {
		r.CheckFn = false
		r.Missing = joinMiss(r.Missing, "placeholder implementation")
	}

	var miss []string
	if !r.Extracted {
		miss = append(miss, "extracted schema")
	}
	if !r.Rule {
		miss = append(miss, "rule file")
	}
	if !r.Test {
		miss = append(miss, "test file")
	}
	if !r.PassFixture {
		miss = append(miss, "pass fixture")
	}
	if !r.FailFixture && !r.FailNA {
		miss = append(miss, "fail fixture")
	}
	if !r.Registered {
		miss = append(miss, "registration")
	}
	if !r.Docs {
		miss = append(miss, "docs")
	}
	if !r.DocsIndex {
		miss = append(miss, "docs index")
	}
	if r.Rule && !r.CheckFn {
		miss = append(miss, "check body")
	}
	if len(miss) > 0 {
		r.Status = "gap"
		r.Missing = joinMiss(r.Missing, strings.Join(miss, ", "))
	} else if r.FailNA {
		r.Status = "covered-empty-schema"
	} else {
		r.Status = "covered"
	}
	return r
}

type extractHit struct {
	pkg       string
	attrs     int
	checkable bool
}

func loadExtracted(dir string) map[string]extractHit {
	out := map[string]extractHit{}
	entries, err := os.ReadDir(dir)
	if err != nil {
		return out
	}
	for _, e := range entries {
		if e.IsDir() || e.Name() == "index.json" || e.Name() == "overlay-index.json" || strings.HasSuffix(e.Name(), ".overlay.json") || !strings.HasSuffix(e.Name(), ".json") {
			continue
		}
		raw, err := os.ReadFile(filepath.Join(dir, e.Name()))
		if err != nil {
			continue
		}
		var pkg extractPkg
		if json.Unmarshal(raw, &pkg) != nil {
			continue
		}
		for _, r := range pkg.Resources {
			kind := r.Kind
			if kind == "" {
				kind = "resource"
			}
			if r.Name == "" || r.Package == "" {
				continue
			}
			attrs := r.Attributes
			if attrs == nil {
				attrs = map[string]interface{}{}
			}
			out[kind+"\x00"+r.Name] = extractHit{pkg: r.Package, attrs: len(attrs), checkable: schemaCheckable(attrs)}
		}
	}
	return out
}

func schemaCheckable(attrs map[string]interface{}) bool {
	for _, raw := range attrs {
		a, ok := raw.(map[string]interface{})
		if !ok {
			continue
		}
		required, _ := a["required"].(bool)
		optional, _ := a["optional"].(bool)
		computed, _ := a["computed"].(bool)
		if required && !computed {
			return true
		}
		if computed && !optional && !required {
			return true
		}
		if v, ok := a["validate"].(map[string]interface{}); ok {
			if _, ok := v["enum"]; ok {
				return true
			}
			if _, ok := v["int_enum"]; ok {
				return true
			}
			if v["regexp"] != nil && v["regexp"] != "" {
				return true
			}
			if v["min"] != nil || v["max"] != nil || v["min_len"] != nil || v["max_len"] != nil {
				return true
			}
			if v["no_zero"] == true {
				return true
			}
			if k, _ := v["kind"].(string); k != "" && k != "invoke" && k != "func" && k != "cloud_data" && k != "no_zero" {
				return true
			}
		}
	}
	return false
}

func findOrphans(root string, types []providerType) []string {
	keep := map[string]bool{}
	for _, t := range types {
		base := fileBase(t.Name, t.Kind)
		keep[base+"_gen.go"] = true
		keep[base+"_gen_test.go"] = true
		keep[base+".md"] = true
	}
	var orphans []string
	rulesDir := filepath.Join(root, "rules")
	entries, _ := os.ReadDir(rulesDir)
	for _, e := range entries {
		name := e.Name()
		if strings.HasSuffix(name, "_gen.go") || strings.HasSuffix(name, "_gen_test.go") {
			if !keep[name] {
				orphans = append(orphans, filepath.Join("rules", name))
			}
		}
	}
	docsDir := filepath.Join(root, "docs", "rules")
	docs, _ := os.ReadDir(docsDir)
	for _, e := range docs {
		name := e.Name()
		if name == "README.md" || !strings.HasSuffix(name, ".md") {
			continue
		}
		if !keep[name] {
			orphans = append(orphans, filepath.Join("docs", "rules", name))
		}
	}
	sort.Strings(orphans)
	return orphans
}

func findDuplicates(types []providerType, regSrc string) []string {
	seen := map[string]int{}
	var dups []string
	for _, t := range types {
		k := t.Kind + ":" + t.Name
		seen[k]++
		if seen[k] == 2 {
			dups = append(dups, k)
		}
	}
	regSeen := map[string]int{}
	for _, m := range regexp.MustCompile(`New(\w+)Rule\(\)`).FindAllStringSubmatch(regSrc, -1) {
		regSeen[m[1]]++
		if regSeen[m[1]] == 2 {
			dups = append(dups, "registration:"+m[1])
		}
	}
	return dups
}

func writeCoverageMD(path string, rows []row, nRes, nData int, gaps []row, orphans, resFiles, dataFiles, pkgs []string, emptySchema, nAttr int) {
	byPkg := map[string][]row{}
	for _, r := range rows {
		byPkg[r.Package] = append(byPkg[r.Package], r)
	}
	pkgNames := make([]string, 0, len(byPkg))
	for p := range byPkg {
		pkgNames = append(pkgNames, p)
	}
	sort.Strings(pkgNames)

	var b strings.Builder
	b.WriteString("# Coverage\n\n")
	b.WriteString("Ground truth is the IBM Cloud Terraform provider `ResourcesMap` and `DataSourcesMap`.\n\n")
	fmt.Fprintf(&b, "| | Count |\n| --- | ---: |\n")
	fmt.Fprintf(&b, "| Service packages | %d |\n", len(pkgs))
	fmt.Fprintf(&b, "| Provider resources discovered | %d |\n", nRes)
	fmt.Fprintf(&b, "| Resources covered | %d |\n", nRes-countKindGaps(gaps, "resource"))
	fmt.Fprintf(&b, "| Resources skipped | %d |\n", countKindGaps(gaps, "resource"))
	fmt.Fprintf(&b, "| Provider data sources discovered | %d |\n", nData)
	fmt.Fprintf(&b, "| Data sources covered | %d |\n", nData-countKindGaps(gaps, "data"))
	fmt.Fprintf(&b, "| Data sources skipped | %d |\n", countKindGaps(gaps, "data"))
	fmt.Fprintf(&b, "| Generated rules | %d |\n", len(rows))
	fmt.Fprintf(&b, "| Attribute rules | %d |\n", nAttr)
	fmt.Fprintf(&b, "| Registered checks | %d |\n", len(rows)+nAttr)
	fmt.Fprintf(&b, "| Empty-schema stubs (fail fixture n/a) | %d |\n", emptySchema)
	fmt.Fprintf(&b, "| Coverage gaps | %d |\n", len(gaps))
	fmt.Fprintf(&b, "| Stale generated files | %d |\n\n", len(orphans))

	if len(gaps) > 0 {
		b.WriteString("## Gaps\n\n| Package | Name | Type | Reason | What would be required |\n| --- | --- | --- | --- | --- |\n")
		for _, g := range gaps {
			fmt.Fprintf(&b, "| `%s` | `%s` | %s | %s | generate the missing artifacts via `make extract generate` |\n",
				g.Package, g.Name, g.Kind, strings.ReplaceAll(g.Missing, "|", "\\|"))
		}
		b.WriteString("\n")
	} else {
		b.WriteString("No skipped constructors. Every provider map entry has extracted schema, rule, test, pass fixture, fail fixture (or empty-schema waiver), registration, and docs.\n\n")
	}

	yn := func(v bool) string {
		if v {
			return "yes"
		}
		return "no"
	}
	failCol := func(r row) string {
		if r.FailNA {
			return "n/a"
		}
		return yn(r.FailFixture)
	}

	for _, pkg := range pkgNames {
		fmt.Fprintf(&b, "## %s\n\n", pkg)
		b.WriteString("| Package | Resource/Data Source | Type | Extracted | Rule | Test | Pass Fixture | Fail Fixture | Registered | Docs | Status |\n")
		b.WriteString("| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |\n")
		items := byPkg[pkg]
		sort.Slice(items, func(i, j int) bool {
			if items[i].Kind != items[j].Kind {
				return items[i].Kind == "resource"
			}
			return items[i].Name < items[j].Name
		})
		for _, r := range items {
			fmt.Fprintf(&b, "| `%s` | `%s` | %s | %s | %s | %s | %s | %s | %s | %s | %s |\n",
				r.Package, r.Name, r.Kind, yn(r.Extracted), yn(r.Rule), yn(r.Test), yn(r.PassFixture), failCol(r), yn(r.Registered), yn(r.Docs), r.Status)
		}
		b.WriteString("\n")
	}

	b.WriteString("## Unregistered source files\n\n")
	b.WriteString("These `resource_ibm_*` / `data_source_ibm_*` files exist under `ibm/service` but are not themselves provider map keys. They are helpers, gen2 variants, or share a constructor registered under another name. They are listed so they cannot be silently ignored.\n\n")
	fmt.Fprintf(&b, "- resource files: %d\n- data source files: %d\n", len(resFiles), len(dataFiles))

	writeOrDie(path, []byte(b.String()))
}

func countKindGaps(gaps []row, kind string) int {
	n := 0
	for _, g := range gaps {
		if g.Kind == kind {
			n++
		}
	}
	return n
}

func patchReadme(path, current string, nRes, nData, nRules, nPkgs, nAttr int) {
	table := fmt.Sprintf("<!-- coverage-summary:start -->\n## Coverage\n\n| | Count |\n| --- | ---: |\n| Service packages | %d |\n| Managed resources | %d |\n| Data sources | %d |\n| Generated rules | %d |\n| Attribute rules | %d |\n| Registered checks | %d |\n| Skipped constructors | 0 |\n\nSee [COVERAGE.md](COVERAGE.md) for the per-type matrix. Run `make coverage` to re-audit against the provider maps.\n<!-- coverage-summary:end -->", nPkgs, nRes, nData, nRules, nAttr, nRules+nAttr)
	re := regexp.MustCompile(`(?s)<!-- coverage-summary:start -->.*?<!-- coverage-summary:end -->`)
	var next string
	if re.MatchString(current) {
		next = re.ReplaceAllString(current, table)
	} else {
		re2 := regexp.MustCompile(`(?s)## Coverage\n\n\| \| Count \|\n\| --- \| ---: \|\n.*?\| Skipped constructors \| \d+ \|\n`)
		if re2.MatchString(current) {
			next = re2.ReplaceAllString(current, table+"\n\n")
		} else {
			next = current
		}
	}
	if next != "" && next != current {
		writeOrDie(path, []byte(next))
	}
}

func readmeHasCounts(readme string, nRes, nData, nRules, skipped int) bool {
	return strings.Contains(readme, fmt.Sprintf("| Managed resources | %d |", nRes)) &&
		strings.Contains(readme, fmt.Sprintf("| Data sources | %d |", nData)) &&
		strings.Contains(readme, fmt.Sprintf("| Generated rules | %d |", nRules)) &&
		strings.Contains(readme, fmt.Sprintf("| Skipped constructors | %d |", skipped))
}

func writeReportJSON(path string, rows []row, nRes, nData int, gaps []row, orphans, resFiles, dataFiles []string) {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		fatalf("mkdir %s: %v", filepath.Dir(path), err)
	}
	payload := map[string]any{
		"totals": map[string]int{
			"resources":      nRes,
			"data_sources":   nData,
			"rules":          len(rows),
			"gaps":           len(gaps),
			"orphans":        len(orphans),
			"resource_files": len(resFiles),
			"data_files":     len(dataFiles),
		},
		"gaps":    gaps,
		"orphans": orphans,
		"items":   rows,
	}
	raw, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		fatalf("encode %s: %v", path, err)
	}
	writeOrDie(path, append(raw, '\n'))
}

func readFile(path string) string {
	b, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	return string(b)
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func joinMiss(a, b string) string {
	if a == "" {
		return b
	}
	if b == "" {
		return a
	}
	return a + "; " + b
}

func writeOrDie(path string, data []byte) {
	if err := os.WriteFile(path, data, 0o644); err != nil {
		fatalf("write %s: %v", path, err)
	}
}

func fatalf(format string, args ...any) {
	fmt.Fprintf(os.Stderr, "coverage: "+format+"\n", args...)
	os.Exit(1)
}
