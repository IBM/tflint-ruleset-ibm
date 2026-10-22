# ibm_iam_trusted_profile_data

Reads ibm_iam_trusted_profile. Name of the trusted profile.

This rule validates `ibm_iam_trusted_profile` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ProfileCount` (`iam-identity-token-api`).

## Configuration

```hcl
rule "ibm_iam_trusted_profile_data" {
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
| `description` | no | — |
| `email` | no | — |
| `entity_tag` | no | — |
| `history` | no | — |
| `iam_id` | no | — |
| `ims_account_id` | no | — |
| `ims_user_id` | no | — |
| `include_activity` | no | — |
| `modified_at` | no | — |
| `name` | no | — |
| `profile_id` | yes | API: Profile identifier |
| `template_id` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_trusted_profile_data_invalid_include_activity` | `include_activity` |
| `ibm_iam_trusted_profile_data_invalid_profile_id` | `profile_id` |

## Example

### Valid

```hcl
data "ibm_iam_trusted_profile" "test" {
  profile_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
