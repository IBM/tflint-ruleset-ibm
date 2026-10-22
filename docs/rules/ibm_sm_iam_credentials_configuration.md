# ibm_sm_iam_credentials_configuration

Manages ibm_sm_iam_credentials_configuration. A human-readable unique name to assign to your configuration.

This rule validates `ibm_sm_iam_credentials_configuration` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `IAMCredentialsConfigurationPrototype` (`secrets-manager-secrets-manager-v2`).

## Configuration

```hcl
rule "ibm_sm_iam_credentials_configuration" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `api_key` | yes | pattern `^(?:[A-Za-z0-9_\-]{4})*(?:[A-Za-z0-9_\-]{2}==|[A-Za-z0-9_\-]{3}=)?$` |
| `config_type` | no | — |
| `created_at` | no | — |
| `created_by` | no | — |
| `disabled` | no | API: This parameter indicates whether the API key configuration is disabled. If it is set to `true`, IAM credentials don't use the configured API key for credentials management. |
| `name` | yes | pattern `^[A-Za-z0-9_][A-Za-z0-9_]*(?:_*-*\.*[A-Za-z0-9]*)*[A-Za-z0-9]+$` |
| `secret_type` | no | — |
| `updated_at` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_sm_iam_credentials_configuration_invalid_api_key` | `api_key` |
| `ibm_sm_iam_credentials_configuration_invalid_disabled` | `disabled` |
| `ibm_sm_iam_credentials_configuration_invalid_name` | `name` |

## Example

### Valid

```hcl
variable "api_key" {
  sensitive = true
}

resource "ibm_sm_iam_credentials_configuration" "test" {
  api_key = var.api_key
  name = "example-name"
}
```
