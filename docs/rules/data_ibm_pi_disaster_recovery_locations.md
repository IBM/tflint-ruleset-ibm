# ibm_pi_disaster_recovery_locations_data

Schema validation for the ibm_pi_disaster_recovery_locations data source.

This rule validates `ibm_pi_disaster_recovery_locations` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pi_disaster_recovery_locations_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `disaster_recovery_locations` | no | — |

## Example

### Valid

```hcl
data "ibm_pi_disaster_recovery_locations" "test" {
}
```
