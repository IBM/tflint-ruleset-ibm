# ibm_code_engine_build_data

Reads ibm_code_engine_build. The name of your build.

This rule validates `ibm_code_engine_build` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `params:patch:/projects/{project_id}/allowed_outbound_destinations/{name}` (`codeengine-v2`).

## Configuration

```hcl
rule "ibm_code_engine_build_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `build_id` | no | — |
| `created_at` | no | — |
| `entity_tag` | no | — |
| `href` | no | — |
| `name` | yes | pattern `^[a-z0-9]([\-a-z0-9]*[a-z0-9])?$` |
| `output_image` | no | — |
| `output_secret` | no | — |
| `project_id` | yes | pattern `^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$` |
| `region` | no | — |
| `resource_type` | no | — |
| `run_build_params` | no | — |
| `source_context_dir` | no | — |
| `source_revision` | no | — |
| `source_secret` | no | — |
| `source_type` | no | — |
| `source_url` | no | — |
| `status` | no | — |
| `status_details` | no | — |
| `strategy_size` | no | — |
| `strategy_spec_file` | no | — |
| `strategy_type` | no | — |
| `timeout` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_code_engine_build_data_invalid_name` | `name` |
| `ibm_code_engine_build_data_invalid_project_id` | `project_id` |

## Example

### Valid

```hcl
data "ibm_code_engine_build" "test" {
  name = "example-name"
  project_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
