# ibm_iam_trusted_profile_identities_data

Schema validation for the ibm_iam_trusted_profile_identities data source.

This rule validates `ibm_iam_trusted_profile_identities` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ProfileCount` (`iam-identity-token-api`).

## Configuration

```hcl
rule "ibm_iam_trusted_profile_identities_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `entity_tag` | no | — |
| `identities` | no | — |
| `profile_id` | yes | API: Profile identifier |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_trusted_profile_identities_data_invalid_profile_id` | `profile_id` |

## Example

### Valid

```hcl
data "ibm_iam_trusted_profile_identities" "test" {
  profile_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
