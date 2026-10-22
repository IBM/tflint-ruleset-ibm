# ibm_schematics_action_data

Reads ibm_schematics_action. The unique name of your action.

This rule validates `ibm_schematics_action` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `params:put:/v2/actions/{action_id}/template_repo_upload` (`schematics-schematics`).

## Configuration

```hcl
rule "ibm_schematics_action_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account` | no | — |
| `action_id` | yes | — |
| `action_inputs` | no | — |
| `action_outputs` | no | — |
| `bastion` | no | — |
| `bastion_credential` | no | — |
| `command_parameter` | no | — |
| `created_at` | no | — |
| `created_by` | no | — |
| `credentials` | no | — |
| `crn` | no | — |
| `description` | no | — |
| `id` | no | — |
| `inventory` | no | — |
| `location` | no | — |
| `name` | no | — |
| `playbook_names` | no | — |
| `resource_group` | no | — |
| `settings` | no | — |
| `source` | no | — |
| `source_created_at` | no | — |
| `source_created_by` | no | — |
| `source_readme_url` | no | — |
| `source_type` | no | — |
| `source_updated_at` | no | — |
| `source_updated_by` | no | — |
| `state` | no | — |
| `sys_lock` | no | — |
| `tags` | no | — |
| `targets_ini` | no | — |
| `updated_at` | no | — |
| `updated_by` | no | — |
| `user_state` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_schematics_action_data_invalid_action_id` | `action_id` |
| `ibm_schematics_action_data_invalid_location` | `location` |

## Example

### Valid

```hcl
data "ibm_schematics_action" "test" {
  action_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
