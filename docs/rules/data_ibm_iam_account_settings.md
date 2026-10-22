# ibm_iam_account_settings_data

Schema validation for the ibm_iam_account_settings data source.

This rule validates `ibm_iam_account_settings` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_iam_account_settings_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_id` | no | — |
| `allowed_ip_addresses` | no | — |
| `entity_tag` | no | — |
| `history` | no | — |
| `include_history` | no | — |
| `max_sessions_per_identity` | no | — |
| `mfa` | no | — |
| `resolve_user_mfa` | no | — |
| `restrict_create_platform_apikey` | no | — |
| `restrict_create_service_id` | no | — |
| `restrict_user_domains` | no | — |
| `restrict_user_list_visibility` | no | — |
| `session_expiration_in_seconds` | no | — |
| `session_invalidation_in_seconds` | no | — |
| `system_access_token_expiration_in_seconds` | no | — |
| `system_refresh_token_expiration_in_seconds` | no | — |
| `user_mfa` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_account_settings_data_invalid_include_history` | `include_history` |
| `ibm_iam_account_settings_data_invalid_resolve_user_mfa` | `resolve_user_mfa` |

## Example

### Valid

```hcl
data "ibm_iam_account_settings" "test" {
}
```
