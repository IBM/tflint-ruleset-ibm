// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package main

// PackageJSON is the per-service extraction artifact consumed by the generator.
type PackageJSON struct {
	Package    string         `json:"package"`
	Resources  []ResourceJSON `json:"resources"`
	Skipped    []SkippedJSON  `json:"skipped"`
	Extracted  int            `json:"extracted"`
	SourceHint string         `json:"source_hint"`
}

// SkippedJSON records a resource the extractor could not fully handle.
type SkippedJSON struct {
	Name   string `json:"name"`
	Reason string `json:"reason"`
}

// ResourceJSON is the ground-truth schema for one Terraform resource.
type ResourceJSON struct {
	Name        string               `json:"name"`
	Function    string               `json:"function"`
	Package     string               `json:"package"`
	Kind        string               `json:"kind"`
	Description string               `json:"description"`
	Attributes  map[string]*AttrJSON `json:"attributes"`
	Validators  []ValidatorJSON      `json:"validators,omitempty"`
}

// AttrJSON describes one schema.Schema entry, including nested Elem structures.
type AttrJSON struct {
	Name        string        `json:"name"`
	Type        string        `json:"type"`
	Required    bool          `json:"required"`
	Optional    bool          `json:"optional"`
	Computed    bool          `json:"computed"`
	ForceNew    bool          `json:"force_new,omitempty"`
	Sensitive   bool          `json:"sensitive,omitempty"`
	MinItems    *int          `json:"min_items,omitempty"`
	MaxItems    *int          `json:"max_items,omitempty"`
	Description string        `json:"description,omitempty"`
	Default     string        `json:"default,omitempty"`
	Validate    *ValidateJSON `json:"validate,omitempty"`
	Elem        *ElemJSON     `json:"elem,omitempty"`
	Conflicts   []string      `json:"conflicts_with,omitempty"`
}

// ElemJSON is either a nested resource schema or a primitive element schema.
type ElemJSON struct {
	Kind       string               `json:"kind"` // "resource" or "schema"
	Type       string               `json:"type,omitempty"`
	Validate   *ValidateJSON        `json:"validate,omitempty"`
	Attributes map[string]*AttrJSON `json:"attributes,omitempty"`
}

// ValidateJSON captures constraints derived from ValidateFunc / ValidateSchema.
type ValidateJSON struct {
	Kind       string   `json:"kind"`
	FuncName   string   `json:"func_name,omitempty"`
	Enum       []string `json:"enum,omitempty"`
	IntEnum    []int    `json:"int_enum,omitempty"`
	Min        *int     `json:"min,omitempty"`
	Max        *int     `json:"max,omitempty"`
	MinLen     *int     `json:"min_len,omitempty"`
	MaxLen     *int     `json:"max_len,omitempty"`
	Regexp     string   `json:"regexp,omitempty"`
	IgnoreCase bool     `json:"ignore_case,omitempty"`
	Nullable   bool     `json:"nullable,omitempty"`
	NoZero     bool     `json:"no_zero,omitempty"`
}

// ValidatorJSON is a raw ValidateSchema entry (before merge onto attributes).
type ValidatorJSON struct {
	Identifier string `json:"identifier"`
	Kind       string `json:"kind"`
	Allowed    string `json:"allowed,omitempty"`
	Min        string `json:"min,omitempty"`
	Max        string `json:"max,omitempty"`
	Regexp     string `json:"regexp,omitempty"`
	MinLen     int    `json:"min_len,omitempty"`
	MaxLen     int    `json:"max_len,omitempty"`
}

// ResourceRef is a Terraform resource name mapped to its constructor.
type ResourceRef struct {
	Name     string
	Package  string
	Function string
	Kind     string
}
