// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

// Client is a placeholder API client used by deep-check rules.
// Schema-validation rules do not require live IBM Cloud credentials.
type Client interface {
	Region() string
}

type nopClient struct {
	region string
}

func (c nopClient) Region() string { return c.region }

// NewClient constructs an IBM Cloud API client from plugin config.
func NewClient(cfg *Config) Client {
	region := "us-south"
	if cfg != nil && cfg.Region != "" {
		region = cfg.Region
	}
	return nopClient{region: region}
}
