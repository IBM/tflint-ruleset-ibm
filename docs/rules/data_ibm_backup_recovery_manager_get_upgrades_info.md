# ibm_backup_recovery_manager_get_upgrades_info_data

Schema validation for the ibm_backup_recovery_manager_get_upgrades_info data source.

This rule validates `ibm_backup_recovery_manager_get_upgrades_info` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_backup_recovery_manager_get_upgrades_info_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cluster_identifiers` | no | — |
| `upgrades_info` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_backup_recovery_manager_get_upgrades_info_data_invalid_cluster_identifiers` | `cluster_identifiers` |

## Example

### Valid

```hcl
data "ibm_backup_recovery_manager_get_upgrades_info" "test" {
}
```
