# ibm_is_vpn_gateway_service_connection_data

Schema validation for the ibm_is_vpn_gateway_service_connection data source.

This rule validates `ibm_is_vpn_gateway_service_connection` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_is_vpn_gateway_service_connection_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `created_at` | no | — |
| `creator` | no | — |
| `id` | no | — |
| `lifecycle_reasons` | no | — |
| `lifecycle_state` | no | — |
| `status` | no | — |
| `status_reasons` | no | — |
| `vpn_gateway` | no | — |
| `vpn_gateway_name` | no | — |
| `vpn_gateway_service_connection` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_vpn_gateway_service_connection_data_invalid_vpn_gateway` | `vpn_gateway` |
| `ibm_is_vpn_gateway_service_connection_data_invalid_vpn_gateway_name` | `vpn_gateway_name` |
| `ibm_is_vpn_gateway_service_connection_data_invalid_vpn_gateway_service_connection` | `vpn_gateway_service_connection` |

## Example

### Valid

```hcl
data "ibm_is_vpn_gateway_service_connection" "test" {
  vpn_gateway_service_connection = "example"
}
```
