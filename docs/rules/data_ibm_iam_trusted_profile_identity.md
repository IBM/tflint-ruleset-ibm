# ibm_iam_trusted_profile_identity_data

Schema validation for the ibm_iam_trusted_profile_identity data source.

This rule validates `ibm_iam_trusted_profile_identity` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ProfileCount` (`iam-identity-token-api`).

## Configuration

```hcl
rule "ibm_iam_trusted_profile_identity_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `accounts` | no | — |
| `description` | no | — |
| `iam_id` | no | — |
| `identifier` | no | — |
| `identifier_id` | yes | — |
| `identity_type` | yes | one of: `crn`, `serviceid`, `user` |
| `profile_id` | yes | API: Profile identifier |
| `type` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_trusted_profile_identity_data_invalid_identifier_id` | `identifier_id` |
| `ibm_iam_trusted_profile_identity_data_invalid_identity_type` | `identity_type` |
| `ibm_iam_trusted_profile_identity_data_invalid_profile_id` | `profile_id` |

## Example

### Valid

```hcl
data "ibm_iam_trusted_profile_identity" "test" {
  identifier_id = "01234567-89ab-cdef-0123-456789abcdef"
  identity_type = "crn"
  profile_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
