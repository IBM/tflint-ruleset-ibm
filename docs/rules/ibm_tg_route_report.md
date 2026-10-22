# ibm_tg_route_report

Schema validation for the ibm_tg_route_report resource.

This rule validates `ibm_tg_route_report` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_tg_route_report" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `connections` | no | — |
| `created_at` | no | — |
| `gateway` | yes | — |
| `overlapping_routes` | no | — |
| `route_report_id` | no | — |
| `status` | no | — |
| `updated_at` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_tg_route_report_invalid_gateway` | `gateway` |

## Example

### Valid

```hcl
resource "ibm_tg_route_report" "test" {
  gateway = "example"
}
```
