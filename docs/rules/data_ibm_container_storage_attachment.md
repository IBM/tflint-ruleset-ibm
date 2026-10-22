# ibm_container_storage_attachment_data

Schema validation for the ibm_container_storage_attachment data source.

This rule validates `ibm_container_storage_attachment` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/v2/storage/createAttachment:post` (`kubernetes-containers-v1-v2`).

## Configuration

```hcl
rule "ibm_container_storage_attachment_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cluster` | yes | — |
| `resource_group_id` | no | — |
| `status` | no | — |
| `volume` | no | — |
| `volume_attachment_id` | yes | — |
| `volume_attachment_name` | no | — |
| `volume_type` | no | — |
| `worker` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_container_storage_attachment_data_invalid_cluster` | `cluster` |
| `ibm_container_storage_attachment_data_invalid_resource_group_id` | `resource_group_id` |
| `ibm_container_storage_attachment_data_invalid_volume_attachment_id` | `volume_attachment_id` |
| `ibm_container_storage_attachment_data_invalid_worker` | `worker` |

## Example

### Valid

```hcl
data "ibm_container_storage_attachment" "test" {
  cluster = "aaaaaaa"
  volume_attachment_id = "01234567-89ab-cdef-0123-456789abcdef"
  worker = "aaaaaaa"
}
```
