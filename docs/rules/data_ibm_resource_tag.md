# ibm_resource_tag_data

Schema validation for the ibm_resource_tag data source.

This rule validates `ibm_resource_tag` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Resource` (`tagging`).

## Configuration

```hcl
rule "ibm_resource_tag_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `resource_id` | no | pattern `^crn:v1(:[a-zA-Z0-9 \-\._~\*\+,;=!$&'\(\)\/\?#\[\]@]*){8}$|^[0-9]+$` |
| `resource_type` | no | pattern `^[A-Za-z0-9_ :.-]*$` |
| `tag_type` | no | one of: `service`, `access`, `user` |
| `tags` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_resource_tag_data_invalid_resource_id` | `resource_id` |
| `ibm_resource_tag_data_invalid_resource_type` | `resource_type` |
| `ibm_resource_tag_data_invalid_tag_type` | `tag_type` |

## Example

### Valid

```hcl
data "ibm_resource_tag" "test" {
}
```
