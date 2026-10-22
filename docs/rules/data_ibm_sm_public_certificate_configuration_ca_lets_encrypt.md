# ibm_sm_public_certificate_configuration_ca_lets_encrypt_data

Reads ibm_sm_public_certificate_configuration_ca_lets_encrypt. The name of the configuration.

This rule validates `ibm_sm_public_certificate_configuration_ca_lets_encrypt` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `SecretGroupPrototype` (`secrets-manager-secrets-manager-v2`).

## Configuration

```hcl
rule "ibm_sm_public_certificate_configuration_ca_lets_encrypt_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `created_at` | no | — |
| `created_by` | no | — |
| `lets_encrypt_environment` | no | — |
| `lets_encrypt_preferred_chain` | no | — |
| `lets_encrypt_private_key` | no | — |
| `name` | yes | pattern `^[A-Za-z0-9_][A-Za-z0-9_]*(?:_*-*\.*[A-Za-z0-9]*)*[A-Za-z0-9]+$` |
| `updated_at` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_sm_public_certificate_configuration_ca_lets_encrypt_data_invalid_name` | `name` |

## Example

### Valid

```hcl
data "ibm_sm_public_certificate_configuration_ca_lets_encrypt" "test" {
  name = "example-name"
}
```
