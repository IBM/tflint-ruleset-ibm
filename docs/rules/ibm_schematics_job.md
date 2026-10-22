# ibm_schematics_job

Manages ibm_schematics_job. Job name, uniquely derived from the related Workspace or Action.

This rule validates `ibm_schematics_job` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Job` (`schematics-schematics`).

## Configuration

```hcl
rule "ibm_schematics_job" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `command_name` | yes | one of: `ansible_playbook_check`, `ansible_playbook_run`, `create_action`, `create_cart`, `create_environment`, `create_workspace`, `delete_action`, `delete_environment`, `delete_workspace`, `environment_init`, `environment_install`, `environment_uninstall`, `patch_action`, `patch_workspace`, `put_action`, `put_environment`, `put_workspace`, `repository_process`, `system_key_delete`, `system_key_disable`, `system_key_enable`, `system_key_restore`, `system_key_rotate`, `workspace_apply`, `workspace_destroy`, `workspace_plan`, `workspace_refresh` |
| `command_object` | yes | one of: `action`, `environment`, `system`, `workspace` |
| `command_object_id` | yes | API: Job command object id (workspace-id, action-id) |
| `command_options` | no | API: Command line options for the command |
| `command_parameter` | no | API: Schematics job command parameter (playbook-name) |
| `description` | no | — |
| `duration` | no | — |
| `end_at` | no | — |
| `location` | no | one of: `us-south`, `us-east`, `eu-gb`, `eu-de` |
| `log_store_url` | no | — |
| `name` | no | — |
| `resource_group` | no | — |
| `results_url` | no | — |
| `start_at` | no | — |
| `state_store_url` | no | — |
| `status` | no | — |
| `submitted_at` | no | — |
| `submitted_by` | no | — |
| `tags` | no | API: User defined tags, while running the job |
| `updated_at` | no | — |
| `bastion` (block) | no | nested block |
| `bastion.host` | no | API: Reference to the Inventory resource definition |
| `bastion.name` | no | API: Bastion Name; the name must be unique |
| `data` (block) | no | nested block |
| `data.job_type` | yes | one of: `repo_download_job`, `workspace_job`, `action_job`, `system_job`, `flow-job` |
| `data.action_job_data` (block) | no | nested block |
| `data.action_job_data.action_name` | no | API: Flow name |
| `data.action_job_data.materialized_inventory` | no | API: Materialized inventory details used by the Action Job, in .ini format |
| `data.action_job_data.updated_at` | no | API: Job status updation timestamp |
| `data.action_job_data.inputs` (block) | no | nested block |
| `data.action_job_data.inputs.link` | no | API: The reference link to the variable value By default the expression points to `$self.value`. |
| `data.action_job_data.inputs.name` | no | API: The name of the variable. For example, `name = "inventory username"`. |
| `data.action_job_data.inputs.value` | no | API: The value for the variable or reference to the value. For example, `value = "&lt;provide your ssh_key_value with \n&gt;"`. |
| `data.action_job_data.inputs.metadata` (block) | no | nested block |
| `data.action_job_data.inputs.metadata.aliases` | no | API: The list of aliases for the variable name. |
| `data.action_job_data.inputs.metadata.default_value` | no | API: Default value for the variable only if the override value is not specified. |
| `data.action_job_data.inputs.metadata.description` | no | API: The description of the meta data. |
| `data.action_job_data.inputs.metadata.group_by` | no | API: The display name of the group this variable belongs to. |
| `data.action_job_data.inputs.metadata.hidden` | no | API: If **true**, the variable is not displayed on UI or Command line. |
| `data.action_job_data.inputs.metadata.immutable` | no | API: Is the variable readonly ? |
| `data.action_job_data.inputs.metadata.matches` | no | API: The regex for the variable value. |
| `data.action_job_data.inputs.metadata.max_length` | no | range 0–10000 |
| `data.action_job_data.inputs.metadata.max_value` | no | range 0–10000 |
| `data.action_job_data.inputs.metadata.min_length` | no | range 0–10000 |
| `data.action_job_data.inputs.metadata.min_value` | no | range 0–10000 |
| `data.action_job_data.inputs.metadata.options` | no | API: The list of possible values for this variable. If type is **integer** or **date**, then the array of string is converted to array of integers or date during the runtime. |
| `data.action_job_data.inputs.metadata.position` | no | range 0–10000 |
| `data.action_job_data.inputs.metadata.secure` | no | API: Is the variable secure or sensitive ? |
| `data.action_job_data.inputs.metadata.source` | no | API: The source of this meta-data. |
| `data.action_job_data.inputs.metadata.type` | no | one of: `boolean`, `string`, `integer`, `date`, `array`, `list`, `map`, `complex`, `link` |
| `data.action_job_data.inventory_record` (block) | no | nested block |
| `data.action_job_data.inventory_record.created_at` | no | API: Inventory creation time |
| `data.action_job_data.inventory_record.created_by` | no | API: Email address of user who created the Inventory |
| `data.action_job_data.inventory_record.description` | no | API: The description of your Inventory. The description can be up to 2048 characters long in size. |
| `data.action_job_data.inventory_record.id` | no | API: Inventory id |
| `data.action_job_data.inventory_record.inventories_ini` | no | API: Input inventory of host and host group for the playbook, in the .ini file format. |
| `data.action_job_data.inventory_record.location` | no | one of: `us-south`, `us-east`, `eu-gb`, `eu-de` |
| `data.action_job_data.inventory_record.name` | no | API: The unique name of your Inventory. The name can be up to 128 characters long and can include alphanumeric characters, spaces, dashes, and underscores. |
| `data.action_job_data.inventory_record.resource_group` | no | API: Resource-group name for the Inventory definition. By default, Inventory will be created in Default Resource Group |
| `data.action_job_data.inventory_record.resource_queries` | no | API: Input resource queries that is used to dynamically generate the inventory of host and host group for the playbook |
| `data.action_job_data.inventory_record.updated_at` | no | API: Inventory updation time |
| `data.action_job_data.inventory_record.updated_by` | no | API: Email address of user who updated the Inventory |
| `data.action_job_data.outputs` (block) | no | nested block |
| `data.action_job_data.outputs.link` | no | API: The reference link to the variable value By default the expression points to `$self.value`. |
| `data.action_job_data.outputs.name` | no | API: The name of the variable. For example, `name = "inventory username"`. |
| `data.action_job_data.outputs.value` | no | API: The value for the variable or reference to the value. For example, `value = "&lt;provide your ssh_key_value with \n&gt;"`. |
| `data.action_job_data.outputs.metadata` (block) | no | nested block |
| `data.action_job_data.outputs.metadata.aliases` | no | API: The list of aliases for the variable name. |
| `data.action_job_data.outputs.metadata.default_value` | no | API: Default value for the variable only if the override value is not specified. |
| `data.action_job_data.outputs.metadata.description` | no | API: The description of the meta data. |
| `data.action_job_data.outputs.metadata.group_by` | no | API: The display name of the group this variable belongs to. |
| `data.action_job_data.outputs.metadata.hidden` | no | API: If **true**, the variable is not displayed on UI or Command line. |
| `data.action_job_data.outputs.metadata.immutable` | no | API: Is the variable readonly ? |
| `data.action_job_data.outputs.metadata.matches` | no | API: The regex for the variable value. |
| `data.action_job_data.outputs.metadata.max_length` | no | range 0–10000 |
| `data.action_job_data.outputs.metadata.max_value` | no | range 0–10000 |
| `data.action_job_data.outputs.metadata.min_length` | no | range 0–10000 |
| `data.action_job_data.outputs.metadata.min_value` | no | range 0–10000 |
| `data.action_job_data.outputs.metadata.options` | no | API: The list of possible values for this variable. If type is **integer** or **date**, then the array of string is converted to array of integers or date during the runtime. |
| `data.action_job_data.outputs.metadata.position` | no | range 0–10000 |
| `data.action_job_data.outputs.metadata.secure` | no | API: Is the variable secure or sensitive ? |
| `data.action_job_data.outputs.metadata.source` | no | API: The source of this meta-data. |
| `data.action_job_data.outputs.metadata.type` | no | one of: `boolean`, `string`, `integer`, `date`, `array`, `list`, `map`, `complex`, `link` |
| `data.action_job_data.settings` (block) | no | nested block |
| `data.action_job_data.settings.link` | no | API: The reference link to the variable value By default the expression points to `$self.value`. |
| `data.action_job_data.settings.name` | no | API: The name of the variable. For example, `name = "inventory username"`. |
| `data.action_job_data.settings.value` | no | API: The value for the variable or reference to the value. For example, `value = "&lt;provide your ssh_key_value with \n&gt;"`. |
| `data.action_job_data.settings.metadata` (block) | no | nested block |
| `data.action_job_data.settings.metadata.aliases` | no | API: The list of aliases for the variable name. |
| `data.action_job_data.settings.metadata.default_value` | no | API: Default value for the variable only if the override value is not specified. |
| `data.action_job_data.settings.metadata.description` | no | API: The description of the meta data. |
| `data.action_job_data.settings.metadata.group_by` | no | API: The display name of the group this variable belongs to. |
| `data.action_job_data.settings.metadata.hidden` | no | API: If **true**, the variable is not displayed on UI or Command line. |
| `data.action_job_data.settings.metadata.immutable` | no | API: Is the variable readonly ? |
| `data.action_job_data.settings.metadata.matches` | no | API: The regex for the variable value. |
| `data.action_job_data.settings.metadata.max_length` | no | range 0–10000 |
| `data.action_job_data.settings.metadata.max_value` | no | range 0–10000 |
| `data.action_job_data.settings.metadata.min_length` | no | range 0–10000 |
| `data.action_job_data.settings.metadata.min_value` | no | range 0–10000 |
| `data.action_job_data.settings.metadata.options` | no | API: The list of possible values for this variable. If type is **integer** or **date**, then the array of string is converted to array of integers or date during the runtime. |
| `data.action_job_data.settings.metadata.position` | no | range 0–10000 |
| `data.action_job_data.settings.metadata.secure` | no | API: Is the variable secure or sensitive ? |
| `data.action_job_data.settings.metadata.source` | no | API: The source of this meta-data. |
| `data.action_job_data.settings.metadata.type` | no | one of: `boolean`, `string`, `integer`, `date`, `array`, `list`, `map`, `complex`, `link` |
| `data.flow_job_data` (block) | no | nested block |
| `data.flow_job_data.flow_id` | no | API: Flow ID |
| `data.flow_job_data.flow_name` | no | API: Flow Name |
| `data.flow_job_data.updated_at` | no | API: Job status updation timestamp |
| `data.flow_job_data.workitems` (block) | no | nested block |
| `data.flow_job_data.workitems.command_object_id` | no | API: command object id |
| `data.flow_job_data.workitems.command_object_name` | no | API: command object name |
| `data.flow_job_data.workitems.layers` | no | API: layer name |
| `data.flow_job_data.workitems.source_type` | no | one of: `local`, `git_hub`, `git_hub_enterprise`, `git_lab`, `ibm_git_lab`, `ibm_cloud_catalog` |
| `data.flow_job_data.workitems.updated_at` | no | API: Job status updation timestamp |
| `data.flow_job_data.workitems.inputs` (block) | no | nested block |
| `data.flow_job_data.workitems.inputs.link` | no | API: The reference link to the variable value By default the expression points to `$self.value`. |
| `data.flow_job_data.workitems.inputs.name` | no | API: The name of the variable. For example, `name = "inventory username"`. |
| `data.flow_job_data.workitems.inputs.value` | no | API: The value for the variable or reference to the value. For example, `value = "&lt;provide your ssh_key_value with \n&gt;"`. |
| `data.flow_job_data.workitems.inputs.metadata` (block) | no | nested block |
| `data.flow_job_data.workitems.last_job` (block) | no | nested block |
| `data.flow_job_data.workitems.last_job.command_name` | no | one of: `workspace_plan`, `workspace_apply`, `workspace_destroy`, `workspace_refresh`, `ansible_playbook_run`, `ansible_playbook_check`, `create_action`, `put_action`, `patch_action`, `delete_action`, `system_key_enable`, `system_key_delete`, `system_key_disable`, `system_key_rotate`, `system_key_restore`, `create_workspace`, `put_workspace`, `patch_workspace`, `delete_workspace`, `create_cart`, `repository_process`, `terraform_commands` |
| `data.flow_job_data.workitems.last_job.command_object` | no | one of: `workspace`, `action`, `system`, `environment` |
| `data.flow_job_data.workitems.last_job.command_object_id` | no | API: Workitem command object id, maps to workspace_id or action_id |
| `data.flow_job_data.workitems.last_job.command_object_name` | no | API: command object name (workspace_name/action_name) |
| `data.flow_job_data.workitems.last_job.job_id` | no | API: Workspace job id |
| `data.flow_job_data.workitems.last_job.job_status` | no | one of: `job_pending`, `job_in_progress`, `job_finished`, `job_failed`, `job_cancelled`, `job_stopped`, `job_stop_in_progress`, `job_ready_to_execute` |
| `data.flow_job_data.workitems.outputs` (block) | no | nested block |
| `data.flow_job_data.workitems.outputs.link` | no | API: The reference link to the variable value By default the expression points to `$self.value`. |
| `data.flow_job_data.workitems.outputs.name` | no | API: The name of the variable. For example, `name = "inventory username"`. |
| `data.flow_job_data.workitems.outputs.value` | no | API: The value for the variable or reference to the value. For example, `value = "&lt;provide your ssh_key_value with \n&gt;"`. |
| `data.flow_job_data.workitems.outputs.metadata` (block) | no | nested block |
| `data.flow_job_data.workitems.settings` (block) | no | nested block |
| `data.flow_job_data.workitems.settings.link` | no | API: The reference link to the variable value By default the expression points to `$self.value`. |
| `data.flow_job_data.workitems.settings.name` | no | API: The name of the variable. For example, `name = "inventory username"`. |
| `data.flow_job_data.workitems.settings.value` | no | API: The value for the variable or reference to the value. For example, `value = "&lt;provide your ssh_key_value with \n&gt;"`. |
| `data.flow_job_data.workitems.settings.metadata` (block) | no | nested block |
| `data.flow_job_data.workitems.source` (block) | no | nested block |
| `data.flow_job_data.workitems.source.source_type` | yes | one of: `local`, `git_hub`, `git_hub_enterprise`, `git_lab`, `ibm_git_lab`, `ibm_cloud_catalog` |
| `data.flow_job_data.workitems.source.catalog` (block) | no | nested block |
| `data.flow_job_data.workitems.source.git` (block) | no | nested block |
| `data.system_job_data` (block) | no | nested block |
| `data.system_job_data.key_id` | no | API: Key ID for which key event is generated |
| `data.system_job_data.schematics_resource_id` | no | API: List of the schematics resource id |
| `data.system_job_data.updated_at` | no | API: Job status updation timestamp |
| `data.workspace_job_data` (block) | no | nested block |
| `data.workspace_job_data.flow_id` | no | API: Flow Id |
| `data.workspace_job_data.flow_name` | no | API: Flow name |
| `data.workspace_job_data.updated_at` | no | API: Job status updation timestamp |
| `data.workspace_job_data.workspace_name` | no | API: Workspace name |
| `data.workspace_job_data.inputs` (block) | no | nested block |
| `data.workspace_job_data.inputs.link` | no | API: The reference link to the variable value By default the expression points to `$self.value`. |
| `data.workspace_job_data.inputs.name` | no | API: The name of the variable. For example, `name = "inventory username"`. |
| `data.workspace_job_data.inputs.value` | no | API: The value for the variable or reference to the value. For example, `value = "&lt;provide your ssh_key_value with \n&gt;"`. |
| `data.workspace_job_data.inputs.metadata` (block) | no | nested block |
| `data.workspace_job_data.inputs.metadata.aliases` | no | API: The list of aliases for the variable name. |
| `data.workspace_job_data.inputs.metadata.default_value` | no | API: Default value for the variable only if the override value is not specified. |
| `data.workspace_job_data.inputs.metadata.description` | no | API: The description of the meta data. |
| `data.workspace_job_data.inputs.metadata.group_by` | no | API: The display name of the group this variable belongs to. |
| `data.workspace_job_data.inputs.metadata.hidden` | no | API: If **true**, the variable is not displayed on UI or Command line. |
| `data.workspace_job_data.inputs.metadata.immutable` | no | API: Is the variable readonly ? |
| `data.workspace_job_data.inputs.metadata.matches` | no | API: The regex for the variable value. |
| `data.workspace_job_data.inputs.metadata.max_length` | no | range 0–10000 |
| `data.workspace_job_data.inputs.metadata.max_value` | no | range 0–10000 |
| `data.workspace_job_data.inputs.metadata.min_length` | no | range 0–10000 |
| `data.workspace_job_data.inputs.metadata.min_value` | no | range 0–10000 |
| `data.workspace_job_data.inputs.metadata.options` | no | API: The list of possible values for this variable. If type is **integer** or **date**, then the array of string is converted to array of integers or date during the runtime. |
| `data.workspace_job_data.inputs.metadata.position` | no | range 0–10000 |
| `data.workspace_job_data.inputs.metadata.secure` | no | API: Is the variable secure or sensitive ? |
| `data.workspace_job_data.inputs.metadata.source` | no | API: The source of this meta-data. |
| `data.workspace_job_data.inputs.metadata.type` | no | one of: `boolean`, `string`, `integer`, `date`, `array`, `list`, `map`, `complex`, `link` |
| `data.workspace_job_data.outputs` (block) | no | nested block |
| `data.workspace_job_data.outputs.link` | no | API: The reference link to the variable value By default the expression points to `$self.value`. |
| `data.workspace_job_data.outputs.name` | no | API: The name of the variable. For example, `name = "inventory username"`. |
| `data.workspace_job_data.outputs.value` | no | API: The value for the variable or reference to the value. For example, `value = "&lt;provide your ssh_key_value with \n&gt;"`. |
| `data.workspace_job_data.outputs.metadata` (block) | no | nested block |
| `data.workspace_job_data.outputs.metadata.aliases` | no | API: The list of aliases for the variable name. |
| `data.workspace_job_data.outputs.metadata.default_value` | no | API: Default value for the variable only if the override value is not specified. |
| `data.workspace_job_data.outputs.metadata.description` | no | API: The description of the meta data. |
| `data.workspace_job_data.outputs.metadata.group_by` | no | API: The display name of the group this variable belongs to. |
| `data.workspace_job_data.outputs.metadata.hidden` | no | API: If **true**, the variable is not displayed on UI or Command line. |
| `data.workspace_job_data.outputs.metadata.immutable` | no | API: Is the variable readonly ? |
| `data.workspace_job_data.outputs.metadata.matches` | no | API: The regex for the variable value. |
| `data.workspace_job_data.outputs.metadata.max_length` | no | range 0–10000 |
| `data.workspace_job_data.outputs.metadata.max_value` | no | range 0–10000 |
| `data.workspace_job_data.outputs.metadata.min_length` | no | range 0–10000 |
| `data.workspace_job_data.outputs.metadata.min_value` | no | range 0–10000 |
| `data.workspace_job_data.outputs.metadata.options` | no | API: The list of possible values for this variable. If type is **integer** or **date**, then the array of string is converted to array of integers or date during the runtime. |
| `data.workspace_job_data.outputs.metadata.position` | no | range 0–10000 |
| `data.workspace_job_data.outputs.metadata.secure` | no | API: Is the variable secure or sensitive ? |
| `data.workspace_job_data.outputs.metadata.source` | no | API: The source of this meta-data. |
| `data.workspace_job_data.outputs.metadata.type` | no | one of: `boolean`, `string`, `integer`, `date`, `array`, `list`, `map`, `complex`, `link` |
| `data.workspace_job_data.settings` (block) | no | nested block |
| `data.workspace_job_data.settings.link` | no | API: The reference link to the variable value By default the expression points to `$self.value`. |
| `data.workspace_job_data.settings.name` | no | API: The name of the variable. For example, `name = "inventory username"`. |
| `data.workspace_job_data.settings.value` | no | API: The value for the variable or reference to the value. For example, `value = "&lt;provide your ssh_key_value with \n&gt;"`. |
| `data.workspace_job_data.settings.metadata` (block) | no | nested block |
| `data.workspace_job_data.settings.metadata.aliases` | no | API: The list of aliases for the variable name. |
| `data.workspace_job_data.settings.metadata.default_value` | no | API: Default value for the variable only if the override value is not specified. |
| `data.workspace_job_data.settings.metadata.description` | no | API: The description of the meta data. |
| `data.workspace_job_data.settings.metadata.group_by` | no | API: The display name of the group this variable belongs to. |
| `data.workspace_job_data.settings.metadata.hidden` | no | API: If **true**, the variable is not displayed on UI or Command line. |
| `data.workspace_job_data.settings.metadata.immutable` | no | API: Is the variable readonly ? |
| `data.workspace_job_data.settings.metadata.matches` | no | API: The regex for the variable value. |
| `data.workspace_job_data.settings.metadata.max_length` | no | range 0–10000 |
| `data.workspace_job_data.settings.metadata.max_value` | no | range 0–10000 |
| `data.workspace_job_data.settings.metadata.min_length` | no | range 0–10000 |
| `data.workspace_job_data.settings.metadata.min_value` | no | range 0–10000 |
| `data.workspace_job_data.settings.metadata.options` | no | API: The list of possible values for this variable. If type is **integer** or **date**, then the array of string is converted to array of integers or date during the runtime. |
| `data.workspace_job_data.settings.metadata.position` | no | range 0–10000 |
| `data.workspace_job_data.settings.metadata.secure` | no | API: Is the variable secure or sensitive ? |
| `data.workspace_job_data.settings.metadata.source` | no | API: The source of this meta-data. |
| `data.workspace_job_data.settings.metadata.type` | no | one of: `boolean`, `string`, `integer`, `date`, `array`, `list`, `map`, `complex`, `link` |
| `data.workspace_job_data.template_data` (block) | no | nested block |
| `data.workspace_job_data.template_data.flow_index` | no | range 0–10000 |
| `data.workspace_job_data.template_data.template_id` | no | API: Template Id |
| `data.workspace_job_data.template_data.template_name` | no | API: Template name |
| `data.workspace_job_data.template_data.updated_at` | no | API: Job status updation timestamp |
| `data.workspace_job_data.template_data.inputs` (block) | no | nested block |
| `data.workspace_job_data.template_data.inputs.link` | no | API: The reference link to the variable value By default the expression points to `$self.value`. |
| `data.workspace_job_data.template_data.inputs.name` | no | API: The name of the variable. For example, `name = "inventory username"`. |
| `data.workspace_job_data.template_data.inputs.value` | no | API: The value for the variable or reference to the value. For example, `value = "&lt;provide your ssh_key_value with \n&gt;"`. |
| `data.workspace_job_data.template_data.inputs.metadata` (block) | no | nested block |
| `data.workspace_job_data.template_data.outputs` (block) | no | nested block |
| `data.workspace_job_data.template_data.outputs.link` | no | API: The reference link to the variable value By default the expression points to `$self.value`. |
| `data.workspace_job_data.template_data.outputs.name` | no | API: The name of the variable. For example, `name = "inventory username"`. |
| `data.workspace_job_data.template_data.outputs.value` | no | API: The value for the variable or reference to the value. For example, `value = "&lt;provide your ssh_key_value with \n&gt;"`. |
| `data.workspace_job_data.template_data.outputs.metadata` (block) | no | nested block |
| `data.workspace_job_data.template_data.settings` (block) | no | nested block |
| `data.workspace_job_data.template_data.settings.link` | no | API: The reference link to the variable value By default the expression points to `$self.value`. |
| `data.workspace_job_data.template_data.settings.name` | no | API: The name of the variable. For example, `name = "inventory username"`. |
| `data.workspace_job_data.template_data.settings.value` | no | API: The value for the variable or reference to the value. For example, `value = "&lt;provide your ssh_key_value with \n&gt;"`. |
| `data.workspace_job_data.template_data.settings.metadata` (block) | no | nested block |
| `job_env_settings` (block) | no | nested block |
| `job_env_settings.link` | no | API: The reference link to the variable value By default the expression points to `$self.value`. |
| `job_env_settings.name` | yes | API: The name of the variable. For example, `name = "inventory username"`. |
| `job_env_settings.value` | yes | API: The value for the variable or reference to the value. For example, `value = "&lt;provide your ssh_key_value with \n&gt;"`. |
| `job_env_settings.metadata` (block) | no | nested block |
| `job_env_settings.metadata.aliases` | no | API: The list of aliases for the variable name. |
| `job_env_settings.metadata.default_value` | no | API: Default value for the variable only if the override value is not specified. |
| `job_env_settings.metadata.description` | no | API: The description of the meta data. |
| `job_env_settings.metadata.group_by` | no | API: The display name of the group this variable belongs to. |
| `job_env_settings.metadata.hidden` | no | API: If **true**, the variable is not displayed on UI or Command line. |
| `job_env_settings.metadata.immutable` | no | API: Is the variable readonly ? |
| `job_env_settings.metadata.matches` | no | API: The regex for the variable value. |
| `job_env_settings.metadata.max_length` | no | range 0–10000 |
| `job_env_settings.metadata.max_value` | no | range 0–10000 |
| `job_env_settings.metadata.min_length` | no | range 0–10000 |
| `job_env_settings.metadata.min_value` | no | range 0–10000 |
| `job_env_settings.metadata.options` | no | API: The list of possible values for this variable. If type is **integer** or **date**, then the array of string is converted to array of integers or date during the runtime. |
| `job_env_settings.metadata.position` | no | range 0–10000 |
| `job_env_settings.metadata.secure` | no | API: Is the variable secure or sensitive ? |
| `job_env_settings.metadata.source` | no | API: The source of this meta-data. |
| `job_env_settings.metadata.type` | yes | one of: `boolean`, `string`, `integer`, `date`, `array`, `list`, `map`, `complex`, `link` |
| `job_inputs` (block) | no | nested block |
| `job_inputs.link` | no | API: The reference link to the variable value By default the expression points to `$self.value`. |
| `job_inputs.name` | yes | API: The name of the variable. For example, `name = "inventory username"`. |
| `job_inputs.value` | yes | API: The value for the variable or reference to the value. For example, `value = "&lt;provide your ssh_key_value with \n&gt;"`. |
| `job_inputs.metadata` (block) | no | nested block |
| `job_inputs.metadata.aliases` | no | API: The list of aliases for the variable name. |
| `job_inputs.metadata.default_value` | no | API: Default value for the variable only if the override value is not specified. |
| `job_inputs.metadata.description` | no | API: The description of the meta data. |
| `job_inputs.metadata.group_by` | no | API: The display name of the group this variable belongs to. |
| `job_inputs.metadata.hidden` | no | API: If **true**, the variable is not displayed on UI or Command line. |
| `job_inputs.metadata.immutable` | no | API: Is the variable readonly ? |
| `job_inputs.metadata.matches` | no | API: The regex for the variable value. |
| `job_inputs.metadata.max_length` | no | range 0–10000 |
| `job_inputs.metadata.max_value` | no | range 0–10000 |
| `job_inputs.metadata.min_length` | no | range 0–10000 |
| `job_inputs.metadata.min_value` | no | range 0–10000 |
| `job_inputs.metadata.options` | no | API: The list of possible values for this variable. If type is **integer** or **date**, then the array of string is converted to array of integers or date during the runtime. |
| `job_inputs.metadata.position` | no | range 0–10000 |
| `job_inputs.metadata.secure` | no | API: Is the variable secure or sensitive ? |
| `job_inputs.metadata.source` | no | API: The source of this meta-data. |
| `job_inputs.metadata.type` | yes | one of: `boolean`, `string`, `integer`, `date`, `array`, `list`, `map`, `complex`, `link` |
| `log_summary` (block) | no | nested block |
| `log_summary.elapsed_time` | no | API: Job log elapsed time (log_analyzed_till - log_start_at) |
| `log_summary.job_id` | no | API: Workspace Id |
| `log_summary.job_type` | no | one of: `repo_download_job`, `workspace_job`, `action_job`, `system_job`, `flow_job` |
| `log_summary.log_analyzed_till` | no | API: Job log update timestamp |
| `log_summary.log_start_at` | no | API: Job log start timestamp |
| `log_summary.action_job` (block) | no | nested block |
| `log_summary.action_job.play_count` | no | API: number of plays in playbook |
| `log_summary.action_job.target_count` | no | API: number of targets or hosts |
| `log_summary.action_job.task_count` | no | API: number of tasks in playbook |
| `log_summary.action_job.recap` (block) | no | nested block |
| `log_summary.action_job.recap.changed` | no | API: Number of changed |
| `log_summary.action_job.recap.failed` | no | API: Number of failed |
| `log_summary.action_job.recap.ok` | no | API: Number of OK |
| `log_summary.action_job.recap.skipped` | no | API: Number of skipped |
| `log_summary.action_job.recap.target` | no | API: List of target or host name |
| `log_summary.action_job.recap.unreachable` | no | API: Number of unreachable |
| `log_summary.flow_job` (block) | no | nested block |
| `log_summary.flow_job.workitems_completed` | no | API: Number of workitems completed successfully |
| `log_summary.flow_job.workitems_failed` | no | API: Number of workitems failed |
| `log_summary.flow_job.workitems_pending` | no | API: Number of workitems pending in the flow |
| `log_summary.flow_job.workitems` (block) | no | nested block |
| `log_summary.flow_job.workitems.job_id` | no | API: workspace JOB ID |
| `log_summary.flow_job.workitems.log_url` | no | API: Log url for job |
| `log_summary.flow_job.workitems.resources_add` | no | API: Number of resources add |
| `log_summary.flow_job.workitems.resources_destroy` | no | API: Number of resources destroy |
| `log_summary.flow_job.workitems.resources_modify` | no | API: Number of resources modify |
| `log_summary.flow_job.workitems.workspace_id` | no | API: workspace ID |
| `log_summary.log_errors` (block) | no | nested block |
| `log_summary.log_errors.error_code` | no | API: Error code in the Log |
| `log_summary.log_errors.error_count` | no | API: Number of occurrence |
| `log_summary.log_errors.error_msg` | no | API: Summary error message in the log |
| `log_summary.repo_download_job` (block) | no | nested block |
| `log_summary.repo_download_job.detected_filetype` | no | API: Detected template or data file type |
| `log_summary.repo_download_job.inputs_count` | no | API: Number of inputs detected |
| `log_summary.repo_download_job.outputs_count` | no | API: Number of outputs detected |
| `log_summary.repo_download_job.quarantined_file_count` | no | API: Number of files quarantined |
| `log_summary.repo_download_job.scanned_file_count` | no | API: Number of files scanned |
| `log_summary.system_job` (block) | no | nested block |
| `log_summary.system_job.failed` | no | API: Number of failed |
| `log_summary.system_job.success` | no | API: Number of passed |
| `log_summary.system_job.target_count` | no | API: number of targets or hosts |
| `log_summary.workspace_job` (block) | no | nested block |
| `log_summary.workspace_job.resources_add` | no | API: Number of resources add |
| `log_summary.workspace_job.resources_destroy` | no | API: Number of resources destroy |
| `log_summary.workspace_job.resources_modify` | no | API: Number of resources modify |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_schematics_job_invalid_command_name` | `command_name` |
| `ibm_schematics_job_invalid_command_object` | `command_object` |
| `ibm_schematics_job_invalid_command_object_id` | `command_object_id` |
| `ibm_schematics_job_invalid_command_options` | `command_options` |
| `ibm_schematics_job_invalid_command_parameter` | `command_parameter` |
| `ibm_schematics_job_invalid_location` | `location` |
| `ibm_schematics_job_invalid_tags` | `tags` |
| `ibm_schematics_job_invalid_bastion__host` | `bastion.host` |
| `ibm_schematics_job_invalid_bastion__name` | `bastion.name` |
| `ibm_schematics_job_invalid_data__job_type` | `data.job_type` |
| `ibm_schematics_job_invalid_data__action_job_data__action_name` | `data.action_job_data.action_name` |
| `ibm_schematics_job_invalid_data__action_job_data__materialized_inventory` | `data.action_job_data.materialized_inventory` |
| `ibm_schematics_job_invalid_data__action_job_data__updated_at` | `data.action_job_data.updated_at` |
| `ibm_schematics_job_invalid_data__action_job_data__inputs__link` | `data.action_job_data.inputs.link` |
| `ibm_schematics_job_invalid_data__action_job_data__inputs__name` | `data.action_job_data.inputs.name` |
| `ibm_schematics_job_invalid_data__action_job_data__inputs__value` | `data.action_job_data.inputs.value` |
| `ibm_schematics_job_invalid_data__action_job_data__inputs__metadata__aliases` | `data.action_job_data.inputs.metadata.aliases` |
| `ibm_schematics_job_invalid_data__action_job_data__inputs__metadata__default_value` | `data.action_job_data.inputs.metadata.default_value` |
| `ibm_schematics_job_invalid_data__action_job_data__inputs__metadata__description` | `data.action_job_data.inputs.metadata.description` |
| `ibm_schematics_job_invalid_data__action_job_data__inputs__metadata__group_by` | `data.action_job_data.inputs.metadata.group_by` |
| `ibm_schematics_job_invalid_data__action_job_data__inputs__metadata__hidden` | `data.action_job_data.inputs.metadata.hidden` |
| `ibm_schematics_job_invalid_data__action_job_data__inputs__metadata__immutable` | `data.action_job_data.inputs.metadata.immutable` |
| `ibm_schematics_job_invalid_data__action_job_data__inputs__metadata__matches` | `data.action_job_data.inputs.metadata.matches` |
| `ibm_schematics_job_invalid_data__action_job_data__inputs__metadata__max_length` | `data.action_job_data.inputs.metadata.max_length` |
| `ibm_schematics_job_invalid_data__action_job_data__inputs__metadata__max_value` | `data.action_job_data.inputs.metadata.max_value` |
| `ibm_schematics_job_invalid_data__action_job_data__inputs__metadata__min_length` | `data.action_job_data.inputs.metadata.min_length` |
| `ibm_schematics_job_invalid_data__action_job_data__inputs__metadata__min_value` | `data.action_job_data.inputs.metadata.min_value` |
| `ibm_schematics_job_invalid_data__action_job_data__inputs__metadata__options` | `data.action_job_data.inputs.metadata.options` |
| `ibm_schematics_job_invalid_data__action_job_data__inputs__metadata__position` | `data.action_job_data.inputs.metadata.position` |
| `ibm_schematics_job_invalid_data__action_job_data__inputs__metadata__secure` | `data.action_job_data.inputs.metadata.secure` |
| `ibm_schematics_job_invalid_data__action_job_data__inputs__metadata__source` | `data.action_job_data.inputs.metadata.source` |
| `ibm_schematics_job_invalid_data__action_job_data__inputs__metadata__type` | `data.action_job_data.inputs.metadata.type` |
| `ibm_schematics_job_invalid_data__action_job_data__inventory_record__created_at` | `data.action_job_data.inventory_record.created_at` |
| `ibm_schematics_job_invalid_data__action_job_data__inventory_record__created_by` | `data.action_job_data.inventory_record.created_by` |
| `ibm_schematics_job_invalid_data__action_job_data__inventory_record__description` | `data.action_job_data.inventory_record.description` |
| `ibm_schematics_job_invalid_data__action_job_data__inventory_record__id` | `data.action_job_data.inventory_record.id` |
| `ibm_schematics_job_invalid_data__action_job_data__inventory_record__inventories_ini` | `data.action_job_data.inventory_record.inventories_ini` |
| `ibm_schematics_job_invalid_data__action_job_data__inventory_record__location` | `data.action_job_data.inventory_record.location` |
| `ibm_schematics_job_invalid_data__action_job_data__inventory_record__name` | `data.action_job_data.inventory_record.name` |
| `ibm_schematics_job_invalid_data__action_job_data__inventory_record__resource_group` | `data.action_job_data.inventory_record.resource_group` |
| `ibm_schematics_job_invalid_data__action_job_data__inventory_record__resource_queries` | `data.action_job_data.inventory_record.resource_queries` |
| `ibm_schematics_job_invalid_data__action_job_data__inventory_record__updated_at` | `data.action_job_data.inventory_record.updated_at` |
| `ibm_schematics_job_invalid_data__action_job_data__inventory_record__updated_by` | `data.action_job_data.inventory_record.updated_by` |
| `ibm_schematics_job_invalid_data__action_job_data__outputs__link` | `data.action_job_data.outputs.link` |
| `ibm_schematics_job_invalid_data__action_job_data__outputs__name` | `data.action_job_data.outputs.name` |
| `ibm_schematics_job_invalid_data__action_job_data__outputs__value` | `data.action_job_data.outputs.value` |
| `ibm_schematics_job_invalid_data__action_job_data__outputs__metadata__aliases` | `data.action_job_data.outputs.metadata.aliases` |
| `ibm_schematics_job_invalid_data__action_job_data__outputs__metadata__default_value` | `data.action_job_data.outputs.metadata.default_value` |
| `ibm_schematics_job_invalid_data__action_job_data__outputs__metadata__description` | `data.action_job_data.outputs.metadata.description` |
| `ibm_schematics_job_invalid_data__action_job_data__outputs__metadata__group_by` | `data.action_job_data.outputs.metadata.group_by` |
| `ibm_schematics_job_invalid_data__action_job_data__outputs__metadata__hidden` | `data.action_job_data.outputs.metadata.hidden` |
| `ibm_schematics_job_invalid_data__action_job_data__outputs__metadata__immutable` | `data.action_job_data.outputs.metadata.immutable` |
| `ibm_schematics_job_invalid_data__action_job_data__outputs__metadata__matches` | `data.action_job_data.outputs.metadata.matches` |
| `ibm_schematics_job_invalid_data__action_job_data__outputs__metadata__max_length` | `data.action_job_data.outputs.metadata.max_length` |
| `ibm_schematics_job_invalid_data__action_job_data__outputs__metadata__max_value` | `data.action_job_data.outputs.metadata.max_value` |
| `ibm_schematics_job_invalid_data__action_job_data__outputs__metadata__min_length` | `data.action_job_data.outputs.metadata.min_length` |
| `ibm_schematics_job_invalid_data__action_job_data__outputs__metadata__min_value` | `data.action_job_data.outputs.metadata.min_value` |
| `ibm_schematics_job_invalid_data__action_job_data__outputs__metadata__options` | `data.action_job_data.outputs.metadata.options` |
| `ibm_schematics_job_invalid_data__action_job_data__outputs__metadata__position` | `data.action_job_data.outputs.metadata.position` |
| `ibm_schematics_job_invalid_data__action_job_data__outputs__metadata__secure` | `data.action_job_data.outputs.metadata.secure` |
| `ibm_schematics_job_invalid_data__action_job_data__outputs__metadata__source` | `data.action_job_data.outputs.metadata.source` |
| `ibm_schematics_job_invalid_data__action_job_data__outputs__metadata__type` | `data.action_job_data.outputs.metadata.type` |
| `ibm_schematics_job_invalid_data__action_job_data__settings__link` | `data.action_job_data.settings.link` |
| `ibm_schematics_job_invalid_data__action_job_data__settings__name` | `data.action_job_data.settings.name` |
| `ibm_schematics_job_invalid_data__action_job_data__settings__value` | `data.action_job_data.settings.value` |
| `ibm_schematics_job_invalid_data__action_job_data__settings__metadata__aliases` | `data.action_job_data.settings.metadata.aliases` |
| `ibm_schematics_job_invalid_data__action_job_data__settings__metadata__default_value` | `data.action_job_data.settings.metadata.default_value` |
| `ibm_schematics_job_invalid_data__action_job_data__settings__metadata__description` | `data.action_job_data.settings.metadata.description` |
| `ibm_schematics_job_invalid_data__action_job_data__settings__metadata__group_by` | `data.action_job_data.settings.metadata.group_by` |
| `ibm_schematics_job_invalid_data__action_job_data__settings__metadata__hidden` | `data.action_job_data.settings.metadata.hidden` |
| `ibm_schematics_job_invalid_data__action_job_data__settings__metadata__immutable` | `data.action_job_data.settings.metadata.immutable` |
| `ibm_schematics_job_invalid_data__action_job_data__settings__metadata__matches` | `data.action_job_data.settings.metadata.matches` |
| `ibm_schematics_job_invalid_data__action_job_data__settings__metadata__max_length` | `data.action_job_data.settings.metadata.max_length` |
| `ibm_schematics_job_invalid_data__action_job_data__settings__metadata__max_value` | `data.action_job_data.settings.metadata.max_value` |
| `ibm_schematics_job_invalid_data__action_job_data__settings__metadata__min_length` | `data.action_job_data.settings.metadata.min_length` |
| `ibm_schematics_job_invalid_data__action_job_data__settings__metadata__min_value` | `data.action_job_data.settings.metadata.min_value` |
| `ibm_schematics_job_invalid_data__action_job_data__settings__metadata__options` | `data.action_job_data.settings.metadata.options` |
| `ibm_schematics_job_invalid_data__action_job_data__settings__metadata__position` | `data.action_job_data.settings.metadata.position` |
| `ibm_schematics_job_invalid_data__action_job_data__settings__metadata__secure` | `data.action_job_data.settings.metadata.secure` |
| `ibm_schematics_job_invalid_data__action_job_data__settings__metadata__source` | `data.action_job_data.settings.metadata.source` |
| `ibm_schematics_job_invalid_data__action_job_data__settings__metadata__type` | `data.action_job_data.settings.metadata.type` |
| `ibm_schematics_job_invalid_data__flow_job_data__flow_id` | `data.flow_job_data.flow_id` |
| `ibm_schematics_job_invalid_data__flow_job_data__flow_name` | `data.flow_job_data.flow_name` |
| `ibm_schematics_job_invalid_data__flow_job_data__updated_at` | `data.flow_job_data.updated_at` |
| `ibm_schematics_job_invalid_data__flow_job_data__workitems__command_object_id` | `data.flow_job_data.workitems.command_object_id` |
| `ibm_schematics_job_invalid_data__flow_job_data__workitems__command_object_name` | `data.flow_job_data.workitems.command_object_name` |
| `ibm_schematics_job_invalid_data__flow_job_data__workitems__layers` | `data.flow_job_data.workitems.layers` |
| `ibm_schematics_job_invalid_data__flow_job_data__workitems__source_type` | `data.flow_job_data.workitems.source_type` |
| `ibm_schematics_job_invalid_data__flow_job_data__workitems__updated_at` | `data.flow_job_data.workitems.updated_at` |
| `ibm_schematics_job_invalid_data__flow_job_data__workitems__inputs__link` | `data.flow_job_data.workitems.inputs.link` |
| `ibm_schematics_job_invalid_data__flow_job_data__workitems__inputs__name` | `data.flow_job_data.workitems.inputs.name` |
| `ibm_schematics_job_invalid_data__flow_job_data__workitems__inputs__value` | `data.flow_job_data.workitems.inputs.value` |
| `ibm_schematics_job_invalid_data__flow_job_data__workitems__last_job__command_name` | `data.flow_job_data.workitems.last_job.command_name` |
| `ibm_schematics_job_invalid_data__flow_job_data__workitems__last_job__command_object` | `data.flow_job_data.workitems.last_job.command_object` |
| `ibm_schematics_job_invalid_data__flow_job_data__workitems__last_job__command_object_id` | `data.flow_job_data.workitems.last_job.command_object_id` |
| `ibm_schematics_job_invalid_data__flow_job_data__workitems__last_job__command_object_name` | `data.flow_job_data.workitems.last_job.command_object_name` |
| `ibm_schematics_job_invalid_data__flow_job_data__workitems__last_job__job_id` | `data.flow_job_data.workitems.last_job.job_id` |
| `ibm_schematics_job_invalid_data__flow_job_data__workitems__last_job__job_status` | `data.flow_job_data.workitems.last_job.job_status` |
| `ibm_schematics_job_invalid_data__flow_job_data__workitems__outputs__link` | `data.flow_job_data.workitems.outputs.link` |
| `ibm_schematics_job_invalid_data__flow_job_data__workitems__outputs__name` | `data.flow_job_data.workitems.outputs.name` |
| `ibm_schematics_job_invalid_data__flow_job_data__workitems__outputs__value` | `data.flow_job_data.workitems.outputs.value` |
| `ibm_schematics_job_invalid_data__flow_job_data__workitems__settings__link` | `data.flow_job_data.workitems.settings.link` |
| `ibm_schematics_job_invalid_data__flow_job_data__workitems__settings__name` | `data.flow_job_data.workitems.settings.name` |
| `ibm_schematics_job_invalid_data__flow_job_data__workitems__settings__value` | `data.flow_job_data.workitems.settings.value` |
| `ibm_schematics_job_invalid_data__flow_job_data__workitems__source__source_type` | `data.flow_job_data.workitems.source.source_type` |
| `ibm_schematics_job_invalid_data__system_job_data__key_id` | `data.system_job_data.key_id` |
| `ibm_schematics_job_invalid_data__system_job_data__schematics_resource_id` | `data.system_job_data.schematics_resource_id` |
| `ibm_schematics_job_invalid_data__system_job_data__updated_at` | `data.system_job_data.updated_at` |
| `ibm_schematics_job_invalid_data__workspace_job_data__flow_id` | `data.workspace_job_data.flow_id` |
| `ibm_schematics_job_invalid_data__workspace_job_data__flow_name` | `data.workspace_job_data.flow_name` |
| `ibm_schematics_job_invalid_data__workspace_job_data__updated_at` | `data.workspace_job_data.updated_at` |
| `ibm_schematics_job_invalid_data__workspace_job_data__workspace_name` | `data.workspace_job_data.workspace_name` |
| `ibm_schematics_job_invalid_data__workspace_job_data__inputs__link` | `data.workspace_job_data.inputs.link` |
| `ibm_schematics_job_invalid_data__workspace_job_data__inputs__name` | `data.workspace_job_data.inputs.name` |
| `ibm_schematics_job_invalid_data__workspace_job_data__inputs__value` | `data.workspace_job_data.inputs.value` |
| `ibm_schematics_job_invalid_data__workspace_job_data__inputs__metadata__aliases` | `data.workspace_job_data.inputs.metadata.aliases` |
| `ibm_schematics_job_invalid_data__workspace_job_data__inputs__metadata__default_value` | `data.workspace_job_data.inputs.metadata.default_value` |
| `ibm_schematics_job_invalid_data__workspace_job_data__inputs__metadata__description` | `data.workspace_job_data.inputs.metadata.description` |
| `ibm_schematics_job_invalid_data__workspace_job_data__inputs__metadata__group_by` | `data.workspace_job_data.inputs.metadata.group_by` |
| `ibm_schematics_job_invalid_data__workspace_job_data__inputs__metadata__hidden` | `data.workspace_job_data.inputs.metadata.hidden` |
| `ibm_schematics_job_invalid_data__workspace_job_data__inputs__metadata__immutable` | `data.workspace_job_data.inputs.metadata.immutable` |
| `ibm_schematics_job_invalid_data__workspace_job_data__inputs__metadata__matches` | `data.workspace_job_data.inputs.metadata.matches` |
| `ibm_schematics_job_invalid_data__workspace_job_data__inputs__metadata__max_length` | `data.workspace_job_data.inputs.metadata.max_length` |
| `ibm_schematics_job_invalid_data__workspace_job_data__inputs__metadata__max_value` | `data.workspace_job_data.inputs.metadata.max_value` |
| `ibm_schematics_job_invalid_data__workspace_job_data__inputs__metadata__min_length` | `data.workspace_job_data.inputs.metadata.min_length` |
| `ibm_schematics_job_invalid_data__workspace_job_data__inputs__metadata__min_value` | `data.workspace_job_data.inputs.metadata.min_value` |
| `ibm_schematics_job_invalid_data__workspace_job_data__inputs__metadata__options` | `data.workspace_job_data.inputs.metadata.options` |
| `ibm_schematics_job_invalid_data__workspace_job_data__inputs__metadata__position` | `data.workspace_job_data.inputs.metadata.position` |
| `ibm_schematics_job_invalid_data__workspace_job_data__inputs__metadata__secure` | `data.workspace_job_data.inputs.metadata.secure` |
| `ibm_schematics_job_invalid_data__workspace_job_data__inputs__metadata__source` | `data.workspace_job_data.inputs.metadata.source` |
| `ibm_schematics_job_invalid_data__workspace_job_data__inputs__metadata__type` | `data.workspace_job_data.inputs.metadata.type` |
| `ibm_schematics_job_invalid_data__workspace_job_data__outputs__link` | `data.workspace_job_data.outputs.link` |
| `ibm_schematics_job_invalid_data__workspace_job_data__outputs__name` | `data.workspace_job_data.outputs.name` |
| `ibm_schematics_job_invalid_data__workspace_job_data__outputs__value` | `data.workspace_job_data.outputs.value` |
| `ibm_schematics_job_invalid_data__workspace_job_data__outputs__metadata__aliases` | `data.workspace_job_data.outputs.metadata.aliases` |
| `ibm_schematics_job_invalid_data__workspace_job_data__outputs__metadata__default_value` | `data.workspace_job_data.outputs.metadata.default_value` |
| `ibm_schematics_job_invalid_data__workspace_job_data__outputs__metadata__description` | `data.workspace_job_data.outputs.metadata.description` |
| `ibm_schematics_job_invalid_data__workspace_job_data__outputs__metadata__group_by` | `data.workspace_job_data.outputs.metadata.group_by` |
| `ibm_schematics_job_invalid_data__workspace_job_data__outputs__metadata__hidden` | `data.workspace_job_data.outputs.metadata.hidden` |
| `ibm_schematics_job_invalid_data__workspace_job_data__outputs__metadata__immutable` | `data.workspace_job_data.outputs.metadata.immutable` |
| `ibm_schematics_job_invalid_data__workspace_job_data__outputs__metadata__matches` | `data.workspace_job_data.outputs.metadata.matches` |
| `ibm_schematics_job_invalid_data__workspace_job_data__outputs__metadata__max_length` | `data.workspace_job_data.outputs.metadata.max_length` |
| `ibm_schematics_job_invalid_data__workspace_job_data__outputs__metadata__max_value` | `data.workspace_job_data.outputs.metadata.max_value` |
| `ibm_schematics_job_invalid_data__workspace_job_data__outputs__metadata__min_length` | `data.workspace_job_data.outputs.metadata.min_length` |
| `ibm_schematics_job_invalid_data__workspace_job_data__outputs__metadata__min_value` | `data.workspace_job_data.outputs.metadata.min_value` |
| `ibm_schematics_job_invalid_data__workspace_job_data__outputs__metadata__options` | `data.workspace_job_data.outputs.metadata.options` |
| `ibm_schematics_job_invalid_data__workspace_job_data__outputs__metadata__position` | `data.workspace_job_data.outputs.metadata.position` |
| `ibm_schematics_job_invalid_data__workspace_job_data__outputs__metadata__secure` | `data.workspace_job_data.outputs.metadata.secure` |
| `ibm_schematics_job_invalid_data__workspace_job_data__outputs__metadata__source` | `data.workspace_job_data.outputs.metadata.source` |
| `ibm_schematics_job_invalid_data__workspace_job_data__outputs__metadata__type` | `data.workspace_job_data.outputs.metadata.type` |
| `ibm_schematics_job_invalid_data__workspace_job_data__settings__link` | `data.workspace_job_data.settings.link` |
| `ibm_schematics_job_invalid_data__workspace_job_data__settings__name` | `data.workspace_job_data.settings.name` |
| `ibm_schematics_job_invalid_data__workspace_job_data__settings__value` | `data.workspace_job_data.settings.value` |
| `ibm_schematics_job_invalid_data__workspace_job_data__settings__metadata__aliases` | `data.workspace_job_data.settings.metadata.aliases` |
| `ibm_schematics_job_invalid_data__workspace_job_data__settings__metadata__default_value` | `data.workspace_job_data.settings.metadata.default_value` |
| `ibm_schematics_job_invalid_data__workspace_job_data__settings__metadata__description` | `data.workspace_job_data.settings.metadata.description` |
| `ibm_schematics_job_invalid_data__workspace_job_data__settings__metadata__group_by` | `data.workspace_job_data.settings.metadata.group_by` |
| `ibm_schematics_job_invalid_data__workspace_job_data__settings__metadata__hidden` | `data.workspace_job_data.settings.metadata.hidden` |
| `ibm_schematics_job_invalid_data__workspace_job_data__settings__metadata__immutable` | `data.workspace_job_data.settings.metadata.immutable` |
| `ibm_schematics_job_invalid_data__workspace_job_data__settings__metadata__matches` | `data.workspace_job_data.settings.metadata.matches` |
| `ibm_schematics_job_invalid_data__workspace_job_data__settings__metadata__max_length` | `data.workspace_job_data.settings.metadata.max_length` |
| `ibm_schematics_job_invalid_data__workspace_job_data__settings__metadata__max_value` | `data.workspace_job_data.settings.metadata.max_value` |
| `ibm_schematics_job_invalid_data__workspace_job_data__settings__metadata__min_length` | `data.workspace_job_data.settings.metadata.min_length` |
| `ibm_schematics_job_invalid_data__workspace_job_data__settings__metadata__min_value` | `data.workspace_job_data.settings.metadata.min_value` |
| `ibm_schematics_job_invalid_data__workspace_job_data__settings__metadata__options` | `data.workspace_job_data.settings.metadata.options` |
| `ibm_schematics_job_invalid_data__workspace_job_data__settings__metadata__position` | `data.workspace_job_data.settings.metadata.position` |
| `ibm_schematics_job_invalid_data__workspace_job_data__settings__metadata__secure` | `data.workspace_job_data.settings.metadata.secure` |
| `ibm_schematics_job_invalid_data__workspace_job_data__settings__metadata__source` | `data.workspace_job_data.settings.metadata.source` |
| `ibm_schematics_job_invalid_data__workspace_job_data__settings__metadata__type` | `data.workspace_job_data.settings.metadata.type` |
| `ibm_schematics_job_invalid_data__workspace_job_data__template_data__flow_index` | `data.workspace_job_data.template_data.flow_index` |
| `ibm_schematics_job_invalid_data__workspace_job_data__template_data__template_id` | `data.workspace_job_data.template_data.template_id` |
| `ibm_schematics_job_invalid_data__workspace_job_data__template_data__template_name` | `data.workspace_job_data.template_data.template_name` |
| `ibm_schematics_job_invalid_data__workspace_job_data__template_data__updated_at` | `data.workspace_job_data.template_data.updated_at` |
| `ibm_schematics_job_invalid_data__workspace_job_data__template_data__inputs__link` | `data.workspace_job_data.template_data.inputs.link` |
| `ibm_schematics_job_invalid_data__workspace_job_data__template_data__inputs__name` | `data.workspace_job_data.template_data.inputs.name` |
| `ibm_schematics_job_invalid_data__workspace_job_data__template_data__inputs__value` | `data.workspace_job_data.template_data.inputs.value` |
| `ibm_schematics_job_invalid_data__workspace_job_data__template_data__outputs__link` | `data.workspace_job_data.template_data.outputs.link` |
| `ibm_schematics_job_invalid_data__workspace_job_data__template_data__outputs__name` | `data.workspace_job_data.template_data.outputs.name` |
| `ibm_schematics_job_invalid_data__workspace_job_data__template_data__outputs__value` | `data.workspace_job_data.template_data.outputs.value` |
| `ibm_schematics_job_invalid_data__workspace_job_data__template_data__settings__link` | `data.workspace_job_data.template_data.settings.link` |
| `ibm_schematics_job_invalid_data__workspace_job_data__template_data__settings__name` | `data.workspace_job_data.template_data.settings.name` |
| `ibm_schematics_job_invalid_data__workspace_job_data__template_data__settings__value` | `data.workspace_job_data.template_data.settings.value` |
| `ibm_schematics_job_invalid_job_env_settings__link` | `job_env_settings.link` |
| `ibm_schematics_job_invalid_job_env_settings__name` | `job_env_settings.name` |
| `ibm_schematics_job_invalid_job_env_settings__value` | `job_env_settings.value` |
| `ibm_schematics_job_invalid_job_env_settings__metadata__aliases` | `job_env_settings.metadata.aliases` |
| `ibm_schematics_job_invalid_job_env_settings__metadata__default_value` | `job_env_settings.metadata.default_value` |
| `ibm_schematics_job_invalid_job_env_settings__metadata__description` | `job_env_settings.metadata.description` |
| `ibm_schematics_job_invalid_job_env_settings__metadata__group_by` | `job_env_settings.metadata.group_by` |
| `ibm_schematics_job_invalid_job_env_settings__metadata__hidden` | `job_env_settings.metadata.hidden` |
| `ibm_schematics_job_invalid_job_env_settings__metadata__immutable` | `job_env_settings.metadata.immutable` |
| `ibm_schematics_job_invalid_job_env_settings__metadata__matches` | `job_env_settings.metadata.matches` |
| `ibm_schematics_job_invalid_job_env_settings__metadata__max_length` | `job_env_settings.metadata.max_length` |
| `ibm_schematics_job_invalid_job_env_settings__metadata__max_value` | `job_env_settings.metadata.max_value` |
| `ibm_schematics_job_invalid_job_env_settings__metadata__min_length` | `job_env_settings.metadata.min_length` |
| `ibm_schematics_job_invalid_job_env_settings__metadata__min_value` | `job_env_settings.metadata.min_value` |
| `ibm_schematics_job_invalid_job_env_settings__metadata__options` | `job_env_settings.metadata.options` |
| `ibm_schematics_job_invalid_job_env_settings__metadata__position` | `job_env_settings.metadata.position` |
| `ibm_schematics_job_invalid_job_env_settings__metadata__secure` | `job_env_settings.metadata.secure` |
| `ibm_schematics_job_invalid_job_env_settings__metadata__source` | `job_env_settings.metadata.source` |
| `ibm_schematics_job_invalid_job_env_settings__metadata__type` | `job_env_settings.metadata.type` |
| `ibm_schematics_job_invalid_job_inputs__link` | `job_inputs.link` |
| `ibm_schematics_job_invalid_job_inputs__name` | `job_inputs.name` |
| `ibm_schematics_job_invalid_job_inputs__value` | `job_inputs.value` |
| `ibm_schematics_job_invalid_job_inputs__metadata__aliases` | `job_inputs.metadata.aliases` |
| `ibm_schematics_job_invalid_job_inputs__metadata__default_value` | `job_inputs.metadata.default_value` |
| `ibm_schematics_job_invalid_job_inputs__metadata__description` | `job_inputs.metadata.description` |
| `ibm_schematics_job_invalid_job_inputs__metadata__group_by` | `job_inputs.metadata.group_by` |
| `ibm_schematics_job_invalid_job_inputs__metadata__hidden` | `job_inputs.metadata.hidden` |
| `ibm_schematics_job_invalid_job_inputs__metadata__immutable` | `job_inputs.metadata.immutable` |
| `ibm_schematics_job_invalid_job_inputs__metadata__matches` | `job_inputs.metadata.matches` |
| `ibm_schematics_job_invalid_job_inputs__metadata__max_length` | `job_inputs.metadata.max_length` |
| `ibm_schematics_job_invalid_job_inputs__metadata__max_value` | `job_inputs.metadata.max_value` |
| `ibm_schematics_job_invalid_job_inputs__metadata__min_length` | `job_inputs.metadata.min_length` |
| `ibm_schematics_job_invalid_job_inputs__metadata__min_value` | `job_inputs.metadata.min_value` |
| `ibm_schematics_job_invalid_job_inputs__metadata__options` | `job_inputs.metadata.options` |
| `ibm_schematics_job_invalid_job_inputs__metadata__position` | `job_inputs.metadata.position` |
| `ibm_schematics_job_invalid_job_inputs__metadata__secure` | `job_inputs.metadata.secure` |
| `ibm_schematics_job_invalid_job_inputs__metadata__source` | `job_inputs.metadata.source` |
| `ibm_schematics_job_invalid_job_inputs__metadata__type` | `job_inputs.metadata.type` |
| `ibm_schematics_job_invalid_log_summary__elapsed_time` | `log_summary.elapsed_time` |
| `ibm_schematics_job_invalid_log_summary__job_id` | `log_summary.job_id` |
| `ibm_schematics_job_invalid_log_summary__job_type` | `log_summary.job_type` |
| `ibm_schematics_job_invalid_log_summary__log_analyzed_till` | `log_summary.log_analyzed_till` |
| `ibm_schematics_job_invalid_log_summary__log_start_at` | `log_summary.log_start_at` |
| `ibm_schematics_job_invalid_log_summary__action_job__play_count` | `log_summary.action_job.play_count` |
| `ibm_schematics_job_invalid_log_summary__action_job__target_count` | `log_summary.action_job.target_count` |
| `ibm_schematics_job_invalid_log_summary__action_job__task_count` | `log_summary.action_job.task_count` |
| `ibm_schematics_job_invalid_log_summary__action_job__recap__changed` | `log_summary.action_job.recap.changed` |
| `ibm_schematics_job_invalid_log_summary__action_job__recap__failed` | `log_summary.action_job.recap.failed` |
| `ibm_schematics_job_invalid_log_summary__action_job__recap__ok` | `log_summary.action_job.recap.ok` |
| `ibm_schematics_job_invalid_log_summary__action_job__recap__skipped` | `log_summary.action_job.recap.skipped` |
| `ibm_schematics_job_invalid_log_summary__action_job__recap__target` | `log_summary.action_job.recap.target` |
| `ibm_schematics_job_invalid_log_summary__action_job__recap__unreachable` | `log_summary.action_job.recap.unreachable` |
| `ibm_schematics_job_invalid_log_summary__flow_job__workitems_completed` | `log_summary.flow_job.workitems_completed` |
| `ibm_schematics_job_invalid_log_summary__flow_job__workitems_failed` | `log_summary.flow_job.workitems_failed` |
| `ibm_schematics_job_invalid_log_summary__flow_job__workitems_pending` | `log_summary.flow_job.workitems_pending` |
| `ibm_schematics_job_invalid_log_summary__flow_job__workitems__job_id` | `log_summary.flow_job.workitems.job_id` |
| `ibm_schematics_job_invalid_log_summary__flow_job__workitems__log_url` | `log_summary.flow_job.workitems.log_url` |
| `ibm_schematics_job_invalid_log_summary__flow_job__workitems__resources_add` | `log_summary.flow_job.workitems.resources_add` |
| `ibm_schematics_job_invalid_log_summary__flow_job__workitems__resources_destroy` | `log_summary.flow_job.workitems.resources_destroy` |
| `ibm_schematics_job_invalid_log_summary__flow_job__workitems__resources_modify` | `log_summary.flow_job.workitems.resources_modify` |
| `ibm_schematics_job_invalid_log_summary__flow_job__workitems__workspace_id` | `log_summary.flow_job.workitems.workspace_id` |
| `ibm_schematics_job_invalid_log_summary__log_errors__error_code` | `log_summary.log_errors.error_code` |
| `ibm_schematics_job_invalid_log_summary__log_errors__error_count` | `log_summary.log_errors.error_count` |
| `ibm_schematics_job_invalid_log_summary__log_errors__error_msg` | `log_summary.log_errors.error_msg` |
| `ibm_schematics_job_invalid_log_summary__repo_download_job__detected_filetype` | `log_summary.repo_download_job.detected_filetype` |
| `ibm_schematics_job_invalid_log_summary__repo_download_job__inputs_count` | `log_summary.repo_download_job.inputs_count` |
| `ibm_schematics_job_invalid_log_summary__repo_download_job__outputs_count` | `log_summary.repo_download_job.outputs_count` |
| `ibm_schematics_job_invalid_log_summary__repo_download_job__quarantined_file_count` | `log_summary.repo_download_job.quarantined_file_count` |
| `ibm_schematics_job_invalid_log_summary__repo_download_job__scanned_file_count` | `log_summary.repo_download_job.scanned_file_count` |
| `ibm_schematics_job_invalid_log_summary__system_job__failed` | `log_summary.system_job.failed` |
| `ibm_schematics_job_invalid_log_summary__system_job__success` | `log_summary.system_job.success` |
| `ibm_schematics_job_invalid_log_summary__system_job__target_count` | `log_summary.system_job.target_count` |
| `ibm_schematics_job_invalid_log_summary__workspace_job__resources_add` | `log_summary.workspace_job.resources_add` |
| `ibm_schematics_job_invalid_log_summary__workspace_job__resources_destroy` | `log_summary.workspace_job.resources_destroy` |
| `ibm_schematics_job_invalid_log_summary__workspace_job__resources_modify` | `log_summary.workspace_job.resources_modify` |

## Example

### Valid

```hcl
resource "ibm_schematics_job" "test" {
  command_name = "ansible_playbook_check"
  command_object = "action"
  command_object_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
