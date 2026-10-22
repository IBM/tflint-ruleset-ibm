# ibm_is_reservation

Manages ibm_is_reservation. Reservation name

This rule validates `ibm_is_reservation` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Reservation` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_reservation" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `affinity_policy` | no | one of: `automatic`, `restricted`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `created_at` | no | — |
| `crn` | no | — |
| `href` | no | — |
| `id` | no | — |
| `lifecycle_state` | no | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `resource_type` | no | — |
| `status` | no | — |
| `status_reasons` | no | — |
| `zone` | yes | zone format |
| `capacity` (block) | yes | nested block |
| `capacity.total` | yes | min 1 |
| `committed_use` (block) | yes | nested block |
| `committed_use.expiration_policy` | no | one of: `release`, `renew`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `committed_use.term` | yes | pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `profile` (block) | yes | nested block |
| `profile.name` | yes | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `profile.resource_type` | yes | one of: `bare_metal_server_profile`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `resource_group` (block) | no | nested block |
| `resource_group.id` | yes | pattern `^[0-9a-f]{32}$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_reservation_invalid_affinity_policy` | `affinity_policy` |
| `ibm_is_reservation_invalid_name` | `name` |
| `ibm_is_reservation_invalid_zone` | `zone` |
| `ibm_is_reservation_invalid_capacity__total` | `capacity.total` |
| `ibm_is_reservation_invalid_committed_use__expiration_policy` | `committed_use.expiration_policy` |
| `ibm_is_reservation_invalid_committed_use__term` | `committed_use.term` |
| `ibm_is_reservation_invalid_profile__name` | `profile.name` |
| `ibm_is_reservation_invalid_profile__resource_type` | `profile.resource_type` |
| `ibm_is_reservation_invalid_resource_group__id` | `resource_group.id` |

## Example

### Valid

```hcl
resource "ibm_is_reservation" "test" {
  zone = "us-south-1"
  capacity {
    total = 1
  }
  committed_use {
    term = "example"
  }
  profile {
    name = "example-name"
    resource_type = "bare_metal_server_profile"
  }
}
```
