// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package rules

import "testing"

func FuzzValidCRN(f *testing.F) {
	f.Add("crn:v1:bluemix:public:is:us-south:a/account:vpc:identifier")
	f.Add("")
	f.Add("not-a-crn")
	f.Fuzz(func(t *testing.T, s string) {
		_ = validCRN(s)
	})
}

func FuzzValidZone(f *testing.F) {
	f.Add("us-south-1")
	f.Add("")
	f.Add("not-a-zone")
	f.Fuzz(func(t *testing.T, s string) {
		_ = validZone(s)
	})
}
