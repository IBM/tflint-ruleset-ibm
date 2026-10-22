# ibm_is_operating_systems_data

Schema validation for the ibm_is_operating_systems data source.

This rule validates `ibm_is_operating_systems` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_is_operating_systems_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `operating_systems` | no | — |

## Example

### Valid

```hcl
data "ibm_is_operating_systems" "test" {
}
```
