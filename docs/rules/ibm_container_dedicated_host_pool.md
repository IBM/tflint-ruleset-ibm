# ibm_container_dedicated_host_pool

Manages ibm_container_dedicated_host_pool. The name of the dedicated host pool

This rule validates `ibm_container_dedicated_host_pool` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `CreateDedicatedHostPool contains the fields in requests to create dedicated host pool instances.` (`kubernetes-containers-v1-v2`).

## Configuration

```hcl
rule "ibm_container_dedicated_host_pool" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `flavor_class` | yes | API: The flavor class value. |
| `host_count` | no | — |
| `metro` | yes | API: The metro value. |
| `name` | yes | API: The name value. |
| `resource_group_id` | no | — |
| `state` | no | — |
| `worker_pools` | no | — |
| `zones` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_container_dedicated_host_pool_invalid_flavor_class` | `flavor_class` |
| `ibm_container_dedicated_host_pool_invalid_metro` | `metro` |
| `ibm_container_dedicated_host_pool_invalid_name` | `name` |
| `ibm_container_dedicated_host_pool_invalid_resource_group_id` | `resource_group_id` |

## Example

### Valid

```hcl
resource "ibm_container_dedicated_host_pool" "test" {
  flavor_class = "aaaaaaa"
  metro = "aaaaaaa"
  name = "aaaaaaa"
}
```
