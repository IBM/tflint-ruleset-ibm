# ibm_cd_tekton_pipeline_trigger_data

Reads ibm_cd_tekton_pipeline_trigger. Trigger name.

This rule validates `ibm_cd_tekton_pipeline_trigger` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Trigger` (`tekton-pipeline`).

## Configuration

```hcl
rule "ibm_cd_tekton_pipeline_trigger_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cron` | no | — |
| `disable_draft_events` | no | — |
| `enable_events_from_forks` | no | — |
| `enabled` | no | — |
| `event_listener` | no | — |
| `events` | no | — |
| `favorite` | no | — |
| `filter` | no | — |
| `href` | no | — |
| `limit_waiting_runs` | no | — |
| `max_concurrent_runs` | no | — |
| `name` | no | — |
| `pipeline_id` | yes | pattern `^[-0-9a-z]+$` |
| `properties` | no | — |
| `secret` | no | — |
| `source` | no | — |
| `tags` | no | — |
| `timezone` | no | — |
| `trigger_id` | yes | — |
| `type` | no | — |
| `webhook_url` | no | — |
| `worker` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cd_tekton_pipeline_trigger_data_invalid_pipeline_id` | `pipeline_id` |
| `ibm_cd_tekton_pipeline_trigger_data_invalid_trigger_id` | `trigger_id` |

## Example

### Valid

```hcl
data "ibm_cd_tekton_pipeline_trigger" "test" {
  pipeline_id = "example-namexxxxxxxxxxxxxxxxxxxxxxxx"
  trigger_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
