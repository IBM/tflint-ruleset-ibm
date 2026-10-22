# ibm_container_vpc_alb

Manages ibm_container_vpc_alb. ALB name

This rule validates `ibm_container_vpc_alb` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_container_vpc_alb" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `alb_id` | yes | — |
| `alb_type` | no | — |
| `cluster` | no | — |
| `disable_deployment` | no | — |
| `enable` | no | — |
| `load_balancer_hostname` | no | — |
| `name` | no | — |
| `resize` | no | — |
| `resource_group_id` | no | — |
| `state` | no | — |
| `status` | no | — |
| `zone` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_container_vpc_alb_invalid_alb_id` | `alb_id` |
| `ibm_container_vpc_alb_invalid_disable_deployment` | `disable_deployment` |
| `ibm_container_vpc_alb_invalid_enable` | `enable` |
| `ibm_container_vpc_alb_invalid_resource_group_id` | `resource_group_id` |

## Example

### Valid

```hcl
resource "ibm_container_vpc_alb" "test" {
  alb_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
