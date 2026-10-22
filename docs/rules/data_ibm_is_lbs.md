# ibm_is_lbs_data

Schema validation for the ibm_is_lbs data source.

This rule validates `ibm_is_lbs` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_is_lbs_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `load_balancers` | no | — |

## Example

### Valid

```hcl
data "ibm_is_lbs" "test" {
}
```
