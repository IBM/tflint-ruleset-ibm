# ibm_cm_account_data

Schema validation for the ibm_cm_account data source.

This rule validates `ibm_cm_account` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Account` (`resource-catalog-private-catalog`).

## Configuration

```hcl
rule "ibm_cm_account_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_filters` | no | — |
| `hide_ibm_cloud_catalog` | no | — |
| `region_filter` | no | — |
| `rev` | no | — |
| `terraform_engines` | no | — |

## Example

### Valid

```hcl
data "ibm_cm_account" "test" {
}
```
