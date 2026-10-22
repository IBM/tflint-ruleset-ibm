# ibm_logs_e2ms_data

Schema validation for the ibm_logs_e2ms data source.

This rule validates `ibm_logs_e2ms` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_logs_e2ms_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `events2metrics` | no | — |

## Example

### Valid

```hcl
data "ibm_logs_e2ms" "test" {
}
```
