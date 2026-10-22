# ibm_logs_dashboard

Manages ibm_logs_dashboard. Display name of the dashboard.

This rule validates `ibm_logs_dashboard` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Dashboard` (`logs-service-api`).

## Configuration

```hcl
rule "ibm_logs_dashboard" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `dashboard_id` | no | — |
| `description` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `href` | no | pattern `^[a-zA-Z0-9]{21}$` |
| `name` | yes | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `relative_time_frame` | no | pattern `^[0-9]+[smhdw]?$` |
| `absolute_time_frame` (block) | no | nested block |
| `absolute_time_frame.from` | no | — |
| `absolute_time_frame.to` | no | API: To is the end of the time frame |
| `annotations` (block) | no | nested block |
| `annotations.enabled` | yes | API: Whether the annotation is enabled. If enabled, annotations are applied. |
| `annotations.href` | no | pattern `^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$` |
| `annotations.id` | yes | pattern `^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$` |
| `annotations.name` | yes | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `annotations.source` (block) | yes | nested block |
| `annotations.source.logs` (block) | no | nested block |
| `annotations.source.logs.message_template` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `annotations.source.logs.label_fields` (block) | no | nested block |
| `annotations.source.logs.label_fields.keypath` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `annotations.source.logs.label_fields.scope` | no | one of: `unspecified`, `user_data`, `label`, `metadata` |
| `annotations.source.logs.lucene_query` (block) | yes | nested block |
| `annotations.source.logs.lucene_query.value` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `annotations.source.logs.strategy` (block) | yes | nested block |
| `annotations.source.logs.strategy.duration` (block) | no | nested block |
| `annotations.source.logs.strategy.instant` (block) | no | nested block |
| `annotations.source.logs.strategy.range` (block) | no | nested block |
| `annotations.source.metrics` (block) | no | nested block |
| `annotations.source.metrics.labels` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `annotations.source.metrics.message_template` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `annotations.source.metrics.promql_query` (block) | no | nested block |
| `annotations.source.metrics.promql_query.value` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `annotations.source.metrics.strategy` (block) | no | nested block |
| `annotations.source.metrics.strategy.start_time_metric` (block) | no | nested block |
| `false` (block) | no | nested block |
| `filters` (block) | no | nested block |
| `filters.collapsed` | no | API: Indicates if the filter's UI representation will be collapsed or expanded. |
| `filters.enabled` | no | API: Indicates if the filter is currently enabled or not. |
| `filters.source` (block) | no | nested block |
| `filters.source.logs` (block) | no | nested block |
| `filters.source.logs.observation_field` (block) | no | nested block |
| `filters.source.logs.observation_field.keypath` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `filters.source.logs.observation_field.scope` | no | one of: `unspecified`, `user_data`, `label`, `metadata` |
| `filters.source.logs.operator` (block) | no | nested block |
| `filters.source.logs.operator.equals` (block) | no | nested block |
| `filters.source.logs.operator.not_equals` (block) | no | nested block |
| `filters.source.metrics` (block) | no | nested block |
| `filters.source.metrics.label` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `filters.source.metrics.operator` (block) | no | nested block |
| `filters.source.metrics.operator.equals` (block) | no | nested block |
| `filters.source.metrics.operator.not_equals` (block) | no | nested block |
| `five_minutes` (block) | no | nested block |
| `folder_id` (block) | no | nested block |
| `folder_id.value` | yes | pattern `^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$` |
| `folder_path` (block) | no | nested block |
| `folder_path.segments` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `layout` (block) | yes | nested block |
| `layout.sections` (block) | no | nested block |
| `layout.sections.href` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `layout.sections.id` (block) | yes | nested block |
| `layout.sections.id.value` | yes | pattern `^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$` |
| `layout.sections.rows` (block) | no | nested block |
| `layout.sections.rows.href` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `layout.sections.rows.appearance` (block) | yes | nested block |
| `layout.sections.rows.appearance.height` | yes | range 0–2147483647 |
| `layout.sections.rows.id` (block) | yes | nested block |
| `layout.sections.rows.id.value` | yes | pattern `^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$` |
| `layout.sections.rows.widgets` (block) | no | nested block |
| `layout.sections.rows.widgets.created_at` | no | API: Creation timestamp |
| `layout.sections.rows.widgets.description` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `layout.sections.rows.widgets.href` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `layout.sections.rows.widgets.title` | yes | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `layout.sections.rows.widgets.updated_at` | no | API: Last update timestamp |
| `layout.sections.rows.widgets.definition` (block) | yes | nested block |
| `layout.sections.rows.widgets.id` (block) | yes | nested block |
| `two_minutes` (block) | no | nested block |
| `variables` (block) | no | nested block |
| `variables.display_name` | yes | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `variables.name` | yes | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `variables.definition` (block) | yes | nested block |
| `variables.definition.multi_select` (block) | no | nested block |
| `variables.definition.multi_select.values_order_direction` | yes | one of: `unspecified`, `asc`, `desc` |
| `variables.definition.multi_select.selection` (block) | yes | nested block |
| `variables.definition.multi_select.selection.all` (block) | no | nested block |
| `variables.definition.multi_select.selection.list` (block) | no | nested block |
| `variables.definition.multi_select.source` (block) | yes | nested block |
| `variables.definition.multi_select.source.constant_list` (block) | no | nested block |
| `variables.definition.multi_select.source.logs_path` (block) | no | nested block |
| `variables.definition.multi_select.source.metric_label` (block) | no | nested block |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_logs_dashboard_invalid_description` | `description` |
| `ibm_logs_dashboard_invalid_href` | `href` |
| `ibm_logs_dashboard_invalid_name` | `name` |
| `ibm_logs_dashboard_invalid_relative_time_frame` | `relative_time_frame` |
| `ibm_logs_dashboard_invalid_absolute_time_frame__from` | `absolute_time_frame.from` |
| `ibm_logs_dashboard_invalid_absolute_time_frame__to` | `absolute_time_frame.to` |
| `ibm_logs_dashboard_invalid_annotations__enabled` | `annotations.enabled` |
| `ibm_logs_dashboard_invalid_annotations__href` | `annotations.href` |
| `ibm_logs_dashboard_invalid_annotations__id` | `annotations.id` |
| `ibm_logs_dashboard_invalid_annotations__name` | `annotations.name` |
| `ibm_logs_dashboard_invalid_annotations__source__logs__message_template` | `annotations.source.logs.message_template` |
| `ibm_logs_dashboard_invalid_annotations__source__logs__label_fields__keypath` | `annotations.source.logs.label_fields.keypath` |
| `ibm_logs_dashboard_invalid_annotations__source__logs__label_fields__scope` | `annotations.source.logs.label_fields.scope` |
| `ibm_logs_dashboard_invalid_annotations__source__logs__lucene_query__value` | `annotations.source.logs.lucene_query.value` |
| `ibm_logs_dashboard_invalid_annotations__source__metrics__labels` | `annotations.source.metrics.labels` |
| `ibm_logs_dashboard_invalid_annotations__source__metrics__message_template` | `annotations.source.metrics.message_template` |
| `ibm_logs_dashboard_invalid_annotations__source__metrics__promql_query__value` | `annotations.source.metrics.promql_query.value` |
| `ibm_logs_dashboard_invalid_filters__collapsed` | `filters.collapsed` |
| `ibm_logs_dashboard_invalid_filters__enabled` | `filters.enabled` |
| `ibm_logs_dashboard_invalid_filters__source__logs__observation_field__keypath` | `filters.source.logs.observation_field.keypath` |
| `ibm_logs_dashboard_invalid_filters__source__logs__observation_field__scope` | `filters.source.logs.observation_field.scope` |
| `ibm_logs_dashboard_invalid_filters__source__metrics__label` | `filters.source.metrics.label` |
| `ibm_logs_dashboard_invalid_folder_id__value` | `folder_id.value` |
| `ibm_logs_dashboard_invalid_folder_path__segments` | `folder_path.segments` |
| `ibm_logs_dashboard_invalid_layout__sections__href` | `layout.sections.href` |
| `ibm_logs_dashboard_invalid_layout__sections__id__value` | `layout.sections.id.value` |
| `ibm_logs_dashboard_invalid_layout__sections__rows__href` | `layout.sections.rows.href` |
| `ibm_logs_dashboard_invalid_layout__sections__rows__appearance__height` | `layout.sections.rows.appearance.height` |
| `ibm_logs_dashboard_invalid_layout__sections__rows__id__value` | `layout.sections.rows.id.value` |
| `ibm_logs_dashboard_invalid_layout__sections__rows__widgets__created_at` | `layout.sections.rows.widgets.created_at` |
| `ibm_logs_dashboard_invalid_layout__sections__rows__widgets__description` | `layout.sections.rows.widgets.description` |
| `ibm_logs_dashboard_invalid_layout__sections__rows__widgets__href` | `layout.sections.rows.widgets.href` |
| `ibm_logs_dashboard_invalid_layout__sections__rows__widgets__title` | `layout.sections.rows.widgets.title` |
| `ibm_logs_dashboard_invalid_layout__sections__rows__widgets__updated_at` | `layout.sections.rows.widgets.updated_at` |
| `ibm_logs_dashboard_invalid_variables__display_name` | `variables.display_name` |
| `ibm_logs_dashboard_invalid_variables__name` | `variables.name` |
| `ibm_logs_dashboard_invalid_variables__definition__multi_select__values_order_direction` | `variables.definition.multi_select.values_order_direction` |

## Example

### Valid

```hcl
resource "ibm_logs_dashboard" "test" {
  name = "example-name"
  layout {
  }
}
```
