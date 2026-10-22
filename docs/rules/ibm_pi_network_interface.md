# ibm_pi_network_interface

Manages ibm_pi_network_interface. Name of the network interface (not unique or indexable).

This rule validates `ibm_pi_network_interface` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pi_network_interface" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `crn` | no | — |
| `external_ip` | no | — |
| `ip_address` | no | — |
| `mac_address` | no | — |
| `name` | no | — |
| `network_interface_id` | no | — |
| `network_security_group_id` | no | — |
| `network_security_group_ids` | no | — |
| `pi_cloud_instance_id` | yes | — |
| `pi_instance_id` | no | — |
| `pi_ip_address` | no | — |
| `pi_name` | no | — |
| `pi_network_id` | yes | — |
| `pi_user_tags` | no | — |
| `status` | no | — |
| `instance` (block) | no | nested block |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_network_interface_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |
| `ibm_pi_network_interface_invalid_pi_instance_id` | `pi_instance_id` |
| `ibm_pi_network_interface_invalid_pi_ip_address` | `pi_ip_address` |
| `ibm_pi_network_interface_invalid_pi_name` | `pi_name` |
| `ibm_pi_network_interface_invalid_pi_network_id` | `pi_network_id` |
| `ibm_pi_network_interface_invalid_pi_user_tags` | `pi_user_tags` |

## Example

### Valid

```hcl
resource "ibm_pi_network_interface" "test" {
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  pi_network_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
