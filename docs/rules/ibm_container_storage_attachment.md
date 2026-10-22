# ibm_container_storage_attachment

Schema validation for the ibm_container_storage_attachment resource.

This rule validates `ibm_container_storage_attachment` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/v2/storage/createAttachment:post` (`kubernetes-containers-v1-v2`).

## Configuration

```hcl
rule "ibm_container_storage_attachment" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cluster` | yes | — |
| `resource_group_id` | no | — |
| `status` | no | — |
| `volume` | yes | — |
| `volume_attachment_id` | no | — |
| `volume_attachment_name` | no | — |
| `volume_type` | no | — |
| `worker` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_container_storage_attachment_invalid_cluster` | `cluster` |
| `ibm_container_storage_attachment_invalid_resource_group_id` | `resource_group_id` |
| `ibm_container_storage_attachment_invalid_volume` | `volume` |
| `ibm_container_storage_attachment_invalid_worker` | `worker` |

## Example

### Valid

```hcl
resource "ibm_container_storage_attachment" "test" {
  cluster = "aaaaaaa"
  volume = "example"
  worker = "aaaaaaa"
}
```
