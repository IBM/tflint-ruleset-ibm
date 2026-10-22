# ibm_code_engine_build

Manages ibm_code_engine_build. The name of the build.

This rule validates `ibm_code_engine_build` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Build` (`codeengine-v2`).

## Configuration

```hcl
rule "ibm_code_engine_build" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `build_id` | no | — |
| `created_at` | no | — |
| `entity_tag` | no | — |
| `etag` | no | — |
| `href` | no | — |
| `name` | yes | pattern `^[a-z0-9]([\-a-z0-9]*[a-z0-9])?$` |
| `output_image` | yes | pattern `^([a-z0-9][a-z0-9\-_.]+[a-z0-9][\/])?([a-z0-9][a-z0-9\-_]+[a-z0-9][\/])?[a-z0-9][a-z0-9\-_.\/]+[a-z0-9](:[\w][\w.\-]{0,127})?(@sha256:[a-fA-F0-9]{64})?$` |
| `output_secret` | yes | pattern `^[a-z0-9]([\-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([\-a-z0-9]*[a-z0-9])?)*$` |
| `project_id` | yes | pattern `^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$` |
| `region` | no | — |
| `resource_type` | no | — |
| `source_context_dir` | no | API: Optional directory in the repository that contains the buildpacks file or the Dockerfile. |
| `source_revision` | no | API: Commit, tag, or branch in the source repository to pull. This field is optional if the `source_type` is `git` and uses the HEAD of default branch if not specified. If the `source_type` value is `local`, this field must be omitted. |
| `source_secret` | no | pattern `^[a-z0-9]([\-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([\-a-z0-9]*[a-z0-9])?)*$` |
| `source_type` | no | one of: `git`, `local` |
| `source_url` | no | pattern `^((https:\/\/[a-z0-9]([\-.]?[a-z0-9])+(:\d{1,5})?)|((ssh:\/\/)?git@[a-z0-9]([\-.]{0,1}[a-z0-9])+(:[a-zA-Z0-9\/][\w\-.]*)?))(\/([\w\-.]|%20)+)*$` |
| `status` | no | — |
| `status_details` | no | — |
| `strategy_size` | no | one of: `large`, `medium`, `small`, `xlarge`, `xxlarge` |
| `strategy_spec_file` | no | API: Optional path to the specification file that is used for build strategies for building an image. |
| `strategy_type` | yes | one of: `buildpacks`, `dockerfile` |
| `timeout` | no | range 1–3600 |
| `run_build_params` (block) | no | nested block |
| `run_build_params.key` | no | pattern `^[\-._a-zA-Z][\-._a-zA-Z0-9]*$` |
| `run_build_params.name` | no | pattern `^[\-._a-zA-Z][\-._a-zA-Z0-9]*$` |
| `run_build_params.reference` | no | pattern `^[a-z0-9]([\-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([\-a-z0-9]*[a-z0-9])?)*$` |
| `run_build_params.type` | yes | one of: `literal`, `config_map_key_reference`, `secret_key_reference`; pattern `^(literal|config_map_key_reference|secret_key_reference)$` |
| `run_build_params.value` | no | API: The literal value of the build param. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_code_engine_build_invalid_name` | `name` |
| `ibm_code_engine_build_invalid_output_image` | `output_image` |
| `ibm_code_engine_build_invalid_output_secret` | `output_secret` |
| `ibm_code_engine_build_invalid_project_id` | `project_id` |
| `ibm_code_engine_build_invalid_source_context_dir` | `source_context_dir` |
| `ibm_code_engine_build_invalid_source_revision` | `source_revision` |
| `ibm_code_engine_build_invalid_source_secret` | `source_secret` |
| `ibm_code_engine_build_invalid_source_type` | `source_type` |
| `ibm_code_engine_build_invalid_source_url` | `source_url` |
| `ibm_code_engine_build_invalid_strategy_size` | `strategy_size` |
| `ibm_code_engine_build_invalid_strategy_spec_file` | `strategy_spec_file` |
| `ibm_code_engine_build_invalid_strategy_type` | `strategy_type` |
| `ibm_code_engine_build_invalid_timeout` | `timeout` |
| `ibm_code_engine_build_invalid_run_build_params__key` | `run_build_params.key` |
| `ibm_code_engine_build_invalid_run_build_params__name` | `run_build_params.name` |
| `ibm_code_engine_build_invalid_run_build_params__reference` | `run_build_params.reference` |
| `ibm_code_engine_build_invalid_run_build_params__type` | `run_build_params.type` |
| `ibm_code_engine_build_invalid_run_build_params__value` | `run_build_params.value` |

## Example

### Valid

```hcl
variable "output_secret" {
  sensitive = true
}

resource "ibm_code_engine_build" "test" {
  name = "example-name"
  output_image = "example-name"
  output_secret = var.output_secret
  project_id = "01234567-89ab-cdef-0123-456789abcdef"
  strategy_type = "buildpacks"
}
```
