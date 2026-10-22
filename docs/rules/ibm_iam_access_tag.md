# ibm_iam_access_tag

Manages ibm_iam_access_tag. Name of the access tag

This rule validates `ibm_iam_access_tag` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Tag` (`tagging`).

## Configuration

```hcl
rule "ibm_iam_access_tag" {
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
| `ibm_iam_access_tag_invalid_name` | `name` |

## Example

### Valid

```hcl
resource "ibm_iam_access_tag" "test" {
  name = "env:prod"
}
```
