# ibm_app_config_integration_en_data

Schema validation for the ibm_app_config_integration_en data source.

This rule validates `ibm_app_config_integration_en` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Integration` (`app-configuration`).

## Configuration

```hcl
rule "ibm_app_config_integration_en_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `created_time` | no | — |
| `en_endpoint` | no | — |
| `en_instance_crn` | no | — |
| `en_source_id` | no | — |
| `guid` | yes | — |
| `href` | no | — |
| `integration_id` | yes | API: Integration id. |
| `integration_type` | no | — |
| `updated_time` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_app_config_integration_en_data_invalid_guid` | `guid` |
| `ibm_app_config_integration_en_data_invalid_integration_id` | `integration_id` |

## Example

### Valid

```hcl
data "ibm_app_config_integration_en" "test" {
  guid = "example"
  integration_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
