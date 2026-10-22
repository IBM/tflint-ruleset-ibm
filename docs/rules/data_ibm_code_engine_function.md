# ibm_code_engine_function_data

Reads ibm_code_engine_function. The name of your function.

This rule validates `ibm_code_engine_function` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `params:patch:/projects/{project_id}/allowed_outbound_destinations/{name}` (`codeengine-v2`).

## Configuration

```hcl
rule "ibm_code_engine_function_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `code_binary` | no | — |
| `code_main` | no | — |
| `code_reference` | no | — |
| `code_secret` | no | — |
| `computed_env_variables` | no | — |
| `created_at` | no | — |
| `endpoint` | no | — |
| `endpoint_internal` | no | — |
| `entity_tag` | no | — |
| `function_id` | no | — |
| `href` | no | — |
| `managed_domain_mappings` | no | — |
| `name` | yes | pattern `^[a-z]([-a-z0-9]*[a-z0-9])?$` |
| `project_id` | yes | pattern `^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$` |
| `region` | no | — |
| `resource_type` | no | — |
| `run_compute_resource_token_enabled` | no | — |
| `run_env_variables` | no | — |
| `runtime` | no | — |
| `scale_concurrency` | no | — |
| `scale_cpu_limit` | no | — |
| `scale_down_delay` | no | — |
| `scale_max_execution_time` | no | — |
| `scale_memory_limit` | no | — |
| `status` | no | — |
| `status_details` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_code_engine_function_data_invalid_name` | `name` |
| `ibm_code_engine_function_data_invalid_project_id` | `project_id` |

## Example

### Valid

```hcl
data "ibm_code_engine_function" "test" {
  name = "example-name"
  project_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
