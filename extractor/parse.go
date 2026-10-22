// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

// pkgParser holds a parsed Go package plus resolved constants.
type pkgParser struct {
	dir    string
	fset   *token.FileSet
	files  []*ast.File
	consts map[string]string
	ints   map[string]int
	funcs  map[string]*ast.FuncDecl
	others map[string]*pkgParser
}

func parsePackageDir(dir string) (*pkgParser, error) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, dir, func(info os.FileInfo) bool {
		name := info.Name()
		if strings.HasSuffix(name, "_test.go") {
			return false
		}
		return strings.HasSuffix(name, ".go")
	}, parser.ParseComments)
	if err != nil {
		return nil, err
	}
	p := &pkgParser{
		dir:    dir,
		fset:   fset,
		consts: map[string]string{},
		ints:   map[string]int{},
		funcs:  map[string]*ast.FuncDecl{},
	}
	pkgNames := make([]string, 0, len(pkgs))
	for name := range pkgs {
		pkgNames = append(pkgNames, name)
	}
	sort.Strings(pkgNames)
	for _, name := range pkgNames {
		pkg := pkgs[name]
		fileNames := make([]string, 0, len(pkg.Files))
		for fn := range pkg.Files {
			fileNames = append(fileNames, fn)
		}
		sort.Strings(fileNames)
		for _, fn := range fileNames {
			f := pkg.Files[fn]
			p.files = append(p.files, f)
			p.collectDecls(f)
		}
	}
	return p, nil
}

func (p *pkgParser) collectDecls(f *ast.File) {
	for _, d := range f.Decls {
		switch decl := d.(type) {
		case *ast.FuncDecl:
			if decl.Name != nil {
				p.funcs[decl.Name.Name] = decl
			}
		case *ast.GenDecl:
			if decl.Tok == token.CONST {
				p.collectConsts(decl)
			}
		}
	}
}

func (p *pkgParser) collectConsts(decl *ast.GenDecl) {
	var lastType ast.Expr
	for _, spec := range decl.Specs {
		vs, ok := spec.(*ast.ValueSpec)
		if !ok {
			continue
		}
		if vs.Type != nil {
			lastType = vs.Type
		}
		_ = lastType
		for i, name := range vs.Names {
			if name == nil {
				continue
			}
			var val ast.Expr
			if i < len(vs.Values) {
				val = vs.Values[i]
			}
			if val == nil {
				continue
			}
			if s, ok := evalString(val, p.consts); ok {
				p.consts[name.Name] = s
			}
			if n, ok := evalInt(val, p.ints); ok {
				p.ints[name.Name] = n
			}
		}
	}
}

func (p *pkgParser) lookupString(name string, locals map[string]string) (string, bool) {
	if locals != nil {
		if v, ok := locals[name]; ok {
			return v, true
		}
	}
	if v, ok := p.consts[name]; ok {
		return v, true
	}
	return "", false
}

func evalString(expr ast.Expr, consts map[string]string) (string, bool) {
	switch e := expr.(type) {
	case *ast.BasicLit:
		if e.Kind == token.STRING {
			s, err := strconv.Unquote(e.Value)
			if err != nil {
				return strings.Trim(e.Value, "`\""), true
			}
			return s, true
		}
	case *ast.Ident:
		if consts != nil {
			if v, ok := consts[e.Name]; ok {
				return v, true
			}
		}
	case *ast.BinaryExpr:
		if e.Op == token.ADD {
			l, ok1 := evalString(e.X, consts)
			r, ok2 := evalString(e.Y, consts)
			if ok1 && ok2 {
				return l + r, true
			}
		}
	}
	return "", false
}

func evalInt(expr ast.Expr, ints map[string]int) (int, bool) {
	switch e := expr.(type) {
	case *ast.BasicLit:
		if e.Kind == token.INT {
			n, err := strconv.Atoi(e.Value)
			if err == nil {
				return n, true
			}
		}
	case *ast.UnaryExpr:
		if e.Op == token.SUB {
			if n, ok := evalInt(e.X, ints); ok {
				return -n, true
			}
		}
	case *ast.Ident:
		if ints != nil {
			if n, ok := ints[e.Name]; ok {
				return n, true
			}
		}
	}
	return 0, false
}

func evalBool(expr ast.Expr) (bool, bool) {
	id, ok := expr.(*ast.Ident)
	if !ok {
		return false, false
	}
	switch id.Name {
	case "true":
		return true, true
	case "false":
		return false, true
	}
	return false, false
}

func (p *pkgParser) resolveKey(expr ast.Expr, locals map[string]string, extras map[string]string) (string, bool) {
	switch e := expr.(type) {
	case *ast.BasicLit:
		return evalString(e, p.consts)
	case *ast.Ident:
		return p.lookupString(e.Name, locals)
	case *ast.SelectorExpr:
		sel := selectorName(e)
		if extras != nil {
			if v, ok := extras[sel]; ok {
				return v, true
			}
			if v, ok := extras[e.Sel.Name]; ok {
				return v, true
			}
		}
		if v, ok := p.consts[e.Sel.Name]; ok {
			return v, true
		}
		return "", false
	}
	return "", false
}

func selectorName(e *ast.SelectorExpr) string {
	if id, ok := e.X.(*ast.Ident); ok {
		return id.Name + "." + e.Sel.Name
	}
	return e.Sel.Name
}

