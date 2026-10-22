# ibm_is_vpn_gateway_advertised_cidrs_data

Schema validation for the ibm_is_vpn_gateway_advertised_cidrs data source.

This rule validates `ibm_is_vpn_gateway_advertised_cidrs` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_is_vpn_gateway_advertised_cidrs_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `advertised_cidrs` | no | — |
| `vpn_gateway` | no | — |
| `vpn_gateway_name` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_vpn_gateway_advertised_cidrs_data_invalid_vpn_gateway` | `vpn_gateway` |
| `ibm_is_vpn_gateway_advertised_cidrs_data_invalid_vpn_gateway_name` | `vpn_gateway_name` |

## Example

### Valid

```hcl
data "ibm_is_vpn_gateway_advertised_cidrs" "test" {
}
```
