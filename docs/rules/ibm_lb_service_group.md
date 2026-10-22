# ibm_lb_service_group

Schema validation for the ibm_lb_service_group resource.

This rule validates `ibm_lb_service_group` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_lb_service_group" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `allocation` | yes | — |
| `load_balancer_id` | yes | — |
| `port` | yes | range 1–65535; port format |
| `routing_method` | yes | — |
| `routing_type` | yes | — |
| `service_group_id` | no | — |
| `tags` | no | — |
| `timeout` | no | — |
| `virtual_server_id` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_lb_service_group_invalid_allocation` | `allocation` |
| `ibm_lb_service_group_invalid_load_balancer_id` | `load_balancer_id` |
| `ibm_lb_service_group_invalid_port` | `port` |
| `ibm_lb_service_group_invalid_routing_method` | `routing_method` |
| `ibm_lb_service_group_invalid_routing_type` | `routing_type` |
| `ibm_lb_service_group_invalid_tags` | `tags` |
| `ibm_lb_service_group_invalid_timeout` | `timeout` |

## Example

### Valid

```hcl
resource "ibm_lb_service_group" "test" {
  allocation = 1
  load_balancer_id = 1
  port = 80
  routing_method = "example"
  routing_type = "example"
}
```
