# ibm_dl_gateways_data

Schema validation for the ibm_dl_gateways data source.

This rule validates `ibm_dl_gateways` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/gateways/{id}/macsec:patch` (`direct_link`).

## Configuration

```hcl
rule "ibm_dl_gateways_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `gateways` | no | — |

## Example

### Valid

```hcl
data "ibm_dl_gateways" "test" {
}
```
