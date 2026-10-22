# ibm_resource_tag

Schema validation for the ibm_resource_tag resource.

This rule validates `ibm_resource_tag` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Tag` (`tagging`).

## Configuration

```hcl
rule "ibm_resource_tag" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_id` | no | — |
| `replace` | no | — |
| `resource_id` | yes | pattern `^crn:v1(:[a-zA-Z0-9 \-\._~\*\+,;=!$&'\(\)\/\?#\[\]@]*){8}$|^[0-9]+$` |
| `resource_type` | no | pattern `^[A-Za-z0-9_ :.-]*$` |
| `tag_type` | no | one of: `service`, `access`, `user` |
| `tags` | no | pattern `^[A-Za-z0-9:_ .-]+$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_resource_tag_invalid_replace` | `replace` |
| `ibm_resource_tag_invalid_resource_id` | `resource_id` |
| `ibm_resource_tag_invalid_resource_type` | `resource_type` |
| `ibm_resource_tag_invalid_tag_type` | `tag_type` |
| `ibm_resource_tag_invalid_tags` | `tags` |

## Example

### Valid

```hcl
resource "ibm_resource_tag" "test" {
  resource_id = "1"
}
```
