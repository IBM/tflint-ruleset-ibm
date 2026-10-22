# ibm_en_destination_ce

Manages ibm_en_destination_ce. The Destintion name.

This rule validates `ibm_en_destination_ce` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `CodeEngineDestinationConfig` (`event-notifications`).

## Configuration

```hcl
rule "ibm_en_destination_ce" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `collect_failed_events` | no | — |
| `description` | no | — |
| `destination_id` | no | — |
| `instance_guid` | yes | — |
| `name` | yes | — |
| `subscription_count` | no | — |
| `subscription_names` | no | — |
| `type` | no | one of: `job`, `application` |
| `updated_at` | no | — |
| `config` (block) | no | nested block |
| `config.params` (block) | no | nested block |
| `config.params.custom_headers` | no | API: Custom headers (Key-Value pair) for webhook call |
| `config.params.job_name` | no | API: name of the code engine job |
| `config.params.project_crn` | no | crn format |
| `config.params.sensitive_headers` | no | API: List of sensitive headers from custom headers |
| `config.params.type` | no | one of: `job`, `application` |
| `config.params.url` | no | pattern `^(?:https?://)?(?:[^@/
]+@)?(?:www.)?([^:/
]+)` |
| `config.params.verb` | no | one of: `get`, `post` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_en_destination_ce_invalid_collect_failed_events` | `collect_failed_events` |
| `ibm_en_destination_ce_invalid_description` | `description` |
| `ibm_en_destination_ce_invalid_instance_guid` | `instance_guid` |
| `ibm_en_destination_ce_invalid_name` | `name` |
| `ibm_en_destination_ce_invalid_type` | `type` |
| `ibm_en_destination_ce_invalid_config__params__custom_headers` | `config.params.custom_headers` |
| `ibm_en_destination_ce_invalid_config__params__job_name` | `config.params.job_name` |
| `ibm_en_destination_ce_invalid_config__params__project_crn` | `config.params.project_crn` |
| `ibm_en_destination_ce_invalid_config__params__sensitive_headers` | `config.params.sensitive_headers` |
| `ibm_en_destination_ce_invalid_config__params__type` | `config.params.type` |
| `ibm_en_destination_ce_invalid_config__params__url` | `config.params.url` |
| `ibm_en_destination_ce_invalid_config__params__verb` | `config.params.verb` |

## Example

### Valid

```hcl
resource "ibm_en_destination_ce" "test" {
  instance_guid = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
  name = "example"
}
```
