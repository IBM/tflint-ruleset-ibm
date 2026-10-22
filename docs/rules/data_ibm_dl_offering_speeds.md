# ibm_dl_offering_speeds_data

Schema validation for the ibm_dl_offering_speeds data source.

This rule validates `ibm_dl_offering_speeds` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `OfferingSpeed` (`direct_link`).

## Configuration

```hcl
rule "ibm_dl_offering_speeds_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `offering_speeds` | no | — |
| `offering_type` | yes | one of: `dedicated`, `connect` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_dl_offering_speeds_data_invalid_offering_type` | `offering_type` |

## Example

### Valid

```hcl
data "ibm_dl_offering_speeds" "test" {
  offering_type = "dedicated"
}
```
