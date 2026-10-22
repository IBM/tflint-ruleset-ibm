# ibm_resource_group

Manages ibm_resource_group. The name of the resource group

This rule validates `ibm_resource_group` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/v2/resource_groups/{id}:patch` (`resource-controller-resource-manager`).

## Configuration

```hcl
rule "ibm_resource_group" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `created_at` | no | — |
| `crn` | no | — |
| `default` | no | — |
| `name` | yes | API: The new name of the resource group. |
| `payment_methods_url` | no | — |
| `quota_id` | no | — |
| `quota_url` | no | — |
| `resource_linkages` | no | — |
| `state` | no | — |
| `tags` | no | — |
| `teams_url` | no | — |
| `updated_at` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_resource_group_invalid_name` | `name` |
| `ibm_resource_group_invalid_tags` | `tags` |

## Example

### Valid

```hcl
resource "ibm_resource_group" "test" {
  name = "example"
}
```
