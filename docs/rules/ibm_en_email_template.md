# ibm_en_email_template

Manages ibm_en_email_template. Subscription name.

This rule validates `ibm_en_email_template` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `TemplateCreate` (`event-notifications`).

## Configuration

```hcl
rule "ibm_en_email_template" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `description` | no | API: The Template description |
| `instance_guid` | yes | — |
| `name` | yes | API: The Message Template |
| `subscription_count` | no | — |
| `subscription_names` | no | — |
| `template_id` | no | — |
| `type` | yes | pattern `^(smtp_custom.notification|smtp_custom.invitation|slack.notification|webhook.notification|pagerduty.notification|event_streams.notification|ibmcejob.notification|ibmceapp.notification|app_configuration.notification)$` |
| `updated_at` | no | — |
| `params` (block) | no | nested block |
| `params.body` | no | API: Template body(Base64 encoded) |
| `params.subject` | no | API: The template subject |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_en_email_template_invalid_description` | `description` |
| `ibm_en_email_template_invalid_instance_guid` | `instance_guid` |
| `ibm_en_email_template_invalid_name` | `name` |
| `ibm_en_email_template_invalid_type` | `type` |
| `ibm_en_email_template_invalid_params__body` | `params.body` |
| `ibm_en_email_template_invalid_params__subject` | `params.subject` |

## Example

### Valid

```hcl
resource "ibm_en_email_template" "test" {
  instance_guid = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
  name = "aaaaaaa"
  type = "pagerduty.notification"
}
```
