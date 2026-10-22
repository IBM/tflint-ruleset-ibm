# ibm_satellite_location_nlb_dns_data

Schema validation for the ibm_satellite_location_nlb_dns data source.

This rule validates `ibm_satellite_location_nlb_dns` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_satellite_location_nlb_dns_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `location` | yes | — |
| `nlb_config` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_satellite_location_nlb_dns_data_invalid_location` | `location` |

## Example

### Valid

```hcl
data "ibm_satellite_location_nlb_dns" "test" {
  location = "example"
}
```
