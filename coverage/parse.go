// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type providerType struct {
	Name     string
	Kind     string // resource | data
	Package  string
	Function string
}

func constructorKind(fn string) string {
	switch {
	case strings.HasPrefix(fn, "Resource"), strings.HasPrefix(fn, "CombinedResource"):
		return "resource"
	case strings.HasPrefix(fn, "Data"), strings.HasPrefix(fn, "CombinedDataSource"):
		return "data"
	default:
		return ""
	}
}

func parseProviderMaps(providerFile string) ([]providerType, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, providerFile, nil, 0)
	if err != nil {
		return nil, err
	}
	var out []providerType
	seen := map[string]bool{}
	ast.Inspect(f, func(n ast.Node) bool {
		kv, ok := n.(*ast.KeyValueExpr)
		if !ok {
			return true
		}
		key, ok := litString(kv.Key)
		if !ok || !strings.HasPrefix(key, "ibm_") {
			return true
		}
		call := findConstructor(kv.Value)
		if call == nil {
			return true
		}
		pkg, fn := callPkgAndFunc(call)
		kind := constructorKind(fn)
		if kind == "" {
			return true
		}
		if strings.HasSuffix(fn, "Validator") {
			return true
		}
		id := kind + "\x00" + key
		if seen[id] {
			return true
		}
		seen[id] = true
		out = append(out, providerType{Name: key, Kind: kind, Package: pkg, Function: fn})
		return true
	})
	return out, nil
}

func findConstructor(expr ast.Expr) *ast.CallExpr {
	call, ok := expr.(*ast.CallExpr)
	if !ok {
		return nil
	}
	_, fn := callPkgAndFunc(call)
	if constructorKind(fn) != "" {
		return call
	}
	for _, arg := range call.Args {
		if inner, ok := arg.(*ast.CallExpr); ok {
			if _, innerFn := callPkgAndFunc(inner); constructorKind(innerFn) != "" {
				return inner
			}
		}
	}
	return nil
}

func callPkgAndFunc(call *ast.CallExpr) (string, string) {
	switch fn := call.Fun.(type) {
	case *ast.SelectorExpr:
		pkg := ""
		if id, ok := fn.X.(*ast.Ident); ok {
			pkg = id.Name
		}
		return pkg, fn.Sel.Name
	case *ast.Ident:
		return "", fn.Name
	}
	return "", ""
}

func litString(expr ast.Expr) (string, bool) {
	lit, ok := expr.(*ast.BasicLit)
	if !ok || lit.Kind != token.STRING {
		return "", false
	}
	s, err := strconv.Unquote(lit.Value)
	if err != nil {
		return strings.Trim(lit.Value, "`\""), true
	}
	return s, true
}

func listServicePackages(serviceRoot string) ([]string, error) {
	entries, err := os.ReadDir(serviceRoot)
	if err != nil {
		return nil, err
	}
	var pkgs []string
	for _, e := range entries {
		if e.IsDir() {
			pkgs = append(pkgs, e.Name())
		}
	}
	return pkgs, nil
}

func listSourceFiles(serviceRoot string) (resFiles, dataFiles []string, err error) {
	err = filepath.Walk(serviceRoot, func(path string, info os.FileInfo, walkErr error) error {
		if walkErr != nil || info.IsDir() {
			return walkErr
		}
		name := info.Name()
		if strings.HasSuffix(name, "_test.go") {
			return nil
		}
		rel, _ := filepath.Rel(serviceRoot, path)
		if strings.HasPrefix(name, "resource_ibm_") && strings.HasSuffix(name, ".go") {
			resFiles = append(resFiles, rel)
		}
		if strings.HasPrefix(name, "data_source_ibm_") && strings.HasSuffix(name, ".go") {
			dataFiles = append(dataFiles, rel)
		}
		return nil
	})
	return
}

func toPascal(name string) string {
	parts := strings.Split(name, "_")
	var b strings.Builder
	for _, p := range parts {
		if p == "" {
			continue
		}
		b.WriteString(strings.ToUpper(p[:1]) + p[1:])
	}
	return b.String()
}

func fileBase(name, kind string) string {
	if kind == "data" {
		return "data_" + name
	}
	return name
}

func ruleName(name, kind string) string {
	if kind == "data" {
		return name + "_data"
	}
	return name
}

func typeName(name, kind string) string {
	if kind == "data" {
		return toPascal(name) + "Data"
	}
	return toPascal(name)
}
