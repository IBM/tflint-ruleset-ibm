# ibm_dl_export_route_filters_data

Schema validation for the ibm_dl_export_route_filters data source.

This rule validates `ibm_dl_export_route_filters` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_dl_export_route_filters_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `export_route_filters` | no | — |
| `gateway` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_dl_export_route_filters_data_invalid_gateway` | `gateway` |

## Example

### Valid

```hcl
data "ibm_dl_export_route_filters" "test" {
  gateway = "example"
}
```
