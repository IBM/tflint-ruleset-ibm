# ibm_code_engine_job_data

Reads ibm_code_engine_job. The name of your job.

This rule validates `ibm_code_engine_job` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `params:patch:/projects/{project_id}/allowed_outbound_destinations/{name}` (`codeengine-v2`).

## Configuration

```hcl
rule "ibm_code_engine_job_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `build` | no | — |
| `build_run` | no | — |
| `computed_env_variables` | no | — |
| `created_at` | no | — |
| `entity_tag` | no | — |
| `href` | no | — |
| `image_reference` | no | — |
| `image_secret` | no | — |
| `job_id` | no | — |
| `name` | yes | pattern `^[a-z0-9]([\-a-z0-9]*[a-z0-9])?$` |
| `project_id` | yes | pattern `^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$` |
| `region` | no | — |
| `resource_type` | no | — |
| `run_arguments` | no | — |
| `run_as_user` | no | — |
| `run_commands` | no | — |
| `run_compute_resource_token_enabled` | no | — |
| `run_env_variables` | no | — |
| `run_mode` | no | — |
| `run_service_account` | no | — |
| `run_volume_mounts` | no | — |
| `scale_array_spec` | no | — |
| `scale_cpu_limit` | no | — |
| `scale_ephemeral_storage_limit` | no | — |
| `scale_max_execution_time` | no | — |
| `scale_memory_limit` | no | — |
| `scale_retry_limit` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_code_engine_job_data_invalid_name` | `name` |
| `ibm_code_engine_job_data_invalid_project_id` | `project_id` |

## Example

### Valid

```hcl
data "ibm_code_engine_job" "test" {
  name = "example-name"
  project_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
