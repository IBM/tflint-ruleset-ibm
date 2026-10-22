# ibm_is_reservation_activate

Manages ibm_is_reservation_activate. Reservation name

This rule validates `ibm_is_reservation_activate` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ReservationCommittedUsePatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_reservation_activate" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `affinity_policy` | no | — |
| `capacity` | no | — |
| `committed_use` | no | — |
| `created_at` | no | — |
| `crn` | no | — |
| `href` | no | — |
| `id` | no | — |
| `lifecycle_state` | no | — |
| `name` | no | — |
| `profile` | no | — |
| `reservation` | yes | — |
| `resource_type` | no | — |
| `status` | no | — |
| `status_reasons` | no | — |
| `zone` | no | — |
| `resource_group` (block) | no | nested block |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_reservation_activate_invalid_reservation` | `reservation` |

## Example

### Valid

```hcl
resource "ibm_is_reservation_activate" "test" {
  reservation = "example"
}
```
