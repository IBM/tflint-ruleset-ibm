# ibm_database

Manages ibm_database. Resource instance name for example, my Database instance

This rule validates `ibm_database` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `leaf-unique` (`cloud-databases-api-cloud-databases-api-v5`).

## Configuration

```hcl
rule "ibm_database" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `adminpassword` | no | — |
| `adminuser` | no | — |
| `async_restore` | no | — |
| `backup_encryption_key_crn` | no | crn format |
| `backup_id` | no | — |
| `configuration` | no | — |
| `configuration_schema` | no | — |
| `deletion_protection` | no | — |
| `groups` | no | — |
| `guid` | no | — |
| `key_protect_instance` | no | — |
| `key_protect_key` | no | — |
| `location` | yes | — |
| `name` | yes | — |
| `offline_restore` | no | — |
| `plan` | yes | one of: `standard`, `standard-gen2`, `enterprise`, `enterprise-gen2`, `enterprise-sharding`, `enterprise-sharding-gen2`, `platinum` |
| `point_in_time_recovery_deployment_id` | no | — |
| `point_in_time_recovery_time` | no | — |
| `remote_leader_id` | no | — |
| `resource_controller_url` | no | — |
| `resource_crn` | no | — |
| `resource_group_id` | no | — |
| `resource_group_name` | no | — |
| `resource_name` | no | — |
| `resource_status` | no | — |
| `service` | yes | one of: `databases-for-etcd`, `databases-for-postgresql`, `databases-for-redis`, `databases-for-valkey`, `databases-for-valkey-cdp-dev`, `databases-for-elasticsearch`, `databases-for-mongodb`, `messages-for-rabbitmq`, `databases-for-mysql`, `databases-for-enterprisedb` |
| `service_endpoints` | no | one of: `public`, `private`, `public-and-private` |
| `skip_initial_backup` | no | — |
| `status` | no | — |
| `tags` | no | pattern `^[A-Za-z0-9:_ .-]+$` |
| `version` | no | — |
| `version_upgrade_skip_backup` | no | — |
| `allowlist` (block) | no | nested block |
| `allowlist.address` | no | cidr format |
| `allowlist.description` | no | — |
| `auto_scaling` (block) | no | nested block |
| `auto_scaling.cpu` (block) | no | nested block |
| `auto_scaling.cpu.rate_increase_percent` | no | — |
| `auto_scaling.cpu.rate_limit_count_per_member` | no | — |
| `auto_scaling.cpu.rate_period_seconds` | no | — |
| `auto_scaling.cpu.rate_units` | no | — |
| `auto_scaling.disk` (block) | no | nested block |
| `auto_scaling.disk.capacity_enabled` | no | — |
| `auto_scaling.disk.free_space_less_than_percent` | no | — |
| `auto_scaling.disk.io_above_percent` | no | — |
| `auto_scaling.disk.io_enabled` | no | — |
| `auto_scaling.disk.io_over_period` | no | — |
| `auto_scaling.disk.rate_increase_percent` | no | — |
| `auto_scaling.disk.rate_limit_mb_per_member` | no | — |
| `auto_scaling.disk.rate_period_seconds` | no | — |
| `auto_scaling.disk.rate_units` | no | — |
| `auto_scaling.memory` (block) | no | nested block |
| `auto_scaling.memory.io_above_percent` | no | — |
| `auto_scaling.memory.io_enabled` | no | — |
| `auto_scaling.memory.io_over_period` | no | — |
| `auto_scaling.memory.rate_increase_percent` | no | — |
| `auto_scaling.memory.rate_limit_mb_per_member` | no | — |
| `auto_scaling.memory.rate_period_seconds` | no | — |
| `auto_scaling.memory.rate_units` | no | — |
| `group` (block) | no | nested block |
| `group.group_id` | yes | — |
| `group.cpu` (block) | no | nested block |
| `group.cpu.allocation_count` | yes | — |
| `group.disk` (block) | no | nested block |
| `group.disk.allocation_mb` | yes | — |
| `group.host_flavor` (block) | no | nested block |
| `group.host_flavor.id` | yes | one of: `multitenant`, `b3c.4x16.encrypted`, `b3c.8x32.encrypted`, `m3c.8x64.encrypted`, `b3c.16x64.encrypted`, `b3c.32x128.encrypted`, `m3c.30x240.encrypted`, `bx3d.4x20`, `bx3d.8x40`, `bx3d.16x80`, `bx3d.32x160`, `bx3d.48x240`, `bxf.16x64`, `bxf.32x128`, `bxf.48x192`, `bxf.4x16`, `bxf.8x32` |
| `group.members` (block) | no | nested block |
| `group.members.allocation_count` | yes | — |
| `group.memory` (block) | no | nested block |
| `group.memory.allocation_mb` | yes | — |
| `logical_replication_slot` (block) | no | nested block |
| `logical_replication_slot.database_name` | yes | — |
| `logical_replication_slot.name` | yes | — |
| `logical_replication_slot.plugin_type` | yes | — |
| `users` (block) | no | nested block |
| `users.name` | yes | — |
| `users.password` | yes | API: Password for new Ops Manager user. Password must be at least 15 characters in length and contain a letter, number and a special character. |
| `users.role` | no | — |
| `users.type` | no | one of: `database`, `ops_manager`, `read_only_replica` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_database_invalid_adminpassword` | `adminpassword` |
| `ibm_database_invalid_async_restore` | `async_restore` |
| `ibm_database_invalid_backup_encryption_key_crn` | `backup_encryption_key_crn` |
| `ibm_database_invalid_backup_id` | `backup_id` |
| `ibm_database_invalid_configuration` | `configuration` |
| `ibm_database_invalid_deletion_protection` | `deletion_protection` |
| `ibm_database_invalid_key_protect_instance` | `key_protect_instance` |
| `ibm_database_invalid_key_protect_key` | `key_protect_key` |
| `ibm_database_invalid_location` | `location` |
| `ibm_database_invalid_name` | `name` |
| `ibm_database_invalid_offline_restore` | `offline_restore` |
| `ibm_database_invalid_plan` | `plan` |
| `ibm_database_invalid_point_in_time_recovery_deployment_id` | `point_in_time_recovery_deployment_id` |
| `ibm_database_invalid_point_in_time_recovery_time` | `point_in_time_recovery_time` |
| `ibm_database_invalid_remote_leader_id` | `remote_leader_id` |
| `ibm_database_invalid_resource_group_id` | `resource_group_id` |
| `ibm_database_invalid_service` | `service` |
| `ibm_database_invalid_service_endpoints` | `service_endpoints` |
| `ibm_database_invalid_skip_initial_backup` | `skip_initial_backup` |
| `ibm_database_invalid_tags` | `tags` |
| `ibm_database_invalid_version` | `version` |
| `ibm_database_invalid_version_upgrade_skip_backup` | `version_upgrade_skip_backup` |
| `ibm_database_invalid_allowlist__address` | `allowlist.address` |
| `ibm_database_invalid_allowlist__description` | `allowlist.description` |
| `ibm_database_invalid_auto_scaling__cpu__rate_increase_percent` | `auto_scaling.cpu.rate_increase_percent` |
| `ibm_database_invalid_auto_scaling__cpu__rate_limit_count_per_member` | `auto_scaling.cpu.rate_limit_count_per_member` |
| `ibm_database_invalid_auto_scaling__cpu__rate_period_seconds` | `auto_scaling.cpu.rate_period_seconds` |
| `ibm_database_invalid_auto_scaling__cpu__rate_units` | `auto_scaling.cpu.rate_units` |
| `ibm_database_invalid_auto_scaling__disk__capacity_enabled` | `auto_scaling.disk.capacity_enabled` |
| `ibm_database_invalid_auto_scaling__disk__free_space_less_than_percent` | `auto_scaling.disk.free_space_less_than_percent` |
| `ibm_database_invalid_auto_scaling__disk__io_above_percent` | `auto_scaling.disk.io_above_percent` |
| `ibm_database_invalid_auto_scaling__disk__io_enabled` | `auto_scaling.disk.io_enabled` |
| `ibm_database_invalid_auto_scaling__disk__io_over_period` | `auto_scaling.disk.io_over_period` |
| `ibm_database_invalid_auto_scaling__disk__rate_increase_percent` | `auto_scaling.disk.rate_increase_percent` |
| `ibm_database_invalid_auto_scaling__disk__rate_limit_mb_per_member` | `auto_scaling.disk.rate_limit_mb_per_member` |
| `ibm_database_invalid_auto_scaling__disk__rate_period_seconds` | `auto_scaling.disk.rate_period_seconds` |
| `ibm_database_invalid_auto_scaling__disk__rate_units` | `auto_scaling.disk.rate_units` |
| `ibm_database_invalid_auto_scaling__memory__io_above_percent` | `auto_scaling.memory.io_above_percent` |
| `ibm_database_invalid_auto_scaling__memory__io_enabled` | `auto_scaling.memory.io_enabled` |
| `ibm_database_invalid_auto_scaling__memory__io_over_period` | `auto_scaling.memory.io_over_period` |
| `ibm_database_invalid_auto_scaling__memory__rate_increase_percent` | `auto_scaling.memory.rate_increase_percent` |
| `ibm_database_invalid_auto_scaling__memory__rate_limit_mb_per_member` | `auto_scaling.memory.rate_limit_mb_per_member` |
| `ibm_database_invalid_auto_scaling__memory__rate_period_seconds` | `auto_scaling.memory.rate_period_seconds` |
| `ibm_database_invalid_auto_scaling__memory__rate_units` | `auto_scaling.memory.rate_units` |
| `ibm_database_invalid_group__group_id` | `group.group_id` |
| `ibm_database_invalid_group__cpu__allocation_count` | `group.cpu.allocation_count` |
| `ibm_database_invalid_group__disk__allocation_mb` | `group.disk.allocation_mb` |
| `ibm_database_invalid_group__host_flavor__id` | `group.host_flavor.id` |
| `ibm_database_invalid_group__members__allocation_count` | `group.members.allocation_count` |
| `ibm_database_invalid_group__memory__allocation_mb` | `group.memory.allocation_mb` |
| `ibm_database_invalid_logical_replication_slot__database_name` | `logical_replication_slot.database_name` |
| `ibm_database_invalid_logical_replication_slot__name` | `logical_replication_slot.name` |
| `ibm_database_invalid_logical_replication_slot__plugin_type` | `logical_replication_slot.plugin_type` |
| `ibm_database_invalid_users__name` | `users.name` |
| `ibm_database_invalid_users__password` | `users.password` |
| `ibm_database_invalid_users__role` | `users.role` |
| `ibm_database_invalid_users__type` | `users.type` |

## Example

### Valid

```hcl
resource "ibm_database" "test" {
  location = "example"
  name = "example"
  plan = "standard"
  service = "databases-for-etcd"
}
```
