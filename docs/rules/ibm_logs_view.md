# ibm_logs_view

Manages ibm_logs_view. View name.

This rule validates `ibm_logs_view` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `View` (`logs-service-api`).

## Configuration

```hcl
rule "ibm_logs_view" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `folder_id` | no | pattern `^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$` |
| `name` | yes | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `tier` | no | one of: `all_logs`, `all_logs_templates`, `priority_insights`, `priority_insights_templates` |
| `view_id` | no | — |
| `filters` (block) | no | nested block |
| `filters.filters` (block) | no | nested block |
| `filters.filters.name` | yes | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `filters.filters.selected_values` | no | API: Filter selected values |
| `search_query` (block) | no | nested block |
| `search_query.query` | yes | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `search_query.syntax_type` | no | one of: `lucene`, `dataprime` |
| `time_selection` (block) | yes | nested block |
| `time_selection.custom_selection` (block) | no | nested block |
| `time_selection.custom_selection.from_time` | yes | API: Custom time selection starting timestamp |
| `time_selection.custom_selection.to_time` | yes | API: Custom time selection ending timestamp |
| `time_selection.quick_selection` (block) | no | nested block |
| `time_selection.quick_selection.caption` | yes | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `time_selection.quick_selection.seconds` | yes | range 0–4294967295 |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_logs_view_invalid_folder_id` | `folder_id` |
| `ibm_logs_view_invalid_name` | `name` |
| `ibm_logs_view_invalid_tier` | `tier` |
| `ibm_logs_view_invalid_filters__filters__name` | `filters.filters.name` |
| `ibm_logs_view_invalid_filters__filters__selected_values` | `filters.filters.selected_values` |
| `ibm_logs_view_invalid_search_query__query` | `search_query.query` |
| `ibm_logs_view_invalid_search_query__syntax_type` | `search_query.syntax_type` |
| `ibm_logs_view_invalid_time_selection__custom_selection__from_time` | `time_selection.custom_selection.from_time` |
| `ibm_logs_view_invalid_time_selection__custom_selection__to_time` | `time_selection.custom_selection.to_time` |
| `ibm_logs_view_invalid_time_selection__quick_selection__caption` | `time_selection.quick_selection.caption` |
| `ibm_logs_view_invalid_time_selection__quick_selection__seconds` | `time_selection.quick_selection.seconds` |

## Example

### Valid

```hcl
resource "ibm_logs_view" "test" {
  name = "example-name"
  time_selection {
  }
}
```
