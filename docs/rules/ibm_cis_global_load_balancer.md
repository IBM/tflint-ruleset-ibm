# ibm_cis_global_load_balancer

Manages ibm_cis_global_load_balancer. name

This rule validates `ibm_cis_global_load_balancer` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `load_balancer_req` (`cis`).

## Configuration

```hcl
rule "ibm_cis_global_load_balancer" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cis_id` | yes | — |
| `created_on` | no | — |
| `default_pool_ids` | yes | — |
| `description` | no | API: Desc |
| `domain_id` | yes | — |
| `enabled` | no | API: Enabled/disabled |
| `fallback_pool_id` | yes | — |
| `glb_id` | no | — |
| `modified_on` | no | — |
| `name` | yes | API: Name |
| `proxied` | no | API: Proxied |
| `session_affinity` | no | one of: `none`, `cookie` |
| `steering_policy` | no | one of: `off`, `geo`, `random`, `dynamic_latency` |
| `ttl` | no | API: TTL |
| `pop_pools` (block) | no | nested block |
| `pop_pools.pool_ids` | yes | — |
| `pop_pools.pop` | yes | — |
| `region_pools` (block) | no | nested block |
| `region_pools.pool_ids` | yes | — |
| `region_pools.region` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cis_global_load_balancer_invalid_cis_id` | `cis_id` |
| `ibm_cis_global_load_balancer_invalid_default_pool_ids` | `default_pool_ids` |
| `ibm_cis_global_load_balancer_invalid_description` | `description` |
| `ibm_cis_global_load_balancer_invalid_domain_id` | `domain_id` |
| `ibm_cis_global_load_balancer_invalid_enabled` | `enabled` |
| `ibm_cis_global_load_balancer_invalid_fallback_pool_id` | `fallback_pool_id` |
| `ibm_cis_global_load_balancer_invalid_name` | `name` |
| `ibm_cis_global_load_balancer_invalid_proxied` | `proxied` |
| `ibm_cis_global_load_balancer_invalid_session_affinity` | `session_affinity` |
| `ibm_cis_global_load_balancer_invalid_steering_policy` | `steering_policy` |
| `ibm_cis_global_load_balancer_invalid_ttl` | `ttl` |
| `ibm_cis_global_load_balancer_invalid_pop_pools__pool_ids` | `pop_pools.pool_ids` |
| `ibm_cis_global_load_balancer_invalid_pop_pools__pop` | `pop_pools.pop` |
| `ibm_cis_global_load_balancer_invalid_region_pools__pool_ids` | `region_pools.pool_ids` |
| `ibm_cis_global_load_balancer_invalid_region_pools__region` | `region_pools.region` |

## Example

### Valid

```hcl
resource "ibm_cis_global_load_balancer" "test" {
  cis_id = "01234567-89ab-cdef-0123-456789abcdef"
  default_pool_ids = "example"
  domain_id = "01234567-89ab-cdef-0123-456789abcdef"
  fallback_pool_id = "01234567-89ab-cdef-0123-456789abcdef"
  name = "example"
}
```
