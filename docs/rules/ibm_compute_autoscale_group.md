# ibm_compute_autoscale_group

Manages ibm_compute_autoscale_group. Name

This rule validates `ibm_compute_autoscale_group` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_compute_autoscale_group" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cooldown` | yes | — |
| `health_check` | no | — |
| `maximum_member_count` | yes | — |
| `minimum_member_count` | yes | — |
| `name` | yes | — |
| `network_vlan_ids` | no | — |
| `port` | no | range 1–65535; port format |
| `regional_group` | yes | — |
| `tags` | no | — |
| `termination_policy` | yes | — |
| `virtual_guest_member_template` | yes | — |
| `virtual_server_id` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_compute_autoscale_group_invalid_cooldown` | `cooldown` |
| `ibm_compute_autoscale_group_invalid_health_check` | `health_check` |
| `ibm_compute_autoscale_group_invalid_maximum_member_count` | `maximum_member_count` |
| `ibm_compute_autoscale_group_invalid_minimum_member_count` | `minimum_member_count` |
| `ibm_compute_autoscale_group_invalid_name` | `name` |
| `ibm_compute_autoscale_group_invalid_network_vlan_ids` | `network_vlan_ids` |
| `ibm_compute_autoscale_group_invalid_port` | `port` |
| `ibm_compute_autoscale_group_invalid_regional_group` | `regional_group` |
| `ibm_compute_autoscale_group_invalid_tags` | `tags` |
| `ibm_compute_autoscale_group_invalid_termination_policy` | `termination_policy` |
| `ibm_compute_autoscale_group_invalid_virtual_guest_member_template` | `virtual_guest_member_template` |
| `ibm_compute_autoscale_group_invalid_virtual_server_id` | `virtual_server_id` |

## Example

### Valid

```hcl
resource "ibm_compute_autoscale_group" "test" {
  cooldown = 1
  maximum_member_count = 1
  minimum_member_count = 1
  name = "example"
  regional_group = "example"
  termination_policy = "example"
  virtual_guest_member_template = "example"
}
```
