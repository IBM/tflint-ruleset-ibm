# ibm_iam_users_data

Schema validation for the ibm_iam_users data source.

This rule validates `ibm_iam_users` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_iam_users_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `users` | no | — |

## Example

### Valid

```hcl
data "ibm_iam_users" "test" {
}
```
