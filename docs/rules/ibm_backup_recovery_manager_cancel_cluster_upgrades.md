# ibm_backup_recovery_manager_cancel_cluster_upgrades

Schema validation for the ibm_backup_recovery_manager_cancel_cluster_upgrades resource.

This rule validates `ibm_backup_recovery_manager_cancel_cluster_upgrades` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_backup_recovery_manager_cancel_cluster_upgrades" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cancelled_upgrade_response_list` | no | — |
| `cluster_identifiers` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_backup_recovery_manager_cancel_cluster_upgrades_invalid_cluster_identifiers` | `cluster_identifiers` |

## Example

### Valid

```hcl
resource "ibm_backup_recovery_manager_cancel_cluster_upgrades" "test" {
  cluster_identifiers = "example"
}
```
