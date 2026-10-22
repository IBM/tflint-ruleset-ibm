# ibm_iam_authorization_policy_detach

Schema validation for the ibm_iam_authorization_policy_detach resource.

This rule validates `ibm_iam_authorization_policy_detach` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_iam_authorization_policy_detach" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `authorization_policy_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_authorization_policy_detach_invalid_authorization_policy_id` | `authorization_policy_id` |

## Example

### Valid

```hcl
resource "ibm_iam_authorization_policy_detach" "test" {
  authorization_policy_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
