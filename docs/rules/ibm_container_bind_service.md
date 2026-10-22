# ibm_container_bind_service

Schema validation for the ibm_container_bind_service resource.

This rule validates `ibm_container_bind_service` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ServiceBindRequest` (`kubernetes-containers-v1-v2`).

## Configuration

```hcl
rule "ibm_container_bind_service" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_guid` | no | — |
| `cluster_name_id` | yes | — |
| `key` | no | API: The key to select. |
| `namespace_id` | yes | — |
| `org_guid` | no | — |
| `region` | no | API: The region value. |
| `resource_group_id` | no | — |
| `role` | no | API: The role value. |
| `service_instance_id` | no | — |
| `service_instance_name` | no | — |
| `space_guid` | no | — |
| `tags` | no | API: The tags value. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_container_bind_service_invalid_account_guid` | `account_guid` |
| `ibm_container_bind_service_invalid_cluster_name_id` | `cluster_name_id` |
| `ibm_container_bind_service_invalid_key` | `key` |
| `ibm_container_bind_service_invalid_namespace_id` | `namespace_id` |
| `ibm_container_bind_service_invalid_org_guid` | `org_guid` |
| `ibm_container_bind_service_invalid_region` | `region` |
| `ibm_container_bind_service_invalid_resource_group_id` | `resource_group_id` |
| `ibm_container_bind_service_invalid_role` | `role` |
| `ibm_container_bind_service_invalid_service_instance_id` | `service_instance_id` |
| `ibm_container_bind_service_invalid_service_instance_name` | `service_instance_name` |
| `ibm_container_bind_service_invalid_space_guid` | `space_guid` |
| `ibm_container_bind_service_invalid_tags` | `tags` |

## Example

### Valid

```hcl
resource "ibm_container_bind_service" "test" {
  cluster_name_id = "01234567-89ab-cdef-0123-456789abcdef"
  namespace_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
