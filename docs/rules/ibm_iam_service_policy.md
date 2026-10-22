# ibm_iam_service_policy

Schema validation for the ibm_iam_service_policy resource.

This rule validates `ibm_iam_service_policy` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Policy` (`iam-policy-management`).

## Configuration

```hcl
rule "ibm_iam_service_policy" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_management` | no | — |
| `description` | no | API: Customer-defined description |
| `iam_id` | no | — |
| `iam_service_id` | no | — |
| `pattern` | no | pattern `^[a-z:-]*$` |
| `roles` | yes | API: A role associated with a policy. |
| `rule_operator` | no | — |
| `tags` | no | API: A tag associated with a resource. |
| `transaction_id` | no | — |
| `resource_attributes` (block) | no | nested block |
| `resource_attributes.name` | yes | API: The name of an attribute. |
| `resource_attributes.operator` | no | API: The operator of an attribute. |
| `resource_attributes.value` | yes | API: The value of an attribute. |
| `resource_tags` (block) | no | nested block |
| `resource_tags.name` | yes | API: The name of an attribute. |
| `resource_tags.operator` | no | API: The operator of an attribute. |
| `resource_tags.value` | yes | API: The value of an attribute. |
| `resources` (block) | no | nested block |
| `resources.attributes` | no | API: An attribute associated with a resource. |
| `resources.region` | no | — |
| `resources.resource` | no | — |
| `resources.resource_group_id` | no | — |
| `resources.resource_instance_id` | no | — |
| `resources.resource_type` | no | — |
| `resources.service` | no | — |
| `resources.service_group_id` | no | — |
| `resources.service_type` | no | — |
| `rule_conditions` (block) | no | nested block |
| `rule_conditions.key` | no | — |
| `rule_conditions.operator` | yes | API: The operator of an attribute. |
| `rule_conditions.value` | no | API: The value of an attribute. |
| `rule_conditions.conditions` (block) | no | nested block |
| `rule_conditions.conditions.key` | yes | — |
| `rule_conditions.conditions.operator` | yes | API: The operator of an attribute. |
| `rule_conditions.conditions.value` | yes | API: The value of an attribute. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_service_policy_invalid_account_management` | `account_management` |
| `ibm_iam_service_policy_invalid_description` | `description` |
| `ibm_iam_service_policy_invalid_iam_id` | `iam_id` |
| `ibm_iam_service_policy_invalid_iam_service_id` | `iam_service_id` |
| `ibm_iam_service_policy_invalid_pattern` | `pattern` |
| `ibm_iam_service_policy_invalid_roles` | `roles` |
| `ibm_iam_service_policy_invalid_rule_operator` | `rule_operator` |
| `ibm_iam_service_policy_invalid_tags` | `tags` |
| `ibm_iam_service_policy_invalid_transaction_id` | `transaction_id` |
| `ibm_iam_service_policy_invalid_resource_attributes__name` | `resource_attributes.name` |
| `ibm_iam_service_policy_invalid_resource_attributes__operator` | `resource_attributes.operator` |
| `ibm_iam_service_policy_invalid_resource_attributes__value` | `resource_attributes.value` |
| `ibm_iam_service_policy_invalid_resource_tags__name` | `resource_tags.name` |
| `ibm_iam_service_policy_invalid_resource_tags__operator` | `resource_tags.operator` |
| `ibm_iam_service_policy_invalid_resource_tags__value` | `resource_tags.value` |
| `ibm_iam_service_policy_invalid_resources__attributes` | `resources.attributes` |
| `ibm_iam_service_policy_invalid_resources__region` | `resources.region` |
| `ibm_iam_service_policy_invalid_resources__resource` | `resources.resource` |
| `ibm_iam_service_policy_invalid_resources__resource_group_id` | `resources.resource_group_id` |
| `ibm_iam_service_policy_invalid_resources__resource_instance_id` | `resources.resource_instance_id` |
| `ibm_iam_service_policy_invalid_resources__resource_type` | `resources.resource_type` |
| `ibm_iam_service_policy_invalid_resources__service` | `resources.service` |
| `ibm_iam_service_policy_invalid_resources__service_group_id` | `resources.service_group_id` |
| `ibm_iam_service_policy_invalid_resources__service_type` | `resources.service_type` |
| `ibm_iam_service_policy_invalid_rule_conditions__key` | `rule_conditions.key` |
| `ibm_iam_service_policy_invalid_rule_conditions__operator` | `rule_conditions.operator` |
| `ibm_iam_service_policy_invalid_rule_conditions__value` | `rule_conditions.value` |
| `ibm_iam_service_policy_invalid_rule_conditions__conditions__key` | `rule_conditions.conditions.key` |
| `ibm_iam_service_policy_invalid_rule_conditions__conditions__operator` | `rule_conditions.conditions.operator` |
| `ibm_iam_service_policy_invalid_rule_conditions__conditions__value` | `rule_conditions.conditions.value` |

## Example

### Valid

```hcl
resource "ibm_iam_service_policy" "test" {
  roles = "example"
}
```
