# ibm_en_topic

Manages ibm_en_topic. Name of the topic.

This rule validates `ibm_en_topic` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Topic` (`event-notifications`).

## Configuration

```hcl
rule "ibm_en_topic" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `description` | no | API: Description of the topic |
| `instance_guid` | yes | — |
| `name` | yes | API: Name of the topic |
| `source_count` | no | — |
| `subscription_count` | no | — |
| `subscriptions` | no | — |
| `topic_id` | no | — |
| `updated_at` | no | — |
| `sources` (block) | no | nested block |
| `sources.id` | yes | API: ID of the source |
| `sources.rules` (block) | no | nested block |
| `sources.rules.enabled` | no | API: Whether the rule is enabled or not |
| `sources.rules.event_type_filter` | no | API: Event type filter |
| `sources.rules.notification_filter` | no | API: Notification filter |
| `sources.rules.event_schedule_filter` (block) | no | nested block |
| `sources.rules.event_schedule_filter.ends_at` | no | API: event schedule end time |
| `sources.rules.event_schedule_filter.expression` | no | API: cron schedule expression |
| `sources.rules.event_schedule_filter.starts_at` | no | API: event schedule start time |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_en_topic_invalid_description` | `description` |
| `ibm_en_topic_invalid_instance_guid` | `instance_guid` |
| `ibm_en_topic_invalid_name` | `name` |
| `ibm_en_topic_invalid_sources__id` | `sources.id` |
| `ibm_en_topic_invalid_sources__rules__enabled` | `sources.rules.enabled` |
| `ibm_en_topic_invalid_sources__rules__event_type_filter` | `sources.rules.event_type_filter` |
| `ibm_en_topic_invalid_sources__rules__notification_filter` | `sources.rules.notification_filter` |
| `ibm_en_topic_invalid_sources__rules__event_schedule_filter__ends_at` | `sources.rules.event_schedule_filter.ends_at` |
| `ibm_en_topic_invalid_sources__rules__event_schedule_filter__expression` | `sources.rules.event_schedule_filter.expression` |
| `ibm_en_topic_invalid_sources__rules__event_schedule_filter__starts_at` | `sources.rules.event_schedule_filter.starts_at` |

## Example

### Valid

```hcl
resource "ibm_en_topic" "test" {
  instance_guid = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
  name = "aaaaaaa"
}
```
