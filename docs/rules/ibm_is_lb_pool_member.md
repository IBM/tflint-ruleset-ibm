# ibm_is_lb_pool_member

Schema validation for the ibm_is_lb_pool_member resource.

This rule validates `ibm_is_lb_pool_member` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `LoadBalancerPoolMemberPrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_lb_pool_member" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `health` | no | — |
| `href` | no | — |
| `lb` | yes | — |
| `pool` | yes | — |
| `port` | yes | range 1–65535; port format |
| `provisioning_status` | no | — |
| `related_crn` | no | — |
| `target_address` | no | — |
| `target_fqdn` | no | — |
| `target_id` | no | — |
| `weight` | no | range 0–100 |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_lb_pool_member_invalid_lb` | `lb` |
| `ibm_is_lb_pool_member_invalid_pool` | `pool` |
| `ibm_is_lb_pool_member_invalid_port` | `port` |
| `ibm_is_lb_pool_member_invalid_target_address` | `target_address` |
| `ibm_is_lb_pool_member_invalid_target_fqdn` | `target_fqdn` |
| `ibm_is_lb_pool_member_invalid_target_id` | `target_id` |
| `ibm_is_lb_pool_member_invalid_weight` | `weight` |

## Example

### Valid

```hcl
resource "ibm_is_lb_pool_member" "test" {
  lb = "example"
  pool = "example"
  port = 80
}
```
