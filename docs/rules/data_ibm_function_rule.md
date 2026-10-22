# ibm_function_rule_data

Reads ibm_function_rule. Name of the rule.

This rule validates `ibm_function_rule` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_function_rule_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `action_name` | no | — |
| `name` | yes | pattern `\A([\w]|[\w][\w@ .-]*[\w@.-]+)\z` |
| `namespace` | yes | — |
| `publish` | no | — |
| `rule_id` | no | — |
| `status` | no | — |
| `trigger_name` | no | — |
| `version` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_function_rule_data_invalid_name` | `name` |
| `ibm_function_rule_data_invalid_namespace` | `namespace` |

## Example

### Valid

```hcl
data "ibm_function_rule" "test" {
  name = "A"
  namespace = "example"
}
```
