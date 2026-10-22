# ibm_satellite_host

Schema validation for the ibm_satellite_host resource.

This rule validates `ibm_satellite_host` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_satellite_host" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cluster` | no | — |
| `host_id` | yes | — |
| `host_provider` | no | — |
| `host_state` | no | — |
| `labels` | no | — |
| `location` | yes | — |
| `wait_till` | no | one of: `location_normal` |
| `worker_pool` | no | — |
| `zone` | no | zone format |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_satellite_host_invalid_cluster` | `cluster` |
| `ibm_satellite_host_invalid_host_id` | `host_id` |
| `ibm_satellite_host_invalid_host_provider` | `host_provider` |
| `ibm_satellite_host_invalid_labels` | `labels` |
| `ibm_satellite_host_invalid_location` | `location` |
| `ibm_satellite_host_invalid_wait_till` | `wait_till` |
| `ibm_satellite_host_invalid_worker_pool` | `worker_pool` |
| `ibm_satellite_host_invalid_zone` | `zone` |

## Example

### Valid

```hcl
resource "ibm_satellite_host" "test" {
  host_id = "01234567-89ab-cdef-0123-456789abcdef"
  location = "example"
}
```
