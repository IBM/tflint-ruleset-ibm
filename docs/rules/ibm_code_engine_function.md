# ibm_code_engine_function

Manages ibm_code_engine_function. The name of the function.

This rule validates `ibm_code_engine_function` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `FunctionPrototype` (`codeengine-v2`).

## Configuration

```hcl
rule "ibm_code_engine_function" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `code_binary` | no | API: Specifies whether the code is binary or not. Defaults to false when `code_reference` is set to a data URL. When `code_reference` is set to a code bundle URL, this field is always true. |
| `code_main` | no | pattern `^[a-zA-Z_][a-zA-Z0-9_]*$` |
| `code_reference` | yes | pattern `^([a-z0-9][a-z0-9\-_.]+[a-z0-9][\/])?([a-z0-9][a-z0-9\-_]+[a-z0-9][\/])?[a-z0-9][a-z0-9\-_.\/]+[a-z0-9](:[\w][\w.\-]{0,127})?(@sha256:[a-fA-F0-9]{64})?$|data:([-\w]+\/[-+\w.]+)?(;?\w+=[-\w]+)*;base64,.*` |
| `code_secret` | no | pattern `^[a-z0-9]([\-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([\-a-z0-9]*[a-z0-9])?)*$` |
| `computed_env_variables` | no | — |
| `created_at` | no | — |
| `endpoint` | no | — |
| `endpoint_internal` | no | — |
| `entity_tag` | no | — |
| `etag` | no | — |
| `function_id` | no | — |
| `href` | no | — |
| `managed_domain_mappings` | no | one of: `local`, `local_private`, `local_public` |
| `name` | yes | pattern `^[a-z]([-a-z0-9]*[a-z0-9])?$` |
| `project_id` | yes | pattern `^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$` |
| `region` | no | — |
| `resource_type` | no | — |
| `run_compute_resource_token_enabled` | no | API: Optional flag to enable the use of a compute resource token mounted to the container file system. |
| `runtime` | yes | pattern `^[a-z]*\-[0-9]*(\.[0-9]*)?$` |
| `scale_concurrency` | no | range 1–100 |
| `scale_cpu_limit` | no | pattern `^([0-9.]+)([eEinumkKMGTPB]*)$` |
| `scale_down_delay` | no | range 0–600 |
| `scale_max_execution_time` | no | range 1–120 |
| `scale_memory_limit` | no | pattern `^([0-9.]+)([eEinumkKMGTPB]*)$` |
| `status` | no | — |
| `status_details` | no | — |
| `run_env_variables` (block) | no | nested block |
| `run_env_variables.key` | no | pattern `^[\-._a-zA-Z0-9]+$` |
| `run_env_variables.name` | no | pattern `^[\-._a-zA-Z0-9]+$` |
| `run_env_variables.prefix` | no | pattern `^[a-zA-Z_][a-zA-Z0-9_]*$` |
| `run_env_variables.reference` | no | pattern `^[a-z0-9]([\-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([\-a-z0-9]*[a-z0-9])?)*$` |
| `run_env_variables.type` | yes | one of: `literal`, `config_map_full_reference`, `secret_full_reference`, `config_map_key_reference`, `secret_key_reference`; pattern `^(literal|config_map_full_reference|secret_full_reference|config_map_key_reference|secret_key_reference)$` |
| `run_env_variables.value` | no | API: The literal value of the environment variable. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_code_engine_function_invalid_code_binary` | `code_binary` |
| `ibm_code_engine_function_invalid_code_main` | `code_main` |
| `ibm_code_engine_function_invalid_code_reference` | `code_reference` |
| `ibm_code_engine_function_invalid_code_secret` | `code_secret` |
| `ibm_code_engine_function_invalid_managed_domain_mappings` | `managed_domain_mappings` |
| `ibm_code_engine_function_invalid_name` | `name` |
| `ibm_code_engine_function_invalid_project_id` | `project_id` |
| `ibm_code_engine_function_invalid_run_compute_resource_token_enabled` | `run_compute_resource_token_enabled` |
| `ibm_code_engine_function_invalid_runtime` | `runtime` |
| `ibm_code_engine_function_invalid_scale_concurrency` | `scale_concurrency` |
| `ibm_code_engine_function_invalid_scale_cpu_limit` | `scale_cpu_limit` |
| `ibm_code_engine_function_invalid_scale_down_delay` | `scale_down_delay` |
| `ibm_code_engine_function_invalid_scale_max_execution_time` | `scale_max_execution_time` |
| `ibm_code_engine_function_invalid_scale_memory_limit` | `scale_memory_limit` |
| `ibm_code_engine_function_invalid_run_env_variables__key` | `run_env_variables.key` |
| `ibm_code_engine_function_invalid_run_env_variables__name` | `run_env_variables.name` |
| `ibm_code_engine_function_invalid_run_env_variables__prefix` | `run_env_variables.prefix` |
| `ibm_code_engine_function_invalid_run_env_variables__reference` | `run_env_variables.reference` |
| `ibm_code_engine_function_invalid_run_env_variables__type` | `run_env_variables.type` |
| `ibm_code_engine_function_invalid_run_env_variables__value` | `run_env_variables.value` |

## Example

### Valid

```hcl
resource "ibm_code_engine_function" "test" {
  code_reference = "example-name"
  name = "example-name"
  project_id = "01234567-89ab-cdef-0123-456789abcdef"
  runtime = "python-3.11"
}
```
