# ibm_is_vpn_gateway_service_connections_data

Schema validation for the ibm_is_vpn_gateway_service_connections data source.

This rule validates `ibm_is_vpn_gateway_service_connections` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_is_vpn_gateway_service_connections_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `service_connections` | no | — |
| `vpn_gateway` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_vpn_gateway_service_connections_data_invalid_vpn_gateway` | `vpn_gateway` |

## Example

### Valid

```hcl
data "ibm_is_vpn_gateway_service_connections" "test" {
  vpn_gateway = "example"
}
```
