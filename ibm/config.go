// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

// Config is plugin configuration for the IBM Cloud ruleset.
type Config struct {
	DeepCheck bool   `hclext:"deep_check,optional"`
	Region    string `hclext:"region,optional"`
}
