# ibm_iam_user_settings

Schema validation for the ibm_iam_user_settings resource.

This rule validates `ibm_iam_user_settings` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `UserMfa` (`iam-identity-token-api`).

## Configuration

```hcl
rule "ibm_iam_user_settings" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `allowed_ip_addresses` | no | — |
| `iam_id` | yes | API: The iam_id of the user. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_user_settings_invalid_allowed_ip_addresses` | `allowed_ip_addresses` |
| `ibm_iam_user_settings_invalid_iam_id` | `iam_id` |

## Example

### Valid

```hcl
resource "ibm_iam_user_settings" "test" {
  iam_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
