// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package rules

import (
	"testing"

	"github.com/terraform-linters/tflint-plugin-sdk/helper"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

func TestRulesRegistered(t *testing.T) {
	if len(Rules) == 0 {
		t.Fatal("no rules registered")
	}
	seen := map[string]bool{}
	for _, rule := range Rules {
		name := rule.Name()
		if name == "" {
			t.Fatal("registered rule with empty name")
		}
		if seen[name] {
			t.Errorf("duplicate rule name %q", name)
		}
		seen[name] = true
		if !rule.Enabled() {
			t.Errorf("%s: expected enabled by default", name)
		}
		if rule.Severity() != tflint.ERROR {
			t.Errorf("%s: expected ERROR severity", name)
		}
		if rule.Link() == "" {
			t.Errorf("%s: empty documentation link", name)
		}
		runner := helper.TestRunner(t, map[string]string{"resource.tf": ""})
		if err := rule.Check(runner); err != nil {
			t.Errorf("%s: Check on empty module: %v", name, err)
		}
	}
	if len(seen) != len(Rules) {
		t.Fatalf("name set %d != rules %d", len(seen), len(Rules))
	}
}
