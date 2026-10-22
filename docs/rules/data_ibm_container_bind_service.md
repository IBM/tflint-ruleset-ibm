# ibm_container_bind_service_data

Schema validation for the ibm_container_bind_service data source.

This rule validates `ibm_container_bind_service` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_container_bind_service_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cluster_name_id` | yes | — |
| `namespace_id` | yes | — |
| `service_instance_id` | no | — |
| `service_instance_name` | no | — |
| `service_key_name` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_container_bind_service_data_invalid_cluster_name_id` | `cluster_name_id` |
| `ibm_container_bind_service_data_invalid_namespace_id` | `namespace_id` |
| `ibm_container_bind_service_data_invalid_service_instance_id` | `service_instance_id` |
| `ibm_container_bind_service_data_invalid_service_instance_name` | `service_instance_name` |

## Example

### Valid

```hcl
data "ibm_container_bind_service" "test" {
  cluster_name_id = "01234567-89ab-cdef-0123-456789abcdef"
  namespace_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
