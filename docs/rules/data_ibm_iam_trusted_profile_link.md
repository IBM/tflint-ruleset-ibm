# ibm_iam_trusted_profile_link_data

Reads ibm_iam_trusted_profile_link. Optional name of the Link.

This rule validates `ibm_iam_trusted_profile_link` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ProfileCount` (`iam-identity-token-api`).

## Configuration

```hcl
rule "ibm_iam_trusted_profile_link_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cr_type` | no | — |
| `created_at` | no | — |
| `entity_tag` | no | — |
| `is_cross_account` | no | — |
| `link` | no | — |
| `link_id` | yes | — |
| `modified_at` | no | — |
| `name` | no | — |
| `profile_id` | yes | API: Profile identifier |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_trusted_profile_link_data_invalid_link_id` | `link_id` |
| `ibm_iam_trusted_profile_link_data_invalid_profile_id` | `profile_id` |

## Example

### Valid

```hcl
data "ibm_iam_trusted_profile_link" "test" {
  link_id = "01234567-89ab-cdef-0123-456789abcdef"
  profile_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
