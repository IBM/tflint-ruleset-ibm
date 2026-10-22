# ibm_pi_network_security_group_action

Schema validation for the ibm_pi_network_security_group_action resource.

This rule validates `ibm_pi_network_security_group_action` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pi_network_security_group_action" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `pi_action` | yes | one of: `disable`, `enable` |
| `pi_cloud_instance_id` | yes | — |
| `state` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_network_security_group_action_invalid_pi_action` | `pi_action` |
| `ibm_pi_network_security_group_action_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |

## Example

### Valid

```hcl
resource "ibm_pi_network_security_group_action" "test" {
  pi_action = "disable"
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
