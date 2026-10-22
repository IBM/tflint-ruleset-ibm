# ibm_backup_recovery_manager_create_cluster_upgrades

Schema validation for the ibm_backup_recovery_manager_create_cluster_upgrades resource.

This rule validates `ibm_backup_recovery_manager_create_cluster_upgrades` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Recovery` (`baas-backup-recovery`).

## Configuration

```hcl
rule "ibm_backup_recovery_manager_create_cluster_upgrades" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `interval_for_rolling_upgrade_in_hours` | no | — |
| `package_url` | no | — |
| `target_version` | no | — |
| `time_stamp_to_upgrade_at_msecs` | no | — |
| `type` | no | one of: `IBM` |
| `upgrade_response_list` | no | — |
| `auth_headers` (block) | no | nested block |
| `auth_headers.key` | yes | API: Specifies the unique key for custom property. |
| `auth_headers.value` | yes | API: Specifies the value for the above custom key. |
| `clusters` (block) | yes | nested block |
| `clusters.cluster_id` | no | — |
| `clusters.cluster_incarnation_id` | no | — |
| `clusters.current_version` | no | — |
| `patch_upgrade_params` (block) | no | nested block |
| `patch_upgrade_params.ignore_pre_checks_failure` | no | — |
| `patch_upgrade_params.package_url` | no | — |
| `patch_upgrade_params.target_version` | no | — |
| `patch_upgrade_params.auth_headers` (block) | no | nested block |
| `patch_upgrade_params.auth_headers.key` | yes | API: Specifies the unique key for custom property. |
| `patch_upgrade_params.auth_headers.value` | yes | API: Specifies the value for the above custom key. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_backup_recovery_manager_create_cluster_upgrades_invalid_interval_for_rolling_upgrade_in_hours` | `interval_for_rolling_upgrade_in_hours` |
| `ibm_backup_recovery_manager_create_cluster_upgrades_invalid_package_url` | `package_url` |
| `ibm_backup_recovery_manager_create_cluster_upgrades_invalid_target_version` | `target_version` |
| `ibm_backup_recovery_manager_create_cluster_upgrades_invalid_time_stamp_to_upgrade_at_msecs` | `time_stamp_to_upgrade_at_msecs` |
| `ibm_backup_recovery_manager_create_cluster_upgrades_invalid_type` | `type` |
| `ibm_backup_recovery_manager_create_cluster_upgrades_invalid_auth_headers__key` | `auth_headers.key` |
| `ibm_backup_recovery_manager_create_cluster_upgrades_invalid_auth_headers__value` | `auth_headers.value` |
| `ibm_backup_recovery_manager_create_cluster_upgrades_invalid_clusters__cluster_id` | `clusters.cluster_id` |
| `ibm_backup_recovery_manager_create_cluster_upgrades_invalid_clusters__cluster_incarnation_id` | `clusters.cluster_incarnation_id` |
| `ibm_backup_recovery_manager_create_cluster_upgrades_invalid_clusters__current_version` | `clusters.current_version` |
| `ibm_backup_recovery_manager_create_cluster_upgrades_invalid_patch_upgrade_params__ignore_pre_checks_failure` | `patch_upgrade_params.ignore_pre_checks_failure` |
| `ibm_backup_recovery_manager_create_cluster_upgrades_invalid_patch_upgrade_params__package_url` | `patch_upgrade_params.package_url` |
| `ibm_backup_recovery_manager_create_cluster_upgrades_invalid_patch_upgrade_params__target_version` | `patch_upgrade_params.target_version` |
| `ibm_backup_recovery_manager_create_cluster_upgrades_invalid_patch_upgrade_params__auth_headers__key` | `patch_upgrade_params.auth_headers.key` |
| `ibm_backup_recovery_manager_create_cluster_upgrades_invalid_patch_upgrade_params__auth_headers__value` | `patch_upgrade_params.auth_headers.value` |

## Example

### Valid

```hcl
resource "ibm_backup_recovery_manager_create_cluster_upgrades" "test" {
  clusters {
  }
}
```
