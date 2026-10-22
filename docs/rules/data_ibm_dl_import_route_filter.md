# ibm_dl_import_route_filter_data

Schema validation for the ibm_dl_import_route_filter data source.

This rule validates `ibm_dl_import_route_filter` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_dl_import_route_filter_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `action` | no | — |
| `before` | no | — |
| `created_at` | no | — |
| `gateway` | yes | — |
| `ge` | no | — |
| `im_filter_id` | yes | — |
| `le` | no | — |
| `prefix` | no | — |
| `updated_at` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_dl_import_route_filter_data_invalid_gateway` | `gateway` |
| `ibm_dl_import_route_filter_data_invalid_im_filter_id` | `im_filter_id` |

## Example

### Valid

```hcl
data "ibm_dl_import_route_filter" "test" {
  gateway = "example"
  im_filter_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
