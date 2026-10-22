# ibm_container_alb_create

Manages ibm_container_alb_create. ALB name

This rule validates `ibm_container_alb_create` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ALB` (`kubernetes-containers-v1-v2`).

## Configuration

```hcl
rule "ibm_container_alb_create" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `alb_id` | no | — |
| `alb_type` | yes | — |
| `cluster` | yes | API: The cluster value. |
| `disable_deployment` | no | — |
| `enable` | no | — |
| `ingress_image` | no | — |
| `ip` | no | — |
| `name` | no | — |
| `nlb_version` | no | — |
| `replicas` | no | — |
| `resize` | no | — |
| `user_ip` | no | — |
| `vlan_id` | yes | API: The private VLAN ID that the subnet CIDR belongs to. |
| `zone` | yes | zone format |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_container_alb_create_invalid_alb_type` | `alb_type` |
| `ibm_container_alb_create_invalid_cluster` | `cluster` |
| `ibm_container_alb_create_invalid_enable` | `enable` |
| `ibm_container_alb_create_invalid_ingress_image` | `ingress_image` |
| `ibm_container_alb_create_invalid_ip` | `ip` |
| `ibm_container_alb_create_invalid_nlb_version` | `nlb_version` |
| `ibm_container_alb_create_invalid_vlan_id` | `vlan_id` |
| `ibm_container_alb_create_invalid_zone` | `zone` |

## Example

### Valid

```hcl
resource "ibm_container_alb_create" "test" {
  alb_type = "example"
  cluster = "aaaaaaa"
  vlan_id = "01234567-89ab-cdef-0123-456789abcdef"
  zone = "us-south-1"
}
```
