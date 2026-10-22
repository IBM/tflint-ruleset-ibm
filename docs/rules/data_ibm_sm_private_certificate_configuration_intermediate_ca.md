# ibm_sm_private_certificate_configuration_intermediate_ca_data

Reads ibm_sm_private_certificate_configuration_intermediate_ca. The name of the configuration.

This rule validates `ibm_sm_private_certificate_configuration_intermediate_ca` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `SecretGroupPrototype` (`secrets-manager-secrets-manager-v2`).

## Configuration

```hcl
rule "ibm_sm_private_certificate_configuration_intermediate_ca_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `alt_names` | no | — |
| `common_name` | no | — |
| `config_type` | no | — |
| `country` | no | — |
| `crl_disable` | no | — |
| `crl_distribution_points_encoded` | no | — |
| `crl_expiry_seconds` | no | — |
| `crypto_key` | no | — |
| `data` | no | — |
| `exclude_cn_from_sans` | no | — |
| `expiration_date` | no | — |
| `format` | no | — |
| `ip_sans` | no | — |
| `issuer` | no | — |
| `issuing_certificates_urls_encoded` | no | — |
| `key_bits` | no | — |
| `key_type` | no | — |
| `locality` | no | — |
| `max_ttl_seconds` | no | — |
| `name` | yes | pattern `^[A-Za-z0-9_][A-Za-z0-9_]*(?:_*-*\.*[A-Za-z0-9]*)*[A-Za-z0-9]+$` |
| `organization` | no | — |
| `other_sans` | no | — |
| `ou` | no | — |
| `postal_code` | no | — |
| `private_key_format` | no | — |
| `province` | no | — |
| `secret_type` | no | — |
| `serial_number` | no | — |
| `signing_method` | no | — |
| `status` | no | — |
| `street_address` | no | — |
| `uri_sans` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_sm_private_certificate_configuration_intermediate_ca_data_invalid_name` | `name` |

## Example

### Valid

```hcl
data "ibm_sm_private_certificate_configuration_intermediate_ca" "test" {
  name = "example-name"
}
```
