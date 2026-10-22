# ibm_config_aggregator_settings

Schema validation for the ibm_config_aggregator_settings resource.

This rule validates `ibm_config_aggregator_settings` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/apprapp/config_aggregator/v1/instances/{guid}/settings:put` (`app-configuration`).

## Configuration

```hcl
rule "ibm_config_aggregator_settings" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `resource_collection_enabled` | yes | API: The field denoting if the resource collection is enabled |
| `resource_collection_regions` | yes | — |
| `trusted_profile_id` | yes | pattern `^[a-zA-Z0-9-]*$` |
| `additional_scope` (block) | no | nested block |
| `additional_scope.enterprise_id` | no | pattern `^[a-zA-Z0-9]+$` |
| `additional_scope.type` | no | pattern `^[a-zA-Z0-9]+$` |
| `additional_scope.profile_template` (block) | no | nested block |
| `additional_scope.profile_template.id` | no | pattern `^[a-zA-Z0-9-=:]+$` |
| `additional_scope.profile_template.trusted_profile_id` | no | pattern `^[a-zA-Z0-9-=:]+$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_config_aggregator_settings_invalid_resource_collection_enabled` | `resource_collection_enabled` |
| `ibm_config_aggregator_settings_invalid_resource_collection_regions` | `resource_collection_regions` |
| `ibm_config_aggregator_settings_invalid_trusted_profile_id` | `trusted_profile_id` |
| `ibm_config_aggregator_settings_invalid_additional_scope__enterprise_id` | `additional_scope.enterprise_id` |
| `ibm_config_aggregator_settings_invalid_additional_scope__type` | `additional_scope.type` |
| `ibm_config_aggregator_settings_invalid_additional_scope__profile_template__id` | `additional_scope.profile_template.id` |
| `ibm_config_aggregator_settings_invalid_additional_scope__profile_template__trusted_profile_id` | `additional_scope.profile_template.trusted_profile_id` |

## Example

### Valid

```hcl
resource "ibm_config_aggregator_settings" "test" {
  resource_collection_enabled = true
  resource_collection_regions = "example"
  trusted_profile_id = "example-namexxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
}
```
