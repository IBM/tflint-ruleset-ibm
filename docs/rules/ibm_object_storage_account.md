# ibm_object_storage_account

Schema validation for the ibm_object_storage_account resource.

This rule validates `ibm_object_storage_account` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_object_storage_account" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `local_note` | no | — |
| `name` | no | — |
| `tags` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_object_storage_account_invalid_local_note` | `local_note` |
| `ibm_object_storage_account_invalid_tags` | `tags` |

## Example

### Valid

```hcl
resource "ibm_object_storage_account" "test" {
}
```
