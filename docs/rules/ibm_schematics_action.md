# ibm_schematics_action

Manages ibm_schematics_action. The unique name of your action.

This rule validates `ibm_schematics_action` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Action` (`schematics-schematics`).

## Configuration

```hcl
rule "ibm_schematics_action" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account` | no | — |
| `command_parameter` | no | API: Schematics job command parameter (playbook-name). |
| `created_at` | no | — |
| `created_by` | no | — |
| `crn` | no | — |
| `description` | no | API: Action description |
| `inventory` | no | API: Target inventory record ID, used by the action or ansible playbook |
| `location` | no | one of: `us-south`, `us-east`, `eu-gb`, `eu-de` |
| `name` | yes | API: The unique name of your action. The name can be up to 128 characters long and can include alphanumeric characters, spaces, dashes, and underscores. **Example** you can use the name to stop action. |
| `playbook_names` | no | — |
| `resource_group` | no | API: Resource-group name for an action. By default, an action is created in `Default` resource group. |
| `source_created_at` | no | — |
| `source_created_by` | no | — |
| `source_readme_url` | no | API: URL of the `README` file, for the source URL. |
| `source_type` | no | one of: `cos_bucket`, `external_scm`, `git_hub`, `git_hub_enterprise`, `git_lab`, `ibm_schematics_action_catalog`, `ibm_git_lab`, `local` |
| `source_updated_at` | no | — |
| `source_updated_by` | no | — |
| `state` | no | — |
| `sys_lock` | no | — |
| `tags` | no | API: Action tags |
| `targets_ini` | no | API: Inventory of host and host group for the playbook in `INI` file format. For example, `"targets_ini": "[webserverhost] 172.22.192.6 [dbhost] 172.22.192.5"`. For more information, about an inventory host group syntax, see [Inventory host groups](https://cloud.ibm.com/docs/schematics?topic=schematics-schematics-cli-reference#schematics-inventory-host-grps). |
| `updated_at` | no | — |
| `updated_by` | no | — |
| `x_github_token` | no | — |
| `action_inputs` (block) | no | nested block |
| `action_inputs.link` | no | API: The reference link to the variable value By default the expression points to `$self.value`. |
| `action_inputs.name` | no | API: The unique name of your action. The name can be up to 128 characters long and can include alphanumeric characters, spaces, dashes, and underscores. **Example** you can use the name to stop action. |
| `action_inputs.value` | no | API: The credential value for the variable or reference to the value. **Note** The SSH key should contain three `\n` in the SSH key value as shown in the example for CLI or API calls. When using Bastion from API you need to add SSH key in both `credentials` and `bastion_credentials`. |
| `action_inputs.metadata` (block) | no | nested block |
| `action_inputs.metadata.aliases` | no | API: The list of aliases for the variable name. |
| `action_inputs.metadata.default_value` | no | API: Default value for the variable only if the override value is not specified. |
| `action_inputs.metadata.description` | no | API: Action description |
| `action_inputs.metadata.group_by` | no | API: The display name of the group this variable belongs to. |
| `action_inputs.metadata.hidden` | no | API: If **true**, the variable is not displayed on UI or Command line. |
| `action_inputs.metadata.immutable` | no | API: Is the variable readonly ? |
| `action_inputs.metadata.matches` | no | API: The regex for the variable value. |
| `action_inputs.metadata.max_length` | no | range 0–10000 |
| `action_inputs.metadata.max_value` | no | range 0–10000 |
| `action_inputs.metadata.min_length` | no | range 0–10000 |
| `action_inputs.metadata.min_value` | no | range 0–10000 |
| `action_inputs.metadata.options` | no | API: The list of possible values for this variable. If type is **integer** or **date**, then the array of string is converted to array of integers or date during the runtime. |
| `action_inputs.metadata.position` | no | range 0–10000 |
| `action_inputs.metadata.secure` | no | API: Is the variable secure or sensitive ? |
| `action_inputs.metadata.source` | no | API: Source of templates, playbooks, or controls |
| `action_inputs.metadata.type` | no | one of: `string`, `link` |
| `action_outputs` (block) | no | nested block |
| `action_outputs.link` | no | API: The reference link to the variable value By default the expression points to `$self.value`. |
| `action_outputs.name` | no | API: The unique name of your action. The name can be up to 128 characters long and can include alphanumeric characters, spaces, dashes, and underscores. **Example** you can use the name to stop action. |
| `action_outputs.value` | no | API: The credential value for the variable or reference to the value. **Note** The SSH key should contain three `\n` in the SSH key value as shown in the example for CLI or API calls. When using Bastion from API you need to add SSH key in both `credentials` and `bastion_credentials`. |
| `action_outputs.metadata` (block) | no | nested block |
| `action_outputs.metadata.aliases` | no | API: The list of aliases for the variable name. |
| `action_outputs.metadata.default_value` | no | API: Default value for the variable only if the override value is not specified. |
| `action_outputs.metadata.description` | no | API: Action description |
| `action_outputs.metadata.group_by` | no | API: The display name of the group this variable belongs to. |
| `action_outputs.metadata.hidden` | no | API: If **true**, the variable is not displayed on UI or Command line. |
| `action_outputs.metadata.immutable` | no | API: Is the variable readonly ? |
| `action_outputs.metadata.matches` | no | API: The regex for the variable value. |
| `action_outputs.metadata.max_length` | no | range 0–10000 |
| `action_outputs.metadata.max_value` | no | range 0–10000 |
| `action_outputs.metadata.min_length` | no | range 0–10000 |
| `action_outputs.metadata.min_value` | no | range 0–10000 |
| `action_outputs.metadata.options` | no | API: The list of possible values for this variable. If type is **integer** or **date**, then the array of string is converted to array of integers or date during the runtime. |
| `action_outputs.metadata.position` | no | range 0–10000 |
| `action_outputs.metadata.secure` | no | API: Is the variable secure or sensitive ? |
| `action_outputs.metadata.source` | no | API: Source of templates, playbooks, or controls |
| `action_outputs.metadata.type` | no | one of: `string`, `link` |
| `bastion` (block) | no | nested block |
| `bastion.host` | no | API: Reference to the Inventory resource definition |
| `bastion.name` | no | API: Bastion Name; the name must be unique |
| `bastion_credential` (block) | no | nested block |
| `bastion_credential.link` | no | API: The reference link to the variable value By default the expression points to `$self.value`. |
| `bastion_credential.name` | no | API: The name of the credential variable. |
| `bastion_credential.value` | no | API: The credential value for the variable or reference to the value. **Note** The SSH key should contain three `\n` in the SSH key value as shown in the example for CLI or API calls. When using Bastion from API you need to add SSH key in both `credentials` and `bastion_credentials`. |
| `bastion_credential.metadata` (block) | no | nested block |
| `bastion_credential.metadata.aliases` | no | API: The list of aliases for the variable name. |
| `bastion_credential.metadata.default_value` | no | API: Default value for the variable only if the override value is not specified. |
| `bastion_credential.metadata.description` | no | API: The description of the meta data. |
| `bastion_credential.metadata.group_by` | no | API: The display name of the group this variable belongs to. |
| `bastion_credential.metadata.hidden` | no | API: If **true**, the variable is not displayed on UI or Command line. |
| `bastion_credential.metadata.immutable` | no | API: Is the variable readonly ? |
| `bastion_credential.metadata.matches` | no | API: The regex for the variable value. |
| `bastion_credential.metadata.max_length` | no | range 0–10000 |
| `bastion_credential.metadata.max_value` | no | range 0–10000 |
| `bastion_credential.metadata.min_length` | no | range 0–10000 |
| `bastion_credential.metadata.min_value` | no | range 0–10000 |
| `bastion_credential.metadata.options` | no | API: The list of possible values for this variable. If type is **integer** or **date**, then the array of string is converted to array of integers or date during the runtime. |
| `bastion_credential.metadata.position` | no | range 0–10000 |
| `bastion_credential.metadata.secure` | no | API: Is the variable secure or sensitive ? |
| `bastion_credential.metadata.source` | no | API: The source of this meta-data. |
| `bastion_credential.metadata.type` | no | one of: `string`, `link` |
| `credentials` (block) | no | nested block |
| `credentials.link` | no | API: The reference link to the variable value By default the expression points to `$self.value`. |
| `credentials.name` | no | API: The name of the credential variable. |
| `credentials.value` | no | API: The credential value for the variable or reference to the value. **Note** The SSH key should contain three `\n` in the SSH key value as shown in the example for CLI or API calls. When using Bastion from API you need to add SSH key in both `credentials` and `bastion_credentials`. |
| `credentials.metadata` (block) | no | nested block |
| `credentials.metadata.aliases` | no | API: The list of aliases for the variable name. |
| `credentials.metadata.default_value` | no | API: Default value for the variable only if the override value is not specified. |
| `credentials.metadata.description` | no | API: The description of the meta data. |
| `credentials.metadata.group_by` | no | API: The display name of the group this variable belongs to. |
| `credentials.metadata.hidden` | no | API: If **true**, the variable is not displayed on UI or Command line. |
| `credentials.metadata.immutable` | no | API: Is the variable readonly ? |
| `credentials.metadata.matches` | no | API: The regex for the variable value. |
| `credentials.metadata.max_length` | no | range 0–10000 |
| `credentials.metadata.max_value` | no | range 0–10000 |
| `credentials.metadata.min_length` | no | range 0–10000 |
| `credentials.metadata.min_value` | no | range 0–10000 |
| `credentials.metadata.options` | no | API: The list of possible values for this variable. If type is **integer** or **date**, then the array of string is converted to array of integers or date during the runtime. |
| `credentials.metadata.position` | no | range 0–10000 |
| `credentials.metadata.secure` | no | API: Is the variable secure or sensitive ? |
| `credentials.metadata.source` | no | API: The source of this meta-data. |
| `credentials.metadata.type` | no | one of: `string`, `link` |
| `settings` (block) | no | nested block |
| `settings.link` | no | API: The reference link to the variable value By default the expression points to `$self.value`. |
| `settings.name` | no | API: The name of the variable. For example, `name = "inventory username"`. |
| `settings.value` | no | API: The value for the variable or reference to the value. For example, `value = "&lt;provide your ssh_key_value with \n&gt;"`. |
| `settings.metadata` (block) | no | nested block |
| `settings.metadata.aliases` | no | API: The list of aliases for the variable name. |
| `settings.metadata.default_value` | no | API: Default value for the variable only if the override value is not specified. |
| `settings.metadata.description` | no | API: The description of the meta data. |
| `settings.metadata.group_by` | no | API: The display name of the group this variable belongs to. |
| `settings.metadata.hidden` | no | API: If **true**, the variable is not displayed on UI or Command line. |
| `settings.metadata.immutable` | no | API: Is the variable readonly ? |
| `settings.metadata.matches` | no | API: The regex for the variable value. |
| `settings.metadata.max_length` | no | range 0–10000 |
| `settings.metadata.max_value` | no | range 0–10000 |
| `settings.metadata.min_length` | no | range 0–10000 |
| `settings.metadata.min_value` | no | range 0–10000 |
| `settings.metadata.options` | no | API: The list of possible values for this variable. If type is **integer** or **date**, then the array of string is converted to array of integers or date during the runtime. |
| `settings.metadata.position` | no | range 0–10000 |
| `settings.metadata.secure` | no | API: Is the variable secure or sensitive ? |
| `settings.metadata.source` | no | API: The source of this meta-data. |
| `settings.metadata.type` | no | one of: `boolean`, `string`, `integer`, `date`, `array`, `list`, `map`, `complex`, `link` |
| `source` (block) | no | nested block |
| `source.source_type` | yes | one of: `local`, `git_hub`, `git_hub_enterprise`, `git_lab`, `ibm_git_lab`, `ibm_cloud_catalog` |
| `source.catalog` (block) | no | nested block |
| `source.catalog.catalog_name` | no | API: The name of the private catalog. |
| `source.catalog.offering_id` | no | API: The ID of an offering in the IBM Cloud Catalog. |
| `source.catalog.offering_kind` | no | API: The type of an offering, in the IBM Cloud Catalog. |
| `source.catalog.offering_name` | no | API: The name of an offering in the IBM Cloud Catalog. |
| `source.catalog.offering_repo_url` | no | API: The repository URL of an offering, in the IBM Cloud Catalog. |
| `source.catalog.offering_version` | no | API: The version of the software template that you chose to install from the IBM Cloud catalog |
| `source.catalog.offering_version_id` | no | API: The ID of an offering version the IBM Cloud Catalog. |
| `source.git` (block) | no | nested block |
| `source.git.computed_git_repo_url` | no | API: The complete URL which is computed by the **git_repo_url**, **git_repo_folder**, and **branch**. |
| `source.git.git_branch` | no | API: The name of the branch that are used to fetch the Git repository. |
| `source.git.git_release` | no | API: The name of the release tag that are used to fetch the Git repository. |
| `source.git.git_repo_folder` | no | API: The name of the folder in the Git repository, that contains the template. |
| `source.git.git_repo_url` | no | API: The URL to the Git repository that can be used to clone the template. |
| `source.git.git_token` | no | API: The Personal Access Token (PAT) to connect to the Git URLs. |
| `user_state` (block) | no | nested block |
| `user_state.set_at` | no | API: When the User who set the state of the Object |
| `user_state.set_by` | no | API: Name of the User who set the state of the Object |
| `user_state.state` | no | one of: `draft`, `live`, `locked`, `disable` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_schematics_action_invalid_command_parameter` | `command_parameter` |
| `ibm_schematics_action_invalid_description` | `description` |
| `ibm_schematics_action_invalid_inventory` | `inventory` |
| `ibm_schematics_action_invalid_location` | `location` |
| `ibm_schematics_action_invalid_name` | `name` |
| `ibm_schematics_action_invalid_resource_group` | `resource_group` |
| `ibm_schematics_action_invalid_source_readme_url` | `source_readme_url` |
| `ibm_schematics_action_invalid_source_type` | `source_type` |
| `ibm_schematics_action_invalid_tags` | `tags` |
| `ibm_schematics_action_invalid_targets_ini` | `targets_ini` |
| `ibm_schematics_action_invalid_x_github_token` | `x_github_token` |
| `ibm_schematics_action_invalid_action_inputs__link` | `action_inputs.link` |
| `ibm_schematics_action_invalid_action_inputs__name` | `action_inputs.name` |
| `ibm_schematics_action_invalid_action_inputs__value` | `action_inputs.value` |
| `ibm_schematics_action_invalid_action_inputs__metadata__aliases` | `action_inputs.metadata.aliases` |
| `ibm_schematics_action_invalid_action_inputs__metadata__default_value` | `action_inputs.metadata.default_value` |
| `ibm_schematics_action_invalid_action_inputs__metadata__description` | `action_inputs.metadata.description` |
| `ibm_schematics_action_invalid_action_inputs__metadata__group_by` | `action_inputs.metadata.group_by` |
| `ibm_schematics_action_invalid_action_inputs__metadata__hidden` | `action_inputs.metadata.hidden` |
| `ibm_schematics_action_invalid_action_inputs__metadata__immutable` | `action_inputs.metadata.immutable` |
| `ibm_schematics_action_invalid_action_inputs__metadata__matches` | `action_inputs.metadata.matches` |
| `ibm_schematics_action_invalid_action_inputs__metadata__max_length` | `action_inputs.metadata.max_length` |
| `ibm_schematics_action_invalid_action_inputs__metadata__max_value` | `action_inputs.metadata.max_value` |
| `ibm_schematics_action_invalid_action_inputs__metadata__min_length` | `action_inputs.metadata.min_length` |
| `ibm_schematics_action_invalid_action_inputs__metadata__min_value` | `action_inputs.metadata.min_value` |
| `ibm_schematics_action_invalid_action_inputs__metadata__options` | `action_inputs.metadata.options` |
| `ibm_schematics_action_invalid_action_inputs__metadata__position` | `action_inputs.metadata.position` |
| `ibm_schematics_action_invalid_action_inputs__metadata__secure` | `action_inputs.metadata.secure` |
| `ibm_schematics_action_invalid_action_inputs__metadata__source` | `action_inputs.metadata.source` |
| `ibm_schematics_action_invalid_action_inputs__metadata__type` | `action_inputs.metadata.type` |
| `ibm_schematics_action_invalid_action_outputs__link` | `action_outputs.link` |
| `ibm_schematics_action_invalid_action_outputs__name` | `action_outputs.name` |
| `ibm_schematics_action_invalid_action_outputs__value` | `action_outputs.value` |
| `ibm_schematics_action_invalid_action_outputs__metadata__aliases` | `action_outputs.metadata.aliases` |
| `ibm_schematics_action_invalid_action_outputs__metadata__default_value` | `action_outputs.metadata.default_value` |
| `ibm_schematics_action_invalid_action_outputs__metadata__description` | `action_outputs.metadata.description` |
| `ibm_schematics_action_invalid_action_outputs__metadata__group_by` | `action_outputs.metadata.group_by` |
| `ibm_schematics_action_invalid_action_outputs__metadata__hidden` | `action_outputs.metadata.hidden` |
| `ibm_schematics_action_invalid_action_outputs__metadata__immutable` | `action_outputs.metadata.immutable` |
| `ibm_schematics_action_invalid_action_outputs__metadata__matches` | `action_outputs.metadata.matches` |
| `ibm_schematics_action_invalid_action_outputs__metadata__max_length` | `action_outputs.metadata.max_length` |
| `ibm_schematics_action_invalid_action_outputs__metadata__max_value` | `action_outputs.metadata.max_value` |
| `ibm_schematics_action_invalid_action_outputs__metadata__min_length` | `action_outputs.metadata.min_length` |
| `ibm_schematics_action_invalid_action_outputs__metadata__min_value` | `action_outputs.metadata.min_value` |
| `ibm_schematics_action_invalid_action_outputs__metadata__options` | `action_outputs.metadata.options` |
| `ibm_schematics_action_invalid_action_outputs__metadata__position` | `action_outputs.metadata.position` |
| `ibm_schematics_action_invalid_action_outputs__metadata__secure` | `action_outputs.metadata.secure` |
| `ibm_schematics_action_invalid_action_outputs__metadata__source` | `action_outputs.metadata.source` |
| `ibm_schematics_action_invalid_action_outputs__metadata__type` | `action_outputs.metadata.type` |
| `ibm_schematics_action_invalid_bastion__host` | `bastion.host` |
| `ibm_schematics_action_invalid_bastion__name` | `bastion.name` |
| `ibm_schematics_action_invalid_bastion_credential__link` | `bastion_credential.link` |
| `ibm_schematics_action_invalid_bastion_credential__name` | `bastion_credential.name` |
| `ibm_schematics_action_invalid_bastion_credential__value` | `bastion_credential.value` |
| `ibm_schematics_action_invalid_bastion_credential__metadata__aliases` | `bastion_credential.metadata.aliases` |
| `ibm_schematics_action_invalid_bastion_credential__metadata__default_value` | `bastion_credential.metadata.default_value` |
| `ibm_schematics_action_invalid_bastion_credential__metadata__description` | `bastion_credential.metadata.description` |
| `ibm_schematics_action_invalid_bastion_credential__metadata__group_by` | `bastion_credential.metadata.group_by` |
| `ibm_schematics_action_invalid_bastion_credential__metadata__hidden` | `bastion_credential.metadata.hidden` |
| `ibm_schematics_action_invalid_bastion_credential__metadata__immutable` | `bastion_credential.metadata.immutable` |
| `ibm_schematics_action_invalid_bastion_credential__metadata__matches` | `bastion_credential.metadata.matches` |
| `ibm_schematics_action_invalid_bastion_credential__metadata__max_length` | `bastion_credential.metadata.max_length` |
| `ibm_schematics_action_invalid_bastion_credential__metadata__max_value` | `bastion_credential.metadata.max_value` |
| `ibm_schematics_action_invalid_bastion_credential__metadata__min_length` | `bastion_credential.metadata.min_length` |
| `ibm_schematics_action_invalid_bastion_credential__metadata__min_value` | `bastion_credential.metadata.min_value` |
| `ibm_schematics_action_invalid_bastion_credential__metadata__options` | `bastion_credential.metadata.options` |
| `ibm_schematics_action_invalid_bastion_credential__metadata__position` | `bastion_credential.metadata.position` |
| `ibm_schematics_action_invalid_bastion_credential__metadata__secure` | `bastion_credential.metadata.secure` |
| `ibm_schematics_action_invalid_bastion_credential__metadata__source` | `bastion_credential.metadata.source` |
| `ibm_schematics_action_invalid_bastion_credential__metadata__type` | `bastion_credential.metadata.type` |
| `ibm_schematics_action_invalid_credentials__link` | `credentials.link` |
| `ibm_schematics_action_invalid_credentials__name` | `credentials.name` |
| `ibm_schematics_action_invalid_credentials__value` | `credentials.value` |
| `ibm_schematics_action_invalid_credentials__metadata__aliases` | `credentials.metadata.aliases` |
| `ibm_schematics_action_invalid_credentials__metadata__default_value` | `credentials.metadata.default_value` |
| `ibm_schematics_action_invalid_credentials__metadata__description` | `credentials.metadata.description` |
| `ibm_schematics_action_invalid_credentials__metadata__group_by` | `credentials.metadata.group_by` |
| `ibm_schematics_action_invalid_credentials__metadata__hidden` | `credentials.metadata.hidden` |
| `ibm_schematics_action_invalid_credentials__metadata__immutable` | `credentials.metadata.immutable` |
| `ibm_schematics_action_invalid_credentials__metadata__matches` | `credentials.metadata.matches` |
| `ibm_schematics_action_invalid_credentials__metadata__max_length` | `credentials.metadata.max_length` |
| `ibm_schematics_action_invalid_credentials__metadata__max_value` | `credentials.metadata.max_value` |
| `ibm_schematics_action_invalid_credentials__metadata__min_length` | `credentials.metadata.min_length` |
| `ibm_schematics_action_invalid_credentials__metadata__min_value` | `credentials.metadata.min_value` |
| `ibm_schematics_action_invalid_credentials__metadata__options` | `credentials.metadata.options` |
| `ibm_schematics_action_invalid_credentials__metadata__position` | `credentials.metadata.position` |
| `ibm_schematics_action_invalid_credentials__metadata__secure` | `credentials.metadata.secure` |
| `ibm_schematics_action_invalid_credentials__metadata__source` | `credentials.metadata.source` |
| `ibm_schematics_action_invalid_credentials__metadata__type` | `credentials.metadata.type` |
| `ibm_schematics_action_invalid_settings__link` | `settings.link` |
| `ibm_schematics_action_invalid_settings__name` | `settings.name` |
| `ibm_schematics_action_invalid_settings__value` | `settings.value` |
| `ibm_schematics_action_invalid_settings__metadata__aliases` | `settings.metadata.aliases` |
| `ibm_schematics_action_invalid_settings__metadata__default_value` | `settings.metadata.default_value` |
| `ibm_schematics_action_invalid_settings__metadata__description` | `settings.metadata.description` |
| `ibm_schematics_action_invalid_settings__metadata__group_by` | `settings.metadata.group_by` |
| `ibm_schematics_action_invalid_settings__metadata__hidden` | `settings.metadata.hidden` |
| `ibm_schematics_action_invalid_settings__metadata__immutable` | `settings.metadata.immutable` |
| `ibm_schematics_action_invalid_settings__metadata__matches` | `settings.metadata.matches` |
| `ibm_schematics_action_invalid_settings__metadata__max_length` | `settings.metadata.max_length` |
| `ibm_schematics_action_invalid_settings__metadata__max_value` | `settings.metadata.max_value` |
| `ibm_schematics_action_invalid_settings__metadata__min_length` | `settings.metadata.min_length` |
| `ibm_schematics_action_invalid_settings__metadata__min_value` | `settings.metadata.min_value` |
| `ibm_schematics_action_invalid_settings__metadata__options` | `settings.metadata.options` |
| `ibm_schematics_action_invalid_settings__metadata__position` | `settings.metadata.position` |
| `ibm_schematics_action_invalid_settings__metadata__secure` | `settings.metadata.secure` |
| `ibm_schematics_action_invalid_settings__metadata__source` | `settings.metadata.source` |
| `ibm_schematics_action_invalid_settings__metadata__type` | `settings.metadata.type` |
| `ibm_schematics_action_invalid_source__source_type` | `source.source_type` |
| `ibm_schematics_action_invalid_source__catalog__catalog_name` | `source.catalog.catalog_name` |
| `ibm_schematics_action_invalid_source__catalog__offering_id` | `source.catalog.offering_id` |
| `ibm_schematics_action_invalid_source__catalog__offering_kind` | `source.catalog.offering_kind` |
| `ibm_schematics_action_invalid_source__catalog__offering_name` | `source.catalog.offering_name` |
| `ibm_schematics_action_invalid_source__catalog__offering_repo_url` | `source.catalog.offering_repo_url` |
| `ibm_schematics_action_invalid_source__catalog__offering_version` | `source.catalog.offering_version` |
| `ibm_schematics_action_invalid_source__catalog__offering_version_id` | `source.catalog.offering_version_id` |
| `ibm_schematics_action_invalid_source__git__computed_git_repo_url` | `source.git.computed_git_repo_url` |
| `ibm_schematics_action_invalid_source__git__git_branch` | `source.git.git_branch` |
| `ibm_schematics_action_invalid_source__git__git_release` | `source.git.git_release` |
| `ibm_schematics_action_invalid_source__git__git_repo_folder` | `source.git.git_repo_folder` |
| `ibm_schematics_action_invalid_source__git__git_repo_url` | `source.git.git_repo_url` |
| `ibm_schematics_action_invalid_source__git__git_token` | `source.git.git_token` |
| `ibm_schematics_action_invalid_user_state__set_at` | `user_state.set_at` |
| `ibm_schematics_action_invalid_user_state__set_by` | `user_state.set_by` |
| `ibm_schematics_action_invalid_user_state__state` | `user_state.state` |

## Example

### Valid

```hcl
resource "ibm_schematics_action" "test" {
  name = "aaaaaaa"
}
```
