# ibm_dl_route_report_data

Schema validation for the ibm_dl_route_report data source.

This rule validates `ibm_dl_route_report` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_dl_route_report_data" {
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
| `route_report` | yes | — |
| `status` | no | — |
| `updated_at` | no | — |
| `virtual_connection_routes` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_dl_route_report_data_invalid_gateway` | `gateway` |
| `ibm_dl_route_report_data_invalid_route_report` | `route_report` |

## Example

### Valid

```hcl
data "ibm_dl_route_report" "test" {
  gateway = "example"
  route_report = "example"
}
```
