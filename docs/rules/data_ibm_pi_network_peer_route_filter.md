# ibm_pi_network_peer_route_filter_data

Schema validation for the ibm_pi_network_peer_route_filter data source.

This rule validates `ibm_pi_network_peer_route_filter` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pi_network_peer_route_filter_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `action` | no | — |
| `creation_date` | no | — |
| `direction` | no | — |
| `error` | no | — |
| `ge` | no | — |
| `index` | no | — |
| `le` | no | — |
| `pi_cloud_instance_id` | yes | — |
| `pi_network_peer_id` | yes | — |
| `pi_route_filter_id` | yes | — |
| `prefix` | no | — |
| `route_filter_id` | no | — |
| `state` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_network_peer_route_filter_data_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |
| `ibm_pi_network_peer_route_filter_data_invalid_pi_network_peer_id` | `pi_network_peer_id` |
| `ibm_pi_network_peer_route_filter_data_invalid_pi_route_filter_id` | `pi_route_filter_id` |

## Example

### Valid

```hcl
data "ibm_pi_network_peer_route_filter" "test" {
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  pi_network_peer_id = "01234567-89ab-cdef-0123-456789abcdef"
  pi_route_filter_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
