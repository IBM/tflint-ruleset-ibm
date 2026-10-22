# ibm_iam_idp_account_setting

Schema validation for the ibm_iam_idp_account_setting resource.

This rule validates `ibm_iam_idp_account_setting` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `AddAccountIdpSettingRequest` (`iam-identity-token-api`).

## Configuration

```hcl
rule "ibm_iam_idp_account_setting" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_id` | yes | — |
| `active` | yes | API: Specifies if the IdP is enabled for usage in the given account context |
| `cloud_user_strategy` | yes | one of: `DYNAMIC`, `STATIC`, `NEVER` |
| `idp_id` | yes | pattern `^[A-Za-z0-9]+$` |
| `idp_name` | no | — |
| `idp_type` | no | — |
| `owner_account` | no | — |
| `owner_account_name` | no | — |
| `ui_default` | yes | API: Specifies if the IdP is used as default in the given account context |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_idp_account_setting_invalid_account_id` | `account_id` |
| `ibm_iam_idp_account_setting_invalid_active` | `active` |
| `ibm_iam_idp_account_setting_invalid_cloud_user_strategy` | `cloud_user_strategy` |
| `ibm_iam_idp_account_setting_invalid_idp_id` | `idp_id` |
| `ibm_iam_idp_account_setting_invalid_ui_default` | `ui_default` |

## Example

### Valid

```hcl
resource "ibm_iam_idp_account_setting" "test" {
  account_id = "01234567-89ab-cdef-0123-456789abcdef"
  active = true
  cloud_user_strategy = "DYNAMIC"
  idp_id = "example"
  ui_default = true
}
```