func schemaTypeName(expr ast.Expr) string {
	sel, ok := expr.(*ast.SelectorExpr)
	if !ok {
		return ""
	}
	switch sel.Sel.Name {
	case "TypeString":
		return "string"
	case "TypeInt":
		return "int"
	case "TypeBool":
		return "bool"
	case "TypeFloat":
		return "float"
	case "TypeList":
		return "list"
	case "TypeSet":
		return "set"
	case "TypeMap":
		return "map"
	}
	return strings.ToLower(strings.TrimPrefix(sel.Sel.Name, "Type"))
}

func unwrapComposite(expr ast.Expr) *ast.CompositeLit {
	switch e := expr.(type) {
	case *ast.CompositeLit:
		return e
	case *ast.UnaryExpr:
		if e.Op == token.AND {
			if c, ok := e.X.(*ast.CompositeLit); ok {
				return c
			}
		}
	}
	return nil
}

func (p *pkgParser) findReturnedResourceLit(fn *ast.FuncDecl) *ast.CompositeLit {
	if fn == nil || fn.Body == nil {
		return nil
	}
	locals := map[string]*ast.CompositeLit{}
	for _, stmt := range fn.Body.List {
		switch s := stmt.(type) {
		case *ast.AssignStmt:
			if len(s.Lhs) == 1 && len(s.Rhs) == 1 {
				if id, ok := s.Lhs[0].(*ast.Ident); ok {
					if c := unwrapComposite(s.Rhs[0]); c != nil && isResourceLit(c) {
						locals[id.Name] = c
					}
					if call, ok := s.Rhs[0].(*ast.CallExpr); ok {
						if name := callFuncName(call); name != "" {
							if other, ok := p.funcs[name]; ok && other != fn {
								if lit := p.findReturnedResourceLit(other); lit != nil {
									locals[id.Name] = lit
								}
							}
						}
					}
				}
			}
		case *ast.ReturnStmt:
			if len(s.Results) == 0 {
				continue
			}
			if c := unwrapComposite(s.Results[0]); c != nil && isResourceLit(c) {
				return c
			}
			if id, ok := s.Results[0].(*ast.Ident); ok {
				if lit, ok := locals[id.Name]; ok {
					return lit
				}
			}
			if call, ok := s.Results[0].(*ast.CallExpr); ok {
				if name := callFuncName(call); name != "" {
					if other, ok := p.funcs[name]; ok && other != fn {
						if lit := p.findReturnedResourceLit(other); lit != nil {
							return lit
						}
					}
				}
				// wrappers like AddSchemaData(&schema.Resource{...}) or AddInstanceFields(ResourceFoo())
				if len(call.Args) > 0 {
					if inner, ok := call.Args[0].(*ast.CallExpr); ok {
						if name := callFuncName(inner); name != "" {
							if other, ok := p.funcs[name]; ok && other != fn {
								if lit := p.findReturnedResourceLit(other); lit != nil {
									return lit
								}
							}
						}
					}
					if c := unwrapComposite(call.Args[0]); c != nil && isResourceLit(c) {
						return c
					}
				}
			}
		}
	}
	return nil
}

func isResourceLit(c *ast.CompositeLit) bool {
	if c == nil {
		return false
	}
	if sel, ok := c.Type.(*ast.SelectorExpr); ok {
		if sel.Sel.Name == "Resource" {
			return true
		}
	}
	// Infer: has Schema field.
	for _, elt := range c.Elts {
		if kv, ok := elt.(*ast.KeyValueExpr); ok {
			if id, ok := kv.Key.(*ast.Ident); ok && id.Name == "Schema" {
				return true
			}
		}
	}
	return false
}

func (p *pkgParser) applyReturnedWrapperMutations(fn *ast.FuncDecl, attrs map[string]*AttrJSON) {
	if fn == nil || fn.Body == nil || attrs == nil {
		return
	}
	for _, stmt := range fn.Body.List {
		ret, ok := stmt.(*ast.ReturnStmt)
		if !ok || len(ret.Results) == 0 {
			continue
		}
		call, ok := ret.Results[0].(*ast.CallExpr)
		if !ok {
			continue
		}
		name := callFuncName(call)
		if name == "" {
			continue
		}
		other, ok := p.funcs[name]
		if !ok {
			continue
		}
		p.mergeSchemaIndexAssigns(other, attrs)
	}
}

func (p *pkgParser) mergeSchemaIndexAssigns(fn *ast.FuncDecl, attrs map[string]*AttrJSON) {
	if fn == nil || fn.Body == nil || attrs == nil {
		return
	}
	ast.Inspect(fn.Body, func(n ast.Node) bool {
		as, ok := n.(*ast.AssignStmt)
		if !ok || len(as.Lhs) == 0 || len(as.Rhs) == 0 {
			return true
		}
		idx, ok := as.Lhs[0].(*ast.IndexExpr)
		if !ok {
			return true
		}
		sel, ok := idx.X.(*ast.SelectorExpr)
		if !ok || sel.Sel.Name != "Schema" {
			return true
		}
		key, ok := p.resolveKey(idx.Index, nil, nil)
		if !ok || key == "" {
			return true
		}
		attr := p.parseAttr(key, as.Rhs[0], nil, 0)
		if attr != nil {
			attrs[key] = attr
		}
		return true
	})
}

func callFuncName(call *ast.CallExpr) string {
	switch f := call.Fun.(type) {
	case *ast.Ident:
		return f.Name
	case *ast.SelectorExpr:
		return f.Sel.Name
	}
	return ""
}

func callPkgAndFunc(call *ast.CallExpr) (pkg, fn string) {
	switch f := call.Fun.(type) {
	case *ast.Ident:
		return "", f.Name
	case *ast.SelectorExpr:
		if id, ok := f.X.(*ast.Ident); ok {
			return id.Name, f.Sel.Name
		}
		return "", f.Sel.Name
	}
	return "", ""
}

