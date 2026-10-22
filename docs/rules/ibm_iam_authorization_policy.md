# ibm_iam_authorization_policy

Schema validation for the ibm_iam_authorization_policy resource.

This rule validates `ibm_iam_authorization_policy` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Policy` (`iam-policy-management`).

## Configuration

```hcl
rule "ibm_iam_authorization_policy" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `description` | no | API: Customer-defined description |
| `roles` | yes | API: A role associated with a policy. |
| `source_resource_group_id` | no | — |
| `source_resource_instance_id` | no | — |
| `source_resource_type` | no | — |
| `source_service_account` | no | — |
| `source_service_name` | no | — |
| `target_resource_group_id` | no | — |
| `target_resource_instance_id` | no | — |
| `target_resource_type` | no | — |
| `target_service_name` | no | — |
| `transaction_id` | no | — |
| `version` | no | — |
| `resource_attributes` (block) | no | nested block |
| `resource_attributes.name` | yes | API: The name of an attribute. |
| `resource_attributes.operator` | no | API: The operator of an attribute. |
| `resource_attributes.value` | yes | API: The value of an attribute. |
| `subject_attributes` (block) | no | nested block |
| `subject_attributes.name` | yes | API: The name of an attribute. |
| `subject_attributes.operator` | no | API: The operator of an attribute. |
| `subject_attributes.value` | yes | API: The value of an attribute. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_authorization_policy_invalid_description` | `description` |
| `ibm_iam_authorization_policy_invalid_roles` | `roles` |
| `ibm_iam_authorization_policy_invalid_source_resource_group_id` | `source_resource_group_id` |
| `ibm_iam_authorization_policy_invalid_source_resource_instance_id` | `source_resource_instance_id` |
| `ibm_iam_authorization_policy_invalid_source_resource_type` | `source_resource_type` |
| `ibm_iam_authorization_policy_invalid_source_service_account` | `source_service_account` |
| `ibm_iam_authorization_policy_invalid_source_service_name` | `source_service_name` |
| `ibm_iam_authorization_policy_invalid_target_resource_group_id` | `target_resource_group_id` |
| `ibm_iam_authorization_policy_invalid_target_resource_instance_id` | `target_resource_instance_id` |
| `ibm_iam_authorization_policy_invalid_target_resource_type` | `target_resource_type` |
| `ibm_iam_authorization_policy_invalid_target_service_name` | `target_service_name` |
| `ibm_iam_authorization_policy_invalid_transaction_id` | `transaction_id` |
| `ibm_iam_authorization_policy_invalid_resource_attributes__name` | `resource_attributes.name` |
| `ibm_iam_authorization_policy_invalid_resource_attributes__operator` | `resource_attributes.operator` |
| `ibm_iam_authorization_policy_invalid_resource_attributes__value` | `resource_attributes.value` |
| `ibm_iam_authorization_policy_invalid_subject_attributes__name` | `subject_attributes.name` |
| `ibm_iam_authorization_policy_invalid_subject_attributes__operator` | `subject_attributes.operator` |
| `ibm_iam_authorization_policy_invalid_subject_attributes__value` | `subject_attributes.value` |

## Example

### Valid

```hcl
resource "ibm_iam_authorization_policy" "test" {
  roles = "example"
}
```
