# ibm_pi_network_peer_data

Reads ibm_pi_network_peer. User defined name.

This rule validates `ibm_pi_network_peer` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pi_network_peer_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `creation_date` | no | — |
| `customer_asn` | no | — |
| `customer_cidr` | no | — |
| `default_export_route_filter` | no | — |
| `default_import_route_filter` | no | — |
| `error` | no | — |
| `ibm_asn` | no | — |
| `ibm_cidr` | no | — |
| `name` | no | — |
| `peer_interface_id` | no | — |
| `pi_cloud_instance_id` | yes | — |
| `pi_network_peer_id` | yes | — |
| `state` | no | — |
| `type` | no | — |
| `updated_date` | no | — |
| `vlan` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_network_peer_data_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |
| `ibm_pi_network_peer_data_invalid_pi_network_peer_id` | `pi_network_peer_id` |

## Example

### Valid

```hcl
data "ibm_pi_network_peer" "test" {
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  pi_network_peer_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
