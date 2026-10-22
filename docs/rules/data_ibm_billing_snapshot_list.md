# ibm_billing_snapshot_list_data

Schema validation for the ibm_billing_snapshot_list data source.

This rule validates `ibm_billing_snapshot_list` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/v1/billing-reports-snapshot-config:patch` (`metering-reporting`).

## Configuration

```hcl
rule "ibm_billing_snapshot_list_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `date_from` | no | — |
| `date_to` | no | — |
| `limit` | no | — |
| `month` | yes | API: Month of captured snapshot. |
| `snapshotcount` | no | — |
| `snapshots` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_billing_snapshot_list_data_invalid_date_from` | `date_from` |
| `ibm_billing_snapshot_list_data_invalid_date_to` | `date_to` |
| `ibm_billing_snapshot_list_data_invalid_limit` | `limit` |
| `ibm_billing_snapshot_list_data_invalid_month` | `month` |

## Example

### Valid

```hcl
data "ibm_billing_snapshot_list" "test" {
  month = "example"
}
```
