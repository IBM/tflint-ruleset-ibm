# ibm_logs_enrichments_data

Schema validation for the ibm_logs_enrichments data source.

This rule validates `ibm_logs_enrichments` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_logs_enrichments_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `enrichments` | no | — |

## Example

### Valid

```hcl
data "ibm_logs_enrichments" "test" {
}
```
