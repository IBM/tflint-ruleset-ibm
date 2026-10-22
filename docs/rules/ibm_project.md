# ibm_project

Schema validation for the ibm_project resource.

This rule validates `ibm_project` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Project update` (`projects`).

## Configuration

```hcl
rule "ibm_project" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `created_at` | no | — |
| `crn` | no | — |
| `cumulative_needs_attention_view` | no | — |
| `cumulative_needs_attention_view_error` | no | — |
| `event_notifications_crn` | no | — |
| `href` | no | — |
| `location` | yes | one of: `ca-tor`, `eu-de`, `eu-gb`, `us-east`, `us-south` |
| `resource_group` | yes | API: The resource group name where the project's data and tools are created. |
| `resource_group_id` | no | — |
| `state` | no | — |
| `configs` (block) | no | nested block |
| `configs.container_state` | no | one of: `approved`, `deleted`, `deleting`, `deleting_failed`, `discarded`, `draft`, `deployed`, `deploying_failed`, `deploying`, `superseded`, `undeploying`, `undeploying_failed`, `validated`, `validating`, `validating_failed`, `working` |
| `configs.container_state_code` | no | one of: `awaiting_input`, `awaiting_prerequisite`, `awaiting_validation`, `awaiting_member_deployment`, `awaiting_stack_setup` |
| `configs.state_code` | no | — |
| `configs.approved_version` (block) | no | nested block |
| `configs.approved_version.container_state` | no | one of: `approved`, `deleted`, `deleting`, `deleting_failed`, `discarded`, `draft`, `deployed`, `deploying_failed`, `deploying`, `superseded`, `undeploying`, `undeploying_failed`, `validated`, `validating`, `validating_failed`, `working` |
| `configs.deployed_version` (block) | no | nested block |
| `configs.deployed_version.container_state` | no | one of: `approved`, `deleted`, `deleting`, `deleting_failed`, `discarded`, `draft`, `deployed`, `deploying_failed`, `deploying`, `superseded`, `undeploying`, `undeploying_failed`, `validated`, `validating`, `validating_failed`, `working` |
| `definition` (block) | yes | nested block |
| `definition.auto_deploy` | no | API: A boolean flag to enable deploying configurations automatically. |
| `definition.auto_deploy_mode` | no | one of: `auto_approval`, `manual_approval` |
| `definition.description` | no | — |
| `definition.destroy_on_delete` | no | — |
| `definition.monitoring_enabled` | no | — |
| `definition.name` | no | — |
| `definition.store` (block) | no | nested block |
| `definition.store.config_directory` | no | API: The directory where project configs are stored. |
| `definition.store.token` | no | — |
| `definition.store.type` | yes | — |
| `definition.store.url` | yes | pattern `^((http(s)?:\/\/)|\/)[a-zA-Z0-9\$\-_\.+!\*'\(\),=&?\/:]+$` |
| `definition.terraform_engine` (block) | no | nested block |
| `definition.terraform_engine.id` | yes | — |
| `definition.terraform_engine.type` | yes | — |
| `environments` (block) | no | nested block |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_project_invalid_location` | `location` |
| `ibm_project_invalid_resource_group` | `resource_group` |
| `ibm_project_invalid_configs__container_state` | `configs.container_state` |
| `ibm_project_invalid_configs__container_state_code` | `configs.container_state_code` |
| `ibm_project_invalid_configs__state_code` | `configs.state_code` |
| `ibm_project_invalid_configs__approved_version__container_state` | `configs.approved_version.container_state` |
| `ibm_project_invalid_configs__deployed_version__container_state` | `configs.deployed_version.container_state` |
| `ibm_project_invalid_definition__auto_deploy` | `definition.auto_deploy` |
| `ibm_project_invalid_definition__auto_deploy_mode` | `definition.auto_deploy_mode` |
| `ibm_project_invalid_definition__description` | `definition.description` |
| `ibm_project_invalid_definition__destroy_on_delete` | `definition.destroy_on_delete` |
| `ibm_project_invalid_definition__monitoring_enabled` | `definition.monitoring_enabled` |
| `ibm_project_invalid_definition__name` | `definition.name` |
| `ibm_project_invalid_definition__store__config_directory` | `definition.store.config_directory` |
| `ibm_project_invalid_definition__store__token` | `definition.store.token` |
| `ibm_project_invalid_definition__store__type` | `definition.store.type` |
| `ibm_project_invalid_definition__store__url` | `definition.store.url` |
| `ibm_project_invalid_definition__terraform_engine__id` | `definition.terraform_engine.id` |
| `ibm_project_invalid_definition__terraform_engine__type` | `definition.terraform_engine.type` |

## Example

### Valid

```hcl
resource "ibm_project" "test" {
  location = "ca-tor"
  resource_group = "aaaaaaa"
  definition {
  }
}
```
