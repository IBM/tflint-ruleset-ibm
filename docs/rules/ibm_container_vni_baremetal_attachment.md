# ibm_container_vni_baremetal_attachment

Schema validation for the ibm_container_vni_baremetal_attachment resource.

This rule validates `ibm_container_vni_baremetal_attachment` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/v2/storage/createAttachment:post` (`kubernetes-containers-v1-v2`).

## Configuration

```hcl
rule "ibm_container_vni_baremetal_attachment" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `auto_delete` | no | — |
| `cluster` | no | — |
| `created_at` | no | — |
| `resource_group_id` | no | — |
| `status` | no | — |
| `vlan_id` | yes | API: The private VLAN ID that the subnet CIDR belongs to. |
| `vni_id` | yes | — |
| `worker` | no | — |
| `worker_id` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_container_vni_baremetal_attachment_invalid_auto_delete` | `auto_delete` |
| `ibm_container_vni_baremetal_attachment_invalid_cluster` | `cluster` |
| `ibm_container_vni_baremetal_attachment_invalid_resource_group_id` | `resource_group_id` |
| `ibm_container_vni_baremetal_attachment_invalid_vlan_id` | `vlan_id` |
| `ibm_container_vni_baremetal_attachment_invalid_vni_id` | `vni_id` |
| `ibm_container_vni_baremetal_attachment_invalid_worker` | `worker` |

## Example

### Valid

```hcl
resource "ibm_container_vni_baremetal_attachment" "test" {
  vlan_id = 1
  vni_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
