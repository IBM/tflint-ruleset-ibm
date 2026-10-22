# ibm_satellite_location_nlb_dns

Schema validation for the ibm_satellite_location_nlb_dns resource.

This rule validates `ibm_satellite_location_nlb_dns` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_satellite_location_nlb_dns" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `ips` | yes | — |
| `location` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_satellite_location_nlb_dns_invalid_ips` | `ips` |
| `ibm_satellite_location_nlb_dns_invalid_location` | `location` |

## Example

### Valid

```hcl
resource "ibm_satellite_location_nlb_dns" "test" {
  ips = "example"
  location = "example"
}
```
