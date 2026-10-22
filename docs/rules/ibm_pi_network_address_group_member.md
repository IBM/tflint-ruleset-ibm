# ibm_pi_network_address_group_member

Manages ibm_pi_network_address_group_member. The name of the Network Address Group.

This rule validates `ibm_pi_network_address_group_member` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pi_network_address_group_member" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `crn` | no | — |
| `members` | no | — |
| `name` | no | — |
| `pi_cidr` | no | cidr format |
| `pi_cloud_instance_id` | yes | — |
| `pi_network_address_group_id` | yes | — |
| `pi_network_address_group_member_id` | no | — |
| `user_tags` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_network_address_group_member_invalid_pi_cidr` | `pi_cidr` |
| `ibm_pi_network_address_group_member_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |
| `ibm_pi_network_address_group_member_invalid_pi_network_address_group_id` | `pi_network_address_group_id` |
| `ibm_pi_network_address_group_member_invalid_pi_network_address_group_member_id` | `pi_network_address_group_member_id` |

## Example

### Valid

```hcl
resource "ibm_pi_network_address_group_member" "test" {
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  pi_network_address_group_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
