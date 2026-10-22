# ibm_pi_placement_group_data

Reads ibm_pi_placement_group. The name of the placement group.

This rule validates `ibm_pi_placement_group` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pi_placement_group_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `crn` | no | — |
| `members` | no | — |
| `name` | no | — |
| `pi_cloud_instance_id` | yes | — |
| `pi_placement_group_id` | no | — |
| `pi_placement_group_name` | no | — |
| `policy` | no | — |
| `user_tags` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_placement_group_data_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |
| `ibm_pi_placement_group_data_invalid_pi_placement_group_id` | `pi_placement_group_id` |
| `ibm_pi_placement_group_data_invalid_pi_placement_group_name` | `pi_placement_group_name` |

## Example

### Valid

```hcl
data "ibm_pi_placement_group" "test" {
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