func (p *pkgParser) extractSchemaMap(lit *ast.CompositeLit, extras map[string]string) map[string]*AttrJSON {
	return p.extractSchemaMapIn(nil, lit, extras, 0)
}

func (p *pkgParser) extractSchemaMapIn(fn *ast.FuncDecl, lit *ast.CompositeLit, extras map[string]string, depth int) map[string]*AttrJSON {
	if lit == nil || depth > 8 {
		return nil
	}
	var schemaExpr ast.Expr
	for _, elt := range lit.Elts {
		kv, ok := elt.(*ast.KeyValueExpr)
		if !ok {
			continue
		}
		id, ok := kv.Key.(*ast.Ident)
		if !ok || id.Name != "Schema" {
			continue
		}
		schemaExpr = kv.Value
	}
	if schemaExpr == nil {
		// Deprecated stubs register a resource literal with no Schema field.
		return map[string]*AttrJSON{}
	}
	return p.parseSchemaRHS(fn, schemaExpr, extras, depth)
}

func (p *pkgParser) parseSchemaMap(expr ast.Expr, extras map[string]string, depth int) map[string]*AttrJSON {
	if depth > 8 {
		return nil
	}
	c := unwrapComposite(expr)
	if c == nil {
		return nil
	}
	out := map[string]*AttrJSON{}
	for _, elt := range c.Elts {
		kv, ok := elt.(*ast.KeyValueExpr)
		if !ok {
			continue
		}
		key, ok := p.resolveKey(kv.Key, nil, extras)
		if !ok || key == "" {
			continue
		}
		attr := p.parseAttr(key, kv.Value, extras, depth)
		if attr != nil {
			out[key] = attr
		}
	}
	return out
}

func (p *pkgParser) parseSchemaRHS(fn *ast.FuncDecl, expr ast.Expr, extras map[string]string, depth int) map[string]*AttrJSON {
	if expr == nil || depth > 8 {
		return nil
	}
	if m := p.parseSchemaMap(expr, extras, depth); m != nil {
		return m
	}
	switch e := expr.(type) {
	case *ast.Ident:
		if fn != nil {
			return p.resolveSchemaIdent(fn, e.Name, extras, depth+1)
		}
		// package-level var assigned from a helper
		if other, ok := p.funcs[e.Name]; ok {
			return p.findReturnedSchemaMap(other, extras, depth+1)
		}
	case *ast.CallExpr:
		name := callFuncName(e)
		if name == "Clone" && len(e.Args) == 1 {
			return p.parseSchemaRHS(fn, e.Args[0], extras, depth+1)
		}
		if other, ok := p.funcs[name]; ok {
			if m := p.findReturnedSchemaMap(other, extras, depth+1); m != nil {
				return m
			}
			if lit := p.findReturnedResourceLit(other); lit != nil {
				return p.extractSchemaMapIn(other, lit, extras, depth+1)
			}
		}
		pkg, fnName := callPkgAndFunc(e)
		if pkg != "" && fnName != "" && p.others != nil {
			if other := p.others[pkg]; other != nil {
				if ofn, ok := other.funcs[fnName]; ok {
					if lit := other.findReturnedResourceLit(ofn); lit != nil {
						return other.extractSchemaMapIn(ofn, lit, extras, depth+1)
					}
					return other.findReturnedSchemaMap(ofn, extras, depth+1)
				}
			}
		}
	case *ast.SelectorExpr:
		if e.Sel.Name == "Schema" {
			return p.parseSchemaRHS(fn, e.X, extras, depth+1)
		}
	case *ast.ParenExpr:
		return p.parseSchemaRHS(fn, e.X, extras, depth+1)
	}
	return nil
}

func (p *pkgParser) findReturnedSchemaMap(fn *ast.FuncDecl, extras map[string]string, depth int) map[string]*AttrJSON {
	if fn == nil || fn.Body == nil || depth > 8 {
		return nil
	}
	for _, stmt := range fn.Body.List {
		ret, ok := stmt.(*ast.ReturnStmt)
		if !ok || len(ret.Results) == 0 {
			continue
		}
		if m := p.parseSchemaRHS(fn, ret.Results[0], extras, depth+1); m != nil {
			return m
		}
	}
	return nil
}

func (p *pkgParser) resolveSchemaIdent(fn *ast.FuncDecl, name string, extras map[string]string, depth int) map[string]*AttrJSON {
	if fn == nil || fn.Body == nil || name == "" || depth > 8 {
		return nil
	}
	var attrs map[string]*AttrJSON
	for _, stmt := range fn.Body.List {
		as, ok := stmt.(*ast.AssignStmt)
		if !ok || len(as.Lhs) == 0 || len(as.Rhs) == 0 {
			continue
		}
		switch lhs := as.Lhs[0].(type) {
		case *ast.Ident:
			if lhs.Name != name {
				continue
			}
			if m := p.parseSchemaRHS(fn, as.Rhs[0], extras, depth+1); m != nil {
				attrs = m
			}
		case *ast.IndexExpr:
			xid, ok := lhs.X.(*ast.Ident)
			if !ok || xid.Name != name || attrs == nil {
				continue
			}
			key, ok := p.resolveKey(lhs.Index, nil, extras)
			if !ok || key == "" {
				continue
			}
			attr := p.parseAttr(key, as.Rhs[0], extras, depth)
			if attr != nil {
				attrs[key] = attr
			}
		}
	}
	return attrs
}

