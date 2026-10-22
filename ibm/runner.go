// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import "github.com/terraform-linters/tflint-plugin-sdk/tflint"

// Runner wraps the SDK runner with IBM Cloud plugin state.
type Runner struct {
	tflint.Runner
	PluginConfig *Config
	Client       Client
}

// NewRunner returns a custom runner for the IBM Cloud ruleset.
func NewRunner(runner tflint.Runner, config *Config, client Client) (*Runner, error) {
	return &Runner{
		Runner:       runner,
		PluginConfig: config,
		Client:       client,
	}, nil
}
