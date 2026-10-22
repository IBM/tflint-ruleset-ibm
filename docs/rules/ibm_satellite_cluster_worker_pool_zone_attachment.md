# ibm_satellite_cluster_worker_pool_zone_attachment

Schema validation for the ibm_satellite_cluster_worker_pool_zone_attachment resource.

This rule validates `ibm_satellite_cluster_worker_pool_zone_attachment` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_satellite_cluster_worker_pool_zone_attachment" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `autobalance_enabled` | no | — |
| `cluster` | yes | — |
| `messages` | no | — |
| `resource_group_id` | no | — |
| `worker_count` | no | — |
| `worker_pool` | yes | — |
| `zone` | yes | zone format |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_satellite_cluster_worker_pool_zone_attachment_invalid_cluster` | `cluster` |
| `ibm_satellite_cluster_worker_pool_zone_attachment_invalid_resource_group_id` | `resource_group_id` |
| `ibm_satellite_cluster_worker_pool_zone_attachment_invalid_worker_pool` | `worker_pool` |
| `ibm_satellite_cluster_worker_pool_zone_attachment_invalid_zone` | `zone` |

## Example

### Valid

```hcl
resource "ibm_satellite_cluster_worker_pool_zone_attachment" "test" {
  cluster = "example"
  worker_pool = "example"
  zone = "us-south-1"
}
```
