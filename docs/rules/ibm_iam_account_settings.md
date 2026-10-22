# ibm_iam_account_settings

Schema validation for the ibm_iam_account_settings resource.

This rule validates `ibm_iam_account_settings` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/v1/accounts/{account_id}/settings/identity:put` (`iam-identity-token-api`).

## Configuration

```hcl
rule "ibm_iam_account_settings" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `allowed_ip_addresses` | no | API: Defines the IP addresses and subnets from which IAM tokens can be created for the account. |
| `entity_tag` | no | — |
| `history` | no | — |
| `if_match` | no | — |
| `include_history` | no | — |
| `max_sessions_per_identity` | no | API: Defines the max allowed sessions per identity required by the account. Valid values: * Any whole number greater than 0 * NOT_SET - To unset account setting and use service default |
| `mfa` | no | one of: `NONE`, `NONE_NO_ROPC`, `TOTP`, `TOTP4ALL`, `LEVEL1`, `LEVEL2`, `LEVEL3` |
| `resolve_user_mfa` | no | — |
| `restrict_create_platform_apikey` | no | one of: `RESTRICTED`, `NOT_RESTRICTED`, `NOT_SET` |
| `restrict_create_service_id` | no | one of: `RESTRICTED`, `NOT_RESTRICTED`, `NOT_SET` |
| `restrict_user_list_visibility` | no | one of: `RESTRICTED`, `NOT_RESTRICTED` |
| `session_expiration_in_seconds` | no | API: Defines the session expiration in seconds for the account. Valid values: * Any whole number between between '900' and '86400' * NOT_SET - To unset account setting and use service default |
| `session_invalidation_in_seconds` | no | API: Defines the period of time in seconds in which a session will be invalidated due to inactivity. Valid values: * Any whole number between '900' and '7200' * NOT_SET - To unset account setting and use service default |
| `system_access_token_expiration_in_seconds` | no | API: Defines the access token expiration in seconds. Valid values: * Any whole number between '900' and '3600' * NOT_SET - To unset account setting and use service default |
| `system_refresh_token_expiration_in_seconds` | no | API: Defines the refresh token expiration in seconds. Valid values: * Any whole number between '900' and '259200' * NOT_SET - To unset account setting and use service default |
| `restrict_user_domains` (block) | no | nested block |
| `restrict_user_domains.invitation_email_allow_patterns` | no | API: The list of allowed email patterns. Wildcard syntax is supported, '*' represents any sequence of zero or more characters in the string, except for '.' and '@'. The sequence ends if a '.' or '@' was found. '**' represents any sequence of zero or more characters in the string - without limit |
| `restrict_user_domains.realm_id` | no | API: The realm that the restrictions apply to |
| `restrict_user_domains.restrict_invitation` | no | API: When true invites will only be possible to the domain patterns provided, otherwise invites are unrestricted |
| `user_mfa` (block) | no | nested block |
| `user_mfa.description` | no | — |
| `user_mfa.email` | no | — |
| `user_mfa.iam_id` | no | API: The iam_id of the user. |
| `user_mfa.mfa` | no | one of: `NONE`, `NONE_NO_ROPC`, `TOTP`, `TOTP4ALL`, `LEVEL1`, `LEVEL2`, `LEVEL3` |
| `user_mfa.name` | no | — |
| `user_mfa.user_name` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_account_settings_invalid_allowed_ip_addresses` | `allowed_ip_addresses` |
| `ibm_iam_account_settings_invalid_entity_tag` | `entity_tag` |
| `ibm_iam_account_settings_invalid_if_match` | `if_match` |
| `ibm_iam_account_settings_invalid_include_history` | `include_history` |
| `ibm_iam_account_settings_invalid_max_sessions_per_identity` | `max_sessions_per_identity` |
| `ibm_iam_account_settings_invalid_mfa` | `mfa` |
| `ibm_iam_account_settings_invalid_resolve_user_mfa` | `resolve_user_mfa` |
| `ibm_iam_account_settings_invalid_restrict_create_platform_apikey` | `restrict_create_platform_apikey` |
| `ibm_iam_account_settings_invalid_restrict_create_service_id` | `restrict_create_service_id` |
| `ibm_iam_account_settings_invalid_restrict_user_list_visibility` | `restrict_user_list_visibility` |
| `ibm_iam_account_settings_invalid_session_expiration_in_seconds` | `session_expiration_in_seconds` |
| `ibm_iam_account_settings_invalid_session_invalidation_in_seconds` | `session_invalidation_in_seconds` |
| `ibm_iam_account_settings_invalid_system_access_token_expiration_in_seconds` | `system_access_token_expiration_in_seconds` |
| `ibm_iam_account_settings_invalid_system_refresh_token_expiration_in_seconds` | `system_refresh_token_expiration_in_seconds` |
| `ibm_iam_account_settings_invalid_restrict_user_domains__invitation_email_allow_patterns` | `restrict_user_domains.invitation_email_allow_patterns` |
| `ibm_iam_account_settings_invalid_restrict_user_domains__realm_id` | `restrict_user_domains.realm_id` |
| `ibm_iam_account_settings_invalid_restrict_user_domains__restrict_invitation` | `restrict_user_domains.restrict_invitation` |
| `ibm_iam_account_settings_invalid_user_mfa__description` | `user_mfa.description` |
| `ibm_iam_account_settings_invalid_user_mfa__email` | `user_mfa.email` |
| `ibm_iam_account_settings_invalid_user_mfa__iam_id` | `user_mfa.iam_id` |
| `ibm_iam_account_settings_invalid_user_mfa__mfa` | `user_mfa.mfa` |
| `ibm_iam_account_settings_invalid_user_mfa__name` | `user_mfa.name` |
| `ibm_iam_account_settings_invalid_user_mfa__user_name` | `user_mfa.user_name` |

## Example

### Valid

```hcl
resource "ibm_iam_account_settings" "test" {
}
```
