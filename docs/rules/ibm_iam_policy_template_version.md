# ibm_iam_policy_template_version

Schema validation for the ibm_iam_policy_template_version resource.

This rule validates `ibm_iam_policy_template_version` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Policy` (`iam-policy-management`).

## Configuration

```hcl
rule "ibm_iam_policy_template_version" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_id` | no | — |
| `committed` | no | — |
| `description` | no | API: Customer-defined description |
| `name` | no | API: The name of an attribute. |
| `template_id` | yes | pattern `^policyTemplate-[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$` |
| `version` | no | — |
| `policy` (block) | yes | nested block |
| `policy.description` | no | API: Customer-defined description |
| `policy.pattern` | no | pattern `^[a-z:-]*$` |
| `policy.roles` | no | API: A role associated with a policy. |
| `policy.rule_operator` | no | — |
| `policy.type` | yes | pattern `^[a-z]+$` |
| `policy.resource` (block) | no | nested block |
| `policy.resource.attributes` (block) | yes | nested block |
| `policy.resource.attributes.key` | yes | — |
| `policy.resource.attributes.operator` | yes | API: The operator of an attribute. |
| `policy.resource.attributes.value` | yes | API: The value of an attribute. |
| `policy.resource.tags` (block) | no | nested block |
| `policy.resource.tags.key` | yes | — |
| `policy.resource.tags.operator` | yes | API: The operator of an attribute. |
| `policy.resource.tags.value` | yes | API: The value of an attribute. |
| `policy.role_template_references` (block) | no | nested block |
| `policy.role_template_references.id` | yes | API: The policy ID. |
| `policy.role_template_references.version` | yes | — |
| `policy.rule_conditions` (block) | no | nested block |
| `policy.rule_conditions.key` | no | — |
| `policy.rule_conditions.operator` | yes | API: The operator of an attribute. |
| `policy.rule_conditions.value` | no | API: The value of an attribute. |
| `policy.rule_conditions.conditions` (block) | no | nested block |
| `policy.rule_conditions.conditions.key` | yes | — |
| `policy.rule_conditions.conditions.operator` | yes | API: The operator of an attribute. |
| `policy.rule_conditions.conditions.value` | yes | API: The value of an attribute. |
| `policy.subject` (block) | no | nested block |
| `policy.subject.attributes` (block) | yes | nested block |
| `policy.subject.attributes.key` | yes | — |
| `policy.subject.attributes.operator` | yes | API: The operator of an attribute. |
| `policy.subject.attributes.value` | yes | API: The value of an attribute. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_policy_template_version_invalid_committed` | `committed` |
| `ibm_iam_policy_template_version_invalid_description` | `description` |
| `ibm_iam_policy_template_version_invalid_name` | `name` |
| `ibm_iam_policy_template_version_invalid_template_id` | `template_id` |
| `ibm_iam_policy_template_version_invalid_policy__description` | `policy.description` |
| `ibm_iam_policy_template_version_invalid_policy__pattern` | `policy.pattern` |
| `ibm_iam_policy_template_version_invalid_policy__roles` | `policy.roles` |
| `ibm_iam_policy_template_version_invalid_policy__rule_operator` | `policy.rule_operator` |
| `ibm_iam_policy_template_version_invalid_policy__type` | `policy.type` |
| `ibm_iam_policy_template_version_invalid_policy__resource__attributes__key` | `policy.resource.attributes.key` |
| `ibm_iam_policy_template_version_invalid_policy__resource__attributes__operator` | `policy.resource.attributes.operator` |
| `ibm_iam_policy_template_version_invalid_policy__resource__attributes__value` | `policy.resource.attributes.value` |
| `ibm_iam_policy_template_version_invalid_policy__resource__tags__key` | `policy.resource.tags.key` |
| `ibm_iam_policy_template_version_invalid_policy__resource__tags__operator` | `policy.resource.tags.operator` |
| `ibm_iam_policy_template_version_invalid_policy__resource__tags__value` | `policy.resource.tags.value` |
| `ibm_iam_policy_template_version_invalid_policy__role_template_references__id` | `policy.role_template_references.id` |
| `ibm_iam_policy_template_version_invalid_policy__role_template_references__version` | `policy.role_template_references.version` |
| `ibm_iam_policy_template_version_invalid_policy__rule_conditions__key` | `policy.rule_conditions.key` |
| `ibm_iam_policy_template_version_invalid_policy__rule_conditions__operator` | `policy.rule_conditions.operator` |
| `ibm_iam_policy_template_version_invalid_policy__rule_conditions__value` | `policy.rule_conditions.value` |
| `ibm_iam_policy_template_version_invalid_policy__rule_conditions__conditions__key` | `policy.rule_conditions.conditions.key` |
| `ibm_iam_policy_template_version_invalid_policy__rule_conditions__conditions__operator` | `policy.rule_conditions.conditions.operator` |
| `ibm_iam_policy_template_version_invalid_policy__rule_conditions__conditions__value` | `policy.rule_conditions.conditions.value` |
| `ibm_iam_policy_template_version_invalid_policy__subject__attributes__key` | `policy.subject.attributes.key` |
| `ibm_iam_policy_template_version_invalid_policy__subject__attributes__operator` | `policy.subject.attributes.operator` |
| `ibm_iam_policy_template_version_invalid_policy__subject__attributes__value` | `policy.subject.attributes.value` |

## Example

### Valid

```hcl
resource "ibm_iam_policy_template_version" "test" {
  template_id = "policyTemplate-01234567-89ab-cdef-0123-456789abcdef"
  policy {
    type = "example"
  }
}
```
