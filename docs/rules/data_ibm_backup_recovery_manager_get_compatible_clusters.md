# ibm_backup_recovery_manager_get_compatible_clusters_data

Schema validation for the ibm_backup_recovery_manager_get_compatible_clusters data source.

This rule validates `ibm_backup_recovery_manager_get_compatible_clusters` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_backup_recovery_manager_get_compatible_clusters_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `compatible_clusters` | no | — |
| `release_version` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_backup_recovery_manager_get_compatible_clusters_data_invalid_release_version` | `release_version` |

## Example

### Valid

```hcl
data "ibm_backup_recovery_manager_get_compatible_clusters" "test" {
}
```
