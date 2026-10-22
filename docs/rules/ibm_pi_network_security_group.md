# ibm_pi_network_security_group

Schema validation for the ibm_pi_network_security_group resource.

This rule validates `ibm_pi_network_security_group` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pi_network_security_group" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `crn` | no | — |
| `default` | no | — |
| `members` | no | — |
| `network_security_group_id` | no | — |
| `pi_cloud_instance_id` | yes | — |
| `pi_name` | yes | — |
| `pi_user_tags` | no | — |
| `rules` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_network_security_group_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |
| `ibm_pi_network_security_group_invalid_pi_name` | `pi_name` |
| `ibm_pi_network_security_group_invalid_pi_user_tags` | `pi_user_tags` |

## Example

### Valid

```hcl
resource "ibm_pi_network_security_group" "test" {
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  pi_name = "example"
}
```
