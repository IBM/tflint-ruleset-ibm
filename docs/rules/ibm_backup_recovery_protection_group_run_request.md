# ibm_backup_recovery_protection_group_run_request

Schema validation for the ibm_backup_recovery_protection_group_run_request resource.

This rule validates `ibm_backup_recovery_protection_group_run_request` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Recovery` (`baas-backup-recovery`).

## Configuration

```hcl
rule "ibm_backup_recovery_protection_group_run_request" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `group_id` | yes | — |
| `run_type` | yes | — |
| `x_ibm_tenant_id` | yes | — |
| `objects` (block) | no | nested block |
| `objects.app_ids` | no | — |
| `objects.id` | yes | API: Specifies the id of the Recovery. |
| `objects.physical_params` (block) | no | nested block |
| `objects.physical_params.metadata_file_path` | no | — |
| `targets_config` (block) | no | nested block |
| `targets_config.use_policy_defaults` | no | — |
| `targets_config.archivals` (block) | no | nested block |
| `targets_config.archivals.archival_target_type` | yes | — |
| `targets_config.archivals.copy_only_fully_successful` | no | — |
| `targets_config.archivals.id` | yes | API: Specifies the id of the Recovery. |
| `targets_config.archivals.retention` (block) | no | nested block |
| `targets_config.archivals.retention.duration` | yes | — |
| `targets_config.archivals.retention.unit` | yes | — |
| `targets_config.archivals.retention.data_lock_config` (block) | no | nested block |
| `targets_config.archivals.retention.data_lock_config.duration` | yes | — |
| `targets_config.archivals.retention.data_lock_config.enable_worm_on_external_target` | no | — |
| `targets_config.archivals.retention.data_lock_config.mode` | yes | one of: `Compliance`, `Administrative` |
| `targets_config.archivals.retention.data_lock_config.unit` | yes | — |
| `targets_config.cloud_replications` (block) | no | nested block |
| `targets_config.cloud_replications.target_type` | yes | — |
| `targets_config.cloud_replications.aws_target` (block) | no | nested block |
| `targets_config.cloud_replications.aws_target.region` | yes | API: Specifies id of the AWS region in which to replicate the Snapshot to. Applicable if replication target is AWS target. |
| `targets_config.cloud_replications.aws_target.source_id` | yes | — |
| `targets_config.cloud_replications.azure_target` (block) | no | nested block |
| `targets_config.cloud_replications.azure_target.resource_group` | no | — |
| `targets_config.cloud_replications.azure_target.source_id` | yes | — |
| `targets_config.cloud_replications.retention` (block) | no | nested block |
| `targets_config.cloud_replications.retention.duration` | yes | — |
| `targets_config.cloud_replications.retention.unit` | yes | — |
| `targets_config.cloud_replications.retention.data_lock_config` (block) | no | nested block |
| `targets_config.cloud_replications.retention.data_lock_config.duration` | yes | — |
| `targets_config.cloud_replications.retention.data_lock_config.enable_worm_on_external_target` | no | — |
| `targets_config.cloud_replications.retention.data_lock_config.mode` | yes | one of: `Compliance`, `Administrative` |
| `targets_config.cloud_replications.retention.data_lock_config.unit` | yes | — |
| `targets_config.replications` (block) | no | nested block |
| `targets_config.replications.id` | yes | API: Specifies the id of the Recovery. |
| `targets_config.replications.retention` (block) | no | nested block |
| `targets_config.replications.retention.duration` | yes | — |
| `targets_config.replications.retention.unit` | yes | — |
| `targets_config.replications.retention.data_lock_config` (block) | no | nested block |
| `targets_config.replications.retention.data_lock_config.duration` | yes | — |
| `targets_config.replications.retention.data_lock_config.enable_worm_on_external_target` | no | — |
| `targets_config.replications.retention.data_lock_config.mode` | yes | one of: `Compliance`, `Administrative` |
| `targets_config.replications.retention.data_lock_config.unit` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_backup_recovery_protection_group_run_request_invalid_group_id` | `group_id` |
| `ibm_backup_recovery_protection_group_run_request_invalid_run_type` | `run_type` |
| `ibm_backup_recovery_protection_group_run_request_invalid_x_ibm_tenant_id` | `x_ibm_tenant_id` |
| `ibm_backup_recovery_protection_group_run_request_invalid_objects__app_ids` | `objects.app_ids` |
| `ibm_backup_recovery_protection_group_run_request_invalid_objects__id` | `objects.id` |
| `ibm_backup_recovery_protection_group_run_request_invalid_objects__physical_params__metadata_file_path` | `objects.physical_params.metadata_file_path` |
| `ibm_backup_recovery_protection_group_run_request_invalid_targets_config__use_policy_defaults` | `targets_config.use_policy_defaults` |
| `ibm_backup_recovery_protection_group_run_request_invalid_targets_config__archivals__archival_target_type` | `targets_config.archivals.archival_target_type` |
| `ibm_backup_recovery_protection_group_run_request_invalid_targets_config__archivals__copy_only_fully_successful` | `targets_config.archivals.copy_only_fully_successful` |
| `ibm_backup_recovery_protection_group_run_request_invalid_targets_config__archivals__id` | `targets_config.archivals.id` |
| `ibm_backup_recovery_protection_group_run_request_invalid_targets_config__archivals__retention__duration` | `targets_config.archivals.retention.duration` |
| `ibm_backup_recovery_protection_group_run_request_invalid_targets_config__archivals__retention__unit` | `targets_config.archivals.retention.unit` |
| `ibm_backup_recovery_protection_group_run_request_invalid_targets_config__archivals__retention__data_lock_config__duration` | `targets_config.archivals.retention.data_lock_config.duration` |
| `ibm_backup_recovery_protection_group_run_request_invalid_targets_config__archivals__retention__data_lock_config__enable_worm_on_external_target` | `targets_config.archivals.retention.data_lock_config.enable_worm_on_external_target` |
| `ibm_backup_recovery_protection_group_run_request_invalid_targets_config__archivals__retention__data_lock_config__mode` | `targets_config.archivals.retention.data_lock_config.mode` |
| `ibm_backup_recovery_protection_group_run_request_invalid_targets_config__archivals__retention__data_lock_config__unit` | `targets_config.archivals.retention.data_lock_config.unit` |
| `ibm_backup_recovery_protection_group_run_request_invalid_targets_config__cloud_replications__target_type` | `targets_config.cloud_replications.target_type` |
| `ibm_backup_recovery_protection_group_run_request_invalid_targets_config__cloud_replications__aws_target__region` | `targets_config.cloud_replications.aws_target.region` |
| `ibm_backup_recovery_protection_group_run_request_invalid_targets_config__cloud_replications__aws_target__source_id` | `targets_config.cloud_replications.aws_target.source_id` |
| `ibm_backup_recovery_protection_group_run_request_invalid_targets_config__cloud_replications__azure_target__resource_group` | `targets_config.cloud_replications.azure_target.resource_group` |
| `ibm_backup_recovery_protection_group_run_request_invalid_targets_config__cloud_replications__azure_target__source_id` | `targets_config.cloud_replications.azure_target.source_id` |
| `ibm_backup_recovery_protection_group_run_request_invalid_targets_config__cloud_replications__retention__duration` | `targets_config.cloud_replications.retention.duration` |
| `ibm_backup_recovery_protection_group_run_request_invalid_targets_config__cloud_replications__retention__unit` | `targets_config.cloud_replications.retention.unit` |
| `ibm_backup_recovery_protection_group_run_request_invalid_targets_config__cloud_replications__retention__data_lock_config__duration` | `targets_config.cloud_replications.retention.data_lock_config.duration` |
| `ibm_backup_recovery_protection_group_run_request_invalid_targets_config__cloud_replications__retention__data_lock_config__enable_worm_on_external_target` | `targets_config.cloud_replications.retention.data_lock_config.enable_worm_on_external_target` |
| `ibm_backup_recovery_protection_group_run_request_invalid_targets_config__cloud_replications__retention__data_lock_config__mode` | `targets_config.cloud_replications.retention.data_lock_config.mode` |
| `ibm_backup_recovery_protection_group_run_request_invalid_targets_config__cloud_replications__retention__data_lock_config__unit` | `targets_config.cloud_replications.retention.data_lock_config.unit` |
| `ibm_backup_recovery_protection_group_run_request_invalid_targets_config__replications__id` | `targets_config.replications.id` |
| `ibm_backup_recovery_protection_group_run_request_invalid_targets_config__replications__retention__duration` | `targets_config.replications.retention.duration` |
| `ibm_backup_recovery_protection_group_run_request_invalid_targets_config__replications__retention__unit` | `targets_config.replications.retention.unit` |
| `ibm_backup_recovery_protection_group_run_request_invalid_targets_config__replications__retention__data_lock_config__duration` | `targets_config.replications.retention.data_lock_config.duration` |
| `ibm_backup_recovery_protection_group_run_request_invalid_targets_config__replications__retention__data_lock_config__enable_worm_on_external_target` | `targets_config.replications.retention.data_lock_config.enable_worm_on_external_target` |
| `ibm_backup_recovery_protection_group_run_request_invalid_targets_config__replications__retention__data_lock_config__mode` | `targets_config.replications.retention.data_lock_config.mode` |
| `ibm_backup_recovery_protection_group_run_request_invalid_targets_config__replications__retention__data_lock_config__unit` | `targets_config.replications.retention.data_lock_config.unit` |

## Example

### Valid

```hcl
resource "ibm_backup_recovery_protection_group_run_request" "test" {
  group_id = "01234567-89ab-cdef-0123-456789abcdef"
  run_type = "example"
  x_ibm_tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
