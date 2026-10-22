# ibm_code_engine_persistent_data_store

Manages ibm_code_engine_persistent_data_store. The name of the persistent data store.

This rule validates `ibm_code_engine_persistent_data_store` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ObjectStorageData` (`codeengine-v2`).

## Configuration

```hcl
rule "ibm_code_engine_persistent_data_store" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `created_at` | no | — |
| `entity_tag` | no | — |
| `etag` | no | — |
| `name` | yes | pattern `^[a-z0-9]([\-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([\-a-z0-9]*[a-z0-9])?)*$` |
| `persistent_data_store_id` | no | — |
| `project_id` | yes | pattern `^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$` |
| `region` | no | — |
| `storage_type` | yes | one of: `object_storage` |
| `data` (block) | no | nested block |
| `data.bucket_location` | no | pattern `^[a-zA-Z0-9-]+$` |
| `data.bucket_name` | no | pattern `^[a-z0-9]([-a-z0-9]*[a-z0-9])?$` |
| `data.secret_name` | no | pattern `^[a-z0-9]([\-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([\-a-z0-9]*[a-z0-9])?)*$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_code_engine_persistent_data_store_invalid_name` | `name` |
| `ibm_code_engine_persistent_data_store_invalid_project_id` | `project_id` |
| `ibm_code_engine_persistent_data_store_invalid_storage_type` | `storage_type` |
| `ibm_code_engine_persistent_data_store_invalid_data__bucket_location` | `data.bucket_location` |
| `ibm_code_engine_persistent_data_store_invalid_data__bucket_name` | `data.bucket_name` |
| `ibm_code_engine_persistent_data_store_invalid_data__secret_name` | `data.secret_name` |

## Example

### Valid

```hcl
resource "ibm_code_engine_persistent_data_store" "test" {
  name = "example-name"
  project_id = "01234567-89ab-cdef-0123-456789abcdef"
  storage_type = "object_storage"
}
```
