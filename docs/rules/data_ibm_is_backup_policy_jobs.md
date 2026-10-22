# ibm_is_backup_policy_jobs_data

Schema validation for the ibm_is_backup_policy_jobs data source.

This rule validates `ibm_is_backup_policy_jobs` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_is_backup_policy_jobs_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `backup_policy_id` | yes | — |
| `backup_policy_plan_id` | no | — |
| `jobs` | no | — |
| `source_id` | no | — |
| `status` | no | — |
| `target_snapshots_crn` | no | — |
| `target_snapshots_id` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_backup_policy_jobs_data_invalid_backup_policy_id` | `backup_policy_id` |
| `ibm_is_backup_policy_jobs_data_invalid_backup_policy_plan_id` | `backup_policy_plan_id` |
| `ibm_is_backup_policy_jobs_data_invalid_source_id` | `source_id` |
| `ibm_is_backup_policy_jobs_data_invalid_status` | `status` |
| `ibm_is_backup_policy_jobs_data_invalid_target_snapshots_crn` | `target_snapshots_crn` |
| `ibm_is_backup_policy_jobs_data_invalid_target_snapshots_id` | `target_snapshots_id` |

## Example

### Valid

```hcl
data "ibm_is_backup_policy_jobs" "test" {
  backup_policy_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
