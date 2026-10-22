# ibm_schematics_policy

Manages ibm_schematics_policy. Name of Schematics customization policy.

This rule validates `ibm_schematics_policy` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Policy` (`schematics-schematics`).

## Configuration

```hcl
rule "ibm_schematics_policy" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account` | no | — |
| `created_at` | no | — |
| `created_by` | no | — |
| `crn` | no | — |
| `description` | no | API: The description of Schematics customization policy. |
| `kind` | no | one of: `agent_assignment_policy` |
| `location` | no | one of: `us-south`, `us-east`, `eu-gb`, `eu-de` |
| `name` | yes | API: Name of Schematics customization policy |
| `resource_group` | no | API: The resource group name for the policy. By default, Policy will be created in `default` Resource Group |
| `tags` | no | API: Tags for the Schematics customization policy |
| `updated_at` | no | — |
| `parameter` (block) | no | nested block |
| `parameter.agent_assignment_policy_parameter` (block) | no | nested block |
| `parameter.agent_assignment_policy_parameter.selector_ids` | no | API: The static selectors of schematics object ids (workspace or action) for the Schematics policy. |
| `parameter.agent_assignment_policy_parameter.selector_kind` | no | one of: `ids`, `scoped` |
| `parameter.agent_assignment_policy_parameter.selector_scope` (block) | no | nested block |
| `parameter.agent_assignment_policy_parameter.selector_scope.kind` | no | one of: `workspace`, `action`, `system`, `environment` |
| `parameter.agent_assignment_policy_parameter.selector_scope.locations` | no | one of: `us-south`, `us-east`, `eu-gb`, `eu-de` |
| `parameter.agent_assignment_policy_parameter.selector_scope.resource_groups` | no | API: The resource group based selector. |
| `parameter.agent_assignment_policy_parameter.selector_scope.tags` | no | API: The tag based selector. |
| `scoped_resources` (block) | no | nested block |
| `scoped_resources.id` | no | API: Schematics resource Id |
| `scoped_resources.kind` | no | one of: `workspace`, `action`, `system`, `environment` |
| `state` (block) | no | nested block |
| `state.state` | no | one of: `draft`, `live`, `locked`, `disable` |
| `target` (block) | no | nested block |
| `target.selector_ids` | no | API: Static selectors of schematics object ids (agent, workspace or action) for the Schematics policy. |
| `target.selector_kind` | no | one of: `ids`, `scoped` |
| `target.selector_scope` (block) | no | nested block |
| `target.selector_scope.kind` | no | one of: `workspace`, `action`, `system`, `environment` |
| `target.selector_scope.locations` | no | one of: `us-south`, `us-east`, `eu-gb`, `eu-de` |
| `target.selector_scope.resource_groups` | no | API: The resource group based selector. |
| `target.selector_scope.tags` | no | API: The tag based selector. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_schematics_policy_invalid_description` | `description` |
| `ibm_schematics_policy_invalid_kind` | `kind` |
| `ibm_schematics_policy_invalid_location` | `location` |
| `ibm_schematics_policy_invalid_name` | `name` |
| `ibm_schematics_policy_invalid_resource_group` | `resource_group` |
| `ibm_schematics_policy_invalid_tags` | `tags` |
| `ibm_schematics_policy_invalid_parameter__agent_assignment_policy_parameter__selector_ids` | `parameter.agent_assignment_policy_parameter.selector_ids` |
| `ibm_schematics_policy_invalid_parameter__agent_assignment_policy_parameter__selector_kind` | `parameter.agent_assignment_policy_parameter.selector_kind` |
| `ibm_schematics_policy_invalid_parameter__agent_assignment_policy_parameter__selector_scope__kind` | `parameter.agent_assignment_policy_parameter.selector_scope.kind` |
| `ibm_schematics_policy_invalid_parameter__agent_assignment_policy_parameter__selector_scope__locations` | `parameter.agent_assignment_policy_parameter.selector_scope.locations` |
| `ibm_schematics_policy_invalid_parameter__agent_assignment_policy_parameter__selector_scope__resource_groups` | `parameter.agent_assignment_policy_parameter.selector_scope.resource_groups` |
| `ibm_schematics_policy_invalid_parameter__agent_assignment_policy_parameter__selector_scope__tags` | `parameter.agent_assignment_policy_parameter.selector_scope.tags` |
| `ibm_schematics_policy_invalid_scoped_resources__id` | `scoped_resources.id` |
| `ibm_schematics_policy_invalid_scoped_resources__kind` | `scoped_resources.kind` |
| `ibm_schematics_policy_invalid_state__state` | `state.state` |
| `ibm_schematics_policy_invalid_target__selector_ids` | `target.selector_ids` |
| `ibm_schematics_policy_invalid_target__selector_kind` | `target.selector_kind` |
| `ibm_schematics_policy_invalid_target__selector_scope__kind` | `target.selector_scope.kind` |
| `ibm_schematics_policy_invalid_target__selector_scope__locations` | `target.selector_scope.locations` |
| `ibm_schematics_policy_invalid_target__selector_scope__resource_groups` | `target.selector_scope.resource_groups` |
| `ibm_schematics_policy_invalid_target__selector_scope__tags` | `target.selector_scope.tags` |

## Example

### Valid

```hcl
resource "ibm_schematics_policy" "test" {
  name = "aaaaaaa"
}
```
