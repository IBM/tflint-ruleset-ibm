# ibm_is_placement_group_data

Reads ibm_is_placement_group. The placement group name.

This rule validates `ibm_is_placement_group` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `BareMetalServerDiskPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_placement_group_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_tags` | no | — |
| `created_at` | no | — |
| `crn` | no | — |
| `href` | no | — |
| `lifecycle_state` | no | — |
| `name` | yes | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `resource_group` | no | — |
| `resource_type` | no | — |
| `strategy` | no | — |
| `tags` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_placement_group_data_invalid_name` | `name` |

## Example

### Valid

```hcl
data "ibm_is_placement_group" "test" {
  name = "example-name"
}
```
