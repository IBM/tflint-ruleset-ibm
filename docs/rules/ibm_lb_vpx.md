# ibm_lb_vpx

Manages ibm_lb_vpx. Name

This rule validates `ibm_lb_vpx` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_lb_vpx" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `datacenter` | yes | — |
| `ip_count` | yes | — |
| `management_ip_address` | no | — |
| `name` | no | — |
| `plan` | yes | — |
| `private_subnet` | no | — |
| `private_vlan_id` | no | — |
| `public_subnet` | no | — |
| `public_vlan_id` | no | — |
| `speed` | yes | — |
| `tags` | no | — |
| `type` | no | — |
| `version` | yes | — |
| `vip_pool` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_lb_vpx_invalid_datacenter` | `datacenter` |
| `ibm_lb_vpx_invalid_ip_count` | `ip_count` |
| `ibm_lb_vpx_invalid_plan` | `plan` |
| `ibm_lb_vpx_invalid_private_subnet` | `private_subnet` |
| `ibm_lb_vpx_invalid_private_vlan_id` | `private_vlan_id` |
| `ibm_lb_vpx_invalid_public_subnet` | `public_subnet` |
| `ibm_lb_vpx_invalid_public_vlan_id` | `public_vlan_id` |
| `ibm_lb_vpx_invalid_speed` | `speed` |
| `ibm_lb_vpx_invalid_tags` | `tags` |
| `ibm_lb_vpx_invalid_version` | `version` |

## Example

### Valid

```hcl
resource "ibm_lb_vpx" "test" {
  datacenter = "example"
  ip_count = 1
  plan = "example"
  speed = 1
  version = "example"
}
```
