# ibm_pi_network_peer

Schema validation for the ibm_pi_network_peer resource.

This rule validates `ibm_pi_network_peer` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pi_network_peer" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `creation_date` | no | — |
| `error` | no | — |
| `peer_id` | no | — |
| `pi_cloud_instance_id` | yes | — |
| `pi_customer_asn` | yes | — |
| `pi_customer_cidr` | yes | cidr format |
| `pi_default_export_route_filter` | no | one of: `allow`, `deny` |
| `pi_default_import_route_filter` | no | one of: `allow`, `deny` |
| `pi_ibm_asn` | yes | — |
| `pi_ibm_cidr` | yes | cidr format |
| `pi_name` | yes | — |
| `pi_peer_interface_id` | yes | — |
| `pi_type` | no | one of: `dcnetwork_bgp` |
| `pi_vlan` | yes | — |
| `state` | no | — |
| `updated_date` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_network_peer_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |
| `ibm_pi_network_peer_invalid_pi_customer_asn` | `pi_customer_asn` |
| `ibm_pi_network_peer_invalid_pi_customer_cidr` | `pi_customer_cidr` |
| `ibm_pi_network_peer_invalid_pi_default_export_route_filter` | `pi_default_export_route_filter` |
| `ibm_pi_network_peer_invalid_pi_default_import_route_filter` | `pi_default_import_route_filter` |
| `ibm_pi_network_peer_invalid_pi_ibm_asn` | `pi_ibm_asn` |
| `ibm_pi_network_peer_invalid_pi_ibm_cidr` | `pi_ibm_cidr` |
| `ibm_pi_network_peer_invalid_pi_name` | `pi_name` |
| `ibm_pi_network_peer_invalid_pi_peer_interface_id` | `pi_peer_interface_id` |
| `ibm_pi_network_peer_invalid_pi_type` | `pi_type` |
| `ibm_pi_network_peer_invalid_pi_vlan` | `pi_vlan` |

## Example

### Valid

```hcl
resource "ibm_pi_network_peer" "test" {
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  pi_customer_asn = 1
  pi_customer_cidr = "10.0.0.0/24"
  pi_ibm_asn = 1
  pi_ibm_cidr = "10.0.0.0/24"
  pi_name = "example"
  pi_peer_interface_id = "01234567-89ab-cdef-0123-456789abcdef"
  pi_vlan = 1
}
```
