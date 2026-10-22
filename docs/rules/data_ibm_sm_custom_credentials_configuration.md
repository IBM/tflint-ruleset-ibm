# ibm_sm_custom_credentials_configuration_data

Reads ibm_sm_custom_credentials_configuration. The name of the configuration.

This rule validates `ibm_sm_custom_credentials_configuration` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `SecretGroupPrototype` (`secrets-manager-secrets-manager-v2`).

## Configuration

```hcl
rule "ibm_sm_custom_credentials_configuration_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `api_key_ref` | no | — |
| `code_engine` | no | — |
| `code_engine_key_ref` | no | — |
| `created_at` | no | — |
| `created_by` | no | — |
| `name` | yes | pattern `^[A-Za-z0-9_][A-Za-z0-9_]*(?:_*-*\.*[A-Za-z0-9]*)*[A-Za-z0-9]+$` |
| `schema` | no | — |
| `task_timeout` | no | — |
| `updated_at` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_sm_custom_credentials_configuration_data_invalid_name` | `name` |

## Example

### Valid

```hcl
data "ibm_sm_custom_credentials_configuration" "test" {
  name = "example-name"
}
```
