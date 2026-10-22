# ibm_container_vni_attachment_data

Schema validation for the ibm_container_vni_attachment data source.

This rule validates `ibm_container_vni_attachment` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/v2/storage/createAttachment:post` (`kubernetes-containers-v1-v2`).

## Configuration

```hcl
rule "ibm_container_vni_attachment_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `auto_delete` | no | — |
| `cluster` | yes | — |
| `created_at` | no | — |
| `mac_address` | no | — |
| `primary_ip_address` | no | — |
| `resource_group_id` | no | — |
| `status` | no | — |
| `vlan_id` | no | — |
| `vni_id` | yes | — |
| `vni_name` | no | — |
| `worker` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_container_vni_attachment_data_invalid_cluster` | `cluster` |
| `ibm_container_vni_attachment_data_invalid_resource_group_id` | `resource_group_id` |
| `ibm_container_vni_attachment_data_invalid_vni_id` | `vni_id` |
| `ibm_container_vni_attachment_data_invalid_worker` | `worker` |

## Example

### Valid

```hcl
data "ibm_container_vni_attachment" "test" {
  cluster = "aaaaaaa"
  vni_id = "01234567-89ab-cdef-0123-456789abcdef"
  worker = "aaaaaaa"
}
```
