# ibm_code_engine_persistent_data_store_data

Reads ibm_code_engine_persistent_data_store. The name of your persistent data store.

This rule validates `ibm_code_engine_persistent_data_store` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `params:patch:/projects/{project_id}/allowed_outbound_destinations/{name}` (`codeengine-v2`).

## Configuration

```hcl
rule "ibm_code_engine_persistent_data_store_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `created_at` | no | — |
| `data` | no | — |
| `entity_tag` | no | — |
| `name` | yes | pattern `^[a-z0-9]([\-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([\-a-z0-9]*[a-z0-9])?)*$` |
| `project_id` | yes | pattern `^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$` |
| `region` | no | — |
| `storage_type` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_code_engine_persistent_data_store_data_invalid_name` | `name` |
| `ibm_code_engine_persistent_data_store_data_invalid_project_id` | `project_id` |

## Example

### Valid

```hcl
data "ibm_code_engine_persistent_data_store" "test" {
  name = "example-name"
  project_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
