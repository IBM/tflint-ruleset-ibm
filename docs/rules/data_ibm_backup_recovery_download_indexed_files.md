# ibm_backup_recovery_download_indexed_files_data

Schema validation for the ibm_backup_recovery_download_indexed_files data source.

This rule validates `ibm_backup_recovery_download_indexed_files` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_backup_recovery_download_indexed_files_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `file_path` | no | — |
| `length` | no | — |
| `nvram_file` | no | — |
| `retry_attempt` | no | — |
| `snapshots_id` | yes | — |
| `start_offset` | no | — |
| `x_ibm_tenant_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_backup_recovery_download_indexed_files_data_invalid_file_path` | `file_path` |
| `ibm_backup_recovery_download_indexed_files_data_invalid_length` | `length` |
| `ibm_backup_recovery_download_indexed_files_data_invalid_nvram_file` | `nvram_file` |
| `ibm_backup_recovery_download_indexed_files_data_invalid_retry_attempt` | `retry_attempt` |
| `ibm_backup_recovery_download_indexed_files_data_invalid_snapshots_id` | `snapshots_id` |
| `ibm_backup_recovery_download_indexed_files_data_invalid_start_offset` | `start_offset` |
| `ibm_backup_recovery_download_indexed_files_data_invalid_x_ibm_tenant_id` | `x_ibm_tenant_id` |

## Example

### Valid

```hcl
data "ibm_backup_recovery_download_indexed_files" "test" {
  snapshots_id = "01234567-89ab-cdef-0123-456789abcdef"
  x_ibm_tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
