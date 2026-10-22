// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// RuleSet is the custom ruleset for the IBM Cloud provider plugin.
type RuleSet struct {
	tflint.BuiltinRuleSet
	config *Config
	client Client
}

// ConfigSchema returns the plugin configuration schema.
func (r *RuleSet) ConfigSchema() *hclext.BodySchema {
	r.config = &Config{}
	return hclext.ImpliedBodySchema(r.config)
}

// ApplyConfig reflects plugin configuration onto the ruleset.
func (r *RuleSet) ApplyConfig(body *hclext.BodyContent) error {
	diags := hclext.DecodeBody(body, nil, r.config)
	if diags.HasErrors() {
		return diags
	}
	r.client = NewClient(r.config)
	return nil
}

// NewRunner injects a custom IBM Cloud runner.
func (r *RuleSet) NewRunner(runner tflint.Runner) (tflint.Runner, error) {
	return NewRunner(runner, r.config, r.client)
}
