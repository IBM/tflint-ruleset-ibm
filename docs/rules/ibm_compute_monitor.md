# ibm_compute_monitor

Schema validation for the ibm_compute_monitor resource.

This rule validates `ibm_compute_monitor` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_compute_monitor" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `guest_id` | yes | — |
| `ip_address` | no | ipv4 format |
| `notified_users` | no | — |
| `query_type_id` | yes | — |
| `response_action_id` | yes | — |
| `tags` | no | — |
| `wait_cycles` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_compute_monitor_invalid_guest_id` | `guest_id` |
| `ibm_compute_monitor_invalid_ip_address` | `ip_address` |
| `ibm_compute_monitor_invalid_notified_users` | `notified_users` |
| `ibm_compute_monitor_invalid_query_type_id` | `query_type_id` |
| `ibm_compute_monitor_invalid_response_action_id` | `response_action_id` |
| `ibm_compute_monitor_invalid_tags` | `tags` |
| `ibm_compute_monitor_invalid_wait_cycles` | `wait_cycles` |

## Example

### Valid

```hcl
resource "ibm_compute_monitor" "test" {
  guest_id = 1
  query_type_id = 1
  response_action_id = 1
}
```
