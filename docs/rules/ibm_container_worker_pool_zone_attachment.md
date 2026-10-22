# ibm_container_worker_pool_zone_attachment

Schema validation for the ibm_container_worker_pool_zone_attachment resource.

This rule validates `ibm_container_worker_pool_zone_attachment` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Worker` (`kubernetes-containers-v1-v2`).

## Configuration

```hcl
rule "ibm_container_worker_pool_zone_attachment" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cluster` | yes | API: The cluster value. |
| `private_vlan_id` | no | — |
| `public_vlan_id` | no | — |
| `region` | no | API: The region value. |
| `resource_group_id` | no | — |
| `wait_till_albs` | no | — |
| `worker_count` | no | — |
| `worker_pool` | yes | — |
| `zone` | yes | zone format |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_container_worker_pool_zone_attachment_invalid_cluster` | `cluster` |
| `ibm_container_worker_pool_zone_attachment_invalid_private_vlan_id` | `private_vlan_id` |
| `ibm_container_worker_pool_zone_attachment_invalid_public_vlan_id` | `public_vlan_id` |
| `ibm_container_worker_pool_zone_attachment_invalid_region` | `region` |
| `ibm_container_worker_pool_zone_attachment_invalid_resource_group_id` | `resource_group_id` |
| `ibm_container_worker_pool_zone_attachment_invalid_wait_till_albs` | `wait_till_albs` |
| `ibm_container_worker_pool_zone_attachment_invalid_worker_pool` | `worker_pool` |
| `ibm_container_worker_pool_zone_attachment_invalid_zone` | `zone` |

## Example

### Valid

```hcl
resource "ibm_container_worker_pool_zone_attachment" "test" {
  cluster = "aaaaaaa"
  worker_pool = "example"
  zone = "us-south-1"
}
```
