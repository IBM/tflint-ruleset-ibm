# ibm_is_security_group_target

Manages ibm_is_security_group_target. Security group target name

This rule validates `ibm_is_security_group_target` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `FloatingIPByTarget` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_security_group_target" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `crn` | no | — |
| `name` | no | — |
| `resource_type` | no | — |
| `security_group` | yes | pattern `^[-0-9a-z_]+$` |
| `target` | yes | pattern `^[-0-9a-z_]+$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_security_group_target_invalid_security_group` | `security_group` |
| `ibm_is_security_group_target_invalid_target` | `target` |

## Example

### Valid

```hcl
resource "ibm_is_security_group_target" "test" {
  security_group = "example-name"
  target = "example-name"
}
```
