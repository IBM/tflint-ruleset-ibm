# ibm_billing_report_snapshot

Schema validation for the ibm_billing_report_snapshot resource.

This rule validates `ibm_billing_report_snapshot` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/v1/billing-reports-snapshot-config:patch` (`metering-reporting`).

## Configuration

```hcl
rule "ibm_billing_report_snapshot" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_type` | no | — |
| `compression` | no | — |
| `content_type` | no | — |
| `cos_bucket` | yes | API: The name of the COS bucket to store the snapshot of the billing reports. |
| `cos_endpoint` | no | — |
| `cos_location` | yes | API: Region of the COS instance. |
| `cos_reports_folder` | no | API: The billing reports root folder to store the billing reports snapshots. |
| `created_at` | no | — |
| `history` | no | — |
| `interval` | yes | one of: `daily` |
| `last_updated_at` | no | — |
| `report_types` | no | one of: `account_summary`, `enterprise_summary`, `account_resource_instance_usage` |
| `state` | no | — |
| `versioning` | no | one of: `new`, `overwrite` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_billing_report_snapshot_invalid_cos_bucket` | `cos_bucket` |
| `ibm_billing_report_snapshot_invalid_cos_location` | `cos_location` |
| `ibm_billing_report_snapshot_invalid_cos_reports_folder` | `cos_reports_folder` |
| `ibm_billing_report_snapshot_invalid_interval` | `interval` |
| `ibm_billing_report_snapshot_invalid_report_types` | `report_types` |
| `ibm_billing_report_snapshot_invalid_versioning` | `versioning` |

## Example

### Valid

```hcl
resource "ibm_billing_report_snapshot" "test" {
  cos_bucket = "example"
  cos_location = "example"
  interval = "daily"
}
```
