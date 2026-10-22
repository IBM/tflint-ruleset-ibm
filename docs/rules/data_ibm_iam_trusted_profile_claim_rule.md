# ibm_iam_trusted_profile_claim_rule_data

Reads ibm_iam_trusted_profile_claim_rule. The optional claim rule name.

This rule validates `ibm_iam_trusted_profile_claim_rule` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ProfileCount` (`iam-identity-token-api`).

## Configuration

```hcl
rule "ibm_iam_trusted_profile_claim_rule_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `conditions` | no | — |
| `cr_type` | no | — |
| `created_at` | no | — |
| `entity_tag` | no | — |
| `expiration` | no | — |
| `modified_at` | no | — |
| `name` | no | — |
| `profile_id` | yes | API: Profile identifier |
| `realm_name` | no | — |
| `rule_id` | yes | — |
| `type` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_trusted_profile_claim_rule_data_invalid_profile_id` | `profile_id` |
| `ibm_iam_trusted_profile_claim_rule_data_invalid_rule_id` | `rule_id` |

## Example

### Valid

```hcl
data "ibm_iam_trusted_profile_claim_rule" "test" {
  profile_id = "01234567-89ab-cdef-0123-456789abcdef"
  rule_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
