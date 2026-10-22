# ibm_backup_recovery_connection_registration_token

Schema validation for the ibm_backup_recovery_connection_registration_token resource.

This rule validates `ibm_backup_recovery_connection_registration_token` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_backup_recovery_connection_registration_token" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `connection_id` | yes | — |
| `registration_token` | no | — |
| `x_ibm_tenant_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_backup_recovery_connection_registration_token_invalid_connection_id` | `connection_id` |
| `ibm_backup_recovery_connection_registration_token_invalid_x_ibm_tenant_id` | `x_ibm_tenant_id` |

## Example

### Valid

```hcl
resource "ibm_backup_recovery_connection_registration_token" "test" {
  connection_id = "01234567-89ab-cdef-0123-456789abcdef"
  x_ibm_tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
