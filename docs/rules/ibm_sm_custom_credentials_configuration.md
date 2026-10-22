# ibm_sm_custom_credentials_configuration

Manages ibm_sm_custom_credentials_configuration. A human-readable unique name to assign to your configuration.

This rule validates `ibm_sm_custom_credentials_configuration` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `CustomCredentialsConfigurationCodeEngine` (`secrets-manager-secrets-manager-v2`).

## Configuration

```hcl
rule "ibm_sm_custom_credentials_configuration" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `api_key_ref` | no | pattern `^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$` |
| `code_engine_key_ref` | no | — |
| `created_at` | no | — |
| `created_by` | no | — |
| `name` | yes | — |
| `schema` | no | — |
| `task_timeout` | no | pattern `^[0-9]+[s,m,h,d]{0,1}$` |
| `updated_at` | no | — |
| `code_engine` (block) | yes | nested block |
| `code_engine.job_name` | yes | pattern `^[a-z0-9]([-a-z0-9]*[a-z0-9])?$` |
| `code_engine.project_id` | yes | pattern `^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$` |
| `code_engine.region` | yes | pattern `^[a-z0-9]([-a-z0-9]*[a-z0-9])?$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_sm_custom_credentials_configuration_invalid_api_key_ref` | `api_key_ref` |
| `ibm_sm_custom_credentials_configuration_invalid_name` | `name` |
| `ibm_sm_custom_credentials_configuration_invalid_task_timeout` | `task_timeout` |
| `ibm_sm_custom_credentials_configuration_invalid_code_engine__job_name` | `code_engine.job_name` |
| `ibm_sm_custom_credentials_configuration_invalid_code_engine__project_id` | `code_engine.project_id` |
| `ibm_sm_custom_credentials_configuration_invalid_code_engine__region` | `code_engine.region` |

## Example

### Valid

```hcl
resource "ibm_sm_custom_credentials_configuration" "test" {
  name = "example"
  code_engine {
    job_name = "example-name"
    project_id = "01234567-89ab-cdef-0123-456789abcdef"
    region = "example-name"
  }
}
```
