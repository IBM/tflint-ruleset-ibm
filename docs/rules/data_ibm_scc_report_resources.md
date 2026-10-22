# ibm_scc_report_resources_data

Schema validation for the ibm_scc_report_resources data source.

This rule validates `ibm_scc_report_resources` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Report` (`security-compliance`).

## Configuration

```hcl
rule "ibm_scc_report_resources_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_id` | no | pattern `^[a-zA-Z0-9_\-.]*$` |
| `component_id` | no | API: The ID of the component. The component_id can be found from the 'service_name' using the Get Services method. |
| `first` | no | — |
| `home_account_id` | no | — |
| `id` | no | API: The ID of the report. |
| `instance_id` | yes | pattern `^[0-9A-Fa-f]{8}-[0-9A-Fa-f]{4}-[0-9A-Fa-f]{4}-[0-9A-Fa-f]{4}-[0-9A-Fa-f]{12}$` |
| `report_id` | yes | — |
| `resource_name` | no | — |
| `resources` | no | — |
| `sort` | no | — |
| `status` | no | one of: `compliant`, `not_compliant`, `unable_to_perform`, `user_evaluation_required`, `not_applicable` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_scc_report_resources_data_invalid_account_id` | `account_id` |
| `ibm_scc_report_resources_data_invalid_component_id` | `component_id` |
| `ibm_scc_report_resources_data_invalid_id` | `id` |
| `ibm_scc_report_resources_data_invalid_instance_id` | `instance_id` |
| `ibm_scc_report_resources_data_invalid_report_id` | `report_id` |
| `ibm_scc_report_resources_data_invalid_resource_name` | `resource_name` |
| `ibm_scc_report_resources_data_invalid_sort` | `sort` |
| `ibm_scc_report_resources_data_invalid_status` | `status` |

## Example

### Valid

```hcl
data "ibm_scc_report_resources" "test" {
  instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  report_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
