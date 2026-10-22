# ibm_iam_account_settings_template

Manages ibm_iam_account_settings_template. The name of the trusted profile template.

This rule validates `ibm_iam_account_settings_template` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/v1/accounts/{account_id}/settings/identity:put` (`iam-identity-token-api`).

## Configuration

```hcl
rule "ibm_iam_account_settings_template" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_id` | no | — |
| `committed` | no | — |
| `created_at` | no | — |
| `created_by_id` | no | — |
| `crn` | no | — |
| `description` | no | — |
| `entity_tag` | no | — |
| `id` | no | — |
| `last_modified_at` | no | — |
| `last_modified_by_id` | no | — |
| `name` | no | — |
| `template_id` | no | — |
| `version` | no | — |
| `account_settings` (block) | no | nested block |
| `account_settings.allowed_ip_addresses` | no | API: Defines the IP addresses and subnets from which IAM tokens can be created for the account. |
| `account_settings.max_sessions_per_identity` | no | API: Defines the max allowed sessions per identity required by the account. Valid values: * Any whole number greater than 0 * NOT_SET - To unset account setting and use service default |
| `account_settings.mfa` | no | one of: `NONE`, `NONE_NO_ROPC`, `TOTP`, `TOTP4ALL`, `LEVEL1`, `LEVEL2`, `LEVEL3` |
| `account_settings.restrict_create_platform_apikey` | no | one of: `RESTRICTED`, `NOT_RESTRICTED`, `NOT_SET` |
| `account_settings.restrict_create_service_id` | no | one of: `RESTRICTED`, `NOT_RESTRICTED`, `NOT_SET` |
| `account_settings.restrict_user_list_visibility` | no | one of: `NOT_RESTRICTED`, `RESTRICTED` |
| `account_settings.session_expiration_in_seconds` | no | API: Defines the session expiration in seconds for the account. Valid values: * Any whole number between between '900' and '86400' * NOT_SET - To unset account setting and use service default |
| `account_settings.session_invalidation_in_seconds` | no | API: Defines the period of time in seconds in which a session will be invalidated due to inactivity. Valid values: * Any whole number between '900' and '7200' * NOT_SET - To unset account setting and use service default |
| `account_settings.system_access_token_expiration_in_seconds` | no | API: Defines the access token expiration in seconds. Valid values: * Any whole number between '900' and '3600' * NOT_SET - To unset account setting and use service default |
| `account_settings.system_refresh_token_expiration_in_seconds` | no | API: Defines the refresh token expiration in seconds. Valid values: * Any whole number between '900' and '259200' * NOT_SET - To unset account setting and use service default |
| `account_settings.restrict_user_domains` (block) | no | nested block |
| `account_settings.restrict_user_domains.account_sufficient` | no | — |
| `account_settings.restrict_user_domains.restrictions` (block) | no | nested block |
| `account_settings.restrict_user_domains.restrictions.invitation_email_allow_patterns` | no | API: The list of allowed email patterns. Wildcard syntax is supported, '*' represents any sequence of zero or more characters in the string, except for '.' and '@'. The sequence ends if a '.' or '@' was found. '**' represents any sequence of zero or more characters in the string - without limit |
| `account_settings.restrict_user_domains.restrictions.realm_id` | yes | API: The realm that the restrictions apply to |
| `account_settings.restrict_user_domains.restrictions.restrict_invitation` | no | API: When true invites will only be possible to the domain patterns provided, otherwise invites are unrestricted |
| `account_settings.user_mfa` (block) | no | nested block |
| `account_settings.user_mfa.iam_id` | yes | API: The iam_id of the user. |
| `account_settings.user_mfa.mfa` | yes | one of: `NONE`, `NONE_NO_ROPC`, `TOTP`, `TOTP4ALL`, `LEVEL1`, `LEVEL2`, `LEVEL3` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_account_settings_template_invalid_account_id` | `account_id` |
| `ibm_iam_account_settings_template_invalid_committed` | `committed` |
| `ibm_iam_account_settings_template_invalid_description` | `description` |
| `ibm_iam_account_settings_template_invalid_name` | `name` |
| `ibm_iam_account_settings_template_invalid_template_id` | `template_id` |
| `ibm_iam_account_settings_template_invalid_account_settings__allowed_ip_addresses` | `account_settings.allowed_ip_addresses` |
| `ibm_iam_account_settings_template_invalid_account_settings__max_sessions_per_identity` | `account_settings.max_sessions_per_identity` |
| `ibm_iam_account_settings_template_invalid_account_settings__mfa` | `account_settings.mfa` |
| `ibm_iam_account_settings_template_invalid_account_settings__restrict_create_platform_apikey` | `account_settings.restrict_create_platform_apikey` |
| `ibm_iam_account_settings_template_invalid_account_settings__restrict_create_service_id` | `account_settings.restrict_create_service_id` |
| `ibm_iam_account_settings_template_invalid_account_settings__restrict_user_list_visibility` | `account_settings.restrict_user_list_visibility` |
| `ibm_iam_account_settings_template_invalid_account_settings__session_expiration_in_seconds` | `account_settings.session_expiration_in_seconds` |
| `ibm_iam_account_settings_template_invalid_account_settings__session_invalidation_in_seconds` | `account_settings.session_invalidation_in_seconds` |
| `ibm_iam_account_settings_template_invalid_account_settings__system_access_token_expiration_in_seconds` | `account_settings.system_access_token_expiration_in_seconds` |
| `ibm_iam_account_settings_template_invalid_account_settings__system_refresh_token_expiration_in_seconds` | `account_settings.system_refresh_token_expiration_in_seconds` |
| `ibm_iam_account_settings_template_invalid_account_settings__restrict_user_domains__account_sufficient` | `account_settings.restrict_user_domains.account_sufficient` |
| `ibm_iam_account_settings_template_invalid_account_settings__restrict_user_domains__restrictions__invitation_email_allow_patterns` | `account_settings.restrict_user_domains.restrictions.invitation_email_allow_patterns` |
| `ibm_iam_account_settings_template_invalid_account_settings__restrict_user_domains__restrictions__realm_id` | `account_settings.restrict_user_domains.restrictions.realm_id` |
| `ibm_iam_account_settings_template_invalid_account_settings__restrict_user_domains__restrictions__restrict_invitation` | `account_settings.restrict_user_domains.restrictions.restrict_invitation` |
| `ibm_iam_account_settings_template_invalid_account_settings__user_mfa__iam_id` | `account_settings.user_mfa.iam_id` |
| `ibm_iam_account_settings_template_invalid_account_settings__user_mfa__mfa` | `account_settings.user_mfa.mfa` |

## Example

### Valid

```hcl
resource "ibm_iam_account_settings_template" "test" {
}
```
