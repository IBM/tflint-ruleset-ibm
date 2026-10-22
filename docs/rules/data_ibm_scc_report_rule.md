# ibm_scc_report_rule_data

Schema validation for the ibm_scc_report_rule data source.

This rule validates `ibm_scc_report_rule` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Report` (`security-compliance`).

## Configuration

```hcl
rule "ibm_scc_report_rule_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_id` | no | — |
| `created_by` | no | — |
| `created_on` | no | — |
| `description` | no | — |
| `id` | no | — |
| `instance_id` | yes | pattern `^[0-9A-Fa-f]{8}-[0-9A-Fa-f]{4}-[0-9A-Fa-f]{4}-[0-9A-Fa-f]{4}-[0-9A-Fa-f]{12}$` |
| `labels` | no | — |
| `report_id` | yes | — |
| `rule_id` | yes | — |
| `type` | no | — |
| `updated_by` | no | — |
| `updated_on` | no | — |
| `version` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_scc_report_rule_data_invalid_instance_id` | `instance_id` |
| `ibm_scc_report_rule_data_invalid_report_id` | `report_id` |
| `ibm_scc_report_rule_data_invalid_rule_id` | `rule_id` |

## Example

### Valid

```hcl
data "ibm_scc_report_rule" "test" {
  instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  report_id = "01234567-89ab-cdef-0123-456789abcdef"
  rule_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
