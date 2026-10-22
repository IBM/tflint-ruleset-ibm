# ibm_en_topic_data

Reads ibm_en_topic. Name of the topic.

This rule validates `ibm_en_topic` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Topic` (`event-notifications`).

## Configuration

```hcl
rule "ibm_en_topic_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `description` | no | — |
| `instance_guid` | yes | — |
| `name` | no | — |
| `source_count` | no | — |
| `sources` | no | — |
| `subscription_count` | no | — |
| `subscriptions` | no | — |
| `topic_id` | yes | pattern `[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}` |
| `updated_at` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_en_topic_data_invalid_instance_guid` | `instance_guid` |
| `ibm_en_topic_data_invalid_topic_id` | `topic_id` |

## Example

### Valid

```hcl
data "ibm_en_topic" "test" {
  instance_guid = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
  topic_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