func (p *pkgParser) parseAttr(name string, expr ast.Expr, extras map[string]string, depth int) *AttrJSON {
	c := unwrapComposite(expr)
	if c == nil {
		return nil
	}
	attr := &AttrJSON{Name: name, Type: "string"}
	for _, elt := range c.Elts {
		kv, ok := elt.(*ast.KeyValueExpr)
		if !ok {
			continue
		}
		field, ok := kv.Key.(*ast.Ident)
		if !ok {
			continue
		}
		switch field.Name {
		case "Type":
			if t := schemaTypeName(kv.Value); t != "" {
				attr.Type = t
			}
		case "Required":
			if v, ok := evalBool(kv.Value); ok {
				attr.Required = v
			}
		case "Optional":
			if v, ok := evalBool(kv.Value); ok {
				attr.Optional = v
			}
		case "Computed":
			if v, ok := evalBool(kv.Value); ok {
				attr.Computed = v
			}
		case "ForceNew":
			if v, ok := evalBool(kv.Value); ok {
				attr.ForceNew = v
			}
		case "Sensitive":
			if v, ok := evalBool(kv.Value); ok {
				attr.Sensitive = v
			}
		case "Description":
			if s, ok := evalString(kv.Value, p.consts); ok {
				attr.Description = s
			}
		case "Default":
			if s, ok := evalString(kv.Value, p.consts); ok {
				attr.Default = s
			} else if n, ok := evalInt(kv.Value, p.ints); ok {
				attr.Default = strconv.Itoa(n)
			} else if b, ok := evalBool(kv.Value); ok {
				attr.Default = strconv.FormatBool(b)
			}
		case "MinItems":
			if n, ok := evalInt(kv.Value, p.ints); ok {
				attr.MinItems = intPtr(n)
			}
		case "MaxItems":
			if n, ok := evalInt(kv.Value, p.ints); ok {
				attr.MaxItems = intPtr(n)
			}
		case "ConflictsWith":
			attr.Conflicts = stringSliceLit(kv.Value, p.consts)
		case "ValidateFunc":
			attr.Validate = p.parseValidateFunc(kv.Value)
		case "Elem":
			attr.Elem = p.parseElem(kv.Value, extras, depth+1)
		}
	}
	if attr.Validate == nil {
		attr.Validate = inferFormatFromName(name, attr.Type)
	} else {
		mergeFormatHeuristic(attr)
	}
	return attr
}

func (p *pkgParser) parseElem(expr ast.Expr, extras map[string]string, depth int) *ElemJSON {
	c := unwrapComposite(expr)
	if c == nil {
		return nil
	}
	kind := "schema"
	if isResourceLit(c) {
		kind = "resource"
	}
	if sel, ok := c.Type.(*ast.SelectorExpr); ok {
		if sel.Sel.Name == "Resource" {
			kind = "resource"
		}
		if sel.Sel.Name == "Schema" {
			kind = "schema"
		}
	}
	elem := &ElemJSON{Kind: kind}
	if kind == "resource" {
		elem.Attributes = p.extractSchemaMapIn(nil, c, extras, depth+1)
		return elem
	}
	// primitive schema
	for _, elt := range c.Elts {
		kv, ok := elt.(*ast.KeyValueExpr)
		if !ok {
			continue
		}
		field, ok := kv.Key.(*ast.Ident)
		if !ok {
			continue
		}
		switch field.Name {
		case "Type":
			elem.Type = schemaTypeName(kv.Value)
		case "ValidateFunc":
			elem.Validate = p.parseValidateFunc(kv.Value)
		case "Elem":
			// list of list — rare; ignore deeper than one extra level
		}
	}
	if elem.Type == "" {
		elem.Type = "string"
	}
	return elem
}

func (p *pkgParser) parseValidateFunc(expr ast.Expr) *ValidateJSON {
	call, ok := expr.(*ast.CallExpr)
	if !ok {
		if id, ok := expr.(*ast.Ident); ok {
			return inferFromFuncName(id.Name)
		}
		if sel, ok := expr.(*ast.SelectorExpr); ok {
			return inferFromFuncName(sel.Sel.Name)
		}
		return nil
	}
	pkg, fn := callPkgAndFunc(call)
	_ = pkg
	switch fn {
	case "InvokeValidator", "InvokeDataSourceValidator":
		v := &ValidateJSON{Kind: "invoke", FuncName: fn}
		if len(call.Args) >= 2 {
			if s, ok := evalString(call.Args[1], p.consts); ok {
				v.FuncName = s
			} else if id, ok := call.Args[1].(*ast.Ident); ok {
				if s, ok := p.lookupString(id.Name, nil); ok {
					v.FuncName = s
				} else {
					v.FuncName = id.Name
				}
			}
		}
		return v
	case "StringInSlice":
		v := &ValidateJSON{Kind: "enum", FuncName: fn}
		if len(call.Args) >= 1 {
			v.Enum = stringSliceLit(call.Args[0], p.consts)
		}
		if len(call.Args) >= 2 {
			if b, ok := evalBool(call.Args[1]); ok {
				v.IgnoreCase = b
			}
		}
		return v
	case "ValidateAllowedStringValues", "ValidateAllowedStringValue":
		v := &ValidateJSON{Kind: "enum", FuncName: fn}
		if len(call.Args) >= 1 {
			v.Enum = stringSliceLit(call.Args[0], p.consts)
		}
		return v
	case "IntBetween":
		v := &ValidateJSON{Kind: "int_range", FuncName: fn}
		if len(call.Args) >= 2 {
			if n, ok := evalInt(call.Args[0], p.ints); ok {
				v.Min = intPtr(n)
			}
			if n, ok := evalInt(call.Args[1], p.ints); ok {
				v.Max = intPtr(n)
			}
		}
		return v
	case "IntAtLeast":
		v := &ValidateJSON{Kind: "int_min", FuncName: fn}
		if len(call.Args) >= 1 {
			if n, ok := evalInt(call.Args[0], p.ints); ok {
				v.Min = intPtr(n)
			}
		}
		return v
	case "IntAtMost":
		v := &ValidateJSON{Kind: "int_max", FuncName: fn}
		if len(call.Args) >= 1 {
			if n, ok := evalInt(call.Args[0], p.ints); ok {
				v.Max = intPtr(n)
			}
		}
		return v
	case "StringLenBetween":
		v := &ValidateJSON{Kind: "string_len", FuncName: fn}
		if len(call.Args) >= 2 {
			if n, ok := evalInt(call.Args[0], p.ints); ok {
				v.MinLen = intPtr(n)
			}
			if n, ok := evalInt(call.Args[1], p.ints); ok {
				v.MaxLen = intPtr(n)
			}
		}
		return v
	case "IsCIDR", "IsCIDRNetwork", "validateCIDRAddress", "ValidateCIDRAddress":
		return &ValidateJSON{Kind: "cidr", FuncName: fn}
	case "IsIPAddress", "IsIPv4Address", "validateIPorCIDR", "ValidateIPorCIDR":
		return &ValidateJSON{Kind: "ipv4", FuncName: fn}
	case "StringMatch":
		v := &ValidateJSON{Kind: "regexp", FuncName: fn}
		if len(call.Args) >= 1 {
			v.Regexp = regexpFromExpr(call.Args[0])
		}
		return v
	default:
		return inferFromFuncName(fn)
	}
}

