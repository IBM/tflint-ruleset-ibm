# ibm_database_point_in_time_recovery_data

Schema validation for the ibm_database_point_in_time_recovery data source.

This rule validates `ibm_database_point_in_time_recovery` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_database_point_in_time_recovery_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `deployment_id` | yes | — |
| `earliest_point_in_time_recovery_time` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_database_point_in_time_recovery_data_invalid_deployment_id` | `deployment_id` |

## Example

### Valid

```hcl
data "ibm_database_point_in_time_recovery" "test" {
  deployment_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
