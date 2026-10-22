# ibm_db2_users_data

Schema validation for the ibm_db2_users data source.

This rule validates `ibm_db2_users` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_db2_users_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `resources` | no | — |
| `users_count` | no | — |
| `x_deployment_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_db2_users_data_invalid_x_deployment_id` | `x_deployment_id` |

## Example

### Valid

```hcl
data "ibm_db2_users" "test" {
  x_deployment_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
