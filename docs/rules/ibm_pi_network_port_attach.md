# ibm_pi_network_port_attach

Schema validation for the ibm_pi_network_port_attach resource.

This rule validates `ibm_pi_network_port_attach` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pi_network_port_attach" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `mac_address` | no | — |
| `macaddress` | no | — |
| `network_port_id` | no | — |
| `pi_cloud_instance_id` | yes | — |
| `pi_instance_id` | yes | — |
| `pi_network_name` | yes | — |
| `pi_network_port_description` | no | — |
| `pi_network_port_ipaddress` | no | — |
| `pi_user_tags` | no | — |
| `public_ip` | no | — |
| `status` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_network_port_attach_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |
| `ibm_pi_network_port_attach_invalid_pi_instance_id` | `pi_instance_id` |
| `ibm_pi_network_port_attach_invalid_pi_network_name` | `pi_network_name` |
| `ibm_pi_network_port_attach_invalid_pi_network_port_description` | `pi_network_port_description` |
| `ibm_pi_network_port_attach_invalid_pi_network_port_ipaddress` | `pi_network_port_ipaddress` |
| `ibm_pi_network_port_attach_invalid_pi_user_tags` | `pi_user_tags` |

## Example

### Valid

```hcl
resource "ibm_pi_network_port_attach" "test" {
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  pi_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  pi_network_name = "example"
}
```
