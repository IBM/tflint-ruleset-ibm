# ibm_backup_recovery_download_files_data

Schema validation for the ibm_backup_recovery_download_files data source.

This rule validates `ibm_backup_recovery_download_files` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_backup_recovery_download_files_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `file_type` | no | — |
| `include_tenants` | no | — |
| `length` | no | — |
| `recovery_download_files_id` | yes | — |
| `source_name` | no | — |
| `start_offset` | no | — |
| `start_time` | no | — |
| `x_ibm_tenant_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_backup_recovery_download_files_data_invalid_file_type` | `file_type` |
| `ibm_backup_recovery_download_files_data_invalid_include_tenants` | `include_tenants` |
| `ibm_backup_recovery_download_files_data_invalid_length` | `length` |
| `ibm_backup_recovery_download_files_data_invalid_recovery_download_files_id` | `recovery_download_files_id` |
| `ibm_backup_recovery_download_files_data_invalid_source_name` | `source_name` |
| `ibm_backup_recovery_download_files_data_invalid_start_offset` | `start_offset` |
| `ibm_backup_recovery_download_files_data_invalid_start_time` | `start_time` |
| `ibm_backup_recovery_download_files_data_invalid_x_ibm_tenant_id` | `x_ibm_tenant_id` |

## Example

### Valid

```hcl
data "ibm_backup_recovery_download_files" "test" {
  recovery_download_files_id = "01234567-89ab-cdef-0123-456789abcdef"
  x_ibm_tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
