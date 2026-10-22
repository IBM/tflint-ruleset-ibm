# ibm_database_task_data

Schema validation for the ibm_database_task data source.

This rule validates `ibm_database_task` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_database_task_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `created_at` | no | — |
| `deployment_id` | no | — |
| `description` | no | — |
| `progress_percent` | no | — |
| `status` | no | — |
| `task_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_database_task_data_invalid_task_id` | `task_id` |

## Example

### Valid

```hcl
data "ibm_database_task" "test" {
  task_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
