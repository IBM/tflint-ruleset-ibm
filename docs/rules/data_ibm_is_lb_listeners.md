# ibm_is_lb_listeners_data

Schema validation for the ibm_is_lb_listeners data source.

This rule validates `ibm_is_lb_listeners` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_is_lb_listeners_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `lb` | yes | — |
| `listeners` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_lb_listeners_data_invalid_lb` | `lb` |

## Example

### Valid

```hcl
data "ibm_is_lb_listeners" "test" {
  lb = "example"
}
```
