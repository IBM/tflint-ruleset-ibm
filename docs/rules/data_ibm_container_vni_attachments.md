# ibm_container_vni_attachments_data

Schema validation for the ibm_container_vni_attachments data source.

This rule validates `ibm_container_vni_attachments` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/v2/storage/createAttachment:post` (`kubernetes-containers-v1-v2`).

## Configuration

```hcl
rule "ibm_container_vni_attachments_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `attachments` | no | — |
| `cluster` | no | — |
| `resource_group_id` | no | — |
| `worker` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_container_vni_attachments_data_invalid_cluster` | `cluster` |
| `ibm_container_vni_attachments_data_invalid_resource_group_id` | `resource_group_id` |
| `ibm_container_vni_attachments_data_invalid_worker` | `worker` |

## Example

### Valid

```hcl
data "ibm_container_vni_attachments" "test" {
}
```
