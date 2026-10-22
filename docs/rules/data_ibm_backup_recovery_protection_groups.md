# ibm_backup_recovery_protection_groups_data

Schema validation for the ibm_backup_recovery_protection_groups data source.

This rule validates `ibm_backup_recovery_protection_groups` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_backup_recovery_protection_groups_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `environments` | no | — |
| `ids` | no | — |
| `include_groups_with_datalock_only` | no | — |
| `include_last_run_info` | no | — |
| `is_active` | no | — |
| `is_deleted` | no | — |
| `is_last_run_sla_violated` | no | — |
| `is_paused` | no | — |
| `last_run_any_status` | no | — |
| `last_run_archival_status` | no | — |
| `last_run_cloud_spin_status` | no | — |
| `last_run_local_backup_status` | no | — |
| `last_run_replication_status` | no | — |
| `names` | no | — |
| `policy_ids` | no | — |
| `protection_groups` | no | — |
| `prune_excluded_source_ids` | no | — |
| `prune_source_ids` | no | — |
| `request_initiator_type` | no | — |
| `source_ids` | no | — |
| `use_cached_data` | no | — |
| `x_ibm_tenant_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_backup_recovery_protection_groups_data_invalid_environments` | `environments` |
| `ibm_backup_recovery_protection_groups_data_invalid_ids` | `ids` |
| `ibm_backup_recovery_protection_groups_data_invalid_include_groups_with_datalock_only` | `include_groups_with_datalock_only` |
| `ibm_backup_recovery_protection_groups_data_invalid_include_last_run_info` | `include_last_run_info` |
| `ibm_backup_recovery_protection_groups_data_invalid_is_active` | `is_active` |
| `ibm_backup_recovery_protection_groups_data_invalid_is_deleted` | `is_deleted` |
| `ibm_backup_recovery_protection_groups_data_invalid_is_last_run_sla_violated` | `is_last_run_sla_violated` |
| `ibm_backup_recovery_protection_groups_data_invalid_is_paused` | `is_paused` |
| `ibm_backup_recovery_protection_groups_data_invalid_last_run_any_status` | `last_run_any_status` |
| `ibm_backup_recovery_protection_groups_data_invalid_last_run_archival_status` | `last_run_archival_status` |
| `ibm_backup_recovery_protection_groups_data_invalid_last_run_cloud_spin_status` | `last_run_cloud_spin_status` |
| `ibm_backup_recovery_protection_groups_data_invalid_last_run_local_backup_status` | `last_run_local_backup_status` |
| `ibm_backup_recovery_protection_groups_data_invalid_last_run_replication_status` | `last_run_replication_status` |
| `ibm_backup_recovery_protection_groups_data_invalid_names` | `names` |
| `ibm_backup_recovery_protection_groups_data_invalid_policy_ids` | `policy_ids` |
| `ibm_backup_recovery_protection_groups_data_invalid_prune_excluded_source_ids` | `prune_excluded_source_ids` |
| `ibm_backup_recovery_protection_groups_data_invalid_prune_source_ids` | `prune_source_ids` |
| `ibm_backup_recovery_protection_groups_data_invalid_request_initiator_type` | `request_initiator_type` |
| `ibm_backup_recovery_protection_groups_data_invalid_source_ids` | `source_ids` |
| `ibm_backup_recovery_protection_groups_data_invalid_use_cached_data` | `use_cached_data` |
| `ibm_backup_recovery_protection_groups_data_invalid_x_ibm_tenant_id` | `x_ibm_tenant_id` |

## Example

### Valid

```hcl
data "ibm_backup_recovery_protection_groups" "test" {
  x_ibm_tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
