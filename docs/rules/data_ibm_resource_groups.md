# ibm_resource_groups_data

Reads ibm_resource_groups. Resource group name

This rule validates `ibm_resource_groups` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/v2/resource_groups/{id}:patch` (`resource-controller-resource-manager`).

## Configuration

```hcl
rule "ibm_resource_groups_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `date` | no | — |
| `include_deleted` | no | — |
| `is_default` | no | — |
| `name` | no | API: The new name of the resource group. |
| `resource_groups` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_resource_groups_data_invalid_date` | `date` |
| `ibm_resource_groups_data_invalid_include_deleted` | `include_deleted` |
| `ibm_resource_groups_data_invalid_is_default` | `is_default` |
| `ibm_resource_groups_data_invalid_name` | `name` |

## Example

### Valid

```hcl
data "ibm_resource_groups" "test" {
}
```
