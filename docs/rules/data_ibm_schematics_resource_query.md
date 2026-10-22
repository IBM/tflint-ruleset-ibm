# ibm_schematics_resource_query_data

Reads ibm_schematics_resource_query. Resource query name.

This rule validates `ibm_schematics_resource_query` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `params:put:/v2/resources_query/{query_id}` (`schematics-schematics`).

## Configuration

```hcl
rule "ibm_schematics_resource_query_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `created_at` | no | — |
| `created_by` | no | — |
| `id` | no | — |
| `location` | no | — |
| `name` | no | — |
| `queries` | no | — |
| `query_id` | yes | — |
| `type` | no | — |
| `updated_at` | no | — |
| `updated_by` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_schematics_resource_query_data_invalid_location` | `location` |
| `ibm_schematics_resource_query_data_invalid_query_id` | `query_id` |

## Example

### Valid

```hcl
data "ibm_schematics_resource_query" "test" {
  query_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
