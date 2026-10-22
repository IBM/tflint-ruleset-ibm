# ibm_is_vpn_gateway_data

Reads ibm_is_vpn_gateway. The user-defined name for this VPN gateway.

This rule validates `ibm_is_vpn_gateway` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_is_vpn_gateway_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_tags` | no | — |
| `advertised_cidrs` | no | — |
| `connections` | no | — |
| `created_at` | no | — |
| `crn` | no | — |
| `health_reasons` | no | — |
| `health_state` | no | — |
| `href` | no | — |
| `lifecycle_reasons` | no | — |
| `lifecycle_state` | no | — |
| `local_asn` | no | — |
| `members` | no | — |
| `mode` | no | — |
| `name` | no | — |
| `resource_group` | no | — |
| `resource_type` | no | — |
| `status` | no | — |
| `subnet` | no | — |
| `tags` | no | — |
| `vpc` | no | — |
| `vpn_gateway` | no | — |
| `vpn_gateway_name` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_vpn_gateway_data_invalid_vpn_gateway` | `vpn_gateway` |
| `ibm_is_vpn_gateway_data_invalid_vpn_gateway_name` | `vpn_gateway_name` |

## Example

### Valid

```hcl
data "ibm_is_vpn_gateway" "test" {
}
```
