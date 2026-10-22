# ibm_is_security_group_rule_data

Reads ibm_is_security_group_rule. The name for this security group rule.

This rule validates `ibm_is_security_group_rule` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_is_security_group_rule_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `code` | no | — |
| `direction` | no | — |
| `href` | no | — |
| `ip_version` | no | — |
| `local` | no | — |
| `name` | no | — |
| `port_max` | no | — |
| `port_min` | no | — |
| `protocol` | no | — |
| `remote` | no | — |
| `security_group` | yes | — |
| `security_group_rule` | yes | — |
| `type` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_security_group_rule_data_invalid_security_group` | `security_group` |
| `ibm_is_security_group_rule_data_invalid_security_group_rule` | `security_group_rule` |

## Example

### Valid

```hcl
data "ibm_is_security_group_rule" "test" {
  security_group = "example"
  security_group_rule = "example"
}
```
