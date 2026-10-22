# ibm_pi_instance_action

Schema validation for the ibm_pi_instance_action resource.

This rule validates `ibm_pi_instance_action` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pi_instance_action" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `health_status` | no | — |
| `pi_action` | yes | one of: `dumprestart`, `hard-reboot`, `immediate-shutdown`, `reset-state`, `start`, `stop`, `soft-reboot` |
| `pi_cloud_instance_id` | yes | — |
| `pi_health_status` | no | one of: `OK`, `WARNING` |
| `pi_instance_id` | yes | — |
| `progress` | no | — |
| `status` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_instance_action_invalid_pi_action` | `pi_action` |
| `ibm_pi_instance_action_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |
| `ibm_pi_instance_action_invalid_pi_health_status` | `pi_health_status` |
| `ibm_pi_instance_action_invalid_pi_instance_id` | `pi_instance_id` |

## Example

### Valid

```hcl
resource "ibm_pi_instance_action" "test" {
  pi_action = "dumprestart"
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  pi_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
