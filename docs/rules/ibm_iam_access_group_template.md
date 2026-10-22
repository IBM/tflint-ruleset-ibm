# ibm_iam_access_group_template

Manages ibm_iam_access_group_template. The name of the access group template.

This rule validates `ibm_iam_access_group_template` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Group` (`iam-access-groups`).

## Configuration

```hcl
rule "ibm_iam_access_group_template" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_id` | no | — |
| `committed` | no | — |
| `created_at` | no | — |
| `created_by_id` | no | — |
| `description` | no | pattern `^[a-zA-Z0-9!@#$%^&*()_+{}:;"'<>,.?\/|\-\s]+$` |
| `href` | no | — |
| `last_modified_at` | no | — |
| `last_modified_by_id` | no | — |
| `name` | yes | pattern `^[a-zA-Z0-9!@#$%^&*()_+{}:;"'<>,.?\/|\-\s]+$` |
| `template_id` | no | — |
| `transaction_id` | no | pattern `^[a-zA-Z0-9_-]+$` |
| `version` | no | — |
| `group` (block) | no | nested block |
| `group.description` | no | API: The group's description - if defined. |
| `group.name` | yes | API: The group's name. |
| `group.action_controls` (block) | no | nested block |
| `group.action_controls.access` (block) | no | nested block |
| `group.action_controls.access.add` | no | — |
| `group.assertions` (block) | no | nested block |
| `group.assertions.action_controls` (block) | no | nested block |
| `group.assertions.action_controls.add` | no | — |
| `group.assertions.action_controls.remove` | no | — |
| `group.assertions.rules` (block) | no | nested block |
| `group.assertions.rules.expiration` | no | min 1 |
| `group.assertions.rules.name` | no | API: The group's name. |
| `group.assertions.rules.realm_name` | no | pattern `^[a-zA-Z0-9!@#$%^&*()_+{}:;"'<>,.?/|\-\s]+$` |
| `group.assertions.rules.action_controls` (block) | no | nested block |
| `group.assertions.rules.action_controls.remove` | no | — |
| `group.assertions.rules.conditions` (block) | no | nested block |
| `group.assertions.rules.conditions.claim` | no | pattern `^[a-zA-Z0-9_-]+$` |
| `group.assertions.rules.conditions.operator` | no | — |
| `group.assertions.rules.conditions.value` | no | pattern `^[a-zA-Z0-9_-]+$` |
| `group.members` (block) | no | nested block |
| `group.members.services` | no | pattern `^[a-zA-Z0-9_-]+$` |
| `group.members.users` | no | pattern `^[a-zA-Z0-9!@#$%^&*()_+{}:;"'<>,.?/|\-\s]+$` |
| `group.members.action_controls` (block) | no | nested block |
| `group.members.action_controls.add` | no | — |
| `group.members.action_controls.remove` | no | — |
| `policy_template_references` (block) | no | nested block |
| `policy_template_references.id` | no | API: The group's access group ID. |
| `policy_template_references.version` | no | pattern `^[0-9]+$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_access_group_template_invalid_committed` | `committed` |
| `ibm_iam_access_group_template_invalid_description` | `description` |
| `ibm_iam_access_group_template_invalid_name` | `name` |
| `ibm_iam_access_group_template_invalid_transaction_id` | `transaction_id` |
| `ibm_iam_access_group_template_invalid_group__description` | `group.description` |
| `ibm_iam_access_group_template_invalid_group__name` | `group.name` |
| `ibm_iam_access_group_template_invalid_group__action_controls__access__add` | `group.action_controls.access.add` |
| `ibm_iam_access_group_template_invalid_group__assertions__action_controls__add` | `group.assertions.action_controls.add` |
| `ibm_iam_access_group_template_invalid_group__assertions__action_controls__remove` | `group.assertions.action_controls.remove` |
| `ibm_iam_access_group_template_invalid_group__assertions__rules__expiration` | `group.assertions.rules.expiration` |
| `ibm_iam_access_group_template_invalid_group__assertions__rules__name` | `group.assertions.rules.name` |
| `ibm_iam_access_group_template_invalid_group__assertions__rules__realm_name` | `group.assertions.rules.realm_name` |
| `ibm_iam_access_group_template_invalid_group__assertions__rules__action_controls__remove` | `group.assertions.rules.action_controls.remove` |
| `ibm_iam_access_group_template_invalid_group__assertions__rules__conditions__claim` | `group.assertions.rules.conditions.claim` |
| `ibm_iam_access_group_template_invalid_group__assertions__rules__conditions__operator` | `group.assertions.rules.conditions.operator` |
| `ibm_iam_access_group_template_invalid_group__assertions__rules__conditions__value` | `group.assertions.rules.conditions.value` |
| `ibm_iam_access_group_template_invalid_group__members__services` | `group.members.services` |
| `ibm_iam_access_group_template_invalid_group__members__users` | `group.members.users` |
| `ibm_iam_access_group_template_invalid_group__members__action_controls__add` | `group.members.action_controls.add` |
| `ibm_iam_access_group_template_invalid_group__members__action_controls__remove` | `group.members.action_controls.remove` |
| `ibm_iam_access_group_template_invalid_policy_template_references__id` | `policy_template_references.id` |
| `ibm_iam_access_group_template_invalid_policy_template_references__version` | `policy_template_references.version` |

## Example

### Valid

```hcl
resource "ibm_iam_access_group_template" "test" {
  name = "example-name"
}
```
