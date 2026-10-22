# ibm_project_config

Schema validation for the ibm_project_config resource.

This rule validates `ibm_project_config` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Project update` (`projects`).

## Configuration

```hcl
rule "ibm_project_config" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `container_state` | no | — |
| `container_state_code` | no | — |
| `created_at` | no | — |
| `deployment_model` | no | — |
| `href` | no | — |
| `is_draft` | no | — |
| `last_saved_at` | no | — |
| `modified_at` | no | — |
| `needs_attention_state` | no | — |
| `outputs` | no | — |
| `project` | no | — |
| `project_config_id` | no | — |
| `project_id` | yes | pattern `^[\.\-0-9a-zA-Z]+$` |
| `references` | no | — |
| `state` | no | — |
| `state_code` | no | — |
| `template_id` | no | — |
| `update_available` | no | — |
| `version` | no | — |
| `approved_version` (block) | no | nested block |
| `approved_version.container_state` | no | one of: `approved`, `deleted`, `deleting`, `deleting_failed`, `discarded`, `draft`, `deployed`, `deploying_failed`, `deploying`, `superseded`, `undeploying`, `undeploying_failed`, `validated`, `validating`, `validating_failed`, `working` |
| `config_error` (block) | no | nested block |
| `config_error.message` | no | — |
| `config_error.details` (block) | no | nested block |
| `definition` (block) | yes | nested block |
| `definition.description` | no | — |
| `definition.environment_id` | no | pattern `^[\.\-0-9a-zA-Z]+$` |
| `definition.inputs` | no | — |
| `definition.locator_id` | no | — |
| `definition.name` | no | — |
| `definition.resource_crns` | no | API: An IBM Cloud resource name that uniquely identifies a resource. |
| `definition.settings` | no | — |
| `definition.authorizations` (block) | no | nested block |
| `definition.authorizations.api_key` | no | API: The IBM Cloud API Key. It can be either raw or pulled from the catalog via a `CRN` or `JSON` blob. |
| `definition.authorizations.method` | no | API: The authorization method. It can authorize by using a trusted profile or an API key in Secrets Manager. |
| `definition.authorizations.trusted_profile_id` | no | API: The trusted profile ID. |
| `definition.compliance_profile` (block) | no | nested block |
| `definition.compliance_profile.attachment_id` | no | pattern `^[\./:a-zA-Z0-9\.\-]+$` |
| `definition.compliance_profile.id` | no | — |
| `definition.compliance_profile.instance_id` | no | pattern `^[\./:a-zA-Z0-9\.\-]+$` |
| `definition.compliance_profile.instance_location` | no | one of: `us-south`, `us-east`, `eu-gb`, `eu-de`, `ca-tor` |
| `definition.compliance_profile.profile_name` | no | API: The name of the compliance profile. |
| `definition.compliance_profile.wp_instance_id` | no | pattern `^[\./:a-zA-Z0-9\.\-]+$` |
| `definition.compliance_profile.wp_instance_location` | no | one of: `us-south`, `us-east`, `eu-gb`, `eu-de`, `ca-tor` |
| `definition.compliance_profile.wp_instance_name` | no | API: The name of the compliance profile. |
| `definition.compliance_profile.wp_policy_id` | no | pattern `^[\./:a-zA-Z0-9\.\-]+$` |
| `definition.compliance_profile.wp_policy_name` | no | API: The name of the compliance profile. |
| `definition.compliance_profile.wp_zone_id` | no | pattern `^[\./:a-zA-Z0-9\.\-]+$` |
| `definition.compliance_profile.wp_zone_name` | no | API: The name of the compliance profile. |
| `definition.members` (block) | no | nested block |
| `definition.members.config_id` | yes | pattern `^[\.\-0-9a-zA-Z]+$` |
| `definition.members.name` | yes | — |
| `definition.uses` (block) | no | nested block |
| `definition.uses.config_id` | yes | pattern `^[\.\-0-9a-zA-Z]+$` |
| `definition.uses.project_id` | yes | pattern `^[\.\-0-9a-zA-Z]+$` |
| `deployed_version` (block) | no | nested block |
| `deployed_version.container_state` | no | one of: `approved`, `deleted`, `deleting`, `deleting_failed`, `discarded`, `draft`, `deployed`, `deploying_failed`, `deploying`, `superseded`, `undeploying`, `undeploying_failed`, `validated`, `validating`, `validating_failed`, `working` |
| `member_of` (block) | no | nested block |
| `schematics` (block) | no | nested block |
| `schematics.workspace_crn` | no | crn format |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_project_config_invalid_project_id` | `project_id` |
| `ibm_project_config_invalid_approved_version__container_state` | `approved_version.container_state` |
| `ibm_project_config_invalid_config_error__message` | `config_error.message` |
| `ibm_project_config_invalid_definition__description` | `definition.description` |
| `ibm_project_config_invalid_definition__environment_id` | `definition.environment_id` |
| `ibm_project_config_invalid_definition__inputs` | `definition.inputs` |
| `ibm_project_config_invalid_definition__locator_id` | `definition.locator_id` |
| `ibm_project_config_invalid_definition__name` | `definition.name` |
| `ibm_project_config_invalid_definition__resource_crns` | `definition.resource_crns` |
| `ibm_project_config_invalid_definition__settings` | `definition.settings` |
| `ibm_project_config_invalid_definition__authorizations__api_key` | `definition.authorizations.api_key` |
| `ibm_project_config_invalid_definition__authorizations__method` | `definition.authorizations.method` |
| `ibm_project_config_invalid_definition__authorizations__trusted_profile_id` | `definition.authorizations.trusted_profile_id` |
| `ibm_project_config_invalid_definition__compliance_profile__attachment_id` | `definition.compliance_profile.attachment_id` |
| `ibm_project_config_invalid_definition__compliance_profile__id` | `definition.compliance_profile.id` |
| `ibm_project_config_invalid_definition__compliance_profile__instance_id` | `definition.compliance_profile.instance_id` |
| `ibm_project_config_invalid_definition__compliance_profile__instance_location` | `definition.compliance_profile.instance_location` |
| `ibm_project_config_invalid_definition__compliance_profile__profile_name` | `definition.compliance_profile.profile_name` |
| `ibm_project_config_invalid_definition__compliance_profile__wp_instance_id` | `definition.compliance_profile.wp_instance_id` |
| `ibm_project_config_invalid_definition__compliance_profile__wp_instance_location` | `definition.compliance_profile.wp_instance_location` |
| `ibm_project_config_invalid_definition__compliance_profile__wp_instance_name` | `definition.compliance_profile.wp_instance_name` |
| `ibm_project_config_invalid_definition__compliance_profile__wp_policy_id` | `definition.compliance_profile.wp_policy_id` |
| `ibm_project_config_invalid_definition__compliance_profile__wp_policy_name` | `definition.compliance_profile.wp_policy_name` |
| `ibm_project_config_invalid_definition__compliance_profile__wp_zone_id` | `definition.compliance_profile.wp_zone_id` |
| `ibm_project_config_invalid_definition__compliance_profile__wp_zone_name` | `definition.compliance_profile.wp_zone_name` |
| `ibm_project_config_invalid_definition__members__config_id` | `definition.members.config_id` |
| `ibm_project_config_invalid_definition__members__name` | `definition.members.name` |
| `ibm_project_config_invalid_definition__uses__config_id` | `definition.uses.config_id` |
| `ibm_project_config_invalid_definition__uses__project_id` | `definition.uses.project_id` |
| `ibm_project_config_invalid_deployed_version__container_state` | `deployed_version.container_state` |
| `ibm_project_config_invalid_schematics__workspace_crn` | `schematics.workspace_crn` |

## Example

### Valid

```hcl
resource "ibm_project_config" "test" {
  project_id = "example-name"
  definition {
  }
}
```
