# ibm_container_vpc_alb_create

Manages ibm_container_vpc_alb_create. ALB name

This rule validates `ibm_container_vpc_alb_create` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ALB` (`kubernetes-containers-v1-v2`).

## Configuration

```hcl
rule "ibm_container_vpc_alb_create" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `alb_id` | no | — |
| `alb_type` | no | — |
| `cluster` | yes | API: The cluster value. |
| `disable_deployment` | no | — |
| `enable` | no | — |
| `load_balancer_hostname` | no | — |
| `name` | no | — |
| `resize` | no | — |
| `resource_group_id` | no | — |
| `state` | no | — |
| `status` | no | — |
| `type` | yes | API: The type of ALB that you want to create. |
| `zone` | yes | zone format |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_container_vpc_alb_create_invalid_cluster` | `cluster` |
| `ibm_container_vpc_alb_create_invalid_enable` | `enable` |
| `ibm_container_vpc_alb_create_invalid_resource_group_id` | `resource_group_id` |
| `ibm_container_vpc_alb_create_invalid_type` | `type` |
| `ibm_container_vpc_alb_create_invalid_zone` | `zone` |

## Example

### Valid

```hcl
resource "ibm_container_vpc_alb_create" "test" {
  cluster = "aaaaaaa"
  type = "aaaaaaa"
  zone = "us-south-1"
}
```
