# ibm_schematics_workspace_data

Reads ibm_schematics_workspace. The name of the workspace.

This rule validates `ibm_schematics_workspace` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `WorkspaceActivities` (`schematics-schematics`).

## Configuration

```hcl
rule "ibm_schematics_workspace_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `applied_shareddata_ids` | no | — |
| `catalog_ref` | no | — |
| `created_at` | no | — |
| `created_by` | no | — |
| `crn` | no | — |
| `description` | no | — |
| `frozen` | no | — |
| `frozen_at` | no | — |
| `frozen_by` | no | — |
| `is_frozen` | no | — |
| `is_locked` | no | — |
| `last_health_check_at` | no | — |
| `location` | no | — |
| `locked` | no | — |
| `locked_by` | no | — |
| `locked_time` | no | — |
| `name` | no | — |
| `resource_controller_url` | no | — |
| `resource_group` | no | — |
| `runtime_data` | no | — |
| `shared_data` | no | — |
| `status` | no | — |
| `status_code` | no | — |
| `status_msg` | no | — |
| `tags` | no | — |
| `template_env_settings` | no | — |
| `template_git_branch` | no | — |
| `template_git_folder` | no | — |
| `template_git_full_url` | no | — |
| `template_git_has_uploadedgitrepotar` | no | — |
| `template_git_release` | no | — |
| `template_git_repo_sha_value` | no | — |
| `template_git_repo_url` | no | — |
| `template_git_url` | no | — |
| `template_init_state_file` | no | — |
| `template_inputs` | no | — |
| `template_ref` | no | — |
| `template_type` | no | — |
| `template_uninstall_script_name` | no | — |
| `template_values` | no | — |
| `updated_at` | no | — |
| `updated_by` | no | — |
| `workspace_id` | yes | API: The ID of the workspace. |
| `template_values_metadata` (block) | no | nested block |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_schematics_workspace_data_invalid_location` | `location` |
| `ibm_schematics_workspace_data_invalid_template_git_has_uploadedgitrepotar` | `template_git_has_uploadedgitrepotar` |
| `ibm_schematics_workspace_data_invalid_workspace_id` | `workspace_id` |

## Example

### Valid

```hcl
data "ibm_schematics_workspace" "test" {
  workspace_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
