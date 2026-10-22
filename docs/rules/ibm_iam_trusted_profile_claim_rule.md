# ibm_iam_trusted_profile_claim_rule

Manages ibm_iam_trusted_profile_claim_rule. Name of the claim rule to be created or updated.

This rule validates `ibm_iam_trusted_profile_claim_rule` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ProfileClaimRuleRequest` (`iam-identity-token-api`).

## Configuration

```hcl
rule "ibm_iam_trusted_profile_claim_rule" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cr_type` | no | API: The compute resource type the rule applies to, required only if type is specified as 'Profile-CR'. Valid values are VSI, PVS, BMS, IKS_SA, ROKS_SA, CE. |
| `created_at` | no | — |
| `entity_tag` | no | — |
| `expiration` | no | API: Session expiration in seconds, only required if type is 'Profile-SAML'. |
| `modified_at` | no | — |
| `name` | no | API: Name of the claim rule to be created or updated |
| `profile_id` | yes | — |
| `realm_name` | no | API: The realm name of the Idp this claim rule applies to. This field is required only if the type is specified as 'Profile-SAML'. |
| `rule_id` | no | — |
| `type` | yes | one of: `Profile-SAML`, `Profile-CR` |
| `conditions` (block) | yes | nested block |
| `conditions.claim` | yes | API: The claim to evaluate against. [Learn more](/docs/account?topic=account-iam-condition-properties&interface=ui#cr-attribute-names). |
| `conditions.operator` | yes | one of: `EQUALS`, `NOT_EQUALS`, `EQUALS_IGNORE_CASE`, `NOT_EQUALS_IGNORE_CASE`, `CONTAINS`, `IN` |
| `conditions.value` | yes | API: The stringified JSON value that the claim is compared to using the operator |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_trusted_profile_claim_rule_invalid_cr_type` | `cr_type` |
| `ibm_iam_trusted_profile_claim_rule_invalid_expiration` | `expiration` |
| `ibm_iam_trusted_profile_claim_rule_invalid_name` | `name` |
| `ibm_iam_trusted_profile_claim_rule_invalid_profile_id` | `profile_id` |
| `ibm_iam_trusted_profile_claim_rule_invalid_realm_name` | `realm_name` |
| `ibm_iam_trusted_profile_claim_rule_invalid_type` | `type` |
| `ibm_iam_trusted_profile_claim_rule_invalid_conditions__claim` | `conditions.claim` |
| `ibm_iam_trusted_profile_claim_rule_invalid_conditions__operator` | `conditions.operator` |
| `ibm_iam_trusted_profile_claim_rule_invalid_conditions__value` | `conditions.value` |

## Example

### Valid

```hcl
resource "ibm_iam_trusted_profile_claim_rule" "test" {
  profile_id = "01234567-89ab-cdef-0123-456789abcdef"
  type = "Profile-SAML"
  conditions {
    claim = "example"
    operator = "EQUALS"
    value = "example"
  }
}
```
