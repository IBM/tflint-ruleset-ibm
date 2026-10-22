// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func main() {
	provider := flag.String("provider", "", "path to local IBM Cloud Terraform provider source root")
	outDir := flag.String("out", "extractor/output", "directory for per-package JSON artifacts")
	packages := flag.String("packages", "", "comma-separated service packages (empty = all)")
	flag.Parse()

	if *provider == "" {
		fmt.Fprintln(os.Stderr, "extractor: -provider is required (local disposable clone of the IBM Cloud Terraform provider)")
		os.Exit(2)
	}
	root := mustAbs(*provider)
	serviceRoot := filepath.Join(root, "ibm", "service")
	providerFile := filepath.Join(root, "ibm", "provider", "provider.go")
	flexDir := filepath.Join(root, "ibm", "flex")

	if !fileExists(providerFile) {
		fmt.Fprintf(os.Stderr, "extractor: provider.go not found at %s\n", providerFile)
		os.Exit(1)
	}
	if !fileExists(serviceRoot) {
		fmt.Fprintf(os.Stderr, "extractor: service packages not found at %s\n", serviceRoot)
		os.Exit(1)
	}

	refs, err := parseProviderResources(providerFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "extractor: parse provider map: %v\n", err)
		os.Exit(1)
	}
	filtered := make([]ResourceRef, 0, len(refs))
	for _, r := range refs {
		if strings.HasSuffix(r.Function, "Validator") {
			continue
		}
		if r.Kind != "resource" && r.Kind != "data" {
			continue
		}
		filtered = append(filtered, r)
	}
	refs = filtered

	want := map[string]bool{}
	if strings.TrimSpace(*packages) != "" {
		for _, p := range strings.Split(*packages, ",") {
			p = strings.TrimSpace(p)
			if p != "" && p != "all" {
				want[p] = true
			}
		}
	}

	byPkg := map[string][]ResourceRef{}
	for _, r := range refs {
		if len(want) > 0 && !want[r.Package] {
			continue
		}
		byPkg[r.Package] = append(byPkg[r.Package], r)
	}

	flexConsts := parseFlexConsts(flexDir)

	others := map[string]*pkgParser{}
	if rcDir := filepath.Join(serviceRoot, "resourcecontroller"); fileExists(rcDir) {
		if rc, err := parsePackageDir(rcDir); err == nil {
			others["resourcecontroller"] = rc
		}
	}

	if err := os.MkdirAll(*outDir, 0o755); err != nil {
		fmt.Fprintf(os.Stderr, "extractor: mkdir: %v\n", err)
		os.Exit(1)
	}

	pkgNames := make([]string, 0, len(byPkg))
	if len(want) > 0 {
		for pkg := range byPkg {
			pkgNames = append(pkgNames, pkg)
		}
		for pkg := range want {
			if _, ok := byPkg[pkg]; !ok {
				pkgNames = append(pkgNames, pkg)
			}
		}
	} else {
		listed, err := listServicePackages(serviceRoot)
		if err != nil {
			fmt.Fprintf(os.Stderr, "extractor: list packages: %v\n", err)
			os.Exit(1)
		}
		seen := map[string]bool{}
		for _, pkg := range listed {
			pkgNames = append(pkgNames, pkg)
			seen[pkg] = true
		}
		for pkg := range byPkg {
			if !seen[pkg] {
				pkgNames = append(pkgNames, pkg)
			}
		}
	}
	sort.Strings(pkgNames)

	var allSkipped []SkippedJSON
	total := 0
	totalResources := 0
	totalData := 0
	for _, pkg := range pkgNames {
		dir := filepath.Join(serviceRoot, pkg)
		if !fileExists(dir) {
			for _, r := range byPkg[pkg] {
				allSkipped = append(allSkipped, SkippedJSON{
					Name:   skipName(r),
					Reason: fmt.Sprintf("service package directory %s is missing", pkg),
				})
			}
			continue
		}
		refs := byPkg[pkg]
		result, skipped := extractPackage(pkg, dir, refs, flexConsts, others)
		allSkipped = append(allSkipped, skipped...)
		nRes := countKind(result.Resources, "resource")
		nData := countKind(result.Resources, "data")
		total += len(result.Resources)
		totalResources += nRes
		totalData += nData
		path := filepath.Join(*outDir, pkg+".json")
		if err := writeJSON(path, result); err != nil {
			fmt.Fprintf(os.Stderr, "extractor: write %s: %v\n", path, err)
			os.Exit(1)
		}
		fmt.Printf("extracted %s: %d resources, %d data sources (%d skipped)\n", pkg, nRes, nData, len(result.Skipped))
	}

	index := map[string]any{
		"packages":       pkgNames,
		"resource_count": totalResources,
		"data_count":     totalData,
		"item_count":     total,
		"skipped":        allSkipped,
	}
	if len(want) == 0 {
		if err := writeJSON(filepath.Join(*outDir, "index.json"), index); err != nil {
			fmt.Fprintf(os.Stderr, "extractor: write index: %v\n", err)
			os.Exit(1)
		}
		if err := writeSkippedMarkdown(filepath.Join(filepath.Dir(*outDir), "..", "SKIPPED.md"), pkgNames, allSkipped, totalResources, totalData); err != nil {
			fmt.Fprintf(os.Stderr, "extractor: write SKIPPED.md: %v\n", err)
			os.Exit(1)
		}
	}
	fmt.Printf("done: %d resources, %d data sources across %d packages (%d skipped)\n", totalResources, totalData, len(pkgNames), len(allSkipped))
}

