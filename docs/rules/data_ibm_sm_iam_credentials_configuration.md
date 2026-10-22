# ibm_sm_iam_credentials_configuration_data

Reads ibm_sm_iam_credentials_configuration. The name of the configuration.

This rule validates `ibm_sm_iam_credentials_configuration` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `SecretGroupPrototype` (`secrets-manager-secrets-manager-v2`).

## Configuration

```hcl
rule "ibm_sm_iam_credentials_configuration_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `api_key` | no | — |
| `config_type` | no | — |
| `created_at` | no | — |
| `created_by` | no | — |
| `disabled` | no | — |
| `name` | yes | pattern `^[A-Za-z0-9_][A-Za-z0-9_]*(?:_*-*\.*[A-Za-z0-9]*)*[A-Za-z0-9]+$` |
| `secret_type` | no | — |
| `updated_at` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_sm_iam_credentials_configuration_data_invalid_name` | `name` |

## Example

### Valid

```hcl
data "ibm_sm_iam_credentials_configuration" "test" {
  name = "example-name"
}
```
