# ibm_pi_instance_ip_data

Schema validation for the ibm_pi_instance_ip data source.

This rule validates `ibm_pi_instance_ip` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pi_instance_ip_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `external_ip` | no | — |
| `ip` | no | — |
| `ipoctet` | no | — |
| `mac_address` | no | — |
| `macaddress` | no | — |
| `network_id` | no | — |
| `network_interface_id` | no | — |
| `network_security_group_ids` | no | — |
| `network_security_groups_href` | no | — |
| `pi_cloud_instance_id` | yes | — |
| `pi_instance_id` | no | — |
| `pi_instance_name` | no | — |
| `pi_network_name` | yes | — |
| `type` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_instance_ip_data_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |
| `ibm_pi_instance_ip_data_invalid_pi_instance_id` | `pi_instance_id` |
| `ibm_pi_instance_ip_data_invalid_pi_instance_name` | `pi_instance_name` |
| `ibm_pi_instance_ip_data_invalid_pi_network_name` | `pi_network_name` |

## Example

### Valid

```hcl
data "ibm_pi_instance_ip" "test" {
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  pi_network_name = "example"
}
```
