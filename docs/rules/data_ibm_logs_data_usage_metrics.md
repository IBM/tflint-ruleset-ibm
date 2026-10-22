# ibm_logs_data_usage_metrics_data

Schema validation for the ibm_logs_data_usage_metrics data source.

This rule validates `ibm_logs_data_usage_metrics` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_logs_data_usage_metrics_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `enabled` | no | — |

## Example

### Valid

```hcl
data "ibm_logs_data_usage_metrics" "test" {
}
```
