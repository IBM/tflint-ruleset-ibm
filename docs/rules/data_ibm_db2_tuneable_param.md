# ibm_db2_tuneable_param_data

Schema validation for the ibm_db2_tuneable_param data source.

This rule validates `ibm_db2_tuneable_param` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_db2_tuneable_param_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `tuneable_param` | no | — |

## Example

### Valid

```hcl
data "ibm_db2_tuneable_param" "test" {
}
```
