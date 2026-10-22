# ibm_tg_gateways_data

Schema validation for the ibm_tg_gateways data source.

This rule validates `ibm_tg_gateways` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_tg_gateways_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `transit_gateways` | no | — |

## Example

### Valid

```hcl
data "ibm_tg_gateways" "test" {
}
```
