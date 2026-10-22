# ibm_is_share_delete_accessor_binding

Schema validation for the ibm_is_share_delete_accessor_binding resource.

This rule validates `ibm_is_share_delete_accessor_binding` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_is_share_delete_accessor_binding" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `accessor_binding` | yes | — |
| `share` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_share_delete_accessor_binding_invalid_accessor_binding` | `accessor_binding` |
| `ibm_is_share_delete_accessor_binding_invalid_share` | `share` |

## Example

### Valid

```hcl
resource "ibm_is_share_delete_accessor_binding" "test" {
  accessor_binding = "example"
  share = "example"
}
```
