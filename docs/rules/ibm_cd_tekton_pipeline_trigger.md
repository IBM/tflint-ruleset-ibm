# ibm_cd_tekton_pipeline_trigger

Manages ibm_cd_tekton_pipeline_trigger. Trigger name.

This rule validates `ibm_cd_tekton_pipeline_trigger` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/tekton_pipelines/{pipeline_id}/triggers/{trigger_id}:patch` (`tekton-pipeline`).

## Configuration

```hcl
rule "ibm_cd_tekton_pipeline_trigger" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cron` | no | pattern `^[-0-9a-zA-Z,\*\/ ]{5,253}$` |
| `disable_draft_events` | no | API: Prevent new pipeline runs from being triggered by events from draft pull requests |
| `enable_events_from_forks` | no | API: Only used for SCM triggers. When enabled, pull request events from forks of the selected repository will trigger a pipeline run |
| `enabled` | no | API: Defines if this trigger is enabled |
| `event_listener` | yes | pattern `^[-0-9a-zA-Z_.]{1,253}$` |
| `events` | no | one of: `push`, `pull_request`, `pull_request_closed` |
| `favorite` | no | API: Mark the trigger as a favorite |
| `filter` | no | API: Either 'events' or 'filter' can be used. Stores the CEL (Common Expression Language) expression value which is used for event filtering against the Git webhook payloads |
| `href` | no | — |
| `limit_waiting_runs` | no | API: Flag that will limit the trigger to a maximum of one waiting run. A newly triggered run will cause any other waiting run(s) to be automatically cancelled |
| `max_concurrent_runs` | no | API: Defines the maximum number of concurrent runs for this trigger. If omitted then the concurrency limit is disabled for this trigger |
| `name` | yes | pattern `^([a-zA-Z0-9]{1,2}|[a-zA-Z0-9][0-9a-zA-Z-_.: \/\(\)\[\]]{1,251}[a-zA-Z0-9])$` |
| `pipeline_id` | yes | pattern `^[-0-9a-z]+$` |
| `properties` | no | — |
| `tags` | no | pattern `^[-0-9a-zA-Z_.]{1,253}$` |
| `timezone` | no | pattern `^[-0-9a-zA-Z+_., \/]{1,253}$` |
| `trigger_id` | no | — |
| `type` | yes | one of: `generic`, `manual`, `scm`, `timer` |
| `webhook_url` | no | — |
| `secret` (block) | no | nested block |
| `secret.algorithm` | no | one of: `md4`, `md5`, `sha1`, `sha256`, `sha384`, `sha512`, `sha512_224`, `sha512_256`, `ripemd160` |
| `secret.key_name` | no | pattern `^[-0-9a-zA-Z_.]{1,253}$` |
| `secret.source` | no | one of: `header`, `payload`, `query` |
| `secret.type` | yes | one of: `token_matches`, `digest_matches`, `internal_validation` |
| `secret.value` | no | API: Secret value, not needed if secret type is `internal_validation` |
| `source` (block) | no | nested block |
| `source.type` | yes | pattern `^git$` |
| `source.properties` (block) | yes | nested block |
| `source.properties.branch` | no | pattern `^[-0-9a-zA-Z_.]{1,253}$` |
| `source.properties.pattern` | no | API: The pattern of Git branch or tag. You can specify a glob pattern such as '!test' or '*master' to match against multiple tags or branches in the repository. The glob pattern used must conform to Bash 4.3 specifications, see bash documentation for more info: https://www.gnu.org/software/bash/manual/bash.html#Pattern-Matching. Only one of branch, pattern, or filter should be specified |
| `source.properties.url` | yes | pattern `^http(s)?:\/\/([^\/?#]*)([^?#]*)(\?([^#]*))?(#(.*))?$` |
| `worker` (block) | no | nested block |
| `worker.id` | yes | pattern `^[-0-9a-zA-Z]{1,253}$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cd_tekton_pipeline_trigger_invalid_cron` | `cron` |
| `ibm_cd_tekton_pipeline_trigger_invalid_disable_draft_events` | `disable_draft_events` |
| `ibm_cd_tekton_pipeline_trigger_invalid_enable_events_from_forks` | `enable_events_from_forks` |
| `ibm_cd_tekton_pipeline_trigger_invalid_enabled` | `enabled` |
| `ibm_cd_tekton_pipeline_trigger_invalid_event_listener` | `event_listener` |
| `ibm_cd_tekton_pipeline_trigger_invalid_events` | `events` |
| `ibm_cd_tekton_pipeline_trigger_invalid_favorite` | `favorite` |
| `ibm_cd_tekton_pipeline_trigger_invalid_filter` | `filter` |
| `ibm_cd_tekton_pipeline_trigger_invalid_limit_waiting_runs` | `limit_waiting_runs` |
| `ibm_cd_tekton_pipeline_trigger_invalid_max_concurrent_runs` | `max_concurrent_runs` |
| `ibm_cd_tekton_pipeline_trigger_invalid_name` | `name` |
| `ibm_cd_tekton_pipeline_trigger_invalid_pipeline_id` | `pipeline_id` |
| `ibm_cd_tekton_pipeline_trigger_invalid_tags` | `tags` |
| `ibm_cd_tekton_pipeline_trigger_invalid_timezone` | `timezone` |
| `ibm_cd_tekton_pipeline_trigger_invalid_type` | `type` |
| `ibm_cd_tekton_pipeline_trigger_invalid_secret__algorithm` | `secret.algorithm` |
| `ibm_cd_tekton_pipeline_trigger_invalid_secret__key_name` | `secret.key_name` |
| `ibm_cd_tekton_pipeline_trigger_invalid_secret__source` | `secret.source` |
| `ibm_cd_tekton_pipeline_trigger_invalid_secret__type` | `secret.type` |
| `ibm_cd_tekton_pipeline_trigger_invalid_secret__value` | `secret.value` |
| `ibm_cd_tekton_pipeline_trigger_invalid_source__type` | `source.type` |
| `ibm_cd_tekton_pipeline_trigger_invalid_source__properties__branch` | `source.properties.branch` |
| `ibm_cd_tekton_pipeline_trigger_invalid_source__properties__pattern` | `source.properties.pattern` |
| `ibm_cd_tekton_pipeline_trigger_invalid_source__properties__url` | `source.properties.url` |
| `ibm_cd_tekton_pipeline_trigger_invalid_worker__id` | `worker.id` |

## Example

### Valid

```hcl
resource "ibm_cd_tekton_pipeline_trigger" "test" {
  event_listener = "example-name"
  name = "example-name"
  pipeline_id = "example-namexxxxxxxxxxxxxxxxxxxxxxxx"
  type = "generic"
}
```
