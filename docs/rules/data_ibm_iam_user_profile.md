# ibm_iam_user_profile_data

Schema validation for the ibm_iam_user_profile data source.

This rule validates `ibm_iam_user_profile` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `UserMfa` (`iam-identity-token-api`).

## Configuration

```hcl
rule "ibm_iam_user_profile_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_id` | no | — |
| `allowed_ip_addresses` | no | — |
| `altphonenumber` | no | — |
| `email` | no | — |
| `firstname` | no | — |
| `iam_id` | yes | API: The iam_id of the user. |
| `ibm_id` | no | — |
| `lastname` | no | — |
| `phonenumber` | no | — |
| `state` | no | — |
| `user_id` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_user_profile_data_invalid_iam_id` | `iam_id` |

## Example

### Valid

```hcl
data "ibm_iam_user_profile" "test" {
  iam_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
