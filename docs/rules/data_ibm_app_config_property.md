# ibm_app_config_property_data

Reads ibm_app_config_property. Property name.

This rule validates `ibm_app_config_property` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Property` (`app-configuration`).

## Configuration

```hcl
rule "ibm_app_config_property_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `collections` | no | — |
| `created_time` | no | — |
| `description` | no | — |
| `environment_id` | yes | API: Environment identifier. Only present for FEATURE, PROPERTY and ENVIRONMENT. Not present for SEGMENT and COLLECTION. |
| `format` | no | — |
| `guid` | yes | — |
| `href` | no | — |
| `include` | no | — |
| `name` | no | — |
| `property_id` | yes | API: Property id. Allowed special characters are dot ( . ), hyphen( - ), underscore ( _ ) only |
| `segment_exists` | no | — |
| `segment_rules` | no | — |
| `tags` | no | — |
| `type` | no | — |
| `updated_time` | no | — |
| `value` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_app_config_property_data_invalid_environment_id` | `environment_id` |
| `ibm_app_config_property_data_invalid_guid` | `guid` |
| `ibm_app_config_property_data_invalid_include` | `include` |
| `ibm_app_config_property_data_invalid_property_id` | `property_id` |

## Example

### Valid

```hcl
data "ibm_app_config_property" "test" {
  environment_id = "01234567-89ab-cdef-0123-456789abcdef"
  guid = "example"
  property_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
