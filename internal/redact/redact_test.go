// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package redact

import (
	"strings"
	"testing"
)

func TestIsSensitive(t *testing.T) {
	cases := map[string]bool{
		"api_key":                  true,
		"ibm_api_key":              true,
		"client_secret":            true,
		"lets_encrypt_private_key": true,
		"admin_password":           true,
		"google_credentials":       true,
		"name":                     false,
		"kms_key_crn":              false,
		"instance_id":              false,
	}
	for name, want := range cases {
		if got := IsSensitive(name, false); got != want {
			t.Errorf("IsSensitive(%q) = %v, want %v", name, got, want)
		}
	}
	if !IsSensitive("name", true) {
		t.Error("provider Sensitive flag must win")
	}
}

func TestDisplay(t *testing.T) {
	if got := Display(true, "s3cr3t"); got != Placeholder {
		t.Errorf("sensitive value shown as %q", got)
	}
	if got := Display(false, "bad"); got != `"bad"` {
		t.Errorf("Display = %s", got)
	}
	long := strings.Repeat("a", MaxShown+10)
	got := Display(false, long)
	if !strings.HasSuffix(got, "...") || strings.Count(got, "a") != MaxShown {
		t.Errorf("long value not truncated: %s", got)
	}
}
