# ibm_database_tasks_data

Schema validation for the ibm_database_tasks data source.

This rule validates `ibm_database_tasks` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_database_tasks_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `deployment_id` | yes | — |
| `tasks` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_database_tasks_data_invalid_deployment_id` | `deployment_id` |

## Example

### Valid

```hcl
data "ibm_database_tasks" "test" {
  deployment_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
