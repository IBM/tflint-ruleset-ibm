# ibm_pi_network_peer_route_filter

Schema validation for the ibm_pi_network_peer_route_filter resource.

This rule validates `ibm_pi_network_peer_route_filter` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pi_network_peer_route_filter" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `creation_date` | no | — |
| `error` | no | — |
| `pi_action` | no | one of: `allow`, `deny` |
| `pi_cloud_instance_id` | yes | — |
| `pi_direction` | yes | one of: `export`, `import` |
| `pi_ge` | no | — |
| `pi_index` | yes | — |
| `pi_le` | no | — |
| `pi_network_peer_id` | yes | — |
| `pi_prefix` | yes | — |
| `route_filter_id` | no | — |
| `state` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_network_peer_route_filter_invalid_pi_action` | `pi_action` |
| `ibm_pi_network_peer_route_filter_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |
| `ibm_pi_network_peer_route_filter_invalid_pi_direction` | `pi_direction` |
| `ibm_pi_network_peer_route_filter_invalid_pi_ge` | `pi_ge` |
| `ibm_pi_network_peer_route_filter_invalid_pi_index` | `pi_index` |
| `ibm_pi_network_peer_route_filter_invalid_pi_le` | `pi_le` |
| `ibm_pi_network_peer_route_filter_invalid_pi_network_peer_id` | `pi_network_peer_id` |
| `ibm_pi_network_peer_route_filter_invalid_pi_prefix` | `pi_prefix` |

## Example

### Valid

```hcl
resource "ibm_pi_network_peer_route_filter" "test" {
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  pi_direction = "export"
  pi_index = 1
  pi_network_peer_id = "01234567-89ab-cdef-0123-456789abcdef"
  pi_prefix = "example"
}
```
