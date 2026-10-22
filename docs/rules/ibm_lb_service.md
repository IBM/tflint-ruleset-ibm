# ibm_lb_service

Schema validation for the ibm_lb_service resource.

This rule validates `ibm_lb_service` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_lb_service" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `enabled` | yes | — |
| `health_check_type` | yes | — |
| `ip_address_id` | yes | — |
| `port` | yes | range 1–65535; port format |
| `service_group_id` | yes | — |
| `tags` | no | — |
| `weight` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_lb_service_invalid_enabled` | `enabled` |
| `ibm_lb_service_invalid_health_check_type` | `health_check_type` |
| `ibm_lb_service_invalid_ip_address_id` | `ip_address_id` |
| `ibm_lb_service_invalid_port` | `port` |
| `ibm_lb_service_invalid_service_group_id` | `service_group_id` |
| `ibm_lb_service_invalid_tags` | `tags` |
| `ibm_lb_service_invalid_weight` | `weight` |

## Example

### Valid

```hcl
resource "ibm_lb_service" "test" {
  enabled = true
  health_check_type = "example"
  ip_address_id = 1
  port = 80
  service_group_id = 1
  weight = 1
}
```
