# ibm_sm_en_registration

Schema validation for the ibm_sm_en_registration resource.

This rule validates `ibm_sm_en_registration` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `NotificationsRegistrationPrototype` (`secrets-manager-secrets-manager-v2`).

## Configuration

```hcl
rule "ibm_sm_en_registration" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `event_notifications_instance_crn` | yes | pattern `^crn:v[0-9](:([A-Za-z0-9-._~!$&'()*+,;=@\/]|%[0-9A-Z]{2})*){8}$` |
| `event_notifications_source_description` | no | API: An optional description for the source that is in your Event Notifications instance. |
| `event_notifications_source_name` | yes | API: The name that is displayed as a source that is in your Event Notifications instance. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_sm_en_registration_invalid_event_notifications_instance_crn` | `event_notifications_instance_crn` |
| `ibm_sm_en_registration_invalid_event_notifications_source_description` | `event_notifications_source_description` |
| `ibm_sm_en_registration_invalid_event_notifications_source_name` | `event_notifications_source_name` |

## Example

### Valid

```hcl
resource "ibm_sm_en_registration" "test" {
  event_notifications_instance_crn = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
  event_notifications_source_name = "aaaaaaa"
}
```
