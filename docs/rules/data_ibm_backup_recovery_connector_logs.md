# ibm_backup_recovery_connector_logs_data

Schema validation for the ibm_backup_recovery_connector_logs data source.

This rule validates `ibm_backup_recovery_connector_logs` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_backup_recovery_connector_logs_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_token` | yes | — |
| `connector_logs` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_backup_recovery_connector_logs_data_invalid_access_token` | `access_token` |

## Example

### Valid

```hcl
variable "access_token" {
  sensitive = true
}

data "ibm_backup_recovery_connector_logs" "test" {
  access_token = var.access_token
}
```
