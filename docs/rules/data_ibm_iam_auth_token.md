# ibm_iam_auth_token_data

Schema validation for the ibm_iam_auth_token data source.

This rule validates `ibm_iam_auth_token` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_iam_auth_token_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `iam_access_token` | no | — |
| `iam_refresh_token` | no | — |
| `uaa_access_token` | no | — |
| `uaa_refresh_token` | no | — |

## Example

### Valid

```hcl
data "ibm_iam_auth_token" "test" {
}
```
