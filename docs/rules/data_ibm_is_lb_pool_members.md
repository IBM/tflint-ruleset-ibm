# ibm_is_lb_pool_members_data

Schema validation for the ibm_is_lb_pool_members data source.

This rule validates `ibm_is_lb_pool_members` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `BareMetalServerReservationAffinityPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_lb_pool_members_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `lb` | yes | — |
| `members` | no | — |
| `pool` | yes | API: Identifies a reservation by a unique property. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_lb_pool_members_data_invalid_lb` | `lb` |
| `ibm_is_lb_pool_members_data_invalid_pool` | `pool` |

## Example

### Valid

```hcl
data "ibm_is_lb_pool_members" "test" {
  lb = "example"
  pool = "example"
}
```
