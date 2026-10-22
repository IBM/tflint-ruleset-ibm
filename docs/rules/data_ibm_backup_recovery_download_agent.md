# ibm_backup_recovery_download_agent_data

Schema validation for the ibm_backup_recovery_download_agent data source.

This rule validates `ibm_backup_recovery_download_agent` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `DownloadAgentRequestParams` (`baas-backup-recovery`).

## Configuration

```hcl
rule "ibm_backup_recovery_download_agent_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `file_path` | yes | — |
| `platform` | yes | one of: `kWindows`, `kLinux` |
| `x_ibm_tenant_id` | yes | — |
| `linux_params` (block) | no | nested block |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_backup_recovery_download_agent_data_invalid_file_path` | `file_path` |
| `ibm_backup_recovery_download_agent_data_invalid_platform` | `platform` |
| `ibm_backup_recovery_download_agent_data_invalid_x_ibm_tenant_id` | `x_ibm_tenant_id` |

## Example

### Valid

```hcl
data "ibm_backup_recovery_download_agent" "test" {
  file_path = "example"
  platform = "kWindows"
  x_ibm_tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
