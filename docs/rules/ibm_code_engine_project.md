# ibm_code_engine_project

Manages ibm_code_engine_project. The name of the project.

This rule validates `ibm_code_engine_project` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ProjectPrototype` (`codeengine-v2`).

## Configuration

```hcl
rule "ibm_code_engine_project" {
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
| `name` | yes | pattern `^([^\x00-\x7F]|[a-zA-Z0-9\-._: ])+$` |
| `project_id` | no | — |
| `region` | no | — |
| `resource_group_id` | no | pattern `^[a-z0-9]*$` |
| `resource_type` | no | — |
| `status` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_code_engine_project_invalid_name` | `name` |
| `ibm_code_engine_project_invalid_resource_group_id` | `resource_group_id` |

## Example

### Valid

```hcl
resource "ibm_code_engine_project" "test" {
  name = "example-name"
}
```
