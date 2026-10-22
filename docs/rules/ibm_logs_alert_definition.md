# ibm_logs_alert_definition

Manages ibm_logs_alert_definition. The name of the alert definition.

This rule validates `ibm_logs_alert_definition` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Alert` (`logs-service-api`).

## Configuration

```hcl
rule "ibm_logs_alert_definition" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `alert_def_id` | no | — |
| `alert_version_id` | no | — |
| `created_time` | no | — |
| `deleted` | no | — |
| `description` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `enabled` | no | — |
| `entity_labels` | no | — |
| `group_by_keys` | no | — |
| `name` | yes | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `phantom_mode` | no | — |
| `priority` | no | one of: `p1`, `p2`, `p3`, `p4`, `p5_or_unspecified` |
| `type` | yes | one of: `flow`, `logs_anomaly`, `logs_immediate_or_unspecified`, `logs_new_value`, `logs_ratio_threshold`, `logs_threshold`, `logs_time_relative_threshold`, `logs_unique_count`, `metric_anomaly`, `metric_threshold` |
| `updated_time` | no | — |
| `active_on` (block) | no | nested block |
| `active_on.day_of_week` | yes | one of: `monday_or_unspecified`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`, `sunday` |
| `active_on.end_time` (block) | yes | nested block |
| `active_on.end_time.hours` | no | range 0–24 |
| `active_on.end_time.minutes` | no | range 0–59 |
| `active_on.start_time` (block) | yes | nested block |
| `active_on.start_time.hours` | no | range 0–24 |
| `active_on.start_time.minutes` | no | range 0–59 |
| `flow` (block) | no | nested block |
| `flow.enforce_suppression` | no | API: Should suppression be enforced on the flow alert |
| `flow.stages` (block) | yes | nested block |
| `flow.stages.timeframe_ms` | yes | pattern `^-?\d{1,19}$` |
| `flow.stages.timeframe_type` | yes | one of: `unspecified`, `up_to` |
| `flow.stages.flow_stages_groups` (block) | yes | nested block |
| `flow.stages.flow_stages_groups.groups` (block) | yes | nested block |
| `flow.stages.flow_stages_groups.groups.alerts_op` | yes | one of: `and_or_unspecified`, `or` |
| `flow.stages.flow_stages_groups.groups.next_op` | yes | one of: `and`, `or` |
| `flow.stages.flow_stages_groups.groups.alert_defs` (block) | yes | nested block |
| `incidents_settings` (block) | no | nested block |
| `incidents_settings.minutes` | no | range 0–59 |
| `incidents_settings.notify_on` | no | one of: `triggered_only`, `triggered_and_resolved` |
| `logs_anomaly` (block) | no | nested block |
| `logs_anomaly.condition_type` | yes | — |
| `logs_anomaly.evaluation_delay_ms` | no | range 0–2147483647 |
| `logs_anomaly.notification_payload_filter` | no | pattern `^[a-zA-Z0-9_.]+$` |
| `logs_anomaly.anomaly_alert_settings` (block) | no | nested block |
| `logs_anomaly.anomaly_alert_settings.percentage_of_deviation` | no | API: The percentage of deviation from the baseline when the alert is triggered. |
| `logs_anomaly.logs_filter` (block) | no | nested block |
| `logs_anomaly.logs_filter.simple_filter` (block) | no | nested block |
| `logs_anomaly.logs_filter.simple_filter.lucene_query` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `logs_anomaly.logs_filter.simple_filter.label_filters` (block) | no | nested block |
| `logs_anomaly.logs_filter.simple_filter.label_filters.severities` | no | one of: `debug_or_unspecified`, `verbose`, `info`, `warning`, `error`, `critical` |
| `logs_anomaly.logs_filter.simple_filter.label_filters.application_name` (block) | no | nested block |
| `logs_anomaly.logs_filter.simple_filter.label_filters.subsystem_name` (block) | no | nested block |
| `logs_anomaly.rules` (block) | yes | nested block |
| `logs_anomaly.rules.condition` (block) | yes | nested block |
| `logs_anomaly.rules.condition.minimum_threshold` | yes | API: The threshold value for the alert condition |
| `logs_anomaly.rules.condition.time_window` (block) | yes | nested block |
| `logs_anomaly.rules.condition.time_window.logs_time_window_specific_value` | yes | one of: `minutes_5_or_unspecified`, `minutes_10`, `minutes_20`, `minutes_15`, `minutes_30`, `hour_1`, `hours_2`, `hours_4`, `hours_6`, `hours_12`, `hours_24`, `hours_36` |
| `logs_immediate` (block) | no | nested block |
| `logs_immediate.notification_payload_filter` | no | pattern `^[a-zA-Z0-9_.]+$` |
| `logs_immediate.logs_filter` (block) | no | nested block |
| `logs_immediate.logs_filter.simple_filter` (block) | no | nested block |
| `logs_immediate.logs_filter.simple_filter.lucene_query` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `logs_immediate.logs_filter.simple_filter.label_filters` (block) | no | nested block |
| `logs_immediate.logs_filter.simple_filter.label_filters.severities` | no | one of: `debug_or_unspecified`, `verbose`, `info`, `warning`, `error`, `critical` |
| `logs_immediate.logs_filter.simple_filter.label_filters.application_name` (block) | no | nested block |
| `logs_immediate.logs_filter.simple_filter.label_filters.subsystem_name` (block) | no | nested block |
| `logs_new_value` (block) | no | nested block |
| `logs_new_value.notification_payload_filter` | no | pattern `^[a-zA-Z0-9_.]+$` |
| `logs_new_value.logs_filter` (block) | no | nested block |
| `logs_new_value.logs_filter.simple_filter` (block) | no | nested block |
| `logs_new_value.logs_filter.simple_filter.lucene_query` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `logs_new_value.logs_filter.simple_filter.label_filters` (block) | no | nested block |
| `logs_new_value.logs_filter.simple_filter.label_filters.severities` | no | one of: `debug_or_unspecified`, `verbose`, `info`, `warning`, `error`, `critical` |
| `logs_new_value.logs_filter.simple_filter.label_filters.application_name` (block) | no | nested block |
| `logs_new_value.logs_filter.simple_filter.label_filters.subsystem_name` (block) | no | nested block |
| `logs_new_value.rules` (block) | yes | nested block |
| `logs_new_value.rules.condition` (block) | yes | nested block |
| `logs_new_value.rules.condition.keypath_to_track` | yes | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `logs_new_value.rules.condition.time_window` (block) | yes | nested block |
| `logs_new_value.rules.condition.time_window.logs_new_value_time_window_specific_value` | yes | one of: `hours_12_or_unspecified`, `hours_24`, `hours_48`, `hours_72`, `week_1`, `month_1`, `months_2`, `months_3` |
| `logs_ratio_threshold` (block) | no | nested block |
| `logs_ratio_threshold.condition_type` | yes | — |
| `logs_ratio_threshold.denominator_alias` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `logs_ratio_threshold.evaluation_delay_ms` | no | range 0–2147483647 |
| `logs_ratio_threshold.group_by_for` | yes | one of: `both_or_unspecified`, `numerator_only`, `denumerator_only` |
| `logs_ratio_threshold.ignore_infinity` | no | API: Should the evaluation ignore infinity value |
| `logs_ratio_threshold.notification_payload_filter` | no | pattern `^[a-zA-Z0-9_.]+$` |
| `logs_ratio_threshold.numerator_alias` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `logs_ratio_threshold.denominator` (block) | yes | nested block |
| `logs_ratio_threshold.denominator.simple_filter` (block) | no | nested block |
| `logs_ratio_threshold.denominator.simple_filter.lucene_query` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `logs_ratio_threshold.denominator.simple_filter.label_filters` (block) | no | nested block |
| `logs_ratio_threshold.denominator.simple_filter.label_filters.severities` | no | one of: `debug_or_unspecified`, `verbose`, `info`, `warning`, `error`, `critical` |
| `logs_ratio_threshold.denominator.simple_filter.label_filters.application_name` (block) | no | nested block |
| `logs_ratio_threshold.denominator.simple_filter.label_filters.subsystem_name` (block) | no | nested block |
| `logs_ratio_threshold.numerator` (block) | yes | nested block |
| `logs_ratio_threshold.numerator.simple_filter` (block) | no | nested block |
| `logs_ratio_threshold.numerator.simple_filter.lucene_query` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `logs_ratio_threshold.numerator.simple_filter.label_filters` (block) | no | nested block |
| `logs_ratio_threshold.numerator.simple_filter.label_filters.severities` | no | one of: `debug_or_unspecified`, `verbose`, `info`, `warning`, `error`, `critical` |
| `logs_ratio_threshold.numerator.simple_filter.label_filters.application_name` (block) | no | nested block |
| `logs_ratio_threshold.numerator.simple_filter.label_filters.subsystem_name` (block) | no | nested block |
| `logs_ratio_threshold.rules` (block) | yes | nested block |
| `logs_ratio_threshold.rules.condition` (block) | yes | nested block |
| `logs_ratio_threshold.rules.condition.threshold` | yes | API: The threshold for the alert condition |
| `logs_ratio_threshold.rules.condition.time_window` (block) | yes | nested block |
| `logs_ratio_threshold.rules.condition.time_window.logs_ratio_time_window_specific_value` | yes | one of: `minutes_5_or_unspecified`, `minutes_10`, `minutes_15`, `minutes_30`, `hour_1`, `hours_2`, `hours_4`, `hours_6`, `hours_12`, `hours_24`, `hours_36` |
| `logs_ratio_threshold.rules.override` (block) | yes | nested block |
| `logs_ratio_threshold.rules.override.priority` | yes | — |
| `logs_ratio_threshold.undetected_values_management` (block) | no | nested block |
| `logs_ratio_threshold.undetected_values_management.auto_retire_timeframe` | yes | one of: `never_or_unspecified`, `minutes_5`, `minutes_10`, `hour_1`, `hours_2`, `hours_6`, `hours_12`, `hours_24` |
| `logs_ratio_threshold.undetected_values_management.trigger_undetected_values` | yes | — |
| `logs_threshold` (block) | no | nested block |
| `logs_threshold.condition_type` | yes | — |
| `logs_threshold.evaluation_delay_ms` | no | range 0–2147483647 |
| `logs_threshold.notification_payload_filter` | no | pattern `^[a-zA-Z0-9_.]+$` |
| `logs_threshold.logs_filter` (block) | no | nested block |
| `logs_threshold.logs_filter.simple_filter` (block) | no | nested block |
| `logs_threshold.logs_filter.simple_filter.lucene_query` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `logs_threshold.logs_filter.simple_filter.label_filters` (block) | no | nested block |
| `logs_threshold.logs_filter.simple_filter.label_filters.severities` | no | one of: `debug_or_unspecified`, `verbose`, `info`, `warning`, `error`, `critical` |
| `logs_threshold.logs_filter.simple_filter.label_filters.application_name` (block) | no | nested block |
| `logs_threshold.logs_filter.simple_filter.label_filters.subsystem_name` (block) | no | nested block |
| `logs_threshold.rules` (block) | yes | nested block |
| `logs_threshold.rules.condition` (block) | yes | nested block |
| `logs_threshold.rules.condition.threshold` | yes | API: The threshold for the alert condition |
| `logs_threshold.rules.condition.time_window` (block) | yes | nested block |
| `logs_threshold.rules.condition.time_window.logs_time_window_specific_value` | yes | one of: `minutes_5_or_unspecified`, `minutes_10`, `minutes_20`, `minutes_15`, `minutes_30`, `hour_1`, `hours_2`, `hours_4`, `hours_6`, `hours_12`, `hours_24`, `hours_36` |
| `logs_threshold.rules.override` (block) | yes | nested block |
| `logs_threshold.rules.override.priority` | yes | — |
| `logs_threshold.undetected_values_management` (block) | no | nested block |
| `logs_threshold.undetected_values_management.auto_retire_timeframe` | yes | one of: `never_or_unspecified`, `minutes_5`, `minutes_10`, `hour_1`, `hours_2`, `hours_6`, `hours_12`, `hours_24` |
| `logs_threshold.undetected_values_management.trigger_undetected_values` | yes | — |
| `logs_time_relative_threshold` (block) | no | nested block |
| `logs_time_relative_threshold.condition_type` | yes | — |
| `logs_time_relative_threshold.evaluation_delay_ms` | no | range 0–2147483647 |
| `logs_time_relative_threshold.ignore_infinity` | no | API: Should the evaluation ignore infinity value |
| `logs_time_relative_threshold.notification_payload_filter` | no | pattern `^[a-zA-Z0-9_.]+$` |
| `logs_time_relative_threshold.logs_filter` (block) | no | nested block |
| `logs_time_relative_threshold.logs_filter.simple_filter` (block) | no | nested block |
| `logs_time_relative_threshold.logs_filter.simple_filter.lucene_query` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `logs_time_relative_threshold.logs_filter.simple_filter.label_filters` (block) | no | nested block |
| `logs_time_relative_threshold.logs_filter.simple_filter.label_filters.severities` | no | one of: `debug_or_unspecified`, `verbose`, `info`, `warning`, `error`, `critical` |
| `logs_time_relative_threshold.logs_filter.simple_filter.label_filters.application_name` (block) | no | nested block |
| `logs_time_relative_threshold.logs_filter.simple_filter.label_filters.subsystem_name` (block) | no | nested block |
| `logs_time_relative_threshold.rules` (block) | yes | nested block |
| `logs_time_relative_threshold.rules.condition` (block) | yes | nested block |
| `logs_time_relative_threshold.rules.condition.compared_to` | yes | one of: `previous_hour_or_unspecified`, `same_hour_yesterday`, `same_hour_last_week`, `yesterday`, `same_day_last_week`, `same_day_last_month` |
| `logs_time_relative_threshold.rules.condition.threshold` | yes | API: The threshold for the alert condition |
| `logs_time_relative_threshold.rules.override` (block) | yes | nested block |
| `logs_time_relative_threshold.rules.override.priority` | yes | — |
| `logs_time_relative_threshold.undetected_values_management` (block) | no | nested block |
| `logs_time_relative_threshold.undetected_values_management.auto_retire_timeframe` | yes | one of: `never_or_unspecified`, `minutes_5`, `minutes_10`, `hour_1`, `hours_2`, `hours_6`, `hours_12`, `hours_24` |
| `logs_time_relative_threshold.undetected_values_management.trigger_undetected_values` | yes | — |
| `logs_unique_count` (block) | no | nested block |
| `logs_unique_count.max_unique_count_per_group_by_key` | no | pattern `^-?\d{1,19}$` |
| `logs_unique_count.notification_payload_filter` | no | pattern `^[a-zA-Z0-9_.]+$` |
| `logs_unique_count.unique_count_keypath` | yes | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `logs_unique_count.logs_filter` (block) | no | nested block |
| `logs_unique_count.logs_filter.simple_filter` (block) | no | nested block |
| `logs_unique_count.logs_filter.simple_filter.lucene_query` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `logs_unique_count.logs_filter.simple_filter.label_filters` (block) | no | nested block |
| `logs_unique_count.logs_filter.simple_filter.label_filters.severities` | no | one of: `debug_or_unspecified`, `verbose`, `info`, `warning`, `error`, `critical` |
| `logs_unique_count.logs_filter.simple_filter.label_filters.application_name` (block) | no | nested block |
| `logs_unique_count.logs_filter.simple_filter.label_filters.subsystem_name` (block) | no | nested block |
| `logs_unique_count.rules` (block) | yes | nested block |
| `logs_unique_count.rules.condition` (block) | yes | nested block |
| `logs_unique_count.rules.condition.max_unique_count` | yes | pattern `^-?\d{1,19}$` |
| `logs_unique_count.rules.condition.time_window` (block) | yes | nested block |
| `logs_unique_count.rules.condition.time_window.logs_unique_value_time_window_specific_value` | yes | one of: `minute_1_or_unspecified`, `minutes_15`, `minutes_20`, `minutes_30`, `hours_1`, `hours_2`, `hours_4`, `hours_6`, `hours_12`, `hours_24`, `minutes_5`, `minutes_10`, `hours_36` |
| `metric_anomaly` (block) | no | nested block |
| `metric_anomaly.condition_type` | yes | — |
| `metric_anomaly.evaluation_delay_ms` | no | range 0–2147483647 |
| `metric_anomaly.anomaly_alert_settings` (block) | no | nested block |
| `metric_anomaly.anomaly_alert_settings.percentage_of_deviation` | no | API: The percentage of deviation from the baseline when the alert is triggered. |
| `metric_anomaly.metric_filter` (block) | yes | nested block |
| `metric_anomaly.metric_filter.promql` | yes | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `metric_anomaly.rules` (block) | yes | nested block |
| `metric_anomaly.rules.condition` (block) | yes | nested block |
| `metric_anomaly.rules.condition.for_over_pct` | no | range 0–4294967295 |
| `metric_anomaly.rules.condition.min_non_null_values_pct` | yes | range 0–4294967295 |
| `metric_anomaly.rules.condition.threshold` | yes | API: The threshold for the alert condition |
| `metric_anomaly.rules.condition.of_the_last` (block) | yes | nested block |
| `metric_anomaly.rules.condition.of_the_last.metric_time_window_dynamic_duration` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `metric_anomaly.rules.condition.of_the_last.metric_time_window_specific_value` | no | one of: `minutes_1_or_unspecified`, `minutes_5`, `minutes_10`, `minutes_15`, `minutes_30`, `hour_1`, `hours_2`, `hours_4`, `hours_6`, `hours_12`, `hours_24`, `minutes_20`, `hours_36` |
| `metric_threshold` (block) | no | nested block |
| `metric_threshold.condition_type` | yes | — |
| `metric_threshold.evaluation_delay_ms` | no | range 0–2147483647 |
| `metric_threshold.metric_filter` (block) | yes | nested block |
| `metric_threshold.metric_filter.promql` | yes | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `metric_threshold.missing_values` (block) | yes | nested block |
| `metric_threshold.missing_values.min_non_null_values_pct` | no | range 0–4294967295 |
| `metric_threshold.missing_values.replace_with_zero` | no | — |
| `metric_threshold.rules` (block) | yes | nested block |
| `metric_threshold.rules.condition` (block) | yes | nested block |
| `metric_threshold.rules.condition.for_over_pct` | yes | range 0–4294967295 |
| `metric_threshold.rules.condition.threshold` | yes | API: The threshold for the alert condition |
| `metric_threshold.rules.condition.of_the_last` (block) | yes | nested block |
| `metric_threshold.rules.condition.of_the_last.metric_time_window_dynamic_duration` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `metric_threshold.rules.condition.of_the_last.metric_time_window_specific_value` | no | one of: `minutes_1_or_unspecified`, `minutes_5`, `minutes_10`, `minutes_15`, `minutes_30`, `hour_1`, `hours_2`, `hours_4`, `hours_6`, `hours_12`, `hours_24`, `minutes_20`, `hours_36` |
| `metric_threshold.rules.override` (block) | yes | nested block |
| `metric_threshold.rules.override.priority` | yes | — |
| `metric_threshold.undetected_values_management` (block) | no | nested block |
| `metric_threshold.undetected_values_management.auto_retire_timeframe` | yes | one of: `never_or_unspecified`, `minutes_5`, `minutes_10`, `hour_1`, `hours_2`, `hours_6`, `hours_12`, `hours_24` |
| `metric_threshold.undetected_values_management.trigger_undetected_values` | yes | — |
| `notification_group` (block) | no | nested block |
| `notification_group.group_by_keys` | no | — |
| `notification_group.webhooks` (block) | no | nested block |
| `notification_group.webhooks.minutes` | no | range 0–59 |
| `notification_group.webhooks.notify_on` | no | one of: `triggered_only`, `triggered_and_resolved` |
| `notification_group.webhooks.integration` (block) | yes | nested block |
| `notification_group.webhooks.integration.integration_id` | no | range 0–4294967295 |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_logs_alert_definition_invalid_deleted` | `deleted` |
| `ibm_logs_alert_definition_invalid_description` | `description` |
| `ibm_logs_alert_definition_invalid_enabled` | `enabled` |
| `ibm_logs_alert_definition_invalid_entity_labels` | `entity_labels` |
| `ibm_logs_alert_definition_invalid_group_by_keys` | `group_by_keys` |
| `ibm_logs_alert_definition_invalid_name` | `name` |
| `ibm_logs_alert_definition_invalid_phantom_mode` | `phantom_mode` |
| `ibm_logs_alert_definition_invalid_priority` | `priority` |
| `ibm_logs_alert_definition_invalid_type` | `type` |
| `ibm_logs_alert_definition_invalid_active_on__day_of_week` | `active_on.day_of_week` |
| `ibm_logs_alert_definition_invalid_active_on__end_time__hours` | `active_on.end_time.hours` |
| `ibm_logs_alert_definition_invalid_active_on__end_time__minutes` | `active_on.end_time.minutes` |
| `ibm_logs_alert_definition_invalid_active_on__start_time__hours` | `active_on.start_time.hours` |
| `ibm_logs_alert_definition_invalid_active_on__start_time__minutes` | `active_on.start_time.minutes` |
| `ibm_logs_alert_definition_invalid_flow__enforce_suppression` | `flow.enforce_suppression` |
| `ibm_logs_alert_definition_invalid_flow__stages__timeframe_ms` | `flow.stages.timeframe_ms` |
| `ibm_logs_alert_definition_invalid_flow__stages__timeframe_type` | `flow.stages.timeframe_type` |
| `ibm_logs_alert_definition_invalid_flow__stages__flow_stages_groups__groups__alerts_op` | `flow.stages.flow_stages_groups.groups.alerts_op` |
| `ibm_logs_alert_definition_invalid_flow__stages__flow_stages_groups__groups__next_op` | `flow.stages.flow_stages_groups.groups.next_op` |
| `ibm_logs_alert_definition_invalid_incidents_settings__minutes` | `incidents_settings.minutes` |
| `ibm_logs_alert_definition_invalid_incidents_settings__notify_on` | `incidents_settings.notify_on` |
| `ibm_logs_alert_definition_invalid_logs_anomaly__condition_type` | `logs_anomaly.condition_type` |
| `ibm_logs_alert_definition_invalid_logs_anomaly__evaluation_delay_ms` | `logs_anomaly.evaluation_delay_ms` |
| `ibm_logs_alert_definition_invalid_logs_anomaly__notification_payload_filter` | `logs_anomaly.notification_payload_filter` |
| `ibm_logs_alert_definition_invalid_logs_anomaly__anomaly_alert_settings__percentage_of_deviation` | `logs_anomaly.anomaly_alert_settings.percentage_of_deviation` |
| `ibm_logs_alert_definition_invalid_logs_anomaly__logs_filter__simple_filter__lucene_query` | `logs_anomaly.logs_filter.simple_filter.lucene_query` |
| `ibm_logs_alert_definition_invalid_logs_anomaly__logs_filter__simple_filter__label_filters__severities` | `logs_anomaly.logs_filter.simple_filter.label_filters.severities` |
| `ibm_logs_alert_definition_invalid_logs_anomaly__rules__condition__minimum_threshold` | `logs_anomaly.rules.condition.minimum_threshold` |
| `ibm_logs_alert_definition_invalid_logs_anomaly__rules__condition__time_window__logs_time_window_specific_value` | `logs_anomaly.rules.condition.time_window.logs_time_window_specific_value` |
| `ibm_logs_alert_definition_invalid_logs_immediate__notification_payload_filter` | `logs_immediate.notification_payload_filter` |
| `ibm_logs_alert_definition_invalid_logs_immediate__logs_filter__simple_filter__lucene_query` | `logs_immediate.logs_filter.simple_filter.lucene_query` |
| `ibm_logs_alert_definition_invalid_logs_immediate__logs_filter__simple_filter__label_filters__severities` | `logs_immediate.logs_filter.simple_filter.label_filters.severities` |
| `ibm_logs_alert_definition_invalid_logs_new_value__notification_payload_filter` | `logs_new_value.notification_payload_filter` |
| `ibm_logs_alert_definition_invalid_logs_new_value__logs_filter__simple_filter__lucene_query` | `logs_new_value.logs_filter.simple_filter.lucene_query` |
| `ibm_logs_alert_definition_invalid_logs_new_value__logs_filter__simple_filter__label_filters__severities` | `logs_new_value.logs_filter.simple_filter.label_filters.severities` |
| `ibm_logs_alert_definition_invalid_logs_new_value__rules__condition__keypath_to_track` | `logs_new_value.rules.condition.keypath_to_track` |
| `ibm_logs_alert_definition_invalid_logs_new_value__rules__condition__time_window__logs_new_value_time_window_specific_value` | `logs_new_value.rules.condition.time_window.logs_new_value_time_window_specific_value` |
| `ibm_logs_alert_definition_invalid_logs_ratio_threshold__condition_type` | `logs_ratio_threshold.condition_type` |
| `ibm_logs_alert_definition_invalid_logs_ratio_threshold__denominator_alias` | `logs_ratio_threshold.denominator_alias` |
| `ibm_logs_alert_definition_invalid_logs_ratio_threshold__evaluation_delay_ms` | `logs_ratio_threshold.evaluation_delay_ms` |
| `ibm_logs_alert_definition_invalid_logs_ratio_threshold__group_by_for` | `logs_ratio_threshold.group_by_for` |
| `ibm_logs_alert_definition_invalid_logs_ratio_threshold__ignore_infinity` | `logs_ratio_threshold.ignore_infinity` |
| `ibm_logs_alert_definition_invalid_logs_ratio_threshold__notification_payload_filter` | `logs_ratio_threshold.notification_payload_filter` |
| `ibm_logs_alert_definition_invalid_logs_ratio_threshold__numerator_alias` | `logs_ratio_threshold.numerator_alias` |
| `ibm_logs_alert_definition_invalid_logs_ratio_threshold__denominator__simple_filter__lucene_query` | `logs_ratio_threshold.denominator.simple_filter.lucene_query` |
| `ibm_logs_alert_definition_invalid_logs_ratio_threshold__denominator__simple_filter__label_filters__severities` | `logs_ratio_threshold.denominator.simple_filter.label_filters.severities` |
| `ibm_logs_alert_definition_invalid_logs_ratio_threshold__numerator__simple_filter__lucene_query` | `logs_ratio_threshold.numerator.simple_filter.lucene_query` |
| `ibm_logs_alert_definition_invalid_logs_ratio_threshold__numerator__simple_filter__label_filters__severities` | `logs_ratio_threshold.numerator.simple_filter.label_filters.severities` |
| `ibm_logs_alert_definition_invalid_logs_ratio_threshold__rules__condition__threshold` | `logs_ratio_threshold.rules.condition.threshold` |
| `ibm_logs_alert_definition_invalid_logs_ratio_threshold__rules__condition__time_window__logs_ratio_time_window_specific_value` | `logs_ratio_threshold.rules.condition.time_window.logs_ratio_time_window_specific_value` |
| `ibm_logs_alert_definition_invalid_logs_ratio_threshold__rules__override__priority` | `logs_ratio_threshold.rules.override.priority` |
| `ibm_logs_alert_definition_invalid_logs_ratio_threshold__undetected_values_management__auto_retire_timeframe` | `logs_ratio_threshold.undetected_values_management.auto_retire_timeframe` |
| `ibm_logs_alert_definition_invalid_logs_ratio_threshold__undetected_values_management__trigger_undetected_values` | `logs_ratio_threshold.undetected_values_management.trigger_undetected_values` |
| `ibm_logs_alert_definition_invalid_logs_threshold__condition_type` | `logs_threshold.condition_type` |
| `ibm_logs_alert_definition_invalid_logs_threshold__evaluation_delay_ms` | `logs_threshold.evaluation_delay_ms` |
| `ibm_logs_alert_definition_invalid_logs_threshold__notification_payload_filter` | `logs_threshold.notification_payload_filter` |
| `ibm_logs_alert_definition_invalid_logs_threshold__logs_filter__simple_filter__lucene_query` | `logs_threshold.logs_filter.simple_filter.lucene_query` |
| `ibm_logs_alert_definition_invalid_logs_threshold__logs_filter__simple_filter__label_filters__severities` | `logs_threshold.logs_filter.simple_filter.label_filters.severities` |
| `ibm_logs_alert_definition_invalid_logs_threshold__rules__condition__threshold` | `logs_threshold.rules.condition.threshold` |
| `ibm_logs_alert_definition_invalid_logs_threshold__rules__condition__time_window__logs_time_window_specific_value` | `logs_threshold.rules.condition.time_window.logs_time_window_specific_value` |
| `ibm_logs_alert_definition_invalid_logs_threshold__rules__override__priority` | `logs_threshold.rules.override.priority` |
| `ibm_logs_alert_definition_invalid_logs_threshold__undetected_values_management__auto_retire_timeframe` | `logs_threshold.undetected_values_management.auto_retire_timeframe` |
| `ibm_logs_alert_definition_invalid_logs_threshold__undetected_values_management__trigger_undetected_values` | `logs_threshold.undetected_values_management.trigger_undetected_values` |
| `ibm_logs_alert_definition_invalid_logs_time_relative_threshold__condition_type` | `logs_time_relative_threshold.condition_type` |
| `ibm_logs_alert_definition_invalid_logs_time_relative_threshold__evaluation_delay_ms` | `logs_time_relative_threshold.evaluation_delay_ms` |
| `ibm_logs_alert_definition_invalid_logs_time_relative_threshold__ignore_infinity` | `logs_time_relative_threshold.ignore_infinity` |
| `ibm_logs_alert_definition_invalid_logs_time_relative_threshold__notification_payload_filter` | `logs_time_relative_threshold.notification_payload_filter` |
| `ibm_logs_alert_definition_invalid_logs_time_relative_threshold__logs_filter__simple_filter__lucene_query` | `logs_time_relative_threshold.logs_filter.simple_filter.lucene_query` |
| `ibm_logs_alert_definition_invalid_logs_time_relative_threshold__logs_filter__simple_filter__label_filters__severities` | `logs_time_relative_threshold.logs_filter.simple_filter.label_filters.severities` |
| `ibm_logs_alert_definition_invalid_logs_time_relative_threshold__rules__condition__compared_to` | `logs_time_relative_threshold.rules.condition.compared_to` |
| `ibm_logs_alert_definition_invalid_logs_time_relative_threshold__rules__condition__threshold` | `logs_time_relative_threshold.rules.condition.threshold` |
| `ibm_logs_alert_definition_invalid_logs_time_relative_threshold__rules__override__priority` | `logs_time_relative_threshold.rules.override.priority` |
| `ibm_logs_alert_definition_invalid_logs_time_relative_threshold__undetected_values_management__auto_retire_timeframe` | `logs_time_relative_threshold.undetected_values_management.auto_retire_timeframe` |
| `ibm_logs_alert_definition_invalid_logs_time_relative_threshold__undetected_values_management__trigger_undetected_values` | `logs_time_relative_threshold.undetected_values_management.trigger_undetected_values` |
| `ibm_logs_alert_definition_invalid_logs_unique_count__max_unique_count_per_group_by_key` | `logs_unique_count.max_unique_count_per_group_by_key` |
| `ibm_logs_alert_definition_invalid_logs_unique_count__notification_payload_filter` | `logs_unique_count.notification_payload_filter` |
| `ibm_logs_alert_definition_invalid_logs_unique_count__unique_count_keypath` | `logs_unique_count.unique_count_keypath` |
| `ibm_logs_alert_definition_invalid_logs_unique_count__logs_filter__simple_filter__lucene_query` | `logs_unique_count.logs_filter.simple_filter.lucene_query` |
| `ibm_logs_alert_definition_invalid_logs_unique_count__logs_filter__simple_filter__label_filters__severities` | `logs_unique_count.logs_filter.simple_filter.label_filters.severities` |
| `ibm_logs_alert_definition_invalid_logs_unique_count__rules__condition__max_unique_count` | `logs_unique_count.rules.condition.max_unique_count` |
| `ibm_logs_alert_definition_invalid_logs_unique_count__rules__condition__time_window__logs_unique_value_time_window_specific_value` | `logs_unique_count.rules.condition.time_window.logs_unique_value_time_window_specific_value` |
| `ibm_logs_alert_definition_invalid_metric_anomaly__condition_type` | `metric_anomaly.condition_type` |
| `ibm_logs_alert_definition_invalid_metric_anomaly__evaluation_delay_ms` | `metric_anomaly.evaluation_delay_ms` |
| `ibm_logs_alert_definition_invalid_metric_anomaly__anomaly_alert_settings__percentage_of_deviation` | `metric_anomaly.anomaly_alert_settings.percentage_of_deviation` |
| `ibm_logs_alert_definition_invalid_metric_anomaly__metric_filter__promql` | `metric_anomaly.metric_filter.promql` |
| `ibm_logs_alert_definition_invalid_metric_anomaly__rules__condition__for_over_pct` | `metric_anomaly.rules.condition.for_over_pct` |
| `ibm_logs_alert_definition_invalid_metric_anomaly__rules__condition__min_non_null_values_pct` | `metric_anomaly.rules.condition.min_non_null_values_pct` |
| `ibm_logs_alert_definition_invalid_metric_anomaly__rules__condition__threshold` | `metric_anomaly.rules.condition.threshold` |
| `ibm_logs_alert_definition_invalid_metric_anomaly__rules__condition__of_the_last__metric_time_window_dynamic_duration` | `metric_anomaly.rules.condition.of_the_last.metric_time_window_dynamic_duration` |
| `ibm_logs_alert_definition_invalid_metric_anomaly__rules__condition__of_the_last__metric_time_window_specific_value` | `metric_anomaly.rules.condition.of_the_last.metric_time_window_specific_value` |
| `ibm_logs_alert_definition_invalid_metric_threshold__condition_type` | `metric_threshold.condition_type` |
| `ibm_logs_alert_definition_invalid_metric_threshold__evaluation_delay_ms` | `metric_threshold.evaluation_delay_ms` |
| `ibm_logs_alert_definition_invalid_metric_threshold__metric_filter__promql` | `metric_threshold.metric_filter.promql` |
| `ibm_logs_alert_definition_invalid_metric_threshold__missing_values__min_non_null_values_pct` | `metric_threshold.missing_values.min_non_null_values_pct` |
| `ibm_logs_alert_definition_invalid_metric_threshold__missing_values__replace_with_zero` | `metric_threshold.missing_values.replace_with_zero` |
| `ibm_logs_alert_definition_invalid_metric_threshold__rules__condition__for_over_pct` | `metric_threshold.rules.condition.for_over_pct` |
| `ibm_logs_alert_definition_invalid_metric_threshold__rules__condition__threshold` | `metric_threshold.rules.condition.threshold` |
| `ibm_logs_alert_definition_invalid_metric_threshold__rules__condition__of_the_last__metric_time_window_dynamic_duration` | `metric_threshold.rules.condition.of_the_last.metric_time_window_dynamic_duration` |
| `ibm_logs_alert_definition_invalid_metric_threshold__rules__condition__of_the_last__metric_time_window_specific_value` | `metric_threshold.rules.condition.of_the_last.metric_time_window_specific_value` |
| `ibm_logs_alert_definition_invalid_metric_threshold__rules__override__priority` | `metric_threshold.rules.override.priority` |
| `ibm_logs_alert_definition_invalid_metric_threshold__undetected_values_management__auto_retire_timeframe` | `metric_threshold.undetected_values_management.auto_retire_timeframe` |
| `ibm_logs_alert_definition_invalid_metric_threshold__undetected_values_management__trigger_undetected_values` | `metric_threshold.undetected_values_management.trigger_undetected_values` |
| `ibm_logs_alert_definition_invalid_notification_group__group_by_keys` | `notification_group.group_by_keys` |
| `ibm_logs_alert_definition_invalid_notification_group__webhooks__minutes` | `notification_group.webhooks.minutes` |
| `ibm_logs_alert_definition_invalid_notification_group__webhooks__notify_on` | `notification_group.webhooks.notify_on` |
| `ibm_logs_alert_definition_invalid_notification_group__webhooks__integration__integration_id` | `notification_group.webhooks.integration.integration_id` |

## Example

### Valid

```hcl
resource "ibm_logs_alert_definition" "test" {
  name = "example-name"
  type = "flow"
}
```
