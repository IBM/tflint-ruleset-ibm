# ibm_en_subscription_msteams

Manages ibm_en_subscription_msteams. Subscription name.

This rule validates `ibm_en_subscription_msteams` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `SubscriptionCreate` (`event-notifications`).

## Configuration

```hcl
rule "ibm_en_subscription_msteams" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `description` | no | API: Subscription description |
| `destination_id` | yes | pattern `[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}` |
| `destination_name` | no | — |
| `destination_type` | no | — |
| `instance_guid` | yes | — |
| `name` | yes | API: Subscription name |
| `subscription_id` | no | — |
| `topic_id` | yes | pattern `[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}` |
| `topic_name` | no | — |
| `updated_at` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_en_subscription_msteams_invalid_description` | `description` |
| `ibm_en_subscription_msteams_invalid_destination_id` | `destination_id` |
| `ibm_en_subscription_msteams_invalid_instance_guid` | `instance_guid` |
| `ibm_en_subscription_msteams_invalid_name` | `name` |
| `ibm_en_subscription_msteams_invalid_topic_id` | `topic_id` |

## Example

### Valid

```hcl
resource "ibm_en_subscription_msteams" "test" {
  destination_id = "01234567-89ab-cdef-0123-456789abcdef"
  instance_guid = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
  name = "aaaaaaa"
  topic_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
