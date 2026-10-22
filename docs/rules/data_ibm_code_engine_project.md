# ibm_code_engine_project_data

Reads ibm_code_engine_project. The name of the project.

This rule validates `ibm_code_engine_project` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `params:post:/projects/{project_id}/apps` (`codeengine-v2`).

## Configuration

```hcl
rule "ibm_code_engine_project_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_id` | no | — |
| `created_at` | no | — |
| `crn` | no | — |
| `href` | no | — |
| `name` | no | — |
| `project_id` | yes | pattern `^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$` |
| `region` | no | — |
| `resource_group_id` | no | — |
| `resource_type` | no | — |
| `status` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_code_engine_project_data_invalid_project_id` | `project_id` |

## Example

### Valid

```hcl
data "ibm_code_engine_project" "test" {
  project_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
