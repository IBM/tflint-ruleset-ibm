# ibm_compute_dedicated_host

Schema validation for the ibm_compute_dedicated_host resource.

This rule validates `ibm_compute_dedicated_host` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_compute_dedicated_host" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cpu_count` | no | — |
| `datacenter` | yes | — |
| `disk_capacity` | no | — |
| `domain` | yes | — |
| `flavor` | no | — |
| `hostname` | yes | — |
| `hourly_billing` | no | — |
| `memory_capacity` | no | — |
| `router_hostname` | yes | — |
| `tags` | no | — |
| `wait_time_minutes` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_compute_dedicated_host_invalid_datacenter` | `datacenter` |
| `ibm_compute_dedicated_host_invalid_domain` | `domain` |
| `ibm_compute_dedicated_host_invalid_flavor` | `flavor` |
| `ibm_compute_dedicated_host_invalid_hostname` | `hostname` |
| `ibm_compute_dedicated_host_invalid_hourly_billing` | `hourly_billing` |
| `ibm_compute_dedicated_host_invalid_router_hostname` | `router_hostname` |
| `ibm_compute_dedicated_host_invalid_tags` | `tags` |
| `ibm_compute_dedicated_host_invalid_wait_time_minutes` | `wait_time_minutes` |

## Example

### Valid

```hcl
resource "ibm_compute_dedicated_host" "test" {
  datacenter = "example"
  domain = "example"
  hostname = "example"
  router_hostname = "example"
}
```
