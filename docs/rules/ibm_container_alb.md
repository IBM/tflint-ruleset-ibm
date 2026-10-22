# ibm_container_alb

Manages ibm_container_alb. ALB name

This rule validates `ibm_container_alb` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ALB` (`kubernetes-containers-v1-v2`).

## Configuration

```hcl
rule "ibm_container_alb" {
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
| `name` | no | — |
| `region` | no | API: The region value. |
| `replicas` | no | — |
| `resize` | no | — |
| `user_ip` | no | — |
| `zone` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_container_alb_invalid_alb_id` | `alb_id` |
| `ibm_container_alb_invalid_disable_deployment` | `disable_deployment` |
| `ibm_container_alb_invalid_enable` | `enable` |
| `ibm_container_alb_invalid_region` | `region` |
| `ibm_container_alb_invalid_user_ip` | `user_ip` |

## Example

### Valid

```hcl
resource "ibm_container_alb" "test" {
  alb_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
