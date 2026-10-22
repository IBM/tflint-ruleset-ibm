# ibm_is_share_data

Reads ibm_is_share. Name of the share.

This rule validates `ibm_is_share` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `BareMetalServerDiskPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_share_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_control_mode` | no | — |
| `access_tags` | no | — |
| `accessor_binding_role` | no | — |
| `accessor_bindings` | no | — |
| `allowed_access_protocols` | no | — |
| `allowed_transit_encryption_modes` | no | — |
| `availability_mode` | no | — |
| `bandwidth` | no | — |
| `created_at` | no | — |
| `crn` | no | — |
| `encryption` | no | — |
| `encryption_key` | no | — |
| `href` | no | — |
| `iops` | no | — |
| `latest_job` | no | — |
| `latest_sync` | no | — |
| `lifecycle_state` | no | — |
| `mount_targets` | no | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `origin_share` | no | — |
| `profile` | no | — |
| `replica_share` | no | — |
| `replication_cron_spec` | no | — |
| `replication_role` | no | — |
| `replication_status` | no | — |
| `replication_status_reasons` | no | — |
| `resource_group` | no | — |
| `resource_type` | no | — |
| `share` | no | — |
| `share_targets` | no | — |
| `size` | no | — |
| `snapshot_count` | no | — |
| `snapshot_size` | no | — |
| `source_share` | no | — |
| `source_snapshot` | no | — |
| `storage_generation` | no | — |
| `tags` | no | — |
| `zone` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_share_data_invalid_name` | `name` |
| `ibm_is_share_data_invalid_share` | `share` |

## Example

### Valid

```hcl
data "ibm_is_share" "test" {
}
```
