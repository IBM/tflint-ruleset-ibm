# ibm_is_backup_policy_job_data

Schema validation for the ibm_is_backup_policy_job data source.

This rule validates `ibm_is_backup_policy_job` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_is_backup_policy_job_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `auto_delete` | no | — |
| `auto_delete_after` | no | — |
| `backup_policy_id` | yes | — |
| `backup_policy_plan` | no | — |
| `completed_at` | no | — |
| `created_at` | no | — |
| `href` | no | — |
| `identifier` | yes | — |
| `job_type` | no | — |
| `resource_type` | no | — |
| `source_instance` | no | — |
| `source_share` | no | — |
| `source_volume` | no | — |
| `status` | no | — |
| `status_reasons` | no | — |
| `target_snapshot` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_backup_policy_job_data_invalid_backup_policy_id` | `backup_policy_id` |
| `ibm_is_backup_policy_job_data_invalid_identifier` | `identifier` |

## Example

### Valid

```hcl
data "ibm_is_backup_policy_job" "test" {
  backup_policy_id = "01234567-89ab-cdef-0123-456789abcdef"
  identifier = "example"
}
```
