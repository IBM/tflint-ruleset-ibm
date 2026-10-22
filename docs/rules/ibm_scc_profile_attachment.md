# ibm_scc_profile_attachment

Manages ibm_scc_profile_attachment. The name of the attachment.

This rule validates `ibm_scc_profile_attachment` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Attachment` (`security-compliance`).

## Configuration

```hcl
rule "ibm_scc_profile_attachment" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_id` | no | — |
| `attachment_id` | no | — |
| `attachment_parameters` | no | — |
| `created_by` | no | — |
| `created_on` | no | — |
| `description` | no | API: The description of the attachment. |
| `instance_id` | yes | pattern `^[0-9A-Fa-f]{8}-[0-9A-Fa-f]{4}-4[0-9A-Fa-f]{3}-[89ABab][0-9A-Fa-f]{3}-[0-9A-Fa-f]{12}$|^$` |
| `last_scan` | no | — |
| `name` | yes | API: The name of the attachment. |
| `next_scan_time` | no | — |
| `profile_attachment_id` | no | — |
| `profile_id` | yes | pattern `^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$` |
| `schedule` | yes | API: The attachment schedule. |
| `scope` | yes | — |
| `status` | yes | — |
| `updated_by` | no | — |
| `updated_on` | no | — |
| `notifications` (block) | no | nested block |
| `notifications.enabled` | no | API: Shows if the notification is enabled or disabled. |
| `notifications.controls` (block) | no | nested block |
| `notifications.controls.failed_control_ids` | no | API: List of controls that triggers a notification should a scan fail. |
| `notifications.controls.threshold_limit` | no | API: The maximum number of not compliant controls before a notification is triggered. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_scc_profile_attachment_invalid_attachment_parameters` | `attachment_parameters` |
| `ibm_scc_profile_attachment_invalid_description` | `description` |
| `ibm_scc_profile_attachment_invalid_instance_id` | `instance_id` |
| `ibm_scc_profile_attachment_invalid_name` | `name` |
| `ibm_scc_profile_attachment_invalid_profile_id` | `profile_id` |
| `ibm_scc_profile_attachment_invalid_schedule` | `schedule` |
| `ibm_scc_profile_attachment_invalid_scope` | `scope` |
| `ibm_scc_profile_attachment_invalid_status` | `status` |
| `ibm_scc_profile_attachment_invalid_notifications__enabled` | `notifications.enabled` |
| `ibm_scc_profile_attachment_invalid_notifications__controls__failed_control_ids` | `notifications.controls.failed_control_ids` |
| `ibm_scc_profile_attachment_invalid_notifications__controls__threshold_limit` | `notifications.controls.threshold_limit` |

## Example

### Valid

```hcl
resource "ibm_scc_profile_attachment" "test" {
  instance_id = "01234567-89ab-4cde-8123-456789abcdef"
  name = "example"
  profile_id = "01234567-89ab-cdef-0123-456789abcdef"
  schedule = "example"
  scope = "example"
  status = "example"
}
```
