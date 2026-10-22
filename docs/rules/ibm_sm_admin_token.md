# ibm_sm_admin_token

Schema validation for the ibm_sm_admin_token resource.

This rule validates `ibm_sm_admin_token` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_sm_admin_token" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `created_at` | no | — |
| `instance_id` | yes | — |
| `token` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_sm_admin_token_invalid_instance_id` | `instance_id` |

## Example

### Valid

```hcl
resource "ibm_sm_admin_token" "test" {
  instance_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
