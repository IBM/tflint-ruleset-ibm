# ibm_database_backups_data

Schema validation for the ibm_database_backups data source.

This rule validates `ibm_database_backups` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Backups` (`cloud-databases-api-cloud-databases-api-v5`).

## Configuration

```hcl
rule "ibm_database_backups_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `backups` | no | — |
| `deployment_id` | no | API: ID of the deployment this backup relates to. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_database_backups_data_invalid_deployment_id` | `deployment_id` |

## Example

### Valid

```hcl
data "ibm_database_backups" "test" {
}
```
