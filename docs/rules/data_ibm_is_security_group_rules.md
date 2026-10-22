# ibm_is_security_group_rules_data

Schema validation for the ibm_is_security_group_rules data source.

This rule validates `ibm_is_security_group_rules` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_is_security_group_rules_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `rules` | no | — |
| `security_group` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_security_group_rules_data_invalid_security_group` | `security_group` |

## Example

### Valid

```hcl
data "ibm_is_security_group_rules" "test" {
  security_group = "example"
}
```
