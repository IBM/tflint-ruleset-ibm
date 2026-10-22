# ibm_is_placement_group

Manages ibm_is_placement_group. The unique user-defined name for this placement group.

This rule validates `ibm_is_placement_group` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `PlacementGroupPrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_placement_group" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_tags` | no | pattern `^([ ]*[A-Za-z0-9:_.-]+[ ]*)+$` |
| `created_at` | no | — |
| `crn` | no | — |
| `href` | no | — |
| `lifecycle_state` | no | — |
| `name` | yes | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `resource_group` | no | API: The resource group to use. If unspecified, the account's [default resource group](https://cloud.ibm.com/apidocs/resource-manager#introduction) will be used. |
| `resource_type` | no | — |
| `strategy` | yes | one of: `host_spread`, `power_spread`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `tags` | no | pattern `^[A-Za-z0-9:_ .-]+$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_placement_group_invalid_access_tags` | `access_tags` |
| `ibm_is_placement_group_invalid_name` | `name` |
| `ibm_is_placement_group_invalid_resource_group` | `resource_group` |
| `ibm_is_placement_group_invalid_strategy` | `strategy` |
| `ibm_is_placement_group_invalid_tags` | `tags` |

## Example

### Valid

```hcl
resource "ibm_is_placement_group" "test" {
  name = "example-name"
  strategy = "host_spread"
}
```
