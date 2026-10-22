# ibm_code_engine_config_map

Manages ibm_code_engine_config_map. The name of the config map.

This rule validates `ibm_code_engine_config_map` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ConfigMapPrototype` (`codeengine-v2`).

## Configuration

```hcl
rule "ibm_code_engine_config_map" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `config_map_id` | no | — |
| `created_at` | no | — |
| `data` | no | API: The key-value pair for the config map. Values must be specified in `KEY=VALUE` format. Each `KEY` field must consist of alphanumeric characters, `-`, `_` or `.` and must not be exceed a max length of 253 characters. Each `VALUE` field can consists of any character and must not be exceed a max length of 1048576 characters. |
| `entity_tag` | no | — |
| `etag` | no | — |
| `href` | no | — |
| `name` | yes | pattern `^[a-z0-9]([\-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([\-a-z0-9]*[a-z0-9])?)*$` |
| `project_id` | yes | pattern `^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$` |
| `region` | no | — |
| `resource_type` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_code_engine_config_map_invalid_data` | `data` |
| `ibm_code_engine_config_map_invalid_name` | `name` |
| `ibm_code_engine_config_map_invalid_project_id` | `project_id` |

## Example

### Valid

```hcl
resource "ibm_code_engine_config_map" "test" {
  name = "example-name"
  project_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
