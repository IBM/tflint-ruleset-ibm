# ibm_pi_network_security_group_member

Manages ibm_pi_network_security_group_member. The name of the network security group.

This rule validates `ibm_pi_network_security_group_member` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pi_network_security_group_member" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `crn` | no | — |
| `default` | no | — |
| `members` | no | — |
| `name` | no | — |
| `network_security_group_member_id` | no | — |
| `pi_cloud_instance_id` | yes | — |
| `pi_network_security_group_id` | yes | — |
| `pi_network_security_group_member_id` | no | — |
| `pi_target` | no | — |
| `pi_type` | no | one of: `ipv4-address`, `network-interface` |
| `user_tags` | no | — |
| `rules` (block) | no | nested block |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_network_security_group_member_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |
| `ibm_pi_network_security_group_member_invalid_pi_network_security_group_id` | `pi_network_security_group_id` |
| `ibm_pi_network_security_group_member_invalid_pi_network_security_group_member_id` | `pi_network_security_group_member_id` |
| `ibm_pi_network_security_group_member_invalid_pi_target` | `pi_target` |
| `ibm_pi_network_security_group_member_invalid_pi_type` | `pi_type` |

## Example

### Valid

```hcl
resource "ibm_pi_network_security_group_member" "test" {
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  pi_network_security_group_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
