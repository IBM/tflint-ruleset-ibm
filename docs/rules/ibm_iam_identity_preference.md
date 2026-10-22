# ibm_iam_identity_preference

Schema validation for the ibm_iam_identity_preference resource.

This rule validates `ibm_iam_identity_preference` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `params:put:/v1/preferences/accounts/{account_id}/identities/{iam_id}/{service}/{preference_id}` (`iam-identity-token-api`).

## Configuration

```hcl
rule "ibm_iam_identity_preference" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_id` | yes | — |
| `iam_id` | yes | — |
| `preference_id` | yes | — |
| `scope` | no | — |
| `service` | yes | — |
| `value_list_of_strings` | no | — |
| `value_string` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_identity_preference_invalid_account_id` | `account_id` |
| `ibm_iam_identity_preference_invalid_iam_id` | `iam_id` |
| `ibm_iam_identity_preference_invalid_preference_id` | `preference_id` |
| `ibm_iam_identity_preference_invalid_service` | `service` |
| `ibm_iam_identity_preference_invalid_value_list_of_strings` | `value_list_of_strings` |
| `ibm_iam_identity_preference_invalid_value_string` | `value_string` |

## Example

### Valid

```hcl
resource "ibm_iam_identity_preference" "test" {
  account_id = "01234567-89ab-cdef-0123-456789abcdef"
  iam_id = "01234567-89ab-cdef-0123-456789abcdef"
  preference_id = "01234567-89ab-cdef-0123-456789abcdef"
  service = "example"
}
```
