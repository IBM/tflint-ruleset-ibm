# ibm_lb_vpx_vip

Manages ibm_lb_vpx_vip. Name

This rule validates `ibm_lb_vpx_vip` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_lb_vpx_vip" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `load_balancing_method` | yes | — |
| `nad_controller_id` | yes | — |
| `name` | yes | — |
| `persistence` | no | — |
| `security_certificate_id` | no | — |
| `source_port` | yes | range 1–65535; port format |
| `tags` | no | — |
| `type` | yes | — |
| `virtual_ip_address` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_lb_vpx_vip_invalid_load_balancing_method` | `load_balancing_method` |
| `ibm_lb_vpx_vip_invalid_nad_controller_id` | `nad_controller_id` |
| `ibm_lb_vpx_vip_invalid_name` | `name` |
| `ibm_lb_vpx_vip_invalid_persistence` | `persistence` |
| `ibm_lb_vpx_vip_invalid_security_certificate_id` | `security_certificate_id` |
| `ibm_lb_vpx_vip_invalid_source_port` | `source_port` |
| `ibm_lb_vpx_vip_invalid_tags` | `tags` |
| `ibm_lb_vpx_vip_invalid_type` | `type` |
| `ibm_lb_vpx_vip_invalid_virtual_ip_address` | `virtual_ip_address` |

## Example

### Valid

```hcl
resource "ibm_lb_vpx_vip" "test" {
  load_balancing_method = "example"
  nad_controller_id = 1
  name = "example"
  source_port = 80
  type = "example"
  virtual_ip_address = "example"
}
```