func inferFromFuncName(name string) *ValidateJSON {
	l := strings.ToLower(name)
	switch {
	case strings.Contains(l, "cidr"):
		return &ValidateJSON{Kind: "cidr", FuncName: name}
	case strings.Contains(l, "ipv4") || strings.Contains(l, "ipaddress") || strings.Contains(l, "ip_address"):
		return &ValidateJSON{Kind: "ipv4", FuncName: name}
	case strings.Contains(l, "crn"):
		return &ValidateJSON{Kind: "crn", FuncName: name}
	case looksLikePortValidator(l):
		return &ValidateJSON{Kind: "port", FuncName: name, Min: intPtr(1), Max: intPtr(65535)}
	case strings.Contains(l, "json"):
		return &ValidateJSON{Kind: "json", FuncName: name}
	case strings.Contains(l, "zone"):
		return &ValidateJSON{Kind: "zone", FuncName: name}
	default:
		if name == "" {
			return nil
		}
		return &ValidateJSON{Kind: "func", FuncName: name}
	}
}

func inferFormatFromName(name, typ string) *ValidateJSON {
	l := strings.ToLower(name)
	if typ != "string" && typ != "int" {
		return nil
	}
	switch {
	case l == "crn" || strings.HasSuffix(l, "_crn"):
		return &ValidateJSON{Kind: "crn", FuncName: "name_heuristic"}
	case strings.Contains(l, "cidr"):
		return &ValidateJSON{Kind: "cidr", FuncName: "name_heuristic"}
	case l == "ipv4_address" || l == "ip_address" || strings.HasSuffix(l, "_ipv4") || l == "ipv4":
		return &ValidateJSON{Kind: "ipv4", FuncName: "name_heuristic"}
	case l == "zone" || l == "zone_name" || strings.HasSuffix(l, "_zone"):
		return &ValidateJSON{Kind: "zone", FuncName: "name_heuristic"}
	case typ == "int" && isTCPPortAttrName(l):
		return &ValidateJSON{Kind: "port", FuncName: "name_heuristic", Min: intPtr(1), Max: intPtr(65535)}
	}
	return nil
}

func isTCPPortAttrName(l string) bool {
	return l == "port" || l == "port_min" || l == "port_max" || strings.HasSuffix(l, "_port")
}

func looksLikePortValidator(l string) bool {
	if strings.Contains(l, "support") || strings.Contains(l, "export") || strings.Contains(l, "report") || strings.Contains(l, "transport") {
		return false
	}
	return strings.Contains(l, "port")
}

func mergeFormatHeuristic(attr *AttrJSON) {
	if attr.Validate == nil {
		return
	}
	if attr.Validate.Kind == "invoke" || attr.Validate.Kind == "func" {
		if h := inferFormatFromName(attr.Name, attr.Type); h != nil && h.Kind != "func" {
			// Keep invoke identifier but attach format kind for generator merge later.
			if attr.Validate.Kind == "func" {
				attr.Validate.Kind = h.Kind
			}
		}
	}
}

func stringSliceLit(expr ast.Expr, consts map[string]string) []string {
	c := unwrapComposite(expr)
	if c == nil {
		if s, ok := evalString(expr, consts); ok {
			return splitCSV(s)
		}
		return nil
	}
	var out []string
	for _, elt := range c.Elts {
		if s, ok := evalString(elt, consts); ok {
			out = append(out, s)
		}
	}
	return out
}

func splitCSV(s string) []string {
	parts := strings.Split(s, ",")
	var out []string
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			out = append(out, p)
		}
	}
	return out
}

func regexpFromExpr(expr ast.Expr) string {
	if call, ok := expr.(*ast.CallExpr); ok {
		if sel, ok := call.Fun.(*ast.SelectorExpr); ok && sel.Sel.Name == "MustCompile" {
			if len(call.Args) > 0 {
				if s, ok := evalString(call.Args[0], nil); ok {
					return s
				}
			}
		}
	}
	if s, ok := evalString(expr, nil); ok {
		return s
	}
	return ""
}

