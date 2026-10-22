# ibm_logs_view_folders_data

Schema validation for the ibm_logs_view_folders data source.

This rule validates `ibm_logs_view_folders` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_logs_view_folders_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `view_folders` | no | — |

## Example

### Valid

```hcl
data "ibm_logs_view_folders" "test" {
}
```
