# ibm_is_ike_policies_data

Schema validation for the ibm_is_ike_policies data source.

This rule validates `ibm_is_ike_policies` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_is_ike_policies_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `ike_policies` | no | — |

## Example

### Valid

```hcl
data "ibm_is_ike_policies" "test" {
}
```
