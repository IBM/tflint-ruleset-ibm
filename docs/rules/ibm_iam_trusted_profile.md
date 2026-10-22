# ibm_iam_trusted_profile

Manages ibm_iam_trusted_profile. Name of the trusted profile.

This rule validates `ibm_iam_trusted_profile` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `UpdateTrustedProfileRequest` (`iam-identity-token-api`).

## Configuration

```hcl
rule "ibm_iam_trusted_profile" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_id` | no | — |
| `activity` | no | — |
| `assignment_id` | no | — |
| `created_at` | no | — |
| `crn` | no | — |
| `description` | no | API: The description of the trusted profile to update. If specified an empty description will clear the description of the trusted profile. If a non empty value is provided the trusted profile will be updated. |
| `email` | no | API: The email of the profile to update. If specified an empty email will clear the email of the profile. If an non empty value is provided the trusted profile will be updated. |
| `entity_tag` | no | — |
| `history` | no | — |
| `iam_id` | no | — |
| `ims_account_id` | no | — |
| `ims_user_id` | no | — |
| `modified_at` | no | — |
| `name` | yes | API: The name of the trusted profile to update. If specified in the request the parameter must not be empty. The name is checked for uniqueness. Failure to this will result in an Error condition. |
| `profile_id` | no | — |
| `template_id` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_trusted_profile_invalid_description` | `description` |
| `ibm_iam_trusted_profile_invalid_email` | `email` |
| `ibm_iam_trusted_profile_invalid_name` | `name` |

## Example

### Valid

```hcl
resource "ibm_iam_trusted_profile" "test" {
  name = "example"
}
```
