# ibm_dl_route_reports_data

Schema validation for the ibm_dl_route_reports data source.

This rule validates `ibm_dl_route_reports` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_dl_route_reports_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `gateway` | yes | — |
| `route_reports` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_dl_route_reports_data_invalid_gateway` | `gateway` |

## Example

### Valid

```hcl
data "ibm_dl_route_reports" "test" {
  gateway = "example"
}
```
