# ibm_pi_dhcp

Schema validation for the ibm_pi_dhcp resource.

This rule validates `ibm_pi_dhcp` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pi_dhcp" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `dhcp_id` | no | — |
| `leases` | no | — |
| `network_id` | no | — |
| `network_name` | no | — |
| `pi_cidr` | no | cidr format |
| `pi_cloud_connection_id` | no | — |
| `pi_cloud_instance_id` | yes | — |
| `pi_dhcp_name` | no | — |
| `pi_dhcp_snat_enabled` | no | — |
| `pi_dns_server` | no | — |
| `status` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_dhcp_invalid_pi_cidr` | `pi_cidr` |
| `ibm_pi_dhcp_invalid_pi_cloud_connection_id` | `pi_cloud_connection_id` |
| `ibm_pi_dhcp_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |
| `ibm_pi_dhcp_invalid_pi_dhcp_name` | `pi_dhcp_name` |
| `ibm_pi_dhcp_invalid_pi_dhcp_snat_enabled` | `pi_dhcp_snat_enabled` |
| `ibm_pi_dhcp_invalid_pi_dns_server` | `pi_dns_server` |

## Example

### Valid

```hcl
resource "ibm_pi_dhcp" "test" {
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
