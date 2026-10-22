# ibm_backup_recovery_protection_source_refresh

Schema validation for the ibm_backup_recovery_protection_source_refresh resource.

This rule validates `ibm_backup_recovery_protection_source_refresh` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_backup_recovery_protection_source_refresh" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `backup_recovery_protection_source_id` | yes | — |
| `x_ibm_tenant_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_backup_recovery_protection_source_refresh_invalid_backup_recovery_protection_source_id` | `backup_recovery_protection_source_id` |
| `ibm_backup_recovery_protection_source_refresh_invalid_x_ibm_tenant_id` | `x_ibm_tenant_id` |

## Example

### Valid

```hcl
resource "ibm_backup_recovery_protection_source_refresh" "test" {
  backup_recovery_protection_source_id = 1
  x_ibm_tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
