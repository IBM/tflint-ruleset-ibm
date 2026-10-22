# ibm_schematics_workspace

Manages ibm_schematics_workspace. The name of your workspace.

This rule validates `ibm_schematics_workspace` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `WorkspaceCreateRequest` (`schematics-schematics`).

## Configuration

```hcl
rule "ibm_schematics_workspace" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `applied_shareddata_ids` | no | API: List of applied shared dataset ID. |
| `created_at` | no | — |
| `created_by` | no | — |
| `crn` | no | — |
| `description` | no | API: The description of the workspace. |
| `frozen` | no | API: If set to true, the workspace is frozen and changes to the workspace are disabled. |
| `frozen_at` | no | API: The timestamp when the workspace was frozen. |
| `frozen_by` | no | API: The user ID that froze the workspace. |
| `last_health_check_at` | no | — |
| `location` | no | API: The location where you want to create your Schematics workspace and run the Schematics jobs. The location that you enter must match the API endpoint that you use. For example, if you use the Frankfurt API endpoint, you must specify `eu-de` as your location. If you use an API endpoint for a geography and you do not specify a location, Schematics determines the location based on availability. |
| `locked` | no | API: If set to true, the workspace is locked and disabled for changes. |
| `locked_by` | no | API: The user ID that initiated a resource-related job, such as applying or destroying resources, that locked the workspace. |
| `locked_time` | no | API: The timestamp when the workspace was locked. |
| `name` | yes | pattern `^[a-zA-Z0-9][a-zA-Z0-9-_ ]*$` |
| `resource_group` | no | API: The ID of the resource group where you want to provision the workspace. |
| `runtime_data` | no | — |
| `status` | no | — |
| `status_code` | no | — |
| `status_msg` | no | — |
| `tags` | no | API: A list of tags that are associated with the workspace. |
| `template_env_settings` | no | — |
| `template_git_branch` | no | — |
| `template_git_folder` | no | — |
| `template_git_has_uploadedgitrepotar` | no | — |
| `template_git_release` | no | — |
| `template_git_repo_sha_value` | no | — |
| `template_git_repo_url` | no | — |
| `template_git_url` | no | — |
| `template_init_state_file` | no | — |
| `template_ref` | no | API: Workspace template ref |
| `template_type` | yes | pattern `^terraform_v(?:1\.4|1\.5|1\.6|1\.7|1\.8|1\.9|1\.10|1\.11|1\.12|1\.13)(?:\.\d+)?$` |
| `template_uninstall_script_name` | no | — |
| `template_values` | no | — |
| `updated_at` | no | — |
| `updated_by` | no | — |
| `x_github_token` | no | — |
| `catalog_ref` (block) | no | nested block |
| `catalog_ref.dry_run` | no | API: Dry run |
| `catalog_ref.item_icon_url` | no | API: The URL to the icon of the software template in the IBM Cloud catalog. |
| `catalog_ref.item_id` | no | API: The ID of the software template that you chose to install from the IBM Cloud catalog. This software is provisioned with Schematics. |
| `catalog_ref.item_name` | no | API: The name of the software that you chose to install from the IBM Cloud catalog. |
| `catalog_ref.item_readme_url` | no | API: The URL to the readme file of the software template in the IBM Cloud catalog. |
| `catalog_ref.item_url` | no | API: The URL to the software template in the IBM Cloud catalog. |
| `catalog_ref.launch_url` | no | API: The URL to the dashboard to access your software. |
| `catalog_ref.offering_version` | no | API: The version of the software template that you chose to install from the IBM Cloud catalog. |
| `catalog_ref.owning_account` | no | API: Owning account ID of the catalog. |
| `shared_data` (block) | no | nested block |
| `shared_data.cluster_created_on` | no | API: Cluster created on. |
| `shared_data.cluster_id` | no | API: The ID of the cluster where you want to provision the resources of all IBM Cloud catalog templates that are included in the catalog offering. |
| `shared_data.cluster_name` | no | API: The cluster name. |
| `shared_data.cluster_type` | no | API: The cluster type. |
| `shared_data.entitlement_keys` | no | API: The entitlement key that you want to use to install IBM Cloud entitled software. |
| `shared_data.namespace` | no | API: The Kubernetes namespace or OpenShift project where the resources of all IBM Cloud catalog templates that are included in the catalog offering are deployed into. |
| `shared_data.region` | no | API: The IBM Cloud region that you want to use for the resources of all IBM Cloud catalog templates that are included in the catalog offering. |
| `shared_data.resource_group_id` | no | API: The ID of the resource group that you want to use for the resources of all IBM Cloud catalog templates that are included in the catalog offering. |
| `shared_data.worker_count` | no | range 0–5000 |
| `shared_data.worker_machine_type` | no | API: The cluster worker type. |
| `template_inputs` (block) | no | nested block |
| `template_inputs.description` | no | API: The description of the workspace. |
| `template_inputs.name` | yes | API: Name of the Service Data |
| `template_inputs.secure` | no | API: Environment variable is secure. |
| `template_inputs.type` | yes | API: Type of the value string, int, bool |
| `template_inputs.use_default` | no | API: Variable uses default value; and is not over-ridden. |
| `template_inputs.value` | yes | API: Values of service data |
| `template_values_metadata` (block) | no | nested block |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_schematics_workspace_invalid_applied_shareddata_ids` | `applied_shareddata_ids` |
| `ibm_schematics_workspace_invalid_description` | `description` |
| `ibm_schematics_workspace_invalid_frozen` | `frozen` |
| `ibm_schematics_workspace_invalid_frozen_at` | `frozen_at` |
| `ibm_schematics_workspace_invalid_frozen_by` | `frozen_by` |
| `ibm_schematics_workspace_invalid_location` | `location` |
| `ibm_schematics_workspace_invalid_locked` | `locked` |
| `ibm_schematics_workspace_invalid_locked_by` | `locked_by` |
| `ibm_schematics_workspace_invalid_locked_time` | `locked_time` |
| `ibm_schematics_workspace_invalid_name` | `name` |
| `ibm_schematics_workspace_invalid_resource_group` | `resource_group` |
| `ibm_schematics_workspace_invalid_tags` | `tags` |
| `ibm_schematics_workspace_invalid_template_env_settings` | `template_env_settings` |
| `ibm_schematics_workspace_invalid_template_git_branch` | `template_git_branch` |
| `ibm_schematics_workspace_invalid_template_git_folder` | `template_git_folder` |
| `ibm_schematics_workspace_invalid_template_git_has_uploadedgitrepotar` | `template_git_has_uploadedgitrepotar` |
| `ibm_schematics_workspace_invalid_template_git_release` | `template_git_release` |
| `ibm_schematics_workspace_invalid_template_git_repo_sha_value` | `template_git_repo_sha_value` |
| `ibm_schematics_workspace_invalid_template_git_repo_url` | `template_git_repo_url` |
| `ibm_schematics_workspace_invalid_template_git_url` | `template_git_url` |
| `ibm_schematics_workspace_invalid_template_init_state_file` | `template_init_state_file` |
| `ibm_schematics_workspace_invalid_template_ref` | `template_ref` |
| `ibm_schematics_workspace_invalid_template_type` | `template_type` |
| `ibm_schematics_workspace_invalid_template_uninstall_script_name` | `template_uninstall_script_name` |
| `ibm_schematics_workspace_invalid_template_values` | `template_values` |
| `ibm_schematics_workspace_invalid_x_github_token` | `x_github_token` |
| `ibm_schematics_workspace_invalid_catalog_ref__dry_run` | `catalog_ref.dry_run` |
| `ibm_schematics_workspace_invalid_catalog_ref__item_icon_url` | `catalog_ref.item_icon_url` |
| `ibm_schematics_workspace_invalid_catalog_ref__item_id` | `catalog_ref.item_id` |
| `ibm_schematics_workspace_invalid_catalog_ref__item_name` | `catalog_ref.item_name` |
| `ibm_schematics_workspace_invalid_catalog_ref__item_readme_url` | `catalog_ref.item_readme_url` |
| `ibm_schematics_workspace_invalid_catalog_ref__item_url` | `catalog_ref.item_url` |
| `ibm_schematics_workspace_invalid_catalog_ref__launch_url` | `catalog_ref.launch_url` |
| `ibm_schematics_workspace_invalid_catalog_ref__offering_version` | `catalog_ref.offering_version` |
| `ibm_schematics_workspace_invalid_catalog_ref__owning_account` | `catalog_ref.owning_account` |
| `ibm_schematics_workspace_invalid_shared_data__cluster_created_on` | `shared_data.cluster_created_on` |
| `ibm_schematics_workspace_invalid_shared_data__cluster_id` | `shared_data.cluster_id` |
| `ibm_schematics_workspace_invalid_shared_data__cluster_name` | `shared_data.cluster_name` |
| `ibm_schematics_workspace_invalid_shared_data__cluster_type` | `shared_data.cluster_type` |
| `ibm_schematics_workspace_invalid_shared_data__entitlement_keys` | `shared_data.entitlement_keys` |
| `ibm_schematics_workspace_invalid_shared_data__namespace` | `shared_data.namespace` |
| `ibm_schematics_workspace_invalid_shared_data__region` | `shared_data.region` |
| `ibm_schematics_workspace_invalid_shared_data__resource_group_id` | `shared_data.resource_group_id` |
| `ibm_schematics_workspace_invalid_shared_data__worker_count` | `shared_data.worker_count` |
| `ibm_schematics_workspace_invalid_shared_data__worker_machine_type` | `shared_data.worker_machine_type` |
| `ibm_schematics_workspace_invalid_template_inputs__description` | `template_inputs.description` |
| `ibm_schematics_workspace_invalid_template_inputs__name` | `template_inputs.name` |
| `ibm_schematics_workspace_invalid_template_inputs__secure` | `template_inputs.secure` |
| `ibm_schematics_workspace_invalid_template_inputs__type` | `template_inputs.type` |
| `ibm_schematics_workspace_invalid_template_inputs__use_default` | `template_inputs.use_default` |
| `ibm_schematics_workspace_invalid_template_inputs__value` | `template_inputs.value` |

## Example

### Valid

```hcl
resource "ibm_schematics_workspace" "test" {
  name = "example-name"
  template_type = "terraform_v1.5"
}
```
