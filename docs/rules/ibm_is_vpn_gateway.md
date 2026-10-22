# ibm_is_vpn_gateway

Manages ibm_is_vpn_gateway. VPN Gateway instance name

This rule validates `ibm_is_vpn_gateway` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `VPNGatewayPrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_vpn_gateway" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_tags` | no | pattern `^([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-]):([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-])$` |
| `advertised_cidrs` | no | — |
| `created_at` | no | — |
| `crn` | no | — |
| `health_reasons` | no | — |
| `health_state` | no | — |
| `lifecycle_reasons` | no | — |
| `lifecycle_state` | no | — |
| `local_asn` | no | range 64512–65534 |
| `members` | no | — |
| `mode` | no | one of: `route`, `policy` |
| `name` | yes | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `private_ip_address` | no | — |
| `private_ip_address2` | no | — |
| `public_ip_address` | no | — |
| `public_ip_address2` | no | — |
| `resource_controller_url` | no | — |
| `resource_crn` | no | — |
| `resource_group` | no | API: The resource group to use. If unspecified, the account's [default resource group](https://cloud.ibm.com/apidocs/resource-manager#introduction) will be used. |
| `resource_group_name` | no | — |
| `resource_name` | no | — |
| `resource_status` | no | — |
| `status` | no | — |
| `subnet` | yes | API: Identifies a subnet by a unique property. |
| `tags` | no | pattern `^[A-Za-z0-9:_ .-]+$` |
| `vpc` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_vpn_gateway_invalid_access_tags` | `access_tags` |
| `ibm_is_vpn_gateway_invalid_local_asn` | `local_asn` |
| `ibm_is_vpn_gateway_invalid_mode` | `mode` |
| `ibm_is_vpn_gateway_invalid_name` | `name` |
| `ibm_is_vpn_gateway_invalid_resource_group` | `resource_group` |
| `ibm_is_vpn_gateway_invalid_subnet` | `subnet` |
| `ibm_is_vpn_gateway_invalid_tags` | `tags` |

## Example

### Valid

```hcl
resource "ibm_is_vpn_gateway" "test" {
  name = "example-name"
  subnet = "example"
}
```
