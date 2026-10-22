# ibm_cd_tekton_pipeline

Manages ibm_cd_tekton_pipeline. String.

This rule validates `ibm_cd_tekton_pipeline` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/tekton_pipelines/{pipeline_id}/triggers/{trigger_id}:patch` (`tekton-pipeline`).

## Configuration

```hcl
rule "ibm_cd_tekton_pipeline" {
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
| `next_build_number` | no | range 1–99999999999999 |
| `pipeline_id` | yes | pattern `^[-0-9a-z]+$` |
| `properties` | no | — |
| `resource_group` | no | — |
| `runs_url` | no | — |
| `status` | no | — |
| `toolchain` | no | — |
| `triggers` | no | — |
| `updated_at` | no | — |
| `worker` (block) | no | nested block |
| `worker.id` | yes | pattern `^[-0-9a-zA-Z]{1,253}$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cd_tekton_pipeline_invalid_enable_notifications` | `enable_notifications` |
| `ibm_cd_tekton_pipeline_invalid_enable_partial_cloning` | `enable_partial_cloning` |
| `ibm_cd_tekton_pipeline_invalid_next_build_number` | `next_build_number` |
| `ibm_cd_tekton_pipeline_invalid_pipeline_id` | `pipeline_id` |
| `ibm_cd_tekton_pipeline_invalid_worker__id` | `worker.id` |

## Example

### Valid

```hcl
resource "ibm_cd_tekton_pipeline" "test" {
  pipeline_id = "example-namexxxxxxxxxxxxxxxxxxxxxxxx"
}
```
