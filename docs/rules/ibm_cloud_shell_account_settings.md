# ibm_cloud_shell_account_settings

Schema validation for the ibm_cloud_shell_account_settings resource.

This rule validates `ibm_cloud_shell_account_settings` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/api/v1/user/accounts/{account_id}/settings:post` (`cloudshell`).

## Configuration

```hcl
rule "ibm_cloud_shell_account_settings" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_id` | yes | pattern `^[a-zA-Z0-9-]*$` |
| `created_at` | no | — |
| `created_by` | no | — |
| `default_enable_new_features` | no | API: You can choose which Cloud Shell features are available in the account and whether any new features are enabled as they become available. The feature settings apply only to the enabled Cloud Shell locations. |
| `default_enable_new_regions` | no | API: Set whether Cloud Shell is enabled in a specific location for the account. The location determines where user and session data are stored. By default, users are routed to the nearest available location. |
| `enabled` | no | API: When enabled, Cloud Shell is available to all users in the account |
| `id` | no | — |
| `rev` | no | pattern `^[a-zA-Z0-9-]*$` |
| `type` | no | — |
| `updated_at` | no | — |
| `updated_by` | no | — |
| `features` (block) | no | nested block |
| `features.enabled` | no | API: State of the feature |
| `features.key` | no | pattern `^[a-zA-Z0-9_.]*$` |
| `regions` (block) | no | nested block |
| `regions.enabled` | no | API: State of the region |
| `regions.key` | no | pattern `^[a-zA-Z-]*$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cloud_shell_account_settings_invalid_account_id` | `account_id` |
| `ibm_cloud_shell_account_settings_invalid_default_enable_new_features` | `default_enable_new_features` |
| `ibm_cloud_shell_account_settings_invalid_default_enable_new_regions` | `default_enable_new_regions` |
| `ibm_cloud_shell_account_settings_invalid_enabled` | `enabled` |
| `ibm_cloud_shell_account_settings_invalid_rev` | `rev` |
| `ibm_cloud_shell_account_settings_invalid_features__enabled` | `features.enabled` |
| `ibm_cloud_shell_account_settings_invalid_features__key` | `features.key` |
| `ibm_cloud_shell_account_settings_invalid_regions__enabled` | `regions.enabled` |
| `ibm_cloud_shell_account_settings_invalid_regions__key` | `regions.key` |

## Example

### Valid

```hcl
resource "ibm_cloud_shell_account_settings" "test" {
  account_id = "example-name"
}
```
