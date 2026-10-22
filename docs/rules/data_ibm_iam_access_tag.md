# ibm_iam_access_tag_data

Reads ibm_iam_access_tag. Name of the access tag to be fetched

This rule validates `ibm_iam_access_tag` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Tag` (`tagging`).

## Configuration

```hcl
rule "ibm_iam_access_tag_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `name` | yes | pattern `^([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-]):([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-])$` |
| `tag_type` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_access_tag_data_invalid_name` | `name` |

## Example

### Valid

```hcl
data "ibm_iam_access_tag" "test" {
  name = "env:prod"
}
```
