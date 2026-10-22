# ibm_pi_network_security_group_data

Reads ibm_pi_network_security_group. The name of the network security group.

This rule validates `ibm_pi_network_security_group` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pi_network_security_group_data" {
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
| `pi_cloud_instance_id` | yes | — |
| `pi_network_security_group_id` | yes | — |
| `rules` | no | — |
| `user_tags` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_network_security_group_data_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |
| `ibm_pi_network_security_group_data_invalid_pi_network_security_group_id` | `pi_network_security_group_id` |

## Example

### Valid

```hcl
data "ibm_pi_network_security_group" "test" {
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  pi_network_security_group_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
