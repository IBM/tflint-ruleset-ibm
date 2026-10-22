# ibm_iam_account_settings_external_interaction

Schema validation for the ibm_iam_account_settings_external_interaction resource.

This rule validates `ibm_iam_account_settings_external_interaction` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `IdentityTypesBase` (`iam-policy-management`).

## Configuration

```hcl
rule "ibm_iam_account_settings_external_interaction" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `accept_language` | no | — |
| `account_id` | yes | — |
| `external_account_identity_interaction` (block) | yes | nested block |
| `external_account_identity_interaction.identity_types` (block) | yes | nested block |
| `external_account_identity_interaction.identity_types.service` (block) | no | nested block |
| `external_account_identity_interaction.identity_types.service.external_allowed_accounts` | yes | pattern `^[A-Za-z0-9-_]+$` |
| `external_account_identity_interaction.identity_types.service.state` | yes | one of: `enabled`, `monitor`, `limited` |
| `external_account_identity_interaction.identity_types.service_id` (block) | no | nested block |
| `external_account_identity_interaction.identity_types.service_id.external_allowed_accounts` | yes | pattern `^[A-Za-z0-9-_]+$` |
| `external_account_identity_interaction.identity_types.service_id.state` | yes | one of: `enabled`, `monitor`, `limited` |
| `external_account_identity_interaction.identity_types.user` (block) | no | nested block |
| `external_account_identity_interaction.identity_types.user.external_allowed_accounts` | yes | pattern `^[A-Za-z0-9-_]+$` |
| `external_account_identity_interaction.identity_types.user.state` | yes | one of: `enabled`, `monitor`, `limited` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_account_settings_external_interaction_invalid_accept_language` | `accept_language` |
| `ibm_iam_account_settings_external_interaction_invalid_account_id` | `account_id` |
| `ibm_iam_account_settings_external_interaction_invalid_external_account_identity_interaction__identity_types__service__external_allowed_accounts` | `external_account_identity_interaction.identity_types.service.external_allowed_accounts` |
| `ibm_iam_account_settings_external_interaction_invalid_external_account_identity_interaction__identity_types__service__state` | `external_account_identity_interaction.identity_types.service.state` |
| `ibm_iam_account_settings_external_interaction_invalid_external_account_identity_interaction__identity_types__service_id__external_allowed_accounts` | `external_account_identity_interaction.identity_types.service_id.external_allowed_accounts` |
| `ibm_iam_account_settings_external_interaction_invalid_external_account_identity_interaction__identity_types__service_id__state` | `external_account_identity_interaction.identity_types.service_id.state` |
| `ibm_iam_account_settings_external_interaction_invalid_external_account_identity_interaction__identity_types__user__external_allowed_accounts` | `external_account_identity_interaction.identity_types.user.external_allowed_accounts` |
| `ibm_iam_account_settings_external_interaction_invalid_external_account_identity_interaction__identity_types__user__state` | `external_account_identity_interaction.identity_types.user.state` |

## Example

### Valid

```hcl
resource "ibm_iam_account_settings_external_interaction" "test" {
  account_id = "01234567-89ab-cdef-0123-456789abcdef"
  external_account_identity_interaction {
    identity_types {
    }
  }
}
```
