# ibm_security_group_rule

Schema validation for the ibm_security_group_rule resource.

This rule validates `ibm_security_group_rule` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_security_group_rule" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `direction` | yes | — |
| `ether_type` | no | — |
| `port_range_max` | no | — |
| `port_range_min` | no | — |
| `protocol` | no | — |
| `remote_group_id` | no | — |
| `remote_ip` | no | — |
| `security_group_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_security_group_rule_invalid_direction` | `direction` |
| `ibm_security_group_rule_invalid_ether_type` | `ether_type` |
| `ibm_security_group_rule_invalid_port_range_max` | `port_range_max` |
| `ibm_security_group_rule_invalid_port_range_min` | `port_range_min` |
| `ibm_security_group_rule_invalid_protocol` | `protocol` |
| `ibm_security_group_rule_invalid_remote_group_id` | `remote_group_id` |
| `ibm_security_group_rule_invalid_remote_ip` | `remote_ip` |
| `ibm_security_group_rule_invalid_security_group_id` | `security_group_id` |

## Example

### Valid

```hcl
resource "ibm_security_group_rule" "test" {
  direction = "example"
  security_group_id = 1
}
```
