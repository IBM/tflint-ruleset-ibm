# ibm_dl_route_report

Schema validation for the ibm_dl_route_report resource.

This rule validates `ibm_dl_route_report` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_dl_route_report" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `advertised_routes` | no | — |
| `created_at` | no | — |
| `gateway` | yes | — |
| `gateway_routes` | no | — |
| `on_prem_routes` | no | — |
| `overlapping_routes` | no | — |
| `route_report_id` | no | — |
| `status` | no | — |
| `updated_at` | no | — |
| `virtual_connection_routes` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_dl_route_report_invalid_gateway` | `gateway` |

## Example

### Valid

```hcl
resource "ibm_dl_route_report" "test" {
  gateway = "example"
}
```
