# ibm_iam_access_group_account_settings

Schema validation for the ibm_iam_access_group_account_settings resource.

This rule validates `ibm_iam_access_group_account_settings` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/v2/groups/settings:patch` (`iam-access-groups`).

## Configuration

```hcl
rule "ibm_iam_access_group_account_settings" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_id` | no | — |
| `public_access_enabled` | yes | API: This flag controls the public access feature within the account. It is set to true by default. Note: When this flag is set to false, all policies within the account attached to the Public Access group will be deleted. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_access_group_account_settings_invalid_public_access_enabled` | `public_access_enabled` |

## Example

### Valid

```hcl
resource "ibm_iam_access_group_account_settings" "test" {
  public_access_enabled = true
}
```
