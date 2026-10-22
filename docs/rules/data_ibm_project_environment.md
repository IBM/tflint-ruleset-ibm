# ibm_project_environment_data

Schema validation for the ibm_project_environment data source.

This rule validates `ibm_project_environment` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Project environment input` (`projects`).

## Configuration

```hcl
rule "ibm_project_environment_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `created_at` | no | — |
| `definition` | no | — |
| `href` | no | — |
| `modified_at` | no | — |
| `project` | no | — |
| `project_environment_id` | yes | — |
| `project_id` | yes | pattern `^[\.\-0-9a-zA-Z]+$` |
| `target_account` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_project_environment_data_invalid_project_environment_id` | `project_environment_id` |
| `ibm_project_environment_data_invalid_project_id` | `project_id` |

## Example

### Valid

```hcl
data "ibm_project_environment" "test" {
  project_environment_id = "01234567-89ab-cdef-0123-456789abcdef"
  project_id = "example-name"
}
```
