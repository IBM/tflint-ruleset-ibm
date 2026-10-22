# ibm_scc_report_controls_data

Schema validation for the ibm_scc_report_controls data source.

This rule validates `ibm_scc_report_controls` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ControlPrototype` (`security-compliance`).

## Configuration

```hcl
rule "ibm_scc_report_controls_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `compliant_count` | no | — |
| `control_category` | no | API: The association of the control |
| `control_description` | no | API: The control description. |
| `control_id` | no | — |
| `control_name` | no | API: The ID of the control library that contains the profile. |
| `controls` | no | — |
| `home_account_id` | no | — |
| `instance_id` | yes | — |
| `not_compliant_count` | no | — |
| `report_id` | yes | — |
| `sort` | no | — |
| `status` | no | API: Details if a control library can be used or not. |
| `total_count` | no | — |
| `unable_to_perform_count` | no | — |
| `user_evaluation_required_count` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_scc_report_controls_data_invalid_control_category` | `control_category` |
| `ibm_scc_report_controls_data_invalid_control_description` | `control_description` |
| `ibm_scc_report_controls_data_invalid_control_id` | `control_id` |
| `ibm_scc_report_controls_data_invalid_control_name` | `control_name` |
| `ibm_scc_report_controls_data_invalid_instance_id` | `instance_id` |
| `ibm_scc_report_controls_data_invalid_report_id` | `report_id` |
| `ibm_scc_report_controls_data_invalid_sort` | `sort` |
| `ibm_scc_report_controls_data_invalid_status` | `status` |

## Example

### Valid

```hcl
data "ibm_scc_report_controls" "test" {
  instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  report_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
