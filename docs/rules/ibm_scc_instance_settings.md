# ibm_scc_instance_settings

Schema validation for the ibm_scc_instance_settings resource.

This rule validates `ibm_scc_instance_settings` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `EventNotificationsPrototype` (`security-compliance`).

## Configuration

```hcl
rule "ibm_scc_instance_settings" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `instance_id` | yes | pattern `^[a-zA-Z0-9 ,\-_]+$` |
| `event_notifications` (block) | yes | nested block |
| `event_notifications.instance_crn` | no | crn format |
| `event_notifications.source_description` | no | API: The description of the source of the Event Notifications. |
| `event_notifications.source_name` | no | API: The name of the source of the Event Notifications. |
| `object_storage` (block) | yes | nested block |
| `object_storage.bucket` | no | API: The connected Cloud Object Storage bucket name. |
| `object_storage.instance_crn` | no | crn format |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_scc_instance_settings_invalid_instance_id` | `instance_id` |
| `ibm_scc_instance_settings_invalid_event_notifications__instance_crn` | `event_notifications.instance_crn` |
| `ibm_scc_instance_settings_invalid_event_notifications__source_description` | `event_notifications.source_description` |
| `ibm_scc_instance_settings_invalid_event_notifications__source_name` | `event_notifications.source_name` |
| `ibm_scc_instance_settings_invalid_object_storage__bucket` | `object_storage.bucket` |
| `ibm_scc_instance_settings_invalid_object_storage__instance_crn` | `object_storage.instance_crn` |

## Example

### Valid

```hcl
resource "ibm_scc_instance_settings" "test" {
  instance_id = "example-name"
  event_notifications {
  }
  object_storage {
  }
}
```
