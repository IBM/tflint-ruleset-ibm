# ibm_schematics_resource_query

Manages ibm_schematics_resource_query. Resource query name.

This rule validates `ibm_schematics_resource_query` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `WorkspaceVariableRequest` (`schematics-schematics`).

## Configuration

```hcl
rule "ibm_schematics_resource_query" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `created_at` | no | — |
| `created_by` | no | — |
| `location` | no | — |
| `name` | no | API: The name of the variable. |
| `type` | no | one of: `vsi` |
| `updated_at` | no | — |
| `updated_by` | no | — |
| `queries` (block) | no | nested block |
| `queries.query_select` | no | API: List of query selection parameters |
| `queries.query_type` | no | one of: `workspaces` |
| `queries.query_condition` (block) | no | nested block |
| `queries.query_condition.description` | no | API: The description of your input variable. |
| `queries.query_condition.name` | no | API: The name of the variable. |
| `queries.query_condition.value` | no | API: Enter the value as a string for the primitive types such as `bool`, `number`, `string`, and `HCL` format for the complex variables, as you provide in a `.tfvars` file. **You need to enter escaped string of `HCL` format for the complex variable value**. For more information, about how to declare variables in a terraform configuration file and provide value to schematics, see [Providing values for t |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_schematics_resource_query_invalid_location` | `location` |
| `ibm_schematics_resource_query_invalid_name` | `name` |
| `ibm_schematics_resource_query_invalid_type` | `type` |
| `ibm_schematics_resource_query_invalid_queries__query_select` | `queries.query_select` |
| `ibm_schematics_resource_query_invalid_queries__query_type` | `queries.query_type` |
| `ibm_schematics_resource_query_invalid_queries__query_condition__description` | `queries.query_condition.description` |
| `ibm_schematics_resource_query_invalid_queries__query_condition__name` | `queries.query_condition.name` |
| `ibm_schematics_resource_query_invalid_queries__query_condition__value` | `queries.query_condition.value` |

## Example

### Valid

```hcl
resource "ibm_schematics_resource_query" "test" {
}
```
