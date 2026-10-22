# ibm_is_vpn_gateway_connection_peer_cidrs_data

Schema validation for the ibm_is_vpn_gateway_connection_peer_cidrs data source.

This rule validates `ibm_is_vpn_gateway_connection_peer_cidrs` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_is_vpn_gateway_connection_peer_cidrs_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cidrs` | no | — |
| `vpn_gateway` | yes | — |
| `vpn_gateway_connection` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_vpn_gateway_connection_peer_cidrs_data_invalid_vpn_gateway` | `vpn_gateway` |
| `ibm_is_vpn_gateway_connection_peer_cidrs_data_invalid_vpn_gateway_connection` | `vpn_gateway_connection` |

## Example

### Valid

```hcl
data "ibm_is_vpn_gateway_connection_peer_cidrs" "test" {
  vpn_gateway = "example"
  vpn_gateway_connection = "example"
}
```
