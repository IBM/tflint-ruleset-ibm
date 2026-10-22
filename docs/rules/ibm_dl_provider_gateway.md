# ibm_dl_provider_gateway

Manages ibm_dl_provider_gateway. The unique user-defined name for this gateway

This rule validates `ibm_dl_provider_gateway` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `GatewayTemplate` (`direct_link_provider_api`).

## Configuration

```hcl
rule "ibm_dl_provider_gateway" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `bgp_asn` | yes | API: BGP ASN. For a 2-byte range, enter a value between 1-64495 or 64999. For a 2-byte or 4-byte range, enter a value between 131072-4199999999. For a 4-byte range, enter a value between 4201000000-4201064511. |
| `bgp_cer_cidr` | no | cidr format |
| `bgp_ibm_asn` | no | — |
| `bgp_ibm_cidr` | no | cidr format |
| `bgp_status` | no | — |
| `created_at` | no | — |
| `crn` | no | — |
| `customer_account_id` | yes | pattern `^[0-9a-f]+$` |
| `name` | yes | pattern `^([a-zA-Z]|[a-zA-Z][-_a-zA-Z0-9]*[a-zA-Z0-9])$` |
| `operational_status` | no | — |
| `port` | yes | API: Select Port Label for the gateway |
| `provider_api_managed` | no | — |
| `resource_controller_url` | no | — |
| `resource_crn` | no | — |
| `resource_group_name` | no | — |
| `resource_name` | no | — |
| `resource_status` | no | — |
| `speed_mbps` | yes | API: Gateway speed in megabits per second |
| `tags` | no | pattern `^[A-Za-z0-9:_ .-]+$` |
| `type` | no | — |
| `vlan` | no | API: VLAN requested for this gateway. VLAN provided should be in the range 2 to 3967. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_dl_provider_gateway_invalid_bgp_asn` | `bgp_asn` |
| `ibm_dl_provider_gateway_invalid_bgp_cer_cidr` | `bgp_cer_cidr` |
| `ibm_dl_provider_gateway_invalid_bgp_ibm_cidr` | `bgp_ibm_cidr` |
| `ibm_dl_provider_gateway_invalid_customer_account_id` | `customer_account_id` |
| `ibm_dl_provider_gateway_invalid_name` | `name` |
| `ibm_dl_provider_gateway_invalid_port` | `port` |
| `ibm_dl_provider_gateway_invalid_speed_mbps` | `speed_mbps` |
| `ibm_dl_provider_gateway_invalid_tags` | `tags` |
| `ibm_dl_provider_gateway_invalid_vlan` | `vlan` |

## Example

### Valid

```hcl
resource "ibm_dl_provider_gateway" "test" {
  bgp_asn = 1
  customer_account_id = "1"
  name = "example-name"
  port = "example"
  speed_mbps = 1
}
```
