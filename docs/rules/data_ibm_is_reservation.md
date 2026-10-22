# ibm_is_reservation_data

Reads ibm_is_reservation. The unique user-defined name for this reservation.

This rule validates `ibm_is_reservation` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `BareMetalServerDiskPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_reservation_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `affinity_policy` | no | — |
| `created_at` | no | — |
| `crn` | no | — |
| `href` | no | — |
| `id` | no | — |
| `identifier` | no | — |
| `lifecycle_state` | no | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `profile` | no | — |
| `resource_group` | no | — |
| `resource_type` | no | — |
| `status` | no | — |
| `status_reasons` | no | — |
| `zone` | no | — |
| `capacity` (block) | no | nested block |
| `committed_use` (block) | no | nested block |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_reservation_data_invalid_identifier` | `identifier` |
| `ibm_is_reservation_data_invalid_name` | `name` |

## Example

### Valid

```hcl
data "ibm_is_reservation" "test" {
}
```
