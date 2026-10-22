# ibm_logs_log_data_retention_tags_data

Schema validation for the ibm_logs_log_data_retention_tags data source.

This rule validates `ibm_logs_log_data_retention_tags` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_logs_log_data_retention_tags_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `tags` | no | — |

## Example

### Valid

```hcl
data "ibm_logs_log_data_retention_tags" "test" {
}
```
