# ibm_sm_secret_groups_data

Schema validation for the ibm_sm_secret_groups data source.

This rule validates `ibm_sm_secret_groups` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_sm_secret_groups_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `secret_groups` | no | — |
| `total_count` | no | — |

## Example

### Valid

```hcl
data "ibm_sm_secret_groups" "test" {
}
```
