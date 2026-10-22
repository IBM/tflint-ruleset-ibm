// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package main

import (
	"github.com/IBM/tflint-ruleset-ibm/ibm"
	"github.com/IBM/tflint-ruleset-ibm/project"
	"github.com/IBM/tflint-ruleset-ibm/rules"
	"github.com/terraform-linters/tflint-plugin-sdk/plugin"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		RuleSet: &ibm.RuleSet{
			BuiltinRuleSet: tflint.BuiltinRuleSet{
				Name:    "ibm",
				Version: project.Version,
				Rules:   rules.Rules,
			},
		},
	})
}
