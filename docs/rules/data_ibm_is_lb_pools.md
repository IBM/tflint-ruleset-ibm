# ibm_is_lb_pools_data

Schema validation for the ibm_is_lb_pools data source.

This rule validates `ibm_is_lb_pools` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_is_lb_pools_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `lb` | yes | — |
| `pools` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_lb_pools_data_invalid_lb` | `lb` |

## Example

### Valid

```hcl
data "ibm_is_lb_pools" "test" {
  lb = "example"
}
```
