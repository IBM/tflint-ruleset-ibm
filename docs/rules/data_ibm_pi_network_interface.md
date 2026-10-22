# ibm_pi_network_interface_data

Reads ibm_pi_network_interface. Name of the Network Interface (not unique or indexable).

This rule validates `ibm_pi_network_interface` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pi_network_interface_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `crn` | no | — |
| `external_ip` | no | — |
| `instance` | no | — |
| `ip_address` | no | — |
| `mac_address` | no | — |
| `name` | no | — |
| `network_interface_id` | no | — |
| `network_security_group_id` | no | — |
| `network_security_group_ids` | no | — |
| `pi_cloud_instance_id` | yes | — |
| `pi_network_id` | yes | — |
| `pi_network_interface_id` | yes | — |
| `status` | no | — |
| `user_tags` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_network_interface_data_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |
| `ibm_pi_network_interface_data_invalid_pi_network_id` | `pi_network_id` |
| `ibm_pi_network_interface_data_invalid_pi_network_interface_id` | `pi_network_interface_id` |

## Example

### Valid

```hcl
data "ibm_pi_network_interface" "test" {
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  pi_network_id = "01234567-89ab-cdef-0123-456789abcdef"
  pi_network_interface_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
