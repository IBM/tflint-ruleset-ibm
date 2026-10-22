# ibm_scc_report_evaluations_data

Schema validation for the ibm_scc_report_evaluations data source.

This rule validates `ibm_scc_report_evaluations` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Report` (`security-compliance`).

## Configuration

```hcl
rule "ibm_scc_report_evaluations_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `assessment_id` | no | API: The ID of the rule to target. A list of rules can be obtained from the list_rules method. |
| `component_id` | no | API: The ID of the component. The component_id can be found from the 'service_name' using the Get Services method. |
| `evaluations` | no | — |
| `first` | no | — |
| `home_account_id` | no | — |
| `instance_id` | yes | pattern `^[0-9A-Fa-f]{8}-[0-9A-Fa-f]{4}-[0-9A-Fa-f]{4}-[0-9A-Fa-f]{4}-[0-9A-Fa-f]{12}$` |
| `report_id` | yes | — |
| `status` | no | one of: `compliant`, `not_compliant`, `unable_to_perform`, `user_evaluation_required`, `not_applicable` |
| `target_id` | no | — |
| `target_name` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_scc_report_evaluations_data_invalid_assessment_id` | `assessment_id` |
| `ibm_scc_report_evaluations_data_invalid_component_id` | `component_id` |
| `ibm_scc_report_evaluations_data_invalid_instance_id` | `instance_id` |
| `ibm_scc_report_evaluations_data_invalid_report_id` | `report_id` |
| `ibm_scc_report_evaluations_data_invalid_status` | `status` |
| `ibm_scc_report_evaluations_data_invalid_target_id` | `target_id` |
| `ibm_scc_report_evaluations_data_invalid_target_name` | `target_name` |

## Example

### Valid

```hcl
data "ibm_scc_report_evaluations" "test" {
  instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  report_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
