# ibm_project_data

Schema validation for the ibm_project data source.

This rule validates `ibm_project` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Project update` (`projects`).

## Configuration

```hcl
rule "ibm_project_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `configs` | no | — |
| `created_at` | no | — |
| `crn` | no | — |
| `cumulative_needs_attention_view` | no | — |
| `cumulative_needs_attention_view_error` | no | — |
| `definition` | no | — |
| `environments` | no | — |
| `event_notifications_crn` | no | — |
| `href` | no | — |
| `location` | no | — |
| `project_id` | yes | pattern `^[\.\-0-9a-zA-Z]+$` |
| `resource_group` | no | — |
| `resource_group_id` | no | — |
| `state` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_project_data_invalid_project_id` | `project_id` |

## Example

### Valid

```hcl
data "ibm_project" "test" {
  project_id = "example-name"
}
```
