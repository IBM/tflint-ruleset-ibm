# ibm_sm_private_certificate_configuration_template_data

Reads ibm_sm_private_certificate_configuration_template. The name of the configuration.

This rule validates `ibm_sm_private_certificate_configuration_template` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `SecretGroupPrototype` (`secrets-manager-secrets-manager-v2`).

## Configuration

```hcl
rule "ibm_sm_private_certificate_configuration_template_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `allow_any_name` | no | — |
| `allow_bare_domains` | no | — |
| `allow_glob_domains` | no | — |
| `allow_ip_sans` | no | — |
| `allow_localhost` | no | — |
| `allow_subdomains` | no | — |
| `allowed_domains` | no | — |
| `allowed_domains_template` | no | — |
| `allowed_other_sans` | no | — |
| `allowed_secret_groups` | no | — |
| `allowed_uri_sans` | no | — |
| `basic_constraints_valid_for_non_ca` | no | — |
| `certificate_authority` | no | — |
| `client_flag` | no | — |
| `code_signing_flag` | no | — |
| `config_type` | no | — |
| `country` | no | — |
| `created_at` | no | — |
| `created_by` | no | — |
| `email_protection_flag` | no | — |
| `enforce_hostnames` | no | — |
| `ext_key_usage` | no | — |
| `ext_key_usage_oids` | no | — |
| `key_bits` | no | — |
| `key_type` | no | — |
| `key_usage` | no | — |
| `locality` | no | — |
| `max_ttl_seconds` | no | — |
| `name` | yes | pattern `^[A-Za-z0-9_][A-Za-z0-9_]*(?:_*-*\.*[A-Za-z0-9]*)*[A-Za-z0-9]+$` |
| `not_before_duration_seconds` | no | — |
| `organization` | no | — |
| `ou` | no | — |
| `policy_identifiers` | no | — |
| `postal_code` | no | — |
| `province` | no | — |
| `require_cn` | no | — |
| `secret_type` | no | — |
| `serial_number` | no | — |
| `server_flag` | no | — |
| `street_address` | no | — |
| `ttl_seconds` | no | — |
| `updated_at` | no | — |
| `use_csr_common_name` | no | — |
| `use_csr_sans` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_sm_private_certificate_configuration_template_data_invalid_name` | `name` |

## Example

### Valid

```hcl
data "ibm_sm_private_certificate_configuration_template" "test" {
  name = "example-name"
}
```