func intPtr(n int) *int { return &n }

func (p *pkgParser) extractValidators() map[string][]ValidatorJSON {
	out := map[string][]ValidatorJSON{}
	funcNames := make([]string, 0, len(p.funcs))
	for name := range p.funcs {
		funcNames = append(funcNames, name)
	}
	sort.Strings(funcNames)
	for _, name := range funcNames {
		fn := p.funcs[name]
		if fn.Body == nil {
			continue
		}
		if !returnsResourceValidator(fn) {
			continue
		}
		locals := map[string]string{}
		var resourceName string
		var entries []ValidatorJSON
		ast.Inspect(fn.Body, func(n ast.Node) bool {
			switch node := n.(type) {
			case *ast.AssignStmt:
				if len(node.Lhs) == 1 && len(node.Rhs) == 1 {
					if id, ok := node.Lhs[0].(*ast.Ident); ok {
						if s, ok := evalString(node.Rhs[0], p.consts); ok {
							locals[id.Name] = s
						}
					}
				}
			case *ast.CompositeLit:
				if !isValidateSchemaLit(node) {
					return true
				}
				v := p.parseValidateSchema(node, locals)
				if v.Identifier != "" {
					entries = append(entries, v)
				}
				if isResourceValidatorLit(node) {
					if name := p.compositeStringField(node, "ResourceName", locals); name != "" {
						resourceName = name
					}
				}
			}
			return true
		})
		if resourceName == "" {
			resourceName = p.scanResourceName(fn, locals)
		}
		if resourceName != "" && len(entries) > 0 {
			out[resourceName] = append(out[resourceName], entries...)
		}
	}
	for resourceName, entries := range out {
		sortValidators(entries)
		out[resourceName] = entries
	}
	return out
}

func returnsResourceValidator(fn *ast.FuncDecl) bool {
	if fn.Type == nil || fn.Type.Results == nil {
		return false
	}
	for _, f := range fn.Type.Results.List {
		switch t := f.Type.(type) {
		case *ast.StarExpr:
			if sel, ok := t.X.(*ast.SelectorExpr); ok && sel.Sel.Name == "ResourceValidator" {
				return true
			}
		case *ast.SelectorExpr:
			if t.Sel.Name == "ResourceValidator" {
				return true
			}
		}
	}
	return strings.HasSuffix(fn.Name.Name, "Validator")
}

func isValidateSchemaLit(c *ast.CompositeLit) bool {
	if sel, ok := c.Type.(*ast.SelectorExpr); ok {
		return sel.Sel.Name == "ValidateSchema"
	}
	// untyped composite with Identifier field
	hasIdent := false
	hasFunc := false
	for _, elt := range c.Elts {
		kv, ok := elt.(*ast.KeyValueExpr)
		if !ok {
			continue
		}
		id, ok := kv.Key.(*ast.Ident)
		if !ok {
			continue
		}
		if id.Name == "Identifier" {
			hasIdent = true
		}
		if id.Name == "ValidateFunctionIdentifier" {
			hasFunc = true
		}
	}
	return hasIdent && hasFunc
}

func isResourceValidatorLit(c *ast.CompositeLit) bool {
	if sel, ok := c.Type.(*ast.SelectorExpr); ok {
		return sel.Sel.Name == "ResourceValidator"
	}
	return false
}

func (p *pkgParser) parseValidateSchema(c *ast.CompositeLit, locals map[string]string) ValidatorJSON {
	v := ValidatorJSON{}
	for _, elt := range c.Elts {
		kv, ok := elt.(*ast.KeyValueExpr)
		if !ok {
			continue
		}
		id, ok := kv.Key.(*ast.Ident)
		if !ok {
			continue
		}
		switch id.Name {
		case "Identifier":
			if s, ok := p.resolveKey(kv.Value, locals, nil); ok {
				v.Identifier = s
			}
		case "ValidateFunctionIdentifier":
			v.Kind = functionIdentifierName(kv.Value)
		case "AllowedValues":
			if s, ok := evalString(kv.Value, mergeMaps(p.consts, locals)); ok {
				v.Allowed = s
			} else if id, ok := kv.Value.(*ast.Ident); ok {
				if s, ok := p.lookupString(id.Name, locals); ok {
					v.Allowed = s
				}
			}
		case "MinValue":
			if s, ok := evalString(kv.Value, p.consts); ok {
				v.Min = s
			} else if n, ok := evalInt(kv.Value, p.ints); ok {
				v.Min = strconv.Itoa(n)
			}
		case "MaxValue":
			if s, ok := evalString(kv.Value, p.consts); ok {
				v.Max = s
			} else if n, ok := evalInt(kv.Value, p.ints); ok {
				v.Max = strconv.Itoa(n)
			}
		case "Regexp":
			if s, ok := evalString(kv.Value, p.consts); ok {
				v.Regexp = s
			}
		case "MinValueLength":
			if n, ok := evalInt(kv.Value, p.ints); ok {
				v.MinLen = n
			}
		case "MaxValueLength":
			if n, ok := evalInt(kv.Value, p.ints); ok {
				v.MaxLen = n
			}
		}
	}
	return v
}

