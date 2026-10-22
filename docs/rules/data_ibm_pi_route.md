# ibm_pi_route_data

Reads ibm_pi_route. Name of the route.

This rule validates `ibm_pi_route` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pi_route_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `action` | no | — |
| `advertise` | no | — |
| `crn` | no | — |
| `destination` | no | — |
| `destination_type` | no | — |
| `enabled` | no | — |
| `name` | no | — |
| `next_hop` | no | — |
| `next_hop_type` | no | — |
| `pi_cloud_instance_id` | yes | — |
| `pi_route_id` | yes | — |
| `state` | no | — |
| `user_tags` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_route_data_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |
| `ibm_pi_route_data_invalid_pi_route_id` | `pi_route_id` |

## Example

### Valid

```hcl
data "ibm_pi_route" "test" {
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  pi_route_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
