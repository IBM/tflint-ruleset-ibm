# ibm_backup_recovery_data_source_connector_patch

Schema validation for the ibm_backup_recovery_data_source_connector_patch resource.

This rule validates `ibm_backup_recovery_data_source_connector_patch` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_backup_recovery_data_source_connector_patch" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cluster_side_ip` | no | — |
| `connection_id` | no | — |
| `connectivity_status` | no | — |
| `connector_id` | yes | — |
| `connector_name` | no | — |
| `software_version` | no | — |
| `tenant_side_ip` | no | — |
| `upgrade_status` | no | — |
| `x_ibm_tenant_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_backup_recovery_data_source_connector_patch_invalid_connector_id` | `connector_id` |
| `ibm_backup_recovery_data_source_connector_patch_invalid_connector_name` | `connector_name` |
| `ibm_backup_recovery_data_source_connector_patch_invalid_x_ibm_tenant_id` | `x_ibm_tenant_id` |

## Example

### Valid

```hcl
resource "ibm_backup_recovery_data_source_connector_patch" "test" {
  connector_id = "01234567-89ab-cdef-0123-456789abcdef"
  x_ibm_tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