func (p *pkgParser) compositeStringField(c *ast.CompositeLit, field string, locals map[string]string) string {
	for _, elt := range c.Elts {
		kv, ok := elt.(*ast.KeyValueExpr)
		if !ok {
			continue
		}
		id, ok := kv.Key.(*ast.Ident)
		if !ok || id.Name != field {
			continue
		}
		if s, ok := p.resolveKey(kv.Value, locals, nil); ok {
			return s
		}
		if s, ok := evalString(kv.Value, p.consts); ok {
			return s
		}
	}
	return ""
}

func (p *pkgParser) scanResourceName(fn *ast.FuncDecl, locals map[string]string) string {
	var name string
	ast.Inspect(fn.Body, func(n ast.Node) bool {
		kv, ok := n.(*ast.KeyValueExpr)
		if !ok {
			return true
		}
		id, ok := kv.Key.(*ast.Ident)
		if !ok || id.Name != "ResourceName" {
			return true
		}
		if s, ok := p.resolveKey(kv.Value, locals, nil); ok {
			name = s
		}
		return true
	})
	return name
}

func functionIdentifierName(expr ast.Expr) string {
	switch e := expr.(type) {
	case *ast.SelectorExpr:
		return e.Sel.Name
	case *ast.Ident:
		return e.Name
	}
	return ""
}

func mergeMaps(a, b map[string]string) map[string]string {
	out := map[string]string{}
	for k, v := range a {
		out[k] = v
	}
	for k, v := range b {
		out[k] = v
	}
	return out
}

func validatorToValidate(v ValidatorJSON) *ValidateJSON {
	out := &ValidateJSON{Kind: "func", FuncName: v.Kind}
	switch v.Kind {
	case "ValidateAllowedStringValue", "ValidateAllowedICDPlanValue":
		out.Kind = "enum"
		out.Enum = splitCSV(v.Allowed)
	case "ValidateAllowedIntValue":
		out.Kind = "int_enum"
		for _, s := range splitCSV(v.Allowed) {
			if n, err := strconv.Atoi(strings.TrimSpace(s)); err == nil {
				out.IntEnum = append(out.IntEnum, n)
			}
		}
	case "IntBetween":
		out.Kind = "int_range"
		if n, err := strconv.Atoi(v.Min); err == nil {
			out.Min = intPtr(n)
		}
		if n, err := strconv.Atoi(v.Max); err == nil {
			out.Max = intPtr(n)
		}
	case "IntAtLeast":
		out.Kind = "int_min"
		if n, err := strconv.Atoi(v.Min); err == nil {
			out.Min = intPtr(n)
		}
	case "IntAtMost":
		out.Kind = "int_max"
		if n, err := strconv.Atoi(v.Max); err == nil {
			out.Max = intPtr(n)
		}
	case "StringLenBetween":
		out.Kind = "string_len"
		if v.MinLen > 0 {
			out.MinLen = intPtr(v.MinLen)
		}
		if v.MaxLen > 0 {
			out.MaxLen = intPtr(v.MaxLen)
		}
	case "ValidateRegexpLen":
		out.Kind = "regexp"
		out.Regexp = v.Regexp
		if v.MinLen > 0 {
			out.MinLen = intPtr(v.MinLen)
		}
		if v.MaxLen > 0 {
			out.MaxLen = intPtr(v.MaxLen)
		}
	case "ValidateRegexp":
		out.Kind = "regexp"
		out.Regexp = v.Regexp
	case "ValidateCIDRAddress":
		out.Kind = "cidr"
	case "ValidateIPorCIDR":
		out.Kind = "ip_or_cidr"
	case "ValidateJSONString", "ValidateJSONParam":
		out.Kind = "json"
	case "ValidateNoZeroValues":
		out.Kind = "no_zero"
		out.NoZero = true
	case "ValidateCloudData":
		out.Kind = "cloud_data"
	default:
		if v.Regexp != "" {
			out.Kind = "regexp"
			out.Regexp = v.Regexp
		} else if v.Allowed != "" {
			out.Kind = "enum"
			out.Enum = splitCSV(v.Allowed)
		}
	}
	return out
}

func mergeValidatorOntoAttrs(attrs map[string]*AttrJSON, vals []ValidatorJSON) {
	for _, v := range vals {
		if v.Identifier == "" || v.Identifier == "id" {
			continue
		}
		converted := validatorToValidate(v)
		if converted == nil || converted.Kind == "cloud_data" {
			continue
		}
		if v.Kind == "ValidateNoZeroValues" {
			converted.NoZero = true
		}
		attr, onElem := findAttrForValidator(attrs, v.Identifier)
		if attr == nil {
			continue
		}
		target := &attr.Validate
		if onElem {
			if attr.Elem == nil {
				attr.Elem = &ElemJSON{Kind: "schema", Type: "string"}
			}
			target = &attr.Elem.Validate
		}
		*target = mergeValidate(*target, converted)
	}
}

