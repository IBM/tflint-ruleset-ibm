# ibm_resource_group_data

Reads ibm_resource_group. Resource group name

This rule validates `ibm_resource_group` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/v2/resource_groups/{id}:patch` (`resource-controller-resource-manager`).

## Configuration

```hcl
rule "ibm_resource_group_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_id` | no | — |
| `created_at` | no | — |
| `crn` | no | — |
| `id` | no | — |
| `is_default` | no | — |
| `name` | no | API: The new name of the resource group. |
| `payment_methods_url` | no | — |
| `quota_id` | no | — |
| `quota_url` | no | — |
| `resource_linkages` | no | — |
| `state` | no | — |
| `teams_url` | no | — |
| `updated_at` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_resource_group_data_invalid_id` | `id` |
| `ibm_resource_group_data_invalid_is_default` | `is_default` |
| `ibm_resource_group_data_invalid_name` | `name` |

## Example

### Valid

```hcl
data "ibm_resource_group" "test" {
}
```
