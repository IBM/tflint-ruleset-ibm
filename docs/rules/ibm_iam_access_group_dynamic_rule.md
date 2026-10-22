# ibm_iam_access_group_dynamic_rule

Manages ibm_iam_access_group_dynamic_rule. The name of the Rule

This rule validates `ibm_iam_access_group_dynamic_rule` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Rule` (`iam-access-groups`).

## Configuration

```hcl
rule "ibm_iam_access_group_dynamic_rule" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_group_id` | yes | API: The group id that the dynamic rule is assigned to. |
| `expiration` | yes | range 1–65535; port format |
| `identity_provider` | yes | — |
| `name` | yes | API: The name of the rule. |
| `rule_id` | no | — |
| `conditions` (block) | yes | nested block |
| `conditions.claim` | yes | API: The claim to evaluate against. This will be found in the `ext` claims of a user's login request. |
| `conditions.operator` | yes | one of: `EQUALS`, `EQUALS_IGNORE_CASE`, `IN`, `NOT_EQUALS_IGNORE_CASE`, `NOT_EQUALS`, `CONTAINS` |
| `conditions.value` | yes | API: The stringified JSON value that the claim is compared to using the operator. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_access_group_dynamic_rule_invalid_access_group_id` | `access_group_id` |
| `ibm_iam_access_group_dynamic_rule_invalid_expiration` | `expiration` |
| `ibm_iam_access_group_dynamic_rule_invalid_identity_provider` | `identity_provider` |
| `ibm_iam_access_group_dynamic_rule_invalid_name` | `name` |
| `ibm_iam_access_group_dynamic_rule_invalid_conditions__claim` | `conditions.claim` |
| `ibm_iam_access_group_dynamic_rule_invalid_conditions__operator` | `conditions.operator` |
| `ibm_iam_access_group_dynamic_rule_invalid_conditions__value` | `conditions.value` |

## Example

### Valid

```hcl
resource "ibm_iam_access_group_dynamic_rule" "test" {
  access_group_id = "01234567-89ab-cdef-0123-456789abcdef"
  expiration = 80
  identity_provider = "example"
  name = "example"
  conditions {
    claim = "example"
    operator = "EQUALS"
    value = "example"
  }
}
```