func mergeValidate(existing *ValidateJSON, incoming *ValidateJSON) *ValidateJSON {
	if incoming == nil {
		return existing
	}
	if existing == nil || existing.Kind == "invoke" || existing.Kind == "func" {
		return incoming
	}
	out := *existing
	if out.FuncName == "" && incoming.FuncName != "" {
		out.FuncName = incoming.FuncName
	}
	if incoming.NoZero {
		out.NoZero = true
	}
	if len(incoming.Enum) > 0 {
		out.Enum = incoming.Enum
		if out.Kind == "invoke" || out.Kind == "func" || out.Kind == "no_zero" {
			out.Kind = "enum"
		}
	}
	if len(incoming.IntEnum) > 0 {
		out.IntEnum = incoming.IntEnum
		if out.Kind == "no_zero" || out.Kind == "invoke" || out.Kind == "func" {
			out.Kind = "int_enum"
		}
	}
	if incoming.Regexp != "" {
		out.Regexp = incoming.Regexp
		if out.Kind == "invoke" || out.Kind == "func" || out.Kind == "no_zero" {
			out.Kind = "regexp"
		}
	}
	if incoming.Min != nil {
		out.Min = incoming.Min
	}
	if incoming.Max != nil {
		out.Max = incoming.Max
	}
	if incoming.MinLen != nil {
		out.MinLen = incoming.MinLen
	}
	if incoming.MaxLen != nil {
		out.MaxLen = incoming.MaxLen
	}
	if incoming.Kind == "int_range" || incoming.Kind == "cidr" || incoming.Kind == "ipv4" || incoming.Kind == "port" || incoming.Kind == "crn" || incoming.Kind == "zone" || incoming.Kind == "ip_or_cidr" {
		if out.Kind == "invoke" || out.Kind == "func" || out.Kind == "no_zero" {
			out.Kind = incoming.Kind
		}
	}
	return &out
}

func findAttrForValidator(attrs map[string]*AttrJSON, identifier string) (*AttrJSON, bool) {
	if a, ok := attrs[identifier]; ok {
		return a, false
	}
	names := make([]string, 0, len(attrs))
	for n := range attrs {
		names = append(names, n)
	}
	sort.Strings(names)
	for _, n := range names {
		a := attrs[n]
		if a.Validate != nil && a.Validate.FuncName == identifier {
			return a, false
		}
		if a.Elem != nil && a.Elem.Validate != nil && a.Elem.Validate.FuncName == identifier {
			return a, true
		}
	}
	return nil, false
}

func sortValidators(vals []ValidatorJSON) {
	sort.SliceStable(vals, func(i, j int) bool {
		a, b := vals[i], vals[j]
		if a.Identifier != b.Identifier {
			return a.Identifier < b.Identifier
		}
		if a.Kind != b.Kind {
			return a.Kind < b.Kind
		}
		if a.Allowed != b.Allowed {
			return a.Allowed < b.Allowed
		}
		if a.Regexp != b.Regexp {
			return a.Regexp < b.Regexp
		}
		if a.Min != b.Min {
			return a.Min < b.Min
		}
		if a.Max != b.Max {
			return a.Max < b.Max
		}
		if a.MinLen != b.MinLen {
			return a.MinLen < b.MinLen
		}
		return a.MaxLen < b.MaxLen
	})
}

func constructorKind(fn string) string {
	switch {
	case strings.HasPrefix(fn, "Resource"), strings.HasPrefix(fn, "CombinedResource"):
		return "resource"
	case strings.HasPrefix(fn, "Data"), strings.HasPrefix(fn, "CombinedDataSource"):
		// DataSource*, Datasource*, DataIBM*, CombinedDataSource*
		return "data"
	default:
		return ""
	}
}

func parseProviderResources(providerFile string) ([]ResourceRef, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, providerFile, nil, 0)
	if err != nil {
		return nil, err
	}
	var refs []ResourceRef
	ast.Inspect(f, func(n ast.Node) bool {
		kv, ok := n.(*ast.KeyValueExpr)
		if !ok {
			return true
		}
		key, ok := evalString(kv.Key, nil)
		if !ok || !strings.HasPrefix(key, "ibm_") {
			return true
		}
		call := extractResourceCall(kv.Value)
		if call == nil {
			return true
		}
		pkg, fn := callPkgAndFunc(call)
		kind := constructorKind(fn)
		if kind == "" {
			return true
		}
		refs = append(refs, ResourceRef{
			Name:     key,
			Package:  pkg,
			Function: fn,
			Kind:     kind,
		})
		return true
	})
	seen := map[string]bool{}
	var out []ResourceRef
	for _, r := range refs {
		id := r.Kind + "\x00" + r.Name
		if seen[id] {
			continue
		}
		seen[id] = true
		out = append(out, r)
	}
	return out, nil
}

func extractResourceCall(expr ast.Expr) *ast.CallExpr {
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

func isWrappedCall(expr ast.Expr) bool {
	call, ok := expr.(*ast.CallExpr)
	if !ok {
		return false
	}
	_, fn := callPkgAndFunc(call)
	return constructorKind(fn) == ""
}

func parseFlexConsts(flexDir string) map[string]string {
	out := map[string]string{}
	p, err := parsePackageDir(flexDir)
	if err != nil {
		return out
	}
	for k, v := range p.consts {
		out[k] = v
		out["flex."+k] = v
	}
	return out
}

func firstSentence(desc string) string {
	desc = strings.TrimSpace(desc)
	if desc == "" {
		return ""
	}
	if i := strings.IndexByte(desc, '.'); i > 0 && i < 180 {
		return desc[:i+1]
	}
	if len(desc) > 180 {
		return desc[:177] + "..."
	}
	return desc
}

func resourceDescription(name, kind string, attrs map[string]*AttrJSON) string {
	verb := "Manages"
	noun := "resource"
	if kind == "data" {
		verb = "Reads"
		noun = "data source"
	}
	if len(attrs) == 0 {
		return fmt.Sprintf("%s %s is registered with an empty schema (deprecated or schema-less).", name, noun)
	}
	if a, ok := attrs["name"]; ok && a.Description != "" {
		return fmt.Sprintf("%s %s. %s", verb, name, firstSentence(a.Description))
	}
	return fmt.Sprintf("Schema validation for the %s %s.", name, noun)
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

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func mustAbs(p string) string {
	a, err := filepath.Abs(p)
	if err != nil {
		return p
	}
	return a
}
