# ibm_metrics_router_settings

Schema validation for the ibm_metrics_router_settings resource.

This rule validates `ibm_metrics_router_settings` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/settings:patch` (`metrics-router-metrics-router-v3`).

## Configuration

```hcl
rule "ibm_metrics_router_settings" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `backup_metadata_region` | no | pattern `^[a-zA-Z0-9 \-_]+$` |
| `permitted_target_regions` | no | pattern `^[a-zA-Z0-9 \-_]+$` |
| `primary_metadata_region` | no | pattern `^[a-zA-Z0-9 \-_]+$` |
| `private_api_endpoint_only` | no | API: If you set this true then you cannot access api through public network. |
| `default_targets` (block) | no | nested block |
| `default_targets.id` | yes | pattern `^[a-zA-Z0-9 \-._:]+$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_metrics_router_settings_invalid_backup_metadata_region` | `backup_metadata_region` |
| `ibm_metrics_router_settings_invalid_permitted_target_regions` | `permitted_target_regions` |
| `ibm_metrics_router_settings_invalid_primary_metadata_region` | `primary_metadata_region` |
| `ibm_metrics_router_settings_invalid_private_api_endpoint_only` | `private_api_endpoint_only` |
| `ibm_metrics_router_settings_invalid_default_targets__id` | `default_targets.id` |

## Example

### Valid

```hcl
resource "ibm_metrics_router_settings" "test" {
}
```
