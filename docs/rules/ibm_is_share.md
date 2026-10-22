# ibm_is_share

Manages ibm_is_share. The unique user-defined name for this file share.

This rule validates `ibm_is_share` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Share` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_share" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_control_mode` | no | one of: `security_group`, `vpc`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `access_tags` | no | pattern `^([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-]):([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-])$` |
| `accessor_binding_role` | no | — |
| `accessor_bindings` | no | — |
| `allowed_access_protocols` | no | one of: `nfs4`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `allowed_transit_encryption_modes` | no | one of: `ipsec`, `none`, `stunnel`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `availability_mode` | no | — |
| `bandwidth` | no | min 1 |
| `created_at` | no | — |
| `crn` | no | — |
| `encryption` | no | — |
| `encryption_key` | no | API: The key used to encrypt this file share. This property will be present if `encryption` is `user_managed`. |
| `href` | no | — |
| `iops` | no | range 100–96000 |
| `last_sync_at` | no | — |
| `latest_job` | no | — |
| `latest_sync` | no | — |
| `lifecycle_reasons` | no | — |
| `lifecycle_state` | no | — |
| `name` | yes | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `profile` | no | API: The [profile](https://cloud.ibm.com/docs/vpc?topic=vpc-file-storage-profiles) for this file share |
| `replication_cron_spec` | no | pattern `^(((\d+,)+\d+|([\d\*]+(\/|-)\d+)|\d+|\*) ?){5}$` |
| `replication_role` | no | — |
| `replication_status` | no | — |
| `replication_status_reasons` | no | — |
| `resource_group` | no | API: The resource group for this file share |
| `resource_type` | no | — |
| `size` | no | range 1–32000 |
| `snapshot_count` | no | — |
| `snapshot_size` | no | — |
| `source_share` | no | API: The source file share for this replica file share. This property will be present when the `replication_role` is `replica`. |
| `source_share_crn` | no | crn format |
| `storage_generation` | no | — |
| `tags` | no | pattern `^[A-Za-z0-9:_ .-]+$` |
| `zone` | no | zone format |
| `initial_owner` (block) | no | nested block |
| `initial_owner.gid` | no | range 0–2147483647 |
| `initial_owner.uid` | no | range 0–2147483647 |
| `mount_targets` (block) | no | nested block |
| `mount_targets.access_protocol` | no | one of: `nfs4`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `mount_targets.name` | yes | pattern `^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `mount_targets.transit_encryption` | no | one of: `ipsec`, `none`, `stunnel`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `mount_targets.vpc` | no | — |
| `mount_targets.virtual_network_interface` (block) | no | nested block |
| `mount_targets.virtual_network_interface.auto_delete` | no | — |
| `mount_targets.virtual_network_interface.id` | no | pattern `^[-0-9a-z_]+$` |
| `mount_targets.virtual_network_interface.name` | no | pattern `^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `mount_targets.virtual_network_interface.protocol_state_filtering_mode` | no | — |
| `mount_targets.virtual_network_interface.resource_group` | no | API: The resource group for this file share |
| `mount_targets.virtual_network_interface.security_groups` | no | — |
| `mount_targets.virtual_network_interface.subnet` | no | — |
| `mount_targets.virtual_network_interface.primary_ip` (block) | no | nested block |
| `mount_targets.virtual_network_interface.primary_ip.address` | no | — |
| `mount_targets.virtual_network_interface.primary_ip.auto_delete` | no | — |
| `mount_targets.virtual_network_interface.primary_ip.name` | no | pattern `^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `mount_targets.virtual_network_interface.primary_ip.reserved_ip` | no | — |
| `origin_share` (block) | no | nested block |
| `origin_share.crn` | no | crn format |
| `origin_share.id` | no | pattern `^[-0-9a-z_]+$` |
| `replica_share` (block) | no | nested block |
| `replica_share.access_tags` | no | — |
| `replica_share.iops` | no | range 100–9007199254740991 |
| `replica_share.name` | yes | pattern `^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `replica_share.profile` | yes | API: The [profile](https://cloud.ibm.com/docs/vpc?topic=vpc-file-storage-profiles) for this file share |
| `replica_share.replication_cron_spec` | yes | pattern `^(((\d+,)+\d+|([\d\*]+(\/|-)\d+)|\d+|\*) ?){5}$` |
| `replica_share.tags` | no | — |
| `replica_share.zone` | no | zone format |
| `replica_share.mount_targets` (block) | no | nested block |
| `replica_share.mount_targets.access_protocol` | no | one of: `nfs4`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `replica_share.mount_targets.name` | no | pattern `^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `replica_share.mount_targets.transit_encryption` | no | one of: `ipsec`, `none`, `stunnel`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `replica_share.mount_targets.vpc` | no | — |
| `replica_share.mount_targets.virtual_network_interface` (block) | no | nested block |
| `replica_share.mount_targets.virtual_network_interface.auto_delete` | no | — |
| `replica_share.mount_targets.virtual_network_interface.id` | no | pattern `^[-0-9a-z_]+$` |
| `replica_share.mount_targets.virtual_network_interface.name` | no | pattern `^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `replica_share.mount_targets.virtual_network_interface.protocol_state_filtering_mode` | no | — |
| `replica_share.mount_targets.virtual_network_interface.resource_group` | no | API: The resource group for this file share |
| `replica_share.mount_targets.virtual_network_interface.security_groups` | no | — |
| `replica_share.mount_targets.virtual_network_interface.subnet` | no | — |
| `replica_share.mount_targets.virtual_network_interface.primary_ip` (block) | no | nested block |
| `replica_share.mount_targets.virtual_network_interface.primary_ip.address` | no | — |
| `replica_share.mount_targets.virtual_network_interface.primary_ip.auto_delete` | no | — |
| `replica_share.mount_targets.virtual_network_interface.primary_ip.name` | no | pattern `^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `replica_share.mount_targets.virtual_network_interface.primary_ip.reserved_ip` | no | — |
| `source_snapshot` (block) | no | nested block |
| `source_snapshot.crn` | no | crn format |
| `source_snapshot.id` | no | pattern `^[-0-9a-z_]+$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_share_invalid_access_control_mode` | `access_control_mode` |
| `ibm_is_share_invalid_access_tags` | `access_tags` |
| `ibm_is_share_invalid_allowed_access_protocols` | `allowed_access_protocols` |
| `ibm_is_share_invalid_allowed_transit_encryption_modes` | `allowed_transit_encryption_modes` |
| `ibm_is_share_invalid_bandwidth` | `bandwidth` |
| `ibm_is_share_invalid_encryption_key` | `encryption_key` |
| `ibm_is_share_invalid_iops` | `iops` |
| `ibm_is_share_invalid_name` | `name` |
| `ibm_is_share_invalid_profile` | `profile` |
| `ibm_is_share_invalid_replication_cron_spec` | `replication_cron_spec` |
| `ibm_is_share_invalid_resource_group` | `resource_group` |
| `ibm_is_share_invalid_size` | `size` |
| `ibm_is_share_invalid_source_share` | `source_share` |
| `ibm_is_share_invalid_source_share_crn` | `source_share_crn` |
| `ibm_is_share_invalid_tags` | `tags` |
| `ibm_is_share_invalid_zone` | `zone` |
| `ibm_is_share_invalid_initial_owner__gid` | `initial_owner.gid` |
| `ibm_is_share_invalid_initial_owner__uid` | `initial_owner.uid` |
| `ibm_is_share_invalid_mount_targets__access_protocol` | `mount_targets.access_protocol` |
| `ibm_is_share_invalid_mount_targets__name` | `mount_targets.name` |
| `ibm_is_share_invalid_mount_targets__transit_encryption` | `mount_targets.transit_encryption` |
| `ibm_is_share_invalid_mount_targets__vpc` | `mount_targets.vpc` |
| `ibm_is_share_invalid_mount_targets__virtual_network_interface__auto_delete` | `mount_targets.virtual_network_interface.auto_delete` |
| `ibm_is_share_invalid_mount_targets__virtual_network_interface__id` | `mount_targets.virtual_network_interface.id` |
| `ibm_is_share_invalid_mount_targets__virtual_network_interface__name` | `mount_targets.virtual_network_interface.name` |
| `ibm_is_share_invalid_mount_targets__virtual_network_interface__protocol_state_filtering_mode` | `mount_targets.virtual_network_interface.protocol_state_filtering_mode` |
| `ibm_is_share_invalid_mount_targets__virtual_network_interface__resource_group` | `mount_targets.virtual_network_interface.resource_group` |
| `ibm_is_share_invalid_mount_targets__virtual_network_interface__security_groups` | `mount_targets.virtual_network_interface.security_groups` |
| `ibm_is_share_invalid_mount_targets__virtual_network_interface__subnet` | `mount_targets.virtual_network_interface.subnet` |
| `ibm_is_share_invalid_mount_targets__virtual_network_interface__primary_ip__address` | `mount_targets.virtual_network_interface.primary_ip.address` |
| `ibm_is_share_invalid_mount_targets__virtual_network_interface__primary_ip__auto_delete` | `mount_targets.virtual_network_interface.primary_ip.auto_delete` |
| `ibm_is_share_invalid_mount_targets__virtual_network_interface__primary_ip__name` | `mount_targets.virtual_network_interface.primary_ip.name` |
| `ibm_is_share_invalid_mount_targets__virtual_network_interface__primary_ip__reserved_ip` | `mount_targets.virtual_network_interface.primary_ip.reserved_ip` |
| `ibm_is_share_invalid_origin_share__crn` | `origin_share.crn` |
| `ibm_is_share_invalid_origin_share__id` | `origin_share.id` |
| `ibm_is_share_invalid_replica_share__access_tags` | `replica_share.access_tags` |
| `ibm_is_share_invalid_replica_share__iops` | `replica_share.iops` |
| `ibm_is_share_invalid_replica_share__name` | `replica_share.name` |
| `ibm_is_share_invalid_replica_share__profile` | `replica_share.profile` |
| `ibm_is_share_invalid_replica_share__replication_cron_spec` | `replica_share.replication_cron_spec` |
| `ibm_is_share_invalid_replica_share__tags` | `replica_share.tags` |
| `ibm_is_share_invalid_replica_share__zone` | `replica_share.zone` |
| `ibm_is_share_invalid_replica_share__mount_targets__access_protocol` | `replica_share.mount_targets.access_protocol` |
| `ibm_is_share_invalid_replica_share__mount_targets__name` | `replica_share.mount_targets.name` |
| `ibm_is_share_invalid_replica_share__mount_targets__transit_encryption` | `replica_share.mount_targets.transit_encryption` |
| `ibm_is_share_invalid_replica_share__mount_targets__vpc` | `replica_share.mount_targets.vpc` |
| `ibm_is_share_invalid_replica_share__mount_targets__virtual_network_interface__auto_delete` | `replica_share.mount_targets.virtual_network_interface.auto_delete` |
| `ibm_is_share_invalid_replica_share__mount_targets__virtual_network_interface__id` | `replica_share.mount_targets.virtual_network_interface.id` |
| `ibm_is_share_invalid_replica_share__mount_targets__virtual_network_interface__name` | `replica_share.mount_targets.virtual_network_interface.name` |
| `ibm_is_share_invalid_replica_share__mount_targets__virtual_network_interface__protocol_state_filtering_mode` | `replica_share.mount_targets.virtual_network_interface.protocol_state_filtering_mode` |
| `ibm_is_share_invalid_replica_share__mount_targets__virtual_network_interface__resource_group` | `replica_share.mount_targets.virtual_network_interface.resource_group` |
| `ibm_is_share_invalid_replica_share__mount_targets__virtual_network_interface__security_groups` | `replica_share.mount_targets.virtual_network_interface.security_groups` |
| `ibm_is_share_invalid_replica_share__mount_targets__virtual_network_interface__subnet` | `replica_share.mount_targets.virtual_network_interface.subnet` |
| `ibm_is_share_invalid_replica_share__mount_targets__virtual_network_interface__primary_ip__address` | `replica_share.mount_targets.virtual_network_interface.primary_ip.address` |
| `ibm_is_share_invalid_replica_share__mount_targets__virtual_network_interface__primary_ip__auto_delete` | `replica_share.mount_targets.virtual_network_interface.primary_ip.auto_delete` |
| `ibm_is_share_invalid_replica_share__mount_targets__virtual_network_interface__primary_ip__name` | `replica_share.mount_targets.virtual_network_interface.primary_ip.name` |
| `ibm_is_share_invalid_replica_share__mount_targets__virtual_network_interface__primary_ip__reserved_ip` | `replica_share.mount_targets.virtual_network_interface.primary_ip.reserved_ip` |
| `ibm_is_share_invalid_source_snapshot__crn` | `source_snapshot.crn` |
| `ibm_is_share_invalid_source_snapshot__id` | `source_snapshot.id` |

## Example

### Valid

```hcl
resource "ibm_is_share" "test" {
  name = "example-name"
}
```
