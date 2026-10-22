# ibm_dl_export_route_filter_data

Schema validation for the ibm_dl_export_route_filter data source.

This rule validates `ibm_dl_export_route_filter` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_dl_export_route_filter_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `action` | no | — |
| `before` | no | — |
| `created_at` | no | — |
| `ex_filter_id` | yes | — |
| `gateway` | yes | — |
| `ge` | no | — |
| `le` | no | — |
| `prefix` | no | — |
| `updated_at` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_dl_export_route_filter_data_invalid_ex_filter_id` | `ex_filter_id` |
| `ibm_dl_export_route_filter_data_invalid_gateway` | `gateway` |

## Example

### Valid

```hcl
data "ibm_dl_export_route_filter" "test" {
  ex_filter_id = "01234567-89ab-cdef-0123-456789abcdef"
  gateway = "example"
}
```
