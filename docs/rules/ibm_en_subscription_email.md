# ibm_en_subscription_email

Manages ibm_en_subscription_email. Subscription name.

This rule validates `ibm_en_subscription_email` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Subscription` (`event-notifications`).

## Configuration

```hcl
rule "ibm_en_subscription_email" {
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
| `from` | no | — |
| `instance_guid` | yes | — |
| `name` | yes | API: Subscription name |
| `subscription_id` | no | — |
| `topic_id` | yes | pattern `[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}` |
| `topic_name` | no | — |
| `updated_at` | no | — |
| `attributes` (block) | no | nested block |
| `attributes.add` | no | API: The email ids or phone numbers to be invited |
| `attributes.add_notification_payload` | no | API: Decision for Notification Payload to be added |
| `attributes.from_name` | no | API: The email name of From |
| `attributes.invited` | no | API: The email id string |
| `attributes.remove` | no | API: The email ids or phone numbers for removal |
| `attributes.reply_to_mail` | no | pattern `^[a-zA-Z0-9+_.-]+@[a-zA-Z0-9.-]+$` |
| `attributes.reply_to_name` | no | API: The email name to reply to |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_en_subscription_email_invalid_description` | `description` |
| `ibm_en_subscription_email_invalid_destination_id` | `destination_id` |
| `ibm_en_subscription_email_invalid_instance_guid` | `instance_guid` |
| `ibm_en_subscription_email_invalid_name` | `name` |
| `ibm_en_subscription_email_invalid_topic_id` | `topic_id` |
| `ibm_en_subscription_email_invalid_attributes__add` | `attributes.add` |
| `ibm_en_subscription_email_invalid_attributes__add_notification_payload` | `attributes.add_notification_payload` |
| `ibm_en_subscription_email_invalid_attributes__from_name` | `attributes.from_name` |
| `ibm_en_subscription_email_invalid_attributes__invited` | `attributes.invited` |
| `ibm_en_subscription_email_invalid_attributes__remove` | `attributes.remove` |
| `ibm_en_subscription_email_invalid_attributes__reply_to_mail` | `attributes.reply_to_mail` |
| `ibm_en_subscription_email_invalid_attributes__reply_to_name` | `attributes.reply_to_name` |

## Example

### Valid

```hcl
resource "ibm_en_subscription_email" "test" {
  destination_id = "01234567-89ab-cdef-0123-456789abcdef"
  instance_guid = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
  name = "aaaaaaa"
  topic_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
