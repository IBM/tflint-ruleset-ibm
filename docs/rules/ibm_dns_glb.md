# ibm_dns_glb

Manages ibm_dns_glb. Name of the load balancer

This rule validates `ibm_dns_glb` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/instances/{instance_id}/dnszones/{dnszone_id}/load_balancers:post` (`dns-svcs`).

## Configuration

```hcl
rule "ibm_dns_glb" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `created_on` | no | — |
| `default_pools` | yes | API: A list of pool IDs ordered by their failover priority. Pools defined here are used by default, or when region_pools are not configured for a given region. |
| `description` | no | API: Descriptive text of the load balancer. |
| `enabled` | no | API: Whether the load balancer is enabled. |
| `fallback_pool` | yes | API: The pool ID to use when all other pools are detected as unhealthy. |
| `glb_id` | no | — |
| `health` | no | — |
| `instance_id` | yes | pattern `^[-0-9a-z_]+$` |
| `modified_on` | no | — |
| `name` | yes | API: Name of the load balancer. |
| `ttl` | no | API: Time to live in second. |
| `zone_id` | yes | pattern `^[-0-9a-z_]+$` |
| `az_pools` (block) | no | nested block |
| `az_pools.availability_zone` | yes | zone format |
| `az_pools.pools` | yes | API: List of load balancer pools. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_dns_glb_invalid_default_pools` | `default_pools` |
| `ibm_dns_glb_invalid_description` | `description` |
| `ibm_dns_glb_invalid_enabled` | `enabled` |
| `ibm_dns_glb_invalid_fallback_pool` | `fallback_pool` |
| `ibm_dns_glb_invalid_instance_id` | `instance_id` |
| `ibm_dns_glb_invalid_name` | `name` |
| `ibm_dns_glb_invalid_ttl` | `ttl` |
| `ibm_dns_glb_invalid_zone_id` | `zone_id` |
| `ibm_dns_glb_invalid_az_pools__availability_zone` | `az_pools.availability_zone` |
| `ibm_dns_glb_invalid_az_pools__pools` | `az_pools.pools` |

## Example

### Valid

```hcl
resource "ibm_dns_glb" "test" {
  default_pools = "example"
  fallback_pool = "example"
  instance_id = "example-namexxxxxxxxxxxxxxxxxxxxxxxx"
  name = "example"
  zone_id = "example-namexxxxxxxxxxxxxxxxxxxxxxxx"
}
```
