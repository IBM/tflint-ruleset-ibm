# ibm_en_subscription_custom_sms_data

Reads ibm_en_subscription_custom_sms. Subscription name.

This rule validates `ibm_en_subscription_custom_sms` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `SMSInviteAttributesItems` (`event-notifications`).

## Configuration

```hcl
rule "ibm_en_subscription_custom_sms_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `attributes` | no | — |
| `description` | no | — |
| `destination_id` | no | — |
| `instance_guid` | yes | — |
| `name` | no | — |
| `subscription_id` | yes | pattern `[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}` |
| `topic_id` | no | — |
| `updated_at` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_en_subscription_custom_sms_data_invalid_instance_guid` | `instance_guid` |
| `ibm_en_subscription_custom_sms_data_invalid_subscription_id` | `subscription_id` |

## Example

### Valid

```hcl
data "ibm_en_subscription_custom_sms" "test" {
  instance_guid = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
  subscription_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
