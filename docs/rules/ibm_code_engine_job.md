# ibm_code_engine_job

Manages ibm_code_engine_job. The name of the job.

This rule validates `ibm_code_engine_job` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Job` (`codeengine-v2`).

## Configuration

```hcl
rule "ibm_code_engine_job" {
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
| `etag` | no | — |
| `href` | no | — |
| `image_reference` | yes | pattern `^([a-z0-9][a-z0-9\-_.]+[a-z0-9][\/])?([a-z0-9][a-z0-9\-_]+[a-z0-9][\/])?[a-z0-9][a-z0-9\-_.\/]+[a-z0-9](:[\w][\w.\-]{0,127})?(@sha256:[a-fA-F0-9]{64})?$` |
| `image_secret` | no | pattern `^[a-z0-9]([\-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([\-a-z0-9]*[a-z0-9])?)*$` |
| `job_id` | no | — |
| `name` | yes | pattern `^[a-z0-9]([\-a-z0-9]*[a-z0-9])?$` |
| `project_id` | yes | pattern `^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$` |
| `region` | no | — |
| `resource_type` | no | — |
| `run_arguments` | no | API: Set arguments for the job that are passed to start job run containers. If not specified an empty string array will be applied and the arguments specified by the container image, will be used to start the container. |
| `run_as_user` | no | API: The user ID (UID) to run the job. |
| `run_commands` | no | API: Set commands for the job that are passed to start job run containers. If not specified an empty string array will be applied and the command specified by the container image, will be used to start the container. |
| `run_compute_resource_token_enabled` | no | API: Optional flag to enable the use of a compute resource token mounted to the container file system. |
| `run_mode` | no | one of: `daemon`, `task`; pattern `^(task|daemon)$` |
| `run_service_account` | no | one of: `default`, `manager`, `none`, `reader`, `writer`; pattern `^(manager|reader|writer|none|default)$` |
| `scale_array_spec` | no | pattern `^(?:[1-9]\d\d\d\d\d\d|[1-9]\d\d\d\d\d|[1-9]\d\d\d\d|[1-9]\d\d\d|[1-9]\d\d|[1-9]?\d)(?:-(?:[1-9]\d\d\d\d\d\d|[1-9]\d\d\d\d\d|[1-9]\d\d\d\d|[1-9]\d\d\d|[1-9]\d\d|[1-9]?\d))?(?:,(?:[1-9]\d\d\d\d\d\d|[1-9]\d\d\d\d\d|[1-9]\d\d\d\d|[1-9]\d\d\d|[1-9]\d\d|[1-9]?\d)(?:-(?:[1-9]\d\d\d\d\d\d|[1-9]\d\d\d\d\d|[1-9]\d\d\d\d|[1-9]\d\d\d|[1-9]\d\d|[1-9]?\d))?)*$` |
| `scale_cpu_limit` | no | pattern `^([0-9.]+)([eEinumkKMGTPB]*)$` |
| `scale_ephemeral_storage_limit` | no | pattern `^([0-9.]+)([eEinumkKMGTPB]*)$` |
| `scale_max_execution_time` | no | range 1–86400 |
| `scale_memory_limit` | no | pattern `^([0-9.]+)([eEinumkKMGTPB]*)$` |
| `scale_retry_limit` | no | range 0–5 |
| `run_env_variables` (block) | no | nested block |
| `run_env_variables.key` | no | pattern `^[\-._a-zA-Z0-9]+$` |
| `run_env_variables.name` | no | pattern `^[\-._a-zA-Z0-9]+$` |
| `run_env_variables.prefix` | no | pattern `^[a-zA-Z_][a-zA-Z0-9_]*$` |
| `run_env_variables.reference` | no | pattern `^[a-z0-9]([\-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([\-a-z0-9]*[a-z0-9])?)*$` |
| `run_env_variables.type` | yes | one of: `literal`, `config_map_full_reference`, `secret_full_reference`, `config_map_key_reference`, `secret_key_reference`; pattern `^(literal|config_map_full_reference|secret_full_reference|config_map_key_reference|secret_key_reference)$` |
| `run_env_variables.value` | no | API: The literal value of the environment variable. |
| `run_volume_mounts` (block) | no | nested block |
| `run_volume_mounts.mount_path` | yes | pattern `^\/([^\/\0]+\/?)+$` |
| `run_volume_mounts.read_only` | no | API: Optional flag for a volume mount of type 'persistent_data_store' to specify whether it is read-only. |
| `run_volume_mounts.reference` | yes | pattern `^[a-z0-9]([\-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([\-a-z0-9]*[a-z0-9])?)*$` |
| `run_volume_mounts.sub_path` | no | API: The path mounted at the mount path. |
| `run_volume_mounts.type` | yes | one of: `config_map`, `persistent_data_store`, `secret`; pattern `^(config_map|persistent_data_store|secret)$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_code_engine_job_invalid_image_reference` | `image_reference` |
| `ibm_code_engine_job_invalid_image_secret` | `image_secret` |
| `ibm_code_engine_job_invalid_name` | `name` |
| `ibm_code_engine_job_invalid_project_id` | `project_id` |
| `ibm_code_engine_job_invalid_run_arguments` | `run_arguments` |
| `ibm_code_engine_job_invalid_run_as_user` | `run_as_user` |
| `ibm_code_engine_job_invalid_run_commands` | `run_commands` |
| `ibm_code_engine_job_invalid_run_compute_resource_token_enabled` | `run_compute_resource_token_enabled` |
| `ibm_code_engine_job_invalid_run_mode` | `run_mode` |
| `ibm_code_engine_job_invalid_run_service_account` | `run_service_account` |
| `ibm_code_engine_job_invalid_scale_array_spec` | `scale_array_spec` |
| `ibm_code_engine_job_invalid_scale_cpu_limit` | `scale_cpu_limit` |
| `ibm_code_engine_job_invalid_scale_ephemeral_storage_limit` | `scale_ephemeral_storage_limit` |
| `ibm_code_engine_job_invalid_scale_max_execution_time` | `scale_max_execution_time` |
| `ibm_code_engine_job_invalid_scale_memory_limit` | `scale_memory_limit` |
| `ibm_code_engine_job_invalid_scale_retry_limit` | `scale_retry_limit` |
| `ibm_code_engine_job_invalid_run_env_variables__key` | `run_env_variables.key` |
| `ibm_code_engine_job_invalid_run_env_variables__name` | `run_env_variables.name` |
| `ibm_code_engine_job_invalid_run_env_variables__prefix` | `run_env_variables.prefix` |
| `ibm_code_engine_job_invalid_run_env_variables__reference` | `run_env_variables.reference` |
| `ibm_code_engine_job_invalid_run_env_variables__type` | `run_env_variables.type` |
| `ibm_code_engine_job_invalid_run_env_variables__value` | `run_env_variables.value` |
| `ibm_code_engine_job_invalid_run_volume_mounts__mount_path` | `run_volume_mounts.mount_path` |
| `ibm_code_engine_job_invalid_run_volume_mounts__read_only` | `run_volume_mounts.read_only` |
| `ibm_code_engine_job_invalid_run_volume_mounts__reference` | `run_volume_mounts.reference` |
| `ibm_code_engine_job_invalid_run_volume_mounts__sub_path` | `run_volume_mounts.sub_path` |
| `ibm_code_engine_job_invalid_run_volume_mounts__type` | `run_volume_mounts.type` |

## Example

### Valid

```hcl
resource "ibm_code_engine_job" "test" {
  image_reference = "example-name"
  name = "example-name"
  project_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
