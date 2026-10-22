# ibm_cloudant

Manages ibm_cloudant. A name for the resource instance

This rule validates `ibm_cloudant` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_cloudant" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_id` | no | — |
| `allow_cleanup` | no | — |
| `capacity` | no | min 1 |
| `created_at` | no | — |
| `created_by` | no | — |
| `crn` | no | — |
| `dashboard_url` | no | — |
| `deleted_at` | no | — |
| `deleted_by` | no | — |
| `enable_cors` | no | — |
| `environment_crn` | no | crn format |
| `extensions` | no | — |
| `guid` | no | — |
| `include_data_events` | no | — |
| `last_operation` | no | — |
| `legacy_credentials` | no | — |
| `location` | yes | — |
| `locked` | no | — |
| `name` | yes | — |
| `onetime_credentials` | no | — |
| `parameters` | no | — |
| `parameters_json` | no | — |
| `plan` | yes | — |
| `plan_history` | no | — |
| `resource_aliases_url` | no | — |
| `resource_bindings_url` | no | — |
| `resource_controller_url` | no | — |
| `resource_crn` | no | — |
| `resource_group_crn` | no | — |
| `resource_group_id` | no | — |
| `resource_group_name` | no | — |
| `resource_id` | no | — |
| `resource_keys_url` | no | — |
| `resource_name` | no | — |
| `resource_plan_id` | no | — |
| `resource_status` | no | — |
| `restored_at` | no | — |
| `restored_by` | no | — |
| `scheduled_reclaim_at` | no | — |
| `scheduled_reclaim_by` | no | — |
| `service` | no | — |
| `service_endpoints` | no | one of: `public`, `private`, `public-and-private` |
| `state` | no | — |
| `status` | no | — |
| `sub_type` | no | — |
| `tags` | no | — |
| `target_crn` | no | — |
| `throughput` | no | — |
| `type` | no | — |
| `update_at` | no | — |
| `update_by` | no | — |
| `cors_config` (block) | no | nested block |
| `cors_config.allow_credentials` | no | — |
| `cors_config.origins` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cloudant_invalid_capacity` | `capacity` |
| `ibm_cloudant_invalid_enable_cors` | `enable_cors` |
| `ibm_cloudant_invalid_environment_crn` | `environment_crn` |
| `ibm_cloudant_invalid_include_data_events` | `include_data_events` |
| `ibm_cloudant_invalid_legacy_credentials` | `legacy_credentials` |
| `ibm_cloudant_invalid_location` | `location` |
| `ibm_cloudant_invalid_name` | `name` |
| `ibm_cloudant_invalid_parameters` | `parameters` |
| `ibm_cloudant_invalid_plan` | `plan` |
| `ibm_cloudant_invalid_resource_group_id` | `resource_group_id` |
| `ibm_cloudant_invalid_service_endpoints` | `service_endpoints` |
| `ibm_cloudant_invalid_tags` | `tags` |
| `ibm_cloudant_invalid_cors_config__allow_credentials` | `cors_config.allow_credentials` |
| `ibm_cloudant_invalid_cors_config__origins` | `cors_config.origins` |

## Example

### Valid

```hcl
resource "ibm_cloudant" "test" {
  location = "example"
  name = "example"
  plan = "example"
}
```
