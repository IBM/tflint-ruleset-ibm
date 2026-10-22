# ibm_cd_tekton_pipeline_data

Reads ibm_cd_tekton_pipeline. String.

This rule validates `ibm_cd_tekton_pipeline` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/tekton_pipelines/{pipeline_id}/triggers/{trigger_id}:patch` (`tekton-pipeline`).

## Configuration

```hcl
rule "ibm_cd_tekton_pipeline_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `build_number` | no | — |
| `created_at` | no | — |
| `definitions` | no | — |
| `enable_notifications` | no | — |
| `enable_partial_cloning` | no | — |
| `enabled` | no | — |
| `href` | no | — |
| `name` | no | — |
| `next_build_number` | no | — |
| `pipeline_id` | yes | pattern `^[-0-9a-z]+$` |
| `properties` | no | — |
| `resource_group` | no | — |
| `runs_url` | no | — |
| `status` | no | — |
| `toolchain` | no | — |
| `triggers` | no | — |
| `updated_at` | no | — |
| `worker` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cd_tekton_pipeline_data_invalid_pipeline_id` | `pipeline_id` |

## Example

### Valid

```hcl
data "ibm_cd_tekton_pipeline" "test" {
  pipeline_id = "example-namexxxxxxxxxxxxxxxxxxxxxxxx"
}
```
