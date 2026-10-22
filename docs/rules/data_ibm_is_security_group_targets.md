# ibm_is_security_group_targets_data

Schema validation for the ibm_is_security_group_targets data source.

This rule validates `ibm_is_security_group_targets` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_is_security_group_targets_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `security_group` | yes | — |
| `targets` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_security_group_targets_data_invalid_security_group` | `security_group` |

## Example

### Valid

```hcl
data "ibm_is_security_group_targets" "test" {
  security_group = "example"
}
```
