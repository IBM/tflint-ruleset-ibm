# ibm_app_config_environment_data

Reads ibm_app_config_environment. Environment name.

This rule validates `ibm_app_config_environment` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Environment` (`app-configuration`).

## Configuration

```hcl
rule "ibm_app_config_environment_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `color_code` | no | — |
| `created_time` | no | — |
| `description` | no | — |
| `environment_id` | yes | API: Environment id. Allowed special characters are dot ( . ), hyphen( - ), underscore ( _ ) only |
| `expand` | no | — |
| `guid` | yes | — |
| `href` | no | — |
| `name` | no | — |
| `tags` | no | — |
| `updated_time` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_app_config_environment_data_invalid_environment_id` | `environment_id` |
| `ibm_app_config_environment_data_invalid_expand` | `expand` |
| `ibm_app_config_environment_data_invalid_guid` | `guid` |

## Example

### Valid

```hcl
data "ibm_app_config_environment" "test" {
  environment_id = "01234567-89ab-cdef-0123-456789abcdef"
  guid = "example"
}
```
