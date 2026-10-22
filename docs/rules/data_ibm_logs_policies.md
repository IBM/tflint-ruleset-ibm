# ibm_logs_policies_data

Schema validation for the ibm_logs_policies data source.

This rule validates `ibm_logs_policies` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_logs_policies_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `enabled_only` | no | — |
| `policies` | no | — |
| `source_type` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_logs_policies_data_invalid_enabled_only` | `enabled_only` |
| `ibm_logs_policies_data_invalid_source_type` | `source_type` |

## Example

### Valid

```hcl
data "ibm_logs_policies" "test" {
}
```
