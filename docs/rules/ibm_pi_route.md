# ibm_pi_route

Schema validation for the ibm_pi_route resource.

This rule validates `ibm_pi_route` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pi_route" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `crn` | no | — |
| `pi_action` | no | one of: `deliver` |
| `pi_advertise` | no | one of: `enable`, `disable` |
| `pi_cloud_instance_id` | yes | — |
| `pi_destination` | yes | — |
| `pi_destination_type` | no | one of: `ipv4-address` |
| `pi_enabled` | no | — |
| `pi_name` | yes | — |
| `pi_next_hop` | yes | — |
| `pi_next_hop_type` | no | one of: `ipv4-address` |
| `pi_user_tags` | no | — |
| `route_id` | no | — |
| `state` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_route_invalid_pi_action` | `pi_action` |
| `ibm_pi_route_invalid_pi_advertise` | `pi_advertise` |
| `ibm_pi_route_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |
| `ibm_pi_route_invalid_pi_destination` | `pi_destination` |
| `ibm_pi_route_invalid_pi_destination_type` | `pi_destination_type` |
| `ibm_pi_route_invalid_pi_enabled` | `pi_enabled` |
| `ibm_pi_route_invalid_pi_name` | `pi_name` |
| `ibm_pi_route_invalid_pi_next_hop` | `pi_next_hop` |
| `ibm_pi_route_invalid_pi_next_hop_type` | `pi_next_hop_type` |
| `ibm_pi_route_invalid_pi_user_tags` | `pi_user_tags` |

## Example

### Valid

```hcl
resource "ibm_pi_route" "test" {
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  pi_destination = "example"
  pi_name = "example"
  pi_next_hop = "example"
}
```
