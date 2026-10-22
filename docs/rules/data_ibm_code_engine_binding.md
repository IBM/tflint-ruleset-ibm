# ibm_code_engine_binding_data

Schema validation for the ibm_code_engine_binding data source.

This rule validates `ibm_code_engine_binding` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `params:post:/projects/{project_id}/allowed_outbound_destinations` (`codeengine-v2`).

## Configuration

```hcl
rule "ibm_code_engine_binding_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `binding_id` | yes | — |
| `component` | no | — |
| `href` | no | — |
| `prefix` | no | — |
| `project_id` | yes | pattern `^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$` |
| `resource_type` | no | — |
| `secret_name` | no | — |
| `status` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_code_engine_binding_data_invalid_binding_id` | `binding_id` |
| `ibm_code_engine_binding_data_invalid_project_id` | `project_id` |

## Example

### Valid

```hcl
data "ibm_code_engine_binding" "test" {
  binding_id = "01234567-89ab-cdef-0123-456789abcdef"
  project_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
