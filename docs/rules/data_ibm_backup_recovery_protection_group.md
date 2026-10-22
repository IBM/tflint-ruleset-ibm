# ibm_backup_recovery_protection_group_data

Reads ibm_backup_recovery_protection_group. Specifies the name of the Protection Group.

This rule validates `ibm_backup_recovery_protection_group` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_backup_recovery_protection_group_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `abort_in_blackouts` | no | — |
| `advanced_configs` | no | — |
| `alert_policy` | no | — |
| `cluster_id` | no | — |
| `description` | no | — |
| `end_time_usecs` | no | — |
| `environment` | no | — |
| `group_id` | no | — |
| `include_last_run_info` | no | — |
| `invalid_entities` | no | — |
| `is_active` | no | — |
| `is_deleted` | no | — |
| `is_paused` | no | — |
| `is_protect_once` | no | — |
| `kubernetes_params` | no | — |
| `last_modified_timestamp_usecs` | no | — |
| `last_run` | no | — |
| `missing_entities` | no | — |
| `mssql_params` | no | — |
| `name` | no | — |
| `num_protected_objects` | no | — |
| `pause_in_blackouts` | no | — |
| `permissions` | no | — |
| `physical_params` | no | — |
| `policy_id` | no | — |
| `priority` | no | — |
| `protection_group_id` | yes | — |
| `prune_excluded_source_ids` | no | — |
| `prune_source_ids` | no | — |
| `qos_policy` | no | — |
| `region_id` | no | — |
| `request_initiator_type` | no | — |
| `sla` | no | — |
| `start_time` | no | — |
| `x_ibm_tenant_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_backup_recovery_protection_group_data_invalid_include_last_run_info` | `include_last_run_info` |
| `ibm_backup_recovery_protection_group_data_invalid_protection_group_id` | `protection_group_id` |
| `ibm_backup_recovery_protection_group_data_invalid_prune_excluded_source_ids` | `prune_excluded_source_ids` |
| `ibm_backup_recovery_protection_group_data_invalid_prune_source_ids` | `prune_source_ids` |
| `ibm_backup_recovery_protection_group_data_invalid_request_initiator_type` | `request_initiator_type` |
| `ibm_backup_recovery_protection_group_data_invalid_x_ibm_tenant_id` | `x_ibm_tenant_id` |

## Example

### Valid

```hcl
data "ibm_backup_recovery_protection_group" "test" {
  protection_group_id = "01234567-89ab-cdef-0123-456789abcdef"
  x_ibm_tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
