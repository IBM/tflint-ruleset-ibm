# ibm_logs_dashboard_data

Reads ibm_logs_dashboard. Display name of the dashboard.

This rule validates `ibm_logs_dashboard` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_logs_dashboard_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `absolute_time_frame` | no | — |
| `annotations` | no | — |
| `dashboard_id` | yes | — |
| `description` | no | — |
| `false` | no | — |
| `filters` | no | — |
| `five_minutes` | no | — |
| `folder_id` | no | — |
| `folder_path` | no | — |
| `href` | no | — |
| `layout` | no | — |
| `name` | no | — |
| `relative_time_frame` | no | — |
| `two_minutes` | no | — |
| `variables` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_logs_dashboard_data_invalid_dashboard_id` | `dashboard_id` |

## Example

### Valid

```hcl
data "ibm_logs_dashboard" "test" {
  dashboard_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
