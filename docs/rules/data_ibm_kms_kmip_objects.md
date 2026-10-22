# ibm_kms_kmip_objects_data

Schema validation for the ibm_kms_kmip_objects data source.

This rule validates `ibm_kms_kmip_objects` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `leaf-unique` (`hs-crypto`).

## Configuration

```hcl
rule "ibm_kms_kmip_objects_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `adapter_id` | no | — |
| `adapter_name` | no | — |
| `endpoint_type` | no | one of: `public`, `private` |
| `instance_id` | yes | API: The v4 UUID used to uniquely identify the instance in IBM Cloud. |
| `limit` | no | range 1–1000 |
| `object_state_filter` | no | — |
| `objects` | no | — |
| `offset` | no | range 0–2147483647 |
| `show_total_count` | no | — |
| `total_count` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_kms_kmip_objects_data_invalid_adapter_id` | `adapter_id` |
| `ibm_kms_kmip_objects_data_invalid_adapter_name` | `adapter_name` |
| `ibm_kms_kmip_objects_data_invalid_endpoint_type` | `endpoint_type` |
| `ibm_kms_kmip_objects_data_invalid_instance_id` | `instance_id` |
| `ibm_kms_kmip_objects_data_invalid_limit` | `limit` |
| `ibm_kms_kmip_objects_data_invalid_object_state_filter` | `object_state_filter` |
| `ibm_kms_kmip_objects_data_invalid_offset` | `offset` |
| `ibm_kms_kmip_objects_data_invalid_show_total_count` | `show_total_count` |

## Example

### Valid

```hcl
data "ibm_kms_kmip_objects" "test" {
  instance_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
