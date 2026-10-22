# ibm_pi_placement_group

Schema validation for the ibm_pi_placement_group resource.

This rule validates `ibm_pi_placement_group` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pi_placement_group" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `crn` | no | — |
| `members` | no | — |
| `pi_cloud_instance_id` | yes | — |
| `pi_placement_group_name` | yes | — |
| `pi_placement_group_policy` | yes | one of: `affinity`, `anti-affinity` |
| `pi_user_tags` | no | — |
| `placement_group_id` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_placement_group_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |
| `ibm_pi_placement_group_invalid_pi_placement_group_name` | `pi_placement_group_name` |
| `ibm_pi_placement_group_invalid_pi_placement_group_policy` | `pi_placement_group_policy` |
| `ibm_pi_placement_group_invalid_pi_user_tags` | `pi_user_tags` |

## Example

### Valid

```hcl
resource "ibm_pi_placement_group" "test" {
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  pi_placement_group_name = "example"
  pi_placement_group_policy = "affinity"
}
```
