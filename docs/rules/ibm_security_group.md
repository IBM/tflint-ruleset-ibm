# ibm_security_group

Manages ibm_security_group. Security group name

This rule validates `ibm_security_group` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_security_group" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `description` | no | — |
| `name` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_security_group_invalid_description` | `description` |
| `ibm_security_group_invalid_name` | `name` |

## Example

### Valid

```hcl
resource "ibm_security_group" "test" {
  name = "example"
}
```
