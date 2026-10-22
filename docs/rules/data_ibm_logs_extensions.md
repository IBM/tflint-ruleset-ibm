# ibm_logs_extensions_data

Schema validation for the ibm_logs_extensions data source.

This rule validates `ibm_logs_extensions` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_logs_extensions_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `deployed` | no | — |
| `extensions` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_logs_extensions_data_invalid_deployed` | `deployed` |

## Example

### Valid

```hcl
data "ibm_logs_extensions" "test" {
}
```
