# ibm_function_rule

Manages ibm_function_rule. Name of rule.

This rule validates `ibm_function_rule` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_function_rule" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `action_name` | yes | — |
| `name` | yes | pattern `\A([\w]|[\w][\w@ .-]*[\w@.-]+)\z` |
| `namespace` | yes | — |
| `publish` | no | — |
| `rule_id` | no | — |
| `status` | no | — |
| `trigger_name` | yes | — |
| `version` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_function_rule_invalid_action_name` | `action_name` |
| `ibm_function_rule_invalid_name` | `name` |
| `ibm_function_rule_invalid_namespace` | `namespace` |
| `ibm_function_rule_invalid_trigger_name` | `trigger_name` |

## Example

### Valid

```hcl
resource "ibm_function_rule" "test" {
  action_name = "example"
  name = "A"
  namespace = "example"
  trigger_name = "example"
}
```