func extractPackage(pkg, dir string, refs []ResourceRef, flexConsts map[string]string, others map[string]*pkgParser) (PackageJSON, []SkippedJSON) {
	out := PackageJSON{
		Package:    pkg,
		Resources:  []ResourceJSON{},
		Skipped:    []SkippedJSON{},
		SourceHint: "derived schema from IBM Cloud Terraform provider service package " + pkg,
	}
	parser, err := parsePackageDir(dir)
	if err != nil {
		for _, r := range refs {
			out.Skipped = append(out.Skipped, SkippedJSON{
				Name:   skipName(r),
				Reason: fmt.Sprintf("failed to parse package AST: %v", err),
			})
		}
		return out, out.Skipped
	}
	parser.others = others
	validators := parser.extractValidators()

	for _, ref := range refs {
		fn, ok := parser.funcs[ref.Function]
		if !ok {
			out.Skipped = append(out.Skipped, SkippedJSON{
				Name:   skipName(ref),
				Reason: fmt.Sprintf("constructor %s not found in package %s", ref.Function, pkg),
			})
			continue
		}
		lit := parser.findReturnedResourceLit(fn)
		if lit == nil {
			out.Skipped = append(out.Skipped, SkippedJSON{
				Name:   skipName(ref),
				Reason: "constructor does not return a static schema.Resource literal (dynamic or generated at runtime)",
			})
			continue
		}
		attrs := parser.extractSchemaMapIn(fn, lit, flexConsts, 0)
		if attrs == nil {
			out.Skipped = append(out.Skipped, SkippedJSON{
				Name:   skipName(ref),
				Reason: "Schema field missing or not a static map literal",
			})
			continue
		}
		parser.applyReturnedWrapperMutations(fn, attrs)
		var vals []ValidatorJSON
		if v, ok := validators[ref.Name]; ok {
			sortValidators(v)
			mergeValidatorOntoAttrs(attrs, v)
			vals = v
		}
		kind := ref.Kind
		if kind == "" {
			kind = "resource"
		}
		res := ResourceJSON{
			Name:        ref.Name,
			Function:    ref.Function,
			Package:     pkg,
			Kind:        kind,
			Description: resourceDescription(ref.Name, kind, attrs),
			Attributes:  attrs,
			Validators:  vals,
		}
		out.Resources = append(out.Resources, res)
	}
	sort.Slice(out.Resources, func(i, j int) bool {
		if out.Resources[i].Kind != out.Resources[j].Kind {
			return out.Resources[i].Kind == "resource"
		}
		return out.Resources[i].Name < out.Resources[j].Name
	})
	sort.Slice(out.Skipped, func(i, j int) bool { return out.Skipped[i].Name < out.Skipped[j].Name })
	out.Extracted = len(out.Resources)
	return out, out.Skipped
}

func writeJSON(path string, v any) error {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	data = append(data, '\n')
	return os.WriteFile(path, data, 0o644)
}

func countKind(items []ResourceJSON, kind string) int {
	n := 0
	for _, r := range items {
		if r.Kind == kind {
			n++
		}
	}
	return n
}

func skipName(r ResourceRef) string {
	if r.Kind == "data" {
		return r.Name + " (data)"
	}
	return r.Name
}

func writeSkippedMarkdown(path string, pkgs []string, skipped []SkippedJSON, resources, data int) error {
	var b strings.Builder
	b.WriteString("# Skipped resources\n\n")
	b.WriteString("The extractor logs resources and data sources it cannot handle instead of dropping them silently.\n\n")
	fmt.Fprintf(&b, "Extracted %d managed resources and %d data sources across %d packages. Skipped %d constructors.\n\n", resources, data, len(pkgs), len(skipped))
	if len(skipped) == 0 {
		b.WriteString("Nothing skipped.\n")
		return os.WriteFile(path, []byte(b.String()), 0o644)
	}
	b.WriteString("| Name | Reason |\n| --- | --- |\n")
	for _, s := range skipped {
		reason := strings.ReplaceAll(s.Reason, "|", "\\|")
		name := strings.ReplaceAll(s.Name, "|", "\\|")
		fmt.Fprintf(&b, "| `%s` | %s |\n", name, reason)
	}
	b.WriteString("\nReasons used by the extractor:\n\n")
	b.WriteString("- constructor not found in the package\n")
	b.WriteString("- constructor does not return a static `schema.Resource` literal (dynamic or runtime-generated schema)\n")
	b.WriteString("- `Schema` field missing or not a static map literal\n")
	b.WriteString("- package directory missing\n")
	b.WriteString("- no managed resources or data sources registered in the provider maps\n")
	return os.WriteFile(path, []byte(b.String()), 0o644)
}
