# ibm_backup_recovery_connector_agents_data

Schema validation for the ibm_backup_recovery_connector_agents data source.

This rule validates `ibm_backup_recovery_connector_agents` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_backup_recovery_connector_agents_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `connection_ids` | no | — |
| `connection_names` | no | — |
| `connector_agents` | no | — |
| `tenant_id` | yes | — |
| `x_ibm_tenant_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_backup_recovery_connector_agents_data_invalid_connection_ids` | `connection_ids` |
| `ibm_backup_recovery_connector_agents_data_invalid_connection_names` | `connection_names` |
| `ibm_backup_recovery_connector_agents_data_invalid_tenant_id` | `tenant_id` |
| `ibm_backup_recovery_connector_agents_data_invalid_x_ibm_tenant_id` | `x_ibm_tenant_id` |

## Example

### Valid

```hcl
data "ibm_backup_recovery_connector_agents" "test" {
  tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
  x_ibm_tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
