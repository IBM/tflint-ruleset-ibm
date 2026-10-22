# ibm_code_engine_config_map_data

Reads ibm_code_engine_config_map. The name of your configmap.

This rule validates `ibm_code_engine_config_map` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `params:put:/projects/{project_id}/config_maps/{name}` (`codeengine-v2`).

## Configuration

```hcl
rule "ibm_code_engine_config_map_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `config_map_id` | no | — |
| `created_at` | no | — |
| `data` | no | — |
| `entity_tag` | no | — |
| `href` | no | — |
| `name` | yes | pattern `^[a-z0-9]([\-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([\-a-z0-9]*[a-z0-9])?)*$` |
| `project_id` | yes | pattern `^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$` |
| `region` | no | — |
| `resource_type` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_code_engine_config_map_data_invalid_name` | `name` |
| `ibm_code_engine_config_map_data_invalid_project_id` | `project_id` |

## Example

### Valid

```hcl
data "ibm_code_engine_config_map" "test" {
  name = "example-name"
  project_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
