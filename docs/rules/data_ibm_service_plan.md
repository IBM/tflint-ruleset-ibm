# ibm_service_plan_data

Schema validation for the ibm_service_plan data source.

This rule validates `ibm_service_plan` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_service_plan_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `plan` | yes | — |
| `service` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_service_plan_data_invalid_plan` | `plan` |
| `ibm_service_plan_data_invalid_service` | `service` |

## Example

### Valid

```hcl
data "ibm_service_plan" "test" {
  plan = "example"
  service = "example"
}
```
