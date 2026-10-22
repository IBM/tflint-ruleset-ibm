// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

// Package redact decides which attribute values may be echoed in lint
// messages. The runtime rules and the generator share it so that generated
// test expectations always match what the rules print.
package redact

import (
	"strconv"
	"strings"
)

// Placeholder replaces the value of a sensitive attribute in messages.
const Placeholder = "(sensitive value)"

// MaxShown is the longest value, in runes, echoed in a message.
const MaxShown = 64

var sensitiveParts = []string{
	"password",
	"passwd",
	"passphrase",
	"secret",
	"token",
	"apikey",
	"api_key",
	"private_key",
	"credential",
	"access_key",
	"auth_key",
	"hmac",
}

// IsSensitive reports whether an attribute value must not be echoed. flagged
// carries the provider's Sensitive schema flag; the name check also covers
// secret-bearing arguments the provider does not mark.
func IsSensitive(name string, flagged bool) bool {
	if flagged {
		return true
	}
	n := strings.ToLower(name)
	for _, p := range sensitiveParts {
		if strings.Contains(n, p) {
			return true
		}
	}
	return false
}

// Display returns the value as it may appear in a message: quoted, cut to
// MaxShown runes, or replaced by Placeholder when sensitive.
func Display(sensitive bool, val string) string {
	if sensitive {
		return Placeholder
	}
	if r := []rune(val); len(r) > MaxShown {
		return strconv.Quote(string(r[:MaxShown])) + "..."
	}
	return strconv.Quote(val)
}
