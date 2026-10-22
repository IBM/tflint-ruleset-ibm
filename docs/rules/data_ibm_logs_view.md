# ibm_logs_view_data

Reads ibm_logs_view. View name.

This rule validates `ibm_logs_view` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_logs_view_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `filters` | no | — |
| `folder_id` | no | — |
| `logs_view_id` | yes | — |
| `name` | no | — |
| `search_query` | no | — |
| `tier` | no | — |
| `time_selection` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_logs_view_data_invalid_logs_view_id` | `logs_view_id` |

## Example

### Valid

```hcl
data "ibm_logs_view" "test" {
  logs_view_id = 1
}
```
