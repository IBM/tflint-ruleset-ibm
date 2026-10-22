// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

type overlayFile struct {
	Package   string            `json:"package"`
	Resources []overlayResource `json:"resources"`
}

type overlayResource struct {
	Name   string                 `json:"name"`
	Kind   string                 `json:"kind"`
	Schema string                 `json:"schema"`
	Spec   string                 `json:"spec"`
	Score  float64                `json:"score"`
	Attrs  map[string]overlayAttr `json:"attrs"`
}

type overlayAttr struct {
	Enum        []string `json:"enum"`
	Pattern     string   `json:"pattern"`
	Format      string   `json:"format"`
	Kind        string   `json:"kind"`
	MinLen      *int     `json:"min_len"`
	MaxLen      *int     `json:"max_len"`
	Min         *int     `json:"min"`
	Max         *int     `json:"max"`
	Description string   `json:"description"`
}

type overlayKey struct {
	Name string
	Kind string
}

func loadOverlays(inDir string) map[overlayKey]overlayResource {
	out := map[overlayKey]overlayResource{}
	entries, err := os.ReadDir(inDir)
	if err != nil {
		return out
	}
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".overlay.json") {
			continue
		}
		raw, err := os.ReadFile(filepath.Join(inDir, e.Name()))
		if err != nil {
			continue
		}
		var f overlayFile
		if json.Unmarshal(raw, &f) != nil {
			continue
		}
		for _, r := range f.Resources {
			kind := r.Kind
			if kind == "" {
				kind = "resource"
			}
			out[overlayKey{Name: r.Name, Kind: kind}] = r
		}
	}
	return out
}

func applyOverlay(spec *resourceSpec, ov overlayResource) {
	if spec == nil || len(ov.Attrs) == 0 {
		return
	}
	spec.APISchema = ov.Schema
	spec.APISpec = ov.Spec
	for path, oa := range ov.Attrs {
		parts := strings.Split(path, ".")
		if len(parts) == 0 {
			continue
		}
		applyOverlayParts(&spec.Attrs, &spec.Blocks, parts, oa)
	}
}

func applyOverlayParts(attrs *[]specAttr, blocks *[]specBlock, parts []string, oa overlayAttr) {
	if len(parts) == 1 {
		for i := range *attrs {
			if (*attrs)[i].Name == parts[0] {
				mergeOverlayAttr(&(*attrs)[i], oa)
				return
			}
		}
		return
	}
	for i := range *blocks {
		if (*blocks)[i].Type == parts[0] {
			applyOverlayParts(&(*blocks)[i].Attrs, &(*blocks)[i].Blocks, parts[1:], oa)
			return
		}
	}
}

func mergeOverlayAttr(a *specAttr, oa overlayAttr) {
	if len(a.Enum) == 0 && len(oa.Enum) > 0 {
		a.Enum = append([]string{}, oa.Enum...)
		if a.Kind == "" {
			a.Kind = "enum"
		}
	}
	if a.Regexp == "" && oa.Pattern != "" {
		a.Regexp = oa.Pattern
		if a.Kind == "" || a.Kind == "enum" && len(a.Enum) == 0 {
			a.Kind = "regexp"
		}
	}
	if a.MinLen == nil && oa.MinLen != nil {
		a.MinLen = oa.MinLen
	}
	if a.MaxLen == nil && oa.MaxLen != nil {
		a.MaxLen = oa.MaxLen
	}
	if a.Min == nil && oa.Min != nil {
		a.Min = oa.Min
	}
	if a.Max == nil && oa.Max != nil {
		a.Max = oa.Max
	}
	if a.Kind == "" && oa.Kind != "" {
		a.Kind = oa.Kind
	}
	if a.APIDescription == "" && oa.Description != "" {
		a.APIDescription = oa.Description
	}
	normalizeConstraints(a)
}
