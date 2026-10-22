# ibm_is_regions_data

Schema validation for the ibm_is_regions data source.

This rule validates `ibm_is_regions` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_is_regions_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `regions` | no | — |

## Example

### Valid

```hcl
data "ibm_is_regions" "test" {
}
```
