# ibm_iam_trusted_profiles_data

Reads ibm_iam_trusted_profiles. Name of the trusted profile to query.

This rule validates `ibm_iam_trusted_profiles` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `CreateServiceIdGroupRequest` (`iam-identity-token-api`).

## Configuration

```hcl
rule "ibm_iam_trusted_profiles_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_id` | yes | API: ID of the account the service ID group belongs to. |
| `include_history` | no | — |
| `name` | no | API: Name of the service ID group. Unique in the account. |
| `profiles` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_trusted_profiles_data_invalid_account_id` | `account_id` |
| `ibm_iam_trusted_profiles_data_invalid_include_history` | `include_history` |
| `ibm_iam_trusted_profiles_data_invalid_name` | `name` |

## Example

### Valid

```hcl
data "ibm_iam_trusted_profiles" "test" {
  account_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
