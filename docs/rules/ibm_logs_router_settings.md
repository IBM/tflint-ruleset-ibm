# ibm_logs_router_settings

Schema validation for the ibm_logs_router_settings resource.

This rule validates `ibm_logs_router_settings` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `leaf-unique` (`logs-router-service-api-logs-router-v1`).

## Configuration

```hcl
rule "ibm_logs_router_settings" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `api_version` | no | — |
| `backup_metadata_region` | no | pattern `^[a-zA-Z0-9 \-_]+$` |
| `permitted_target_regions` | no | — |
| `primary_metadata_region` | no | pattern `^[a-zA-Z0-9 \-_]+$` |
| `private_api_endpoint_only` | no | — |
| `default_targets` (block) | no | nested block |
| `default_targets.id` | yes | uuid format |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_logs_router_settings_invalid_backup_metadata_region` | `backup_metadata_region` |
| `ibm_logs_router_settings_invalid_permitted_target_regions` | `permitted_target_regions` |
| `ibm_logs_router_settings_invalid_primary_metadata_region` | `primary_metadata_region` |
| `ibm_logs_router_settings_invalid_private_api_endpoint_only` | `private_api_endpoint_only` |
| `ibm_logs_router_settings_invalid_default_targets__id` | `default_targets.id` |

## Example

### Valid

```hcl
resource "ibm_logs_router_settings" "test" {
}
```
