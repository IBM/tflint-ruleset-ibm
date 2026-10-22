# ibm_database_backup_data

Schema validation for the ibm_database_backup data source.

This rule validates `ibm_database_backup` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_database_backup_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `backup_id` | yes | — |
| `created_at` | no | — |
| `deployment_id` | no | — |
| `download_link` | no | — |
| `is_downloadable` | no | — |
| `is_restorable` | no | — |
| `status` | no | — |
| `type` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_database_backup_data_invalid_backup_id` | `backup_id` |

## Example

### Valid

```hcl
data "ibm_database_backup" "test" {
  backup_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
