# ibm_scc_report_violation_drift_data

Schema validation for the ibm_scc_report_violation_drift data source.

This rule validates `ibm_scc_report_violation_drift` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Report` (`security-compliance`).

## Configuration

```hcl
rule "ibm_scc_report_violation_drift_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `data_points` | no | — |
| `home_account_id` | no | — |
| `instance_id` | yes | pattern `^[0-9A-Fa-f]{8}-[0-9A-Fa-f]{4}-[0-9A-Fa-f]{4}-[0-9A-Fa-f]{4}-[0-9A-Fa-f]{12}$` |
| `report_id` | yes | — |
| `scan_time_duration` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_scc_report_violation_drift_data_invalid_instance_id` | `instance_id` |
| `ibm_scc_report_violation_drift_data_invalid_report_id` | `report_id` |
| `ibm_scc_report_violation_drift_data_invalid_scan_time_duration` | `scan_time_duration` |

## Example

### Valid

```hcl
data "ibm_scc_report_violation_drift" "test" {
  instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  report_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
