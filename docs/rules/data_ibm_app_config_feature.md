# ibm_app_config_feature_data

Reads ibm_app_config_feature. Feature name.

This rule validates `ibm_app_config_feature` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Feature` (`app-configuration`).

## Configuration

```hcl
rule "ibm_app_config_feature_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `collections` | no | — |
| `created_time` | no | — |
| `description` | no | — |
| `disabled_value` | no | — |
| `enabled` | no | — |
| `enabled_value` | no | — |
| `environment_id` | yes | API: Environment identifier. Only present for FEATURE, PROPERTY and ENVIRONMENT. Not present for SEGMENT and COLLECTION. |
| `feature_id` | yes | API: Feature id. Allowed special characters are dot ( . ), hyphen( - ), underscore ( _ ) only |
| `format` | no | — |
| `guid` | yes | — |
| `href` | no | — |
| `includes` | no | — |
| `name` | no | — |
| `rollout_percentage` | no | — |
| `segment_exists` | no | — |
| `segment_rules` | no | — |
| `tags` | no | — |
| `type` | no | — |
| `updated_time` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_app_config_feature_data_invalid_environment_id` | `environment_id` |
| `ibm_app_config_feature_data_invalid_feature_id` | `feature_id` |
| `ibm_app_config_feature_data_invalid_guid` | `guid` |
| `ibm_app_config_feature_data_invalid_includes` | `includes` |

## Example

### Valid

```hcl
data "ibm_app_config_feature" "test" {
  environment_id = "01234567-89ab-cdef-0123-456789abcdef"
  feature_id = "01234567-89ab-cdef-0123-456789abcdef"
  guid = "example"
}
```
