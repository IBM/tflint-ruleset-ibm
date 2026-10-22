# ibm_backup_recovery_source_registrations_data

Schema validation for the ibm_backup_recovery_source_registrations data source.

This rule validates `ibm_backup_recovery_source_registrations` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_backup_recovery_source_registrations_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `encryption_key` | no | — |
| `ids` | no | — |
| `ignore_tenant_migration_in_progress_check` | no | — |
| `include_external_metadata` | no | — |
| `include_source_credentials` | no | — |
| `registrations` | no | — |
| `use_cached_data` | no | — |
| `x_ibm_tenant_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_backup_recovery_source_registrations_data_invalid_encryption_key` | `encryption_key` |
| `ibm_backup_recovery_source_registrations_data_invalid_ids` | `ids` |
| `ibm_backup_recovery_source_registrations_data_invalid_ignore_tenant_migration_in_progress_check` | `ignore_tenant_migration_in_progress_check` |
| `ibm_backup_recovery_source_registrations_data_invalid_include_external_metadata` | `include_external_metadata` |
| `ibm_backup_recovery_source_registrations_data_invalid_include_source_credentials` | `include_source_credentials` |
| `ibm_backup_recovery_source_registrations_data_invalid_use_cached_data` | `use_cached_data` |
| `ibm_backup_recovery_source_registrations_data_invalid_x_ibm_tenant_id` | `x_ibm_tenant_id` |

## Example

### Valid

```hcl
data "ibm_backup_recovery_source_registrations" "test" {
  x_ibm_tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
