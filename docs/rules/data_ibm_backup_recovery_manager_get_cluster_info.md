# ibm_backup_recovery_manager_get_cluster_info_data

Schema validation for the ibm_backup_recovery_manager_get_cluster_info data source.

This rule validates `ibm_backup_recovery_manager_get_cluster_info` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_backup_recovery_manager_get_cluster_info_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cohesity_clusters` | no | — |
| `sp_clusters` | no | — |

## Example

### Valid

```hcl
data "ibm_backup_recovery_manager_get_cluster_info" "test" {
}
```
