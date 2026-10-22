# ibm_enterprise_account

Manages ibm_enterprise_account. The name of the account.

This rule validates `ibm_enterprise_account` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Account` (`enterprise-apis-enterprise`).

## Configuration

```hcl
rule "ibm_enterprise_account" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_id` | no | — |
| `created_at` | no | — |
| `created_by` | no | — |
| `crn` | no | — |
| `enterprise_account_id` | no | API: The enterprise account ID. |
| `enterprise_id` | no | API: The enterprise ID that the account is a part of. |
| `enterprise_path` | no | — |
| `iam_apikey` | no | — |
| `iam_apikey_id` | no | — |
| `iam_service_id` | no | — |
| `is_enterprise_account` | no | — |
| `name` | no | API: The name of the account. |
| `owner_email` | no | — |
| `owner_iam_id` | no | API: The IAM ID of the owner of the account. |
| `paid` | no | — |
| `parent` | yes | API: The CRN of the parent of the account. |
| `state` | no | — |
| `updated_at` | no | — |
| `updated_by` | no | — |
| `url` | no | — |
| `options` (block) | no | nested block |
| `options.create_iam_service_id_with_apikey_and_owner_policies` | no | — |
| `traits` (block) | no | nested block |
| `traits.enterprise_iam_managed` | no | — |
| `traits.mfa` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_enterprise_account_invalid_account_id` | `account_id` |
| `ibm_enterprise_account_invalid_enterprise_account_id` | `enterprise_account_id` |
| `ibm_enterprise_account_invalid_enterprise_id` | `enterprise_id` |
| `ibm_enterprise_account_invalid_name` | `name` |
| `ibm_enterprise_account_invalid_owner_iam_id` | `owner_iam_id` |
| `ibm_enterprise_account_invalid_parent` | `parent` |
| `ibm_enterprise_account_invalid_options__create_iam_service_id_with_apikey_and_owner_policies` | `options.create_iam_service_id_with_apikey_and_owner_policies` |
| `ibm_enterprise_account_invalid_traits__enterprise_iam_managed` | `traits.enterprise_iam_managed` |
| `ibm_enterprise_account_invalid_traits__mfa` | `traits.mfa` |

## Example

### Valid

```hcl
resource "ibm_enterprise_account" "test" {
  parent = "example"
}
```
