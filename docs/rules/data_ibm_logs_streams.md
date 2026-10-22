# ibm_logs_streams_data

Schema validation for the ibm_logs_streams data source.

This rule validates `ibm_logs_streams` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_logs_streams_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `streams` | no | — |

## Example

### Valid

```hcl
data "ibm_logs_streams" "test" {
}
```
