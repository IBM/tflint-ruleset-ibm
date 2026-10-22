# ibm_scc_latest_reports_data

Schema validation for the ibm_scc_latest_reports data source.

This rule validates `ibm_scc_latest_reports` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_scc_latest_reports_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `controls_summary` | no | — |
| `evaluations_summary` | no | — |
| `home_account_id` | no | — |
| `instance_id` | yes | — |
| `reports` | no | — |
| `score` | no | — |
| `sort` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_scc_latest_reports_data_invalid_instance_id` | `instance_id` |
| `ibm_scc_latest_reports_data_invalid_sort` | `sort` |

## Example

### Valid

```hcl
data "ibm_scc_latest_reports" "test" {
  instance_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
