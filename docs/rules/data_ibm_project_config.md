# ibm_project_config_data

Schema validation for the ibm_project_config data source.

This rule validates `ibm_project_config` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Project update` (`projects`).

## Configuration

```hcl
rule "ibm_project_config_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `approved_version` | no | — |
| `config_error` | no | — |
| `container_state` | no | — |
| `container_state_code` | no | — |
| `created_at` | no | — |
| `definition` | no | — |
| `deployed_version` | no | — |
| `deployment_model` | no | — |
| `href` | no | — |
| `is_draft` | no | — |
| `last_saved_at` | no | — |
| `member_of` | no | — |
| `modified_at` | no | — |
| `needs_attention_state` | no | — |
| `outputs` | no | — |
| `project` | no | — |
| `project_config_id` | yes | — |
| `project_id` | yes | pattern `^[\.\-0-9a-zA-Z]+$` |
| `references` | no | — |
| `schematics` | no | — |
| `state` | no | — |
| `state_code` | no | — |
| `template_id` | no | — |
| `update_available` | no | — |
| `version` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_project_config_data_invalid_project_config_id` | `project_config_id` |
| `ibm_project_config_data_invalid_project_id` | `project_id` |

## Example

### Valid

```hcl
data "ibm_project_config" "test" {
  project_config_id = "01234567-89ab-cdef-0123-456789abcdef"
  project_id = "example-name"
}
```
