# ibm_container_dedicated_host

Schema validation for the ibm_container_dedicated_host resource.

This rule validates `ibm_container_dedicated_host` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `CreateDedicatedHost contains the fields in requests to create dedicated host instances.` (`kubernetes-containers-v1-v2`).

## Configuration

```hcl
rule "ibm_container_dedicated_host" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `flavor` | yes | API: The flavor value. |
| `host_id` | no | — |
| `host_pool_id` | yes | — |
| `life_cycle` | no | — |
| `placement_enabled` | no | — |
| `resources` | no | — |
| `workers` | no | — |
| `zone` | yes | zone format |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_container_dedicated_host_invalid_flavor` | `flavor` |
| `ibm_container_dedicated_host_invalid_host_pool_id` | `host_pool_id` |
| `ibm_container_dedicated_host_invalid_placement_enabled` | `placement_enabled` |
| `ibm_container_dedicated_host_invalid_zone` | `zone` |

## Example

### Valid

```hcl
resource "ibm_container_dedicated_host" "test" {
  flavor = "aaaaaaa"
  host_pool_id = "01234567-89ab-cdef-0123-456789abcdef"
  zone = "us-south-1"
}
```
