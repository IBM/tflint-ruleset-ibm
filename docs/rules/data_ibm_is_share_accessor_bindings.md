# ibm_is_share_accessor_bindings_data

Schema validation for the ibm_is_share_accessor_bindings data source.

This rule validates `ibm_is_share_accessor_bindings` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_is_share_accessor_bindings_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `accessor_bindings` | no | — |
| `share` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_share_accessor_bindings_data_invalid_share` | `share` |

## Example

### Valid

```hcl
data "ibm_is_share_accessor_bindings" "test" {
  share = "example"
}
```
