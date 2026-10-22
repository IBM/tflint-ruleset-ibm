# ibm_tg_connection_prefix_filter_data

Schema validation for the ibm_tg_connection_prefix_filter data source.

This rule validates `ibm_tg_connection_prefix_filter` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `RouteReportOverlappingRoute` (`transit-gateway`).

## Configuration

```hcl
rule "ibm_tg_connection_prefix_filter_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `action` | no | — |
| `before` | no | — |
| `connection_id` | yes | pattern `^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$` |
| `created_at` | no | — |
| `filter_id` | yes | — |
| `gateway` | yes | — |
| `ge` | no | — |
| `le` | no | — |
| `prefix` | no | — |
| `updated_at` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_tg_connection_prefix_filter_data_invalid_connection_id` | `connection_id` |
| `ibm_tg_connection_prefix_filter_data_invalid_filter_id` | `filter_id` |
| `ibm_tg_connection_prefix_filter_data_invalid_gateway` | `gateway` |

## Example

### Valid

```hcl
data "ibm_tg_connection_prefix_filter" "test" {
  connection_id = "01234567-89ab-cdef-0123-456789abcdef"
  filter_id = "01234567-89ab-cdef-0123-456789abcdef"
  gateway = "example"
}
```
