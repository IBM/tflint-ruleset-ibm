# ibm_cr_namespace

Manages ibm_cr_namespace. The name of the namespace.

This rule validates `ibm_cr_namespace` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_cr_namespace" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account` | no | — |
| `created_date` | no | — |
| `created_on` | no | — |
| `crn` | no | — |
| `name` | yes | pattern `^[a-z0-9]+[a-z0-9_-]+[a-z0-9]+$` |
| `resource_created_date` | no | — |
| `resource_group_id` | no | — |
| `tags` | no | — |
| `updated_date` | no | — |
| `updated_on` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cr_namespace_invalid_name` | `name` |
| `ibm_cr_namespace_invalid_resource_group_id` | `resource_group_id` |
| `ibm_cr_namespace_invalid_tags` | `tags` |

## Example

### Valid

```hcl
resource "ibm_cr_namespace" "test" {
  name = "example-name"
}
```
