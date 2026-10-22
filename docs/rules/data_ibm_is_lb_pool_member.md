# ibm_is_lb_pool_member_data

Schema validation for the ibm_is_lb_pool_member data source.

This rule validates `ibm_is_lb_pool_member` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `BareMetalServerReservationAffinityPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_lb_pool_member_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `created_at` | no | — |
| `health` | no | — |
| `href` | no | — |
| `lb` | yes | — |
| `member` | yes | — |
| `pool` | yes | API: Identifies a reservation by a unique property. |
| `port` | no | — |
| `provisioning_status` | no | — |
| `target` | no | — |
| `weight` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_lb_pool_member_data_invalid_lb` | `lb` |
| `ibm_is_lb_pool_member_data_invalid_member` | `member` |
| `ibm_is_lb_pool_member_data_invalid_pool` | `pool` |

## Example

### Valid

```hcl
data "ibm_is_lb_pool_member" "test" {
  lb = "example"
  member = "example"
  pool = "example"
}
```
