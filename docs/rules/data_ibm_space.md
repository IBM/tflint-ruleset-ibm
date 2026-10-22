# ibm_space_data

Reads ibm_space. Space name, for example dev

This rule validates `ibm_space` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_space_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `auditors` | no | — |
| `developers` | no | — |
| `managers` | no | — |
| `name` | no | — |
| `org` | yes | — |
| `space` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_space_data_invalid_name` | `name` |
| `ibm_space_data_invalid_org` | `org` |
| `ibm_space_data_invalid_space` | `space` |

## Example

### Valid

```hcl
data "ibm_space" "test" {
  org = "example"
}
```
