# ibm_tg_locations_data

Schema validation for the ibm_tg_locations data source.

This rule validates `ibm_tg_locations` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_tg_locations_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `locations` | no | — |

## Example

### Valid

```hcl
data "ibm_tg_locations" "test" {
}
```
