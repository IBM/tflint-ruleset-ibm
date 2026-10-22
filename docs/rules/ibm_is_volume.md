# ibm_is_volume

Manages ibm_is_volume. Volume name

This rule validates `ibm_is_volume` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Volume` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_volume" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_tags` | no | pattern `^([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-]):([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-])$` |
| `adjustable_capacity_states` | no | — |
| `adjustable_iops_states` | no | — |
| `bandwidth` | no | range 100–8192 |
| `capacity` | no | min 10 |
| `catalog_offering` | no | — |
| `crn` | no | — |
| `delete_all_snapshots` | no | — |
| `encryption_key` | no | API: The root key used to wrap the data encryption key for the volume. This property will be present for volumes with an `encryption` type of `user_managed`. |
| `encryption_type` | no | — |
| `health_reasons` | no | — |
| `health_state` | no | — |
| `iops` | no | min 100 |
| `name` | yes | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `operating_system` | no | — |
| `profile` | yes | one of: `general-purpose`, `5iops-tier`, `10iops-tier`, `custom`, `sdp` |
| `resource_controller_url` | no | — |
| `resource_crn` | no | — |
| `resource_group` | no | API: The resource group for this volume |
| `resource_group_name` | no | — |
| `resource_name` | no | — |
| `resource_status` | no | — |
| `source_snapshot` | no | pattern `^[-0-9a-z_]+$` |
| `source_snapshot_crn` | no | crn format |
| `status` | no | — |
| `status_reasons` | no | — |
| `storage_generation` | no | — |
| `tags` | no | pattern `^[A-Za-z0-9:_ .-]+$` |
| `zone` | yes | zone format |
| `allowed_use` (block) | no | nested block |
| `allowed_use.api_version` | no | pattern `^\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12]\d|3[01])$` |
| `allowed_use.bare_metal_server` | no | pattern `^([a-zA-Z_][a-zA-Z0-9_]*|[-+*/%]|&&|\|\||!|==|!=|<|<=|>|>=|~|\bin\b|\(|\)|\[|\]|,|\.|"|'|"|'|\s+|\d+)+$` |
| `allowed_use.instance` | no | pattern `^([a-zA-Z_][a-zA-Z0-9_]*|[-+*/%]|&&|\|\||!|==|!=|<|<=|>|>=|~|\bin\b|\(|\)|\[|\]|,|\.|"|'|"|'|\s+|\d+)+$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_volume_invalid_access_tags` | `access_tags` |
| `ibm_is_volume_invalid_bandwidth` | `bandwidth` |
| `ibm_is_volume_invalid_capacity` | `capacity` |
| `ibm_is_volume_invalid_delete_all_snapshots` | `delete_all_snapshots` |
| `ibm_is_volume_invalid_encryption_key` | `encryption_key` |
| `ibm_is_volume_invalid_iops` | `iops` |
| `ibm_is_volume_invalid_name` | `name` |
| `ibm_is_volume_invalid_profile` | `profile` |
| `ibm_is_volume_invalid_resource_group` | `resource_group` |
| `ibm_is_volume_invalid_source_snapshot` | `source_snapshot` |
| `ibm_is_volume_invalid_source_snapshot_crn` | `source_snapshot_crn` |
| `ibm_is_volume_invalid_tags` | `tags` |
| `ibm_is_volume_invalid_zone` | `zone` |
| `ibm_is_volume_invalid_allowed_use__api_version` | `allowed_use.api_version` |
| `ibm_is_volume_invalid_allowed_use__bare_metal_server` | `allowed_use.bare_metal_server` |
| `ibm_is_volume_invalid_allowed_use__instance` | `allowed_use.instance` |

## Example

### Valid

```hcl
resource "ibm_is_volume" "test" {
  name = "example-name"
  profile = "general-purpose"
  zone = "us-south-1"
}
```
