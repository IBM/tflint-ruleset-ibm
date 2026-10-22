# ibm_notification_distribution_list_destination

Schema validation for the ibm_notification_distribution_list_destination resource.

This rule validates `ibm_notification_distribution_list_destination` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `AddDestinationPrototype` (`platform-notifications`).

## Configuration

```hcl
rule "ibm_notification_distribution_list_destination" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_id` | yes | pattern `^[0-9a-zA-Z]{1,100}$` |
| `destination_id` | no | pattern `^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$` |
| `destination_type` | yes | one of: `event_notifications` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_notification_distribution_list_destination_invalid_account_id` | `account_id` |
| `ibm_notification_distribution_list_destination_invalid_destination_id` | `destination_id` |
| `ibm_notification_distribution_list_destination_invalid_destination_type` | `destination_type` |

## Example

### Valid

```hcl
resource "ibm_notification_distribution_list_destination" "test" {
  account_id = "example"
  destination_type = "event_notifications"
}
```
