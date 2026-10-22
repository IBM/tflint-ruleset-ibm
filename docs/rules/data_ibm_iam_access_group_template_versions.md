# ibm_iam_access_group_template_versions_data

Schema validation for the ibm_iam_access_group_template_versions data source.

This rule validates `ibm_iam_access_group_template_versions` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `params:post:/v1/group_templates/{template_id}/versions` (`iam-access-groups`).

## Configuration

```hcl
rule "ibm_iam_access_group_template_versions_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `first` | no | — |
| `group_template_versions` | no | — |
| `last` | no | — |
| `previous` | no | — |
| `template_id` | yes | pattern `^[a-zA-Z0-9_-]+$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_access_group_template_versions_data_invalid_template_id` | `template_id` |

## Example

### Valid

```hcl
data "ibm_iam_access_group_template_versions" "test" {
  template_id = "example-name"
}
```
