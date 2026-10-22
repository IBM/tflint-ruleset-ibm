# ibm_backup_recovery_data_source_connection

Schema validation for the ibm_backup_recovery_data_source_connection resource.

This rule validates `ibm_backup_recovery_data_source_connection` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_backup_recovery_data_source_connection" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `connection_env_type` | no | one of: `kIksClassic`, `kIksVpc`, `kRoksClassic`, `kRoksVpc` |
| `connection_id` | no | — |
| `connection_name` | yes | — |
| `connector_ids` | no | — |
| `registration_token` | no | — |
| `tenant_id` | no | — |
| `upgrading_connector_id` | no | — |
| `x_ibm_tenant_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_backup_recovery_data_source_connection_invalid_connection_env_type` | `connection_env_type` |
| `ibm_backup_recovery_data_source_connection_invalid_connection_name` | `connection_name` |
| `ibm_backup_recovery_data_source_connection_invalid_x_ibm_tenant_id` | `x_ibm_tenant_id` |

## Example

### Valid

```hcl
resource "ibm_backup_recovery_data_source_connection" "test" {
  connection_name = "example"
  x_ibm_tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
