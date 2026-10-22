# ibm_lb_vpx_service

Manages ibm_lb_vpx_service. name

This rule validates `ibm_lb_vpx_service` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_lb_vpx_service" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `connection_limit` | yes | — |
| `destination_ip_address` | yes | — |
| `destination_port` | yes | range 1–65535; port format |
| `health_check` | yes | — |
| `name` | yes | — |
| `tags` | no | — |
| `usip` | no | — |
| `vip_id` | yes | — |
| `weight` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_lb_vpx_service_invalid_connection_limit` | `connection_limit` |
| `ibm_lb_vpx_service_invalid_destination_ip_address` | `destination_ip_address` |
| `ibm_lb_vpx_service_invalid_destination_port` | `destination_port` |
| `ibm_lb_vpx_service_invalid_health_check` | `health_check` |
| `ibm_lb_vpx_service_invalid_name` | `name` |
| `ibm_lb_vpx_service_invalid_tags` | `tags` |
| `ibm_lb_vpx_service_invalid_usip` | `usip` |
| `ibm_lb_vpx_service_invalid_vip_id` | `vip_id` |
| `ibm_lb_vpx_service_invalid_weight` | `weight` |

## Example

### Valid

```hcl
resource "ibm_lb_vpx_service" "test" {
  connection_limit = 1
  destination_ip_address = "example"
  destination_port = 80
  health_check = "example"
  name = "example"
  vip_id = "01234567-89ab-cdef-0123-456789abcdef"
  weight = 1
}
```
