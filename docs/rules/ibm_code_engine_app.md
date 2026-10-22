# ibm_code_engine_app

Manages ibm_code_engine_app. The name of the app.

This rule validates `ibm_code_engine_app` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `App` (`codeengine-v2`).

## Configuration

```hcl
rule "ibm_code_engine_app" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `app_id` | no | — |
| `build` | no | — |
| `build_run` | no | — |
| `computed_env_variables` | no | — |
| `created_at` | no | — |
| `endpoint` | no | — |
| `endpoint_internal` | no | — |
| `entity_tag` | no | — |
| `etag` | no | — |
| `href` | no | — |
| `image_port` | no | range 0–65535 |
| `image_reference` | yes | pattern `^([a-z0-9][a-z0-9\-_.]+[a-z0-9][\/])?([a-z0-9][a-z0-9\-_]+[a-z0-9][\/])?[a-z0-9][a-z0-9\-_.\/]+[a-z0-9](:[\w][\w.\-]{0,127})?(@sha256:[a-fA-F0-9]{64})?$` |
| `image_secret` | no | pattern `^[a-z0-9]([\-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([\-a-z0-9]*[a-z0-9])?)*$` |
| `managed_domain_mappings` | no | one of: `local`, `local_private`, `local_public` |
| `name` | yes | pattern `^[a-z]([-a-z0-9]*[a-z0-9])?$` |
| `project_id` | yes | pattern `^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$` |
| `region` | no | — |
| `resource_type` | no | — |
| `run_arguments` | no | API: Optional arguments for the app that are passed to start the container. If not specified an empty string array will be applied and the arguments specified by the container image, will be used to start the container. |
| `run_as_user` | no | API: Optional user ID (UID) to run the app. |
| `run_commands` | no | API: Optional commands for the app that are passed to start the container. If not specified an empty string array will be applied and the command specified by the container image, will be used to start the container. |
| `run_compute_resource_token_enabled` | no | API: Optional flag to enable the use of a compute resource token mounted to the container file system. |
| `run_service_account` | no | one of: `default`, `manager`, `none`, `reader`, `writer`; pattern `^(manager|reader|writer|none|default)$` |
| `scale_concurrency` | no | range 1–1000 |
| `scale_concurrency_target` | no | range 1–1000 |
| `scale_cpu_limit` | no | pattern `^([0-9.]+)([eEinumkKMGTPB]*)$` |
| `scale_down_delay` | no | range 0–3600 |
| `scale_ephemeral_storage_limit` | no | pattern `^([0-9.]+)([eEinumkKMGTPB]*)$` |
| `scale_initial_instances` | no | API: Optional initial number of instances that are created upon app creation or app update. |
| `scale_max_instances` | no | API: Optional maximum number of instances for this app. If you set this value to `0`, this property does not set a upper scaling limit. However, the app scaling is still limited by the project quota for instances. See [Limits and quotas for Code Engine](https://cloud.ibm.com/docs/codeengine?topic=codeengine-limits). |
| `scale_memory_limit` | no | pattern `^([0-9.]+)([eEinumkKMGTPB]*)$` |
| `scale_min_instances` | no | API: Optional minimum number of instances for this app. If you set this value to `0`, the app will scale down to zero, if not hit by any request for some time. |
| `scale_request_timeout` | no | API: Optional amount of time in seconds that is allowed for a running app to respond to a request. |
| `status` | no | — |
| `status_details` | no | — |
| `probe_liveness` (block) | no | nested block |
| `probe_liveness.failure_threshold` | no | range 1–10 |
| `probe_liveness.initial_delay` | no | range 1–10 |
| `probe_liveness.interval` | no | range 1–60 |
| `probe_liveness.path` | no | pattern `^\/(([a-zA-Z0-9-._~!$&'()*+,;=:@]|%[a-fA-F0-9]{2})+(\/([a-zA-Z0-9-._~!$&'()*+,;=:@]|%[a-fA-F0-9]{2})*)*)?(\?([a-zA-Z0-9-._~!$&'()*+,;=:@\/?]|%[a-fA-F0-9]{2})*)?$` |
| `probe_liveness.port` | no | range 1–65535; port format |
| `probe_liveness.timeout` | no | range 1–3600 |
| `probe_liveness.type` | yes | one of: `tcp`, `http` |
| `probe_readiness` (block) | no | nested block |
| `probe_readiness.failure_threshold` | no | range 1–10 |
| `probe_readiness.initial_delay` | no | range 1–10 |
| `probe_readiness.interval` | no | range 1–60 |
| `probe_readiness.path` | no | pattern `^\/(([a-zA-Z0-9-._~!$&'()*+,;=:@]|%[a-fA-F0-9]{2})+(\/([a-zA-Z0-9-._~!$&'()*+,;=:@]|%[a-fA-F0-9]{2})*)*)?(\?([a-zA-Z0-9-._~!$&'()*+,;=:@\/?]|%[a-fA-F0-9]{2})*)?$` |
| `probe_readiness.port` | no | range 1–65535; port format |
| `probe_readiness.timeout` | no | range 1–3600 |
| `probe_readiness.type` | yes | one of: `tcp`, `http` |
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
| `ibm_code_engine_app_invalid_image_port` | `image_port` |
| `ibm_code_engine_app_invalid_image_reference` | `image_reference` |
| `ibm_code_engine_app_invalid_image_secret` | `image_secret` |
| `ibm_code_engine_app_invalid_managed_domain_mappings` | `managed_domain_mappings` |
| `ibm_code_engine_app_invalid_name` | `name` |
| `ibm_code_engine_app_invalid_project_id` | `project_id` |
| `ibm_code_engine_app_invalid_run_arguments` | `run_arguments` |
| `ibm_code_engine_app_invalid_run_as_user` | `run_as_user` |
| `ibm_code_engine_app_invalid_run_commands` | `run_commands` |
| `ibm_code_engine_app_invalid_run_compute_resource_token_enabled` | `run_compute_resource_token_enabled` |
| `ibm_code_engine_app_invalid_run_service_account` | `run_service_account` |
| `ibm_code_engine_app_invalid_scale_concurrency` | `scale_concurrency` |
| `ibm_code_engine_app_invalid_scale_concurrency_target` | `scale_concurrency_target` |
| `ibm_code_engine_app_invalid_scale_cpu_limit` | `scale_cpu_limit` |
| `ibm_code_engine_app_invalid_scale_down_delay` | `scale_down_delay` |
| `ibm_code_engine_app_invalid_scale_ephemeral_storage_limit` | `scale_ephemeral_storage_limit` |
| `ibm_code_engine_app_invalid_scale_initial_instances` | `scale_initial_instances` |
| `ibm_code_engine_app_invalid_scale_max_instances` | `scale_max_instances` |
| `ibm_code_engine_app_invalid_scale_memory_limit` | `scale_memory_limit` |
| `ibm_code_engine_app_invalid_scale_min_instances` | `scale_min_instances` |
| `ibm_code_engine_app_invalid_scale_request_timeout` | `scale_request_timeout` |
| `ibm_code_engine_app_invalid_probe_liveness__failure_threshold` | `probe_liveness.failure_threshold` |
| `ibm_code_engine_app_invalid_probe_liveness__initial_delay` | `probe_liveness.initial_delay` |
| `ibm_code_engine_app_invalid_probe_liveness__interval` | `probe_liveness.interval` |
| `ibm_code_engine_app_invalid_probe_liveness__path` | `probe_liveness.path` |
| `ibm_code_engine_app_invalid_probe_liveness__port` | `probe_liveness.port` |
| `ibm_code_engine_app_invalid_probe_liveness__timeout` | `probe_liveness.timeout` |
| `ibm_code_engine_app_invalid_probe_liveness__type` | `probe_liveness.type` |
| `ibm_code_engine_app_invalid_probe_readiness__failure_threshold` | `probe_readiness.failure_threshold` |
| `ibm_code_engine_app_invalid_probe_readiness__initial_delay` | `probe_readiness.initial_delay` |
| `ibm_code_engine_app_invalid_probe_readiness__interval` | `probe_readiness.interval` |
| `ibm_code_engine_app_invalid_probe_readiness__path` | `probe_readiness.path` |
| `ibm_code_engine_app_invalid_probe_readiness__port` | `probe_readiness.port` |
| `ibm_code_engine_app_invalid_probe_readiness__timeout` | `probe_readiness.timeout` |
| `ibm_code_engine_app_invalid_probe_readiness__type` | `probe_readiness.type` |
| `ibm_code_engine_app_invalid_run_env_variables__key` | `run_env_variables.key` |
| `ibm_code_engine_app_invalid_run_env_variables__name` | `run_env_variables.name` |
| `ibm_code_engine_app_invalid_run_env_variables__prefix` | `run_env_variables.prefix` |
| `ibm_code_engine_app_invalid_run_env_variables__reference` | `run_env_variables.reference` |
| `ibm_code_engine_app_invalid_run_env_variables__type` | `run_env_variables.type` |
| `ibm_code_engine_app_invalid_run_env_variables__value` | `run_env_variables.value` |
| `ibm_code_engine_app_invalid_run_volume_mounts__mount_path` | `run_volume_mounts.mount_path` |
| `ibm_code_engine_app_invalid_run_volume_mounts__read_only` | `run_volume_mounts.read_only` |
| `ibm_code_engine_app_invalid_run_volume_mounts__reference` | `run_volume_mounts.reference` |
| `ibm_code_engine_app_invalid_run_volume_mounts__sub_path` | `run_volume_mounts.sub_path` |
| `ibm_code_engine_app_invalid_run_volume_mounts__type` | `run_volume_mounts.type` |

## Example

### Valid

```hcl
resource "ibm_code_engine_app" "test" {
  image_reference = "example-name"
  name = "example-name"
  project_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
