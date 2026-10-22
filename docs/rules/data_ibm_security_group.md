# ibm_security_group_data

Reads ibm_security_group. The name of the security group

This rule validates `ibm_security_group` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_security_group_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `description` | no | — |
| `most_recent` | no | — |
| `name` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_security_group_data_invalid_description` | `description` |
| `ibm_security_group_data_invalid_most_recent` | `most_recent` |
| `ibm_security_group_data_invalid_name` | `name` |

## Example

### Valid

```hcl
data "ibm_security_group" "test" {
  name = "example"
}
```
