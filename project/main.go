// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package project

import "fmt"

// Version is the ruleset version. Release builds set it from the git tag
// with -ldflags "-X github.com/IBM/tflint-ruleset-ibm/project.Version=...".
var Version = "0.1.0"

// ReferenceLink returns the documentation URL for a generated rule.
func ReferenceLink(name string) string {
	return fmt.Sprintf("https://github.com/IBM/tflint-ruleset-ibm/blob/v%s/docs/rules/%s.md", Version, name)
}
