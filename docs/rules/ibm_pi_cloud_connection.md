# ibm_pi_cloud_connection

Schema validation for the ibm_pi_cloud_connection resource.

This rule validates `ibm_pi_cloud_connection` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pi_cloud_connection" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cloud_connection_id` | no | — |
| `connection_mode` | no | — |
| `gre_source_address` | no | — |
| `ibm_ip_address` | no | — |
| `pi_cloud_connection_classic_enabled` | no | — |
| `pi_cloud_connection_global_routing` | no | — |
| `pi_cloud_connection_gre_cidr` | no | cidr format |
| `pi_cloud_connection_gre_destination_address` | no | — |
| `pi_cloud_connection_metered` | no | — |
| `pi_cloud_connection_name` | yes | — |
| `pi_cloud_connection_networks` | no | — |
| `pi_cloud_connection_speed` | yes | — |
| `pi_cloud_connection_transit_enabled` | no | — |
| `pi_cloud_connection_vpc_crns` | no | — |
| `pi_cloud_connection_vpc_enabled` | no | — |
| `pi_cloud_instance_id` | yes | — |
| `port` | no | — |
| `status` | no | — |
| `user_ip_address` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_cloud_connection_invalid_pi_cloud_connection_classic_enabled` | `pi_cloud_connection_classic_enabled` |
| `ibm_pi_cloud_connection_invalid_pi_cloud_connection_global_routing` | `pi_cloud_connection_global_routing` |
| `ibm_pi_cloud_connection_invalid_pi_cloud_connection_gre_cidr` | `pi_cloud_connection_gre_cidr` |
| `ibm_pi_cloud_connection_invalid_pi_cloud_connection_gre_destination_address` | `pi_cloud_connection_gre_destination_address` |
| `ibm_pi_cloud_connection_invalid_pi_cloud_connection_metered` | `pi_cloud_connection_metered` |
| `ibm_pi_cloud_connection_invalid_pi_cloud_connection_name` | `pi_cloud_connection_name` |
| `ibm_pi_cloud_connection_invalid_pi_cloud_connection_networks` | `pi_cloud_connection_networks` |
| `ibm_pi_cloud_connection_invalid_pi_cloud_connection_speed` | `pi_cloud_connection_speed` |
| `ibm_pi_cloud_connection_invalid_pi_cloud_connection_transit_enabled` | `pi_cloud_connection_transit_enabled` |
| `ibm_pi_cloud_connection_invalid_pi_cloud_connection_vpc_crns` | `pi_cloud_connection_vpc_crns` |
| `ibm_pi_cloud_connection_invalid_pi_cloud_connection_vpc_enabled` | `pi_cloud_connection_vpc_enabled` |
| `ibm_pi_cloud_connection_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |

## Example

### Valid

```hcl
resource "ibm_pi_cloud_connection" "test" {
  pi_cloud_connection_name = "example"
  pi_cloud_connection_speed = 1
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
