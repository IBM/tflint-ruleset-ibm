# ibm_logs_alert

Manages ibm_logs_alert. Alert name.

This rule validates `ibm_logs_alert` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Alert` (`logs-service-api`).

## Configuration

```hcl
rule "ibm_logs_alert" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `alert_id` | no | — |
| `description` | no | pattern `^[A-Za-z0-9_\-\s]+$` |
| `is_active` | yes | API: Alert is active |
| `meta_labels_strings` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `name` | yes | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `notification_payload_filters` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `severity` | yes | one of: `critical`, `error`, `info_or_unspecified`, `warning` |
| `unique_identifier` | no | — |
| `active_when` (block) | no | nested block |
| `active_when.timeframes` (block) | yes | nested block |
| `active_when.timeframes.days_of_week` | yes | one of: `monday_or_unspecified`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`, `sunday` |
| `active_when.timeframes.range` (block) | yes | nested block |
| `active_when.timeframes.range.end` (block) | yes | nested block |
| `active_when.timeframes.range.end.hours` | no | range 0–24 |
| `active_when.timeframes.range.end.minutes` | no | range 0–59 |
| `active_when.timeframes.range.end.seconds` | no | range 0–59 |
| `active_when.timeframes.range.start` (block) | yes | nested block |
| `active_when.timeframes.range.start.hours` | no | range 0–24 |
| `active_when.timeframes.range.start.minutes` | no | range 0–59 |
| `active_when.timeframes.range.start.seconds` | no | range 0–59 |
| `condition` (block) | yes | nested block |
| `condition.flow` (block) | no | nested block |
| `condition.flow.enforce_suppression` | no | API: Should suppression be enforced on the flow alert |
| `condition.flow.parameters` (block) | no | nested block |
| `condition.flow.parameters.cardinality_fields` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `condition.flow.parameters.group_by` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `condition.flow.parameters.ignore_infinity` | no | API: Should the evaluation ignore infinity value |
| `condition.flow.parameters.relative_timeframe` | no | one of: `hour_or_unspecified`, `day`, `week`, `month` |
| `condition.flow.parameters.threshold` | yes | API: The threshold for the alert condition |
| `condition.flow.parameters.timeframe` | yes | one of: `timeframe_5_min_or_unspecified`, `timeframe_10_min`, `timeframe_20_min`, `timeframe_30_min`, `timeframe_1_h`, `timeframe_2_h`, `timeframe_3_h`, `timeframe_4_h`, `timeframe_6_h`, `timeframe_12_h`, `timeframe_24_h`, `timeframe_48_h`, `timeframe_72_h`, `timeframe_1_w`, `timeframe_1_m`, `timeframe_2_m`, `timeframe_3_m`, `timeframe_15_min`, `timeframe_1_min`, `timeframe_2_min`, `timeframe_36_h` |
| `condition.flow.parameters.metric_alert_parameters` (block) | no | nested block |
| `condition.flow.parameters.metric_alert_parameters.arithmetic_operator` | yes | one of: `avg_or_unspecified`, `min`, `max`, `sum`, `count`, `percentile` |
| `condition.flow.parameters.metric_alert_parameters.arithmetic_operator_modifier` | no | range 0–4294967295 |
| `condition.flow.parameters.metric_alert_parameters.metric_field` | yes | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `condition.flow.parameters.metric_alert_parameters.metric_source` | yes | one of: `logs2metrics_or_unspecified`, `prometheus` |
| `condition.flow.parameters.metric_alert_parameters.non_null_percentage` | no | range 0–4294967295 |
| `condition.flow.parameters.metric_alert_parameters.sample_threshold_percentage` | no | range 0–4294967295 |
| `condition.flow.parameters.metric_alert_parameters.swap_null_values` | no | API: Indicates whether null values are replaced with zero values. |
| `condition.flow.parameters.metric_alert_promql_parameters` (block) | no | nested block |
| `condition.flow.parameters.metric_alert_promql_parameters.arithmetic_operator_modifier` | no | range 0–4294967295 |
| `condition.flow.parameters.metric_alert_promql_parameters.non_null_percentage` | no | range 0–4294967295 |
| `condition.flow.parameters.metric_alert_promql_parameters.promql_text` | yes | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `condition.flow.parameters.metric_alert_promql_parameters.sample_threshold_percentage` | yes | range 0–4294967295 |
| `condition.flow.parameters.metric_alert_promql_parameters.swap_null_values` | no | API: Indicates whether null values are replaced with zero values. |
| `condition.flow.parameters.related_extended_data` (block) | no | nested block |
| `condition.flow.parameters.related_extended_data.cleanup_deadman_duration` | no | one of: `cleanup_deadman_duration_never_or_unspecified`, `cleanup_deadman_duration_5min`, `cleanup_deadman_duration_10min`, `cleanup_deadman_duration_1h`, `cleanup_deadman_duration_2h`, `cleanup_deadman_duration_6h`, `cleanup_deadman_duration_12h`, `cleanup_deadman_duration_24h` |
| `condition.flow.parameters.related_extended_data.should_trigger_deadman` | no | API: Should we trigger deadman |
| `condition.flow.stages` (block) | no | nested block |
| `condition.flow.stages.groups` (block) | no | nested block |
| `condition.flow.stages.groups.next_op` | no | one of: `and`, `or` |
| `condition.flow.stages.groups.alerts` (block) | no | nested block |
| `condition.flow.stages.timeframe` (block) | no | nested block |
| `condition.flow.stages.timeframe.ms` | no | range 0–4294967295 |
| `condition.immediate` (block) | no | nested block |
| `condition.less_than` (block) | no | nested block |
| `condition.less_than.parameters` (block) | yes | nested block |
| `condition.less_than.parameters.cardinality_fields` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `condition.less_than.parameters.group_by` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `condition.less_than.parameters.ignore_infinity` | no | API: Should the evaluation ignore infinity value |
| `condition.less_than.parameters.relative_timeframe` | no | one of: `hour_or_unspecified`, `day`, `week`, `month` |
| `condition.less_than.parameters.threshold` | yes | API: The threshold for the alert condition |
| `condition.less_than.parameters.timeframe` | yes | one of: `timeframe_5_min_or_unspecified`, `timeframe_10_min`, `timeframe_20_min`, `timeframe_30_min`, `timeframe_1_h`, `timeframe_2_h`, `timeframe_3_h`, `timeframe_4_h`, `timeframe_6_h`, `timeframe_12_h`, `timeframe_24_h`, `timeframe_48_h`, `timeframe_72_h`, `timeframe_1_w`, `timeframe_1_m`, `timeframe_2_m`, `timeframe_3_m`, `timeframe_15_min`, `timeframe_1_min`, `timeframe_2_min`, `timeframe_36_h` |
| `condition.less_than.parameters.metric_alert_parameters` (block) | no | nested block |
| `condition.less_than.parameters.metric_alert_parameters.arithmetic_operator` | yes | one of: `avg_or_unspecified`, `min`, `max`, `sum`, `count`, `percentile` |
| `condition.less_than.parameters.metric_alert_parameters.arithmetic_operator_modifier` | no | range 0–4294967295 |
| `condition.less_than.parameters.metric_alert_parameters.metric_field` | yes | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `condition.less_than.parameters.metric_alert_parameters.metric_source` | yes | one of: `logs2metrics_or_unspecified`, `prometheus` |
| `condition.less_than.parameters.metric_alert_parameters.non_null_percentage` | no | range 0–4294967295 |
| `condition.less_than.parameters.metric_alert_parameters.sample_threshold_percentage` | no | range 0–4294967295 |
| `condition.less_than.parameters.metric_alert_parameters.swap_null_values` | no | API: Indicates whether null values are replaced with zero values. |
| `condition.less_than.parameters.metric_alert_promql_parameters` (block) | no | nested block |
| `condition.less_than.parameters.metric_alert_promql_parameters.arithmetic_operator_modifier` | no | range 0–4294967295 |
| `condition.less_than.parameters.metric_alert_promql_parameters.non_null_percentage` | no | range 0–4294967295 |
| `condition.less_than.parameters.metric_alert_promql_parameters.promql_text` | yes | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `condition.less_than.parameters.metric_alert_promql_parameters.sample_threshold_percentage` | yes | range 0–4294967295 |
| `condition.less_than.parameters.metric_alert_promql_parameters.swap_null_values` | no | API: Indicates whether null values are replaced with zero values. |
| `condition.less_than.parameters.related_extended_data` (block) | no | nested block |
| `condition.less_than.parameters.related_extended_data.cleanup_deadman_duration` | no | one of: `cleanup_deadman_duration_never_or_unspecified`, `cleanup_deadman_duration_5min`, `cleanup_deadman_duration_10min`, `cleanup_deadman_duration_1h`, `cleanup_deadman_duration_2h`, `cleanup_deadman_duration_6h`, `cleanup_deadman_duration_12h`, `cleanup_deadman_duration_24h` |
| `condition.less_than.parameters.related_extended_data.should_trigger_deadman` | no | API: Should we trigger deadman |
| `condition.less_than_usual` (block) | no | nested block |
| `condition.less_than_usual.parameters` (block) | yes | nested block |
| `condition.less_than_usual.parameters.cardinality_fields` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `condition.less_than_usual.parameters.group_by` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `condition.less_than_usual.parameters.ignore_infinity` | no | API: Should the evaluation ignore infinity value |
| `condition.less_than_usual.parameters.relative_timeframe` | no | one of: `hour_or_unspecified`, `day`, `week`, `month` |
| `condition.less_than_usual.parameters.threshold` | yes | API: The threshold for the alert condition |
| `condition.less_than_usual.parameters.timeframe` | yes | one of: `timeframe_5_min_or_unspecified`, `timeframe_10_min`, `timeframe_20_min`, `timeframe_30_min`, `timeframe_1_h`, `timeframe_2_h`, `timeframe_3_h`, `timeframe_4_h`, `timeframe_6_h`, `timeframe_12_h`, `timeframe_24_h`, `timeframe_48_h`, `timeframe_72_h`, `timeframe_1_w`, `timeframe_1_m`, `timeframe_2_m`, `timeframe_3_m`, `timeframe_15_min`, `timeframe_1_min`, `timeframe_2_min`, `timeframe_36_h` |
| `condition.less_than_usual.parameters.metric_alert_parameters` (block) | no | nested block |
| `condition.less_than_usual.parameters.metric_alert_parameters.arithmetic_operator` | yes | one of: `avg_or_unspecified`, `min`, `max`, `sum`, `count`, `percentile` |
| `condition.less_than_usual.parameters.metric_alert_parameters.arithmetic_operator_modifier` | no | range 0–4294967295 |
| `condition.less_than_usual.parameters.metric_alert_parameters.metric_field` | yes | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `condition.less_than_usual.parameters.metric_alert_parameters.metric_source` | yes | one of: `logs2metrics_or_unspecified`, `prometheus` |
| `condition.less_than_usual.parameters.metric_alert_parameters.non_null_percentage` | no | range 0–4294967295 |
| `condition.less_than_usual.parameters.metric_alert_parameters.sample_threshold_percentage` | no | range 0–4294967295 |
| `condition.less_than_usual.parameters.metric_alert_parameters.swap_null_values` | no | API: Indicates whether null values are replaced with zero values. |
| `condition.less_than_usual.parameters.metric_alert_promql_parameters` (block) | no | nested block |
| `condition.less_than_usual.parameters.metric_alert_promql_parameters.arithmetic_operator_modifier` | no | range 0–4294967295 |
| `condition.less_than_usual.parameters.metric_alert_promql_parameters.non_null_percentage` | no | range 0–4294967295 |
| `condition.less_than_usual.parameters.metric_alert_promql_parameters.promql_text` | yes | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `condition.less_than_usual.parameters.metric_alert_promql_parameters.sample_threshold_percentage` | yes | range 0–4294967295 |
| `condition.less_than_usual.parameters.metric_alert_promql_parameters.swap_null_values` | no | API: Indicates whether null values are replaced with zero values. |
| `condition.less_than_usual.parameters.related_extended_data` (block) | no | nested block |
| `condition.less_than_usual.parameters.related_extended_data.cleanup_deadman_duration` | no | one of: `cleanup_deadman_duration_never_or_unspecified`, `cleanup_deadman_duration_5min`, `cleanup_deadman_duration_10min`, `cleanup_deadman_duration_1h`, `cleanup_deadman_duration_2h`, `cleanup_deadman_duration_6h`, `cleanup_deadman_duration_12h`, `cleanup_deadman_duration_24h` |
| `condition.less_than_usual.parameters.related_extended_data.should_trigger_deadman` | no | API: Should we trigger deadman |
| `condition.more_than` (block) | no | nested block |
| `condition.more_than.evaluation_window` | no | one of: `rolling_or_unspecified`, `dynamic` |
| `condition.more_than.parameters` (block) | yes | nested block |
| `condition.more_than.parameters.cardinality_fields` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `condition.more_than.parameters.group_by` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `condition.more_than.parameters.ignore_infinity` | no | API: Should the evaluation ignore infinity value |
| `condition.more_than.parameters.relative_timeframe` | no | one of: `hour_or_unspecified`, `day`, `week`, `month` |
| `condition.more_than.parameters.threshold` | yes | API: The threshold for the alert condition |
| `condition.more_than.parameters.timeframe` | yes | one of: `timeframe_5_min_or_unspecified`, `timeframe_10_min`, `timeframe_20_min`, `timeframe_30_min`, `timeframe_1_h`, `timeframe_2_h`, `timeframe_3_h`, `timeframe_4_h`, `timeframe_6_h`, `timeframe_12_h`, `timeframe_24_h`, `timeframe_48_h`, `timeframe_72_h`, `timeframe_1_w`, `timeframe_1_m`, `timeframe_2_m`, `timeframe_3_m`, `timeframe_15_min`, `timeframe_1_min`, `timeframe_2_min`, `timeframe_36_h` |
| `condition.more_than.parameters.metric_alert_parameters` (block) | no | nested block |
| `condition.more_than.parameters.metric_alert_parameters.arithmetic_operator` | yes | one of: `avg_or_unspecified`, `min`, `max`, `sum`, `count`, `percentile` |
| `condition.more_than.parameters.metric_alert_parameters.arithmetic_operator_modifier` | no | range 0–4294967295 |
| `condition.more_than.parameters.metric_alert_parameters.metric_field` | yes | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `condition.more_than.parameters.metric_alert_parameters.metric_source` | yes | one of: `logs2metrics_or_unspecified`, `prometheus` |
| `condition.more_than.parameters.metric_alert_parameters.non_null_percentage` | no | range 0–4294967295 |
| `condition.more_than.parameters.metric_alert_parameters.sample_threshold_percentage` | no | range 0–4294967295 |
| `condition.more_than.parameters.metric_alert_parameters.swap_null_values` | no | API: Indicates whether null values are replaced with zero values. |
| `condition.more_than.parameters.metric_alert_promql_parameters` (block) | no | nested block |
| `condition.more_than.parameters.metric_alert_promql_parameters.arithmetic_operator_modifier` | no | range 0–4294967295 |
| `condition.more_than.parameters.metric_alert_promql_parameters.non_null_percentage` | no | range 0–4294967295 |
| `condition.more_than.parameters.metric_alert_promql_parameters.promql_text` | yes | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `condition.more_than.parameters.metric_alert_promql_parameters.sample_threshold_percentage` | yes | range 0–4294967295 |
| `condition.more_than.parameters.metric_alert_promql_parameters.swap_null_values` | no | API: Indicates whether null values are replaced with zero values. |
| `condition.more_than.parameters.related_extended_data` (block) | no | nested block |
| `condition.more_than.parameters.related_extended_data.cleanup_deadman_duration` | no | one of: `cleanup_deadman_duration_never_or_unspecified`, `cleanup_deadman_duration_5min`, `cleanup_deadman_duration_10min`, `cleanup_deadman_duration_1h`, `cleanup_deadman_duration_2h`, `cleanup_deadman_duration_6h`, `cleanup_deadman_duration_12h`, `cleanup_deadman_duration_24h` |
| `condition.more_than.parameters.related_extended_data.should_trigger_deadman` | no | API: Should we trigger deadman |
| `condition.more_than_usual` (block) | no | nested block |
| `condition.more_than_usual.parameters` (block) | yes | nested block |
| `condition.more_than_usual.parameters.cardinality_fields` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `condition.more_than_usual.parameters.group_by` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `condition.more_than_usual.parameters.ignore_infinity` | no | API: Should the evaluation ignore infinity value |
| `condition.more_than_usual.parameters.relative_timeframe` | no | one of: `hour_or_unspecified`, `day`, `week`, `month` |
| `condition.more_than_usual.parameters.threshold` | yes | API: The threshold for the alert condition |
| `condition.more_than_usual.parameters.timeframe` | yes | one of: `timeframe_5_min_or_unspecified`, `timeframe_10_min`, `timeframe_20_min`, `timeframe_30_min`, `timeframe_1_h`, `timeframe_2_h`, `timeframe_3_h`, `timeframe_4_h`, `timeframe_6_h`, `timeframe_12_h`, `timeframe_24_h`, `timeframe_48_h`, `timeframe_72_h`, `timeframe_1_w`, `timeframe_1_m`, `timeframe_2_m`, `timeframe_3_m`, `timeframe_15_min`, `timeframe_1_min`, `timeframe_2_min`, `timeframe_36_h` |
| `condition.more_than_usual.parameters.metric_alert_parameters` (block) | no | nested block |
| `condition.more_than_usual.parameters.metric_alert_parameters.arithmetic_operator` | yes | one of: `avg_or_unspecified`, `min`, `max`, `sum`, `count`, `percentile` |
| `condition.more_than_usual.parameters.metric_alert_parameters.arithmetic_operator_modifier` | no | range 0–4294967295 |
| `condition.more_than_usual.parameters.metric_alert_parameters.metric_field` | yes | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `condition.more_than_usual.parameters.metric_alert_parameters.metric_source` | yes | one of: `logs2metrics_or_unspecified`, `prometheus` |
| `condition.more_than_usual.parameters.metric_alert_parameters.non_null_percentage` | no | range 0–4294967295 |
| `condition.more_than_usual.parameters.metric_alert_parameters.sample_threshold_percentage` | no | range 0–4294967295 |
| `condition.more_than_usual.parameters.metric_alert_parameters.swap_null_values` | no | API: Indicates whether null values are replaced with zero values. |
| `condition.more_than_usual.parameters.metric_alert_promql_parameters` (block) | no | nested block |
| `condition.more_than_usual.parameters.metric_alert_promql_parameters.arithmetic_operator_modifier` | no | range 0–4294967295 |
| `condition.more_than_usual.parameters.metric_alert_promql_parameters.non_null_percentage` | no | range 0–4294967295 |
| `condition.more_than_usual.parameters.metric_alert_promql_parameters.promql_text` | yes | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `condition.more_than_usual.parameters.metric_alert_promql_parameters.sample_threshold_percentage` | yes | range 0–4294967295 |
| `condition.more_than_usual.parameters.metric_alert_promql_parameters.swap_null_values` | no | API: Indicates whether null values are replaced with zero values. |
| `condition.more_than_usual.parameters.related_extended_data` (block) | no | nested block |
| `condition.more_than_usual.parameters.related_extended_data.cleanup_deadman_duration` | no | one of: `cleanup_deadman_duration_never_or_unspecified`, `cleanup_deadman_duration_5min`, `cleanup_deadman_duration_10min`, `cleanup_deadman_duration_1h`, `cleanup_deadman_duration_2h`, `cleanup_deadman_duration_6h`, `cleanup_deadman_duration_12h`, `cleanup_deadman_duration_24h` |
| `condition.more_than_usual.parameters.related_extended_data.should_trigger_deadman` | no | API: Should we trigger deadman |
| `condition.new_value` (block) | no | nested block |
| `condition.new_value.parameters` (block) | yes | nested block |
| `condition.new_value.parameters.cardinality_fields` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `condition.new_value.parameters.group_by` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `condition.new_value.parameters.ignore_infinity` | no | API: Should the evaluation ignore infinity value |
| `condition.new_value.parameters.relative_timeframe` | no | one of: `hour_or_unspecified`, `day`, `week`, `month` |
| `condition.new_value.parameters.threshold` | yes | API: The threshold for the alert condition |
| `condition.new_value.parameters.timeframe` | yes | one of: `timeframe_5_min_or_unspecified`, `timeframe_10_min`, `timeframe_20_min`, `timeframe_30_min`, `timeframe_1_h`, `timeframe_2_h`, `timeframe_3_h`, `timeframe_4_h`, `timeframe_6_h`, `timeframe_12_h`, `timeframe_24_h`, `timeframe_48_h`, `timeframe_72_h`, `timeframe_1_w`, `timeframe_1_m`, `timeframe_2_m`, `timeframe_3_m`, `timeframe_15_min`, `timeframe_1_min`, `timeframe_2_min`, `timeframe_36_h` |
| `condition.new_value.parameters.metric_alert_parameters` (block) | no | nested block |
| `condition.new_value.parameters.metric_alert_parameters.arithmetic_operator` | yes | one of: `avg_or_unspecified`, `min`, `max`, `sum`, `count`, `percentile` |
| `condition.new_value.parameters.metric_alert_parameters.arithmetic_operator_modifier` | no | range 0–4294967295 |
| `condition.new_value.parameters.metric_alert_parameters.metric_field` | yes | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `condition.new_value.parameters.metric_alert_parameters.metric_source` | yes | one of: `logs2metrics_or_unspecified`, `prometheus` |
| `condition.new_value.parameters.metric_alert_parameters.non_null_percentage` | no | range 0–4294967295 |
| `condition.new_value.parameters.metric_alert_parameters.sample_threshold_percentage` | no | range 0–4294967295 |
| `condition.new_value.parameters.metric_alert_parameters.swap_null_values` | no | API: Indicates whether null values are replaced with zero values. |
| `condition.new_value.parameters.metric_alert_promql_parameters` (block) | no | nested block |
| `condition.new_value.parameters.metric_alert_promql_parameters.arithmetic_operator_modifier` | no | range 0–4294967295 |
| `condition.new_value.parameters.metric_alert_promql_parameters.non_null_percentage` | no | range 0–4294967295 |
| `condition.new_value.parameters.metric_alert_promql_parameters.promql_text` | yes | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `condition.new_value.parameters.metric_alert_promql_parameters.sample_threshold_percentage` | yes | range 0–4294967295 |
| `condition.new_value.parameters.metric_alert_promql_parameters.swap_null_values` | no | API: Indicates whether null values are replaced with zero values. |
| `condition.new_value.parameters.related_extended_data` (block) | no | nested block |
| `condition.new_value.parameters.related_extended_data.cleanup_deadman_duration` | no | one of: `cleanup_deadman_duration_never_or_unspecified`, `cleanup_deadman_duration_5min`, `cleanup_deadman_duration_10min`, `cleanup_deadman_duration_1h`, `cleanup_deadman_duration_2h`, `cleanup_deadman_duration_6h`, `cleanup_deadman_duration_12h`, `cleanup_deadman_duration_24h` |
| `condition.new_value.parameters.related_extended_data.should_trigger_deadman` | no | API: Should we trigger deadman |
| `condition.unique_count` (block) | no | nested block |
| `condition.unique_count.parameters` (block) | yes | nested block |
| `condition.unique_count.parameters.cardinality_fields` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `condition.unique_count.parameters.group_by` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `condition.unique_count.parameters.ignore_infinity` | no | API: Should the evaluation ignore infinity value |
| `condition.unique_count.parameters.relative_timeframe` | no | one of: `hour_or_unspecified`, `day`, `week`, `month` |
| `condition.unique_count.parameters.threshold` | yes | API: The threshold for the alert condition |
| `condition.unique_count.parameters.timeframe` | yes | one of: `timeframe_5_min_or_unspecified`, `timeframe_10_min`, `timeframe_20_min`, `timeframe_30_min`, `timeframe_1_h`, `timeframe_2_h`, `timeframe_3_h`, `timeframe_4_h`, `timeframe_6_h`, `timeframe_12_h`, `timeframe_24_h`, `timeframe_48_h`, `timeframe_72_h`, `timeframe_1_w`, `timeframe_1_m`, `timeframe_2_m`, `timeframe_3_m`, `timeframe_15_min`, `timeframe_1_min`, `timeframe_2_min`, `timeframe_36_h` |
| `condition.unique_count.parameters.metric_alert_parameters` (block) | no | nested block |
| `condition.unique_count.parameters.metric_alert_parameters.arithmetic_operator` | yes | one of: `avg_or_unspecified`, `min`, `max`, `sum`, `count`, `percentile` |
| `condition.unique_count.parameters.metric_alert_parameters.arithmetic_operator_modifier` | no | range 0–4294967295 |
| `condition.unique_count.parameters.metric_alert_parameters.metric_field` | yes | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `condition.unique_count.parameters.metric_alert_parameters.metric_source` | yes | one of: `logs2metrics_or_unspecified`, `prometheus` |
| `condition.unique_count.parameters.metric_alert_parameters.non_null_percentage` | no | range 0–4294967295 |
| `condition.unique_count.parameters.metric_alert_parameters.sample_threshold_percentage` | no | range 0–4294967295 |
| `condition.unique_count.parameters.metric_alert_parameters.swap_null_values` | no | API: Indicates whether null values are replaced with zero values. |
| `condition.unique_count.parameters.metric_alert_promql_parameters` (block) | no | nested block |
| `condition.unique_count.parameters.metric_alert_promql_parameters.arithmetic_operator_modifier` | no | range 0–4294967295 |
| `condition.unique_count.parameters.metric_alert_promql_parameters.non_null_percentage` | no | range 0–4294967295 |
| `condition.unique_count.parameters.metric_alert_promql_parameters.promql_text` | yes | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `condition.unique_count.parameters.metric_alert_promql_parameters.sample_threshold_percentage` | yes | range 0–4294967295 |
| `condition.unique_count.parameters.metric_alert_promql_parameters.swap_null_values` | no | API: Indicates whether null values are replaced with zero values. |
| `condition.unique_count.parameters.related_extended_data` (block) | no | nested block |
| `condition.unique_count.parameters.related_extended_data.cleanup_deadman_duration` | no | one of: `cleanup_deadman_duration_never_or_unspecified`, `cleanup_deadman_duration_5min`, `cleanup_deadman_duration_10min`, `cleanup_deadman_duration_1h`, `cleanup_deadman_duration_2h`, `cleanup_deadman_duration_6h`, `cleanup_deadman_duration_12h`, `cleanup_deadman_duration_24h` |
| `condition.unique_count.parameters.related_extended_data.should_trigger_deadman` | no | API: Should we trigger deadman |
| `expiration` (block) | no | nested block |
| `expiration.day` | no | range 1–31 |
| `expiration.month` | no | range 1–12 |
| `expiration.year` | no | range 0–2147483647 |
| `filters` (block) | yes | nested block |
| `filters.alias` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `filters.filter_type` | no | one of: `text_or_unspecified`, `template`, `ratio`, `unique_count`, `time_relative`, `metric`, `flow` |
| `filters.severities` | no | one of: `debug_or_unspecified`, `verbose`, `info`, `warning`, `error`, `critical` |
| `filters.text` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `filters.metadata` (block) | no | nested block |
| `filters.metadata.applications` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `filters.metadata.subsystems` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `filters.ratio_alerts` (block) | no | nested block |
| `filters.ratio_alerts.alias` | yes | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `filters.ratio_alerts.applications` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `filters.ratio_alerts.group_by` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `filters.ratio_alerts.severities` | no | one of: `debug_or_unspecified`, `verbose`, `info`, `warning`, `error`, `critical` |
| `filters.ratio_alerts.subsystems` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `filters.ratio_alerts.text` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `incident_settings` (block) | no | nested block |
| `incident_settings.notify_on` | no | one of: `triggered_only`, `triggered_and_resolved` |
| `incident_settings.retriggering_period_seconds` | no | range 0–4294967295 |
| `incident_settings.use_as_notification_settings` | no | API: Use these settings for all notificaion webhook |
| `meta_labels` (block) | no | nested block |
| `meta_labels.key` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `meta_labels.value` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `notification_groups` (block) | no | nested block |
| `notification_groups.group_by_fields` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `notification_groups.notifications` (block) | no | nested block |
| `notification_groups.notifications.integration_id` | no | range 0–4294967295 |
| `notification_groups.notifications.notify_on` | no | one of: `triggered_only`, `triggered_and_resolved` |
| `notification_groups.notifications.retriggering_period_seconds` | no | range 0–4294967295 |
| `notification_groups.notifications.recipients` (block) | no | nested block |
| `notification_groups.notifications.recipients.emails` | no | pattern `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_logs_alert_invalid_description` | `description` |
| `ibm_logs_alert_invalid_is_active` | `is_active` |
| `ibm_logs_alert_invalid_meta_labels_strings` | `meta_labels_strings` |
| `ibm_logs_alert_invalid_name` | `name` |
| `ibm_logs_alert_invalid_notification_payload_filters` | `notification_payload_filters` |
| `ibm_logs_alert_invalid_severity` | `severity` |
| `ibm_logs_alert_invalid_active_when__timeframes__days_of_week` | `active_when.timeframes.days_of_week` |
| `ibm_logs_alert_invalid_active_when__timeframes__range__end__hours` | `active_when.timeframes.range.end.hours` |
| `ibm_logs_alert_invalid_active_when__timeframes__range__end__minutes` | `active_when.timeframes.range.end.minutes` |
| `ibm_logs_alert_invalid_active_when__timeframes__range__end__seconds` | `active_when.timeframes.range.end.seconds` |
| `ibm_logs_alert_invalid_active_when__timeframes__range__start__hours` | `active_when.timeframes.range.start.hours` |
| `ibm_logs_alert_invalid_active_when__timeframes__range__start__minutes` | `active_when.timeframes.range.start.minutes` |
| `ibm_logs_alert_invalid_active_when__timeframes__range__start__seconds` | `active_when.timeframes.range.start.seconds` |
| `ibm_logs_alert_invalid_condition__flow__enforce_suppression` | `condition.flow.enforce_suppression` |
| `ibm_logs_alert_invalid_condition__flow__parameters__cardinality_fields` | `condition.flow.parameters.cardinality_fields` |
| `ibm_logs_alert_invalid_condition__flow__parameters__group_by` | `condition.flow.parameters.group_by` |
| `ibm_logs_alert_invalid_condition__flow__parameters__ignore_infinity` | `condition.flow.parameters.ignore_infinity` |
| `ibm_logs_alert_invalid_condition__flow__parameters__relative_timeframe` | `condition.flow.parameters.relative_timeframe` |
| `ibm_logs_alert_invalid_condition__flow__parameters__threshold` | `condition.flow.parameters.threshold` |
| `ibm_logs_alert_invalid_condition__flow__parameters__timeframe` | `condition.flow.parameters.timeframe` |
| `ibm_logs_alert_invalid_condition__flow__parameters__metric_alert_parameters__arithmetic_operator` | `condition.flow.parameters.metric_alert_parameters.arithmetic_operator` |
| `ibm_logs_alert_invalid_condition__flow__parameters__metric_alert_parameters__arithmetic_operator_modifier` | `condition.flow.parameters.metric_alert_parameters.arithmetic_operator_modifier` |
| `ibm_logs_alert_invalid_condition__flow__parameters__metric_alert_parameters__metric_field` | `condition.flow.parameters.metric_alert_parameters.metric_field` |
| `ibm_logs_alert_invalid_condition__flow__parameters__metric_alert_parameters__metric_source` | `condition.flow.parameters.metric_alert_parameters.metric_source` |
| `ibm_logs_alert_invalid_condition__flow__parameters__metric_alert_parameters__non_null_percentage` | `condition.flow.parameters.metric_alert_parameters.non_null_percentage` |
| `ibm_logs_alert_invalid_condition__flow__parameters__metric_alert_parameters__sample_threshold_percentage` | `condition.flow.parameters.metric_alert_parameters.sample_threshold_percentage` |
| `ibm_logs_alert_invalid_condition__flow__parameters__metric_alert_parameters__swap_null_values` | `condition.flow.parameters.metric_alert_parameters.swap_null_values` |
| `ibm_logs_alert_invalid_condition__flow__parameters__metric_alert_promql_parameters__arithmetic_operator_modifier` | `condition.flow.parameters.metric_alert_promql_parameters.arithmetic_operator_modifier` |
| `ibm_logs_alert_invalid_condition__flow__parameters__metric_alert_promql_parameters__non_null_percentage` | `condition.flow.parameters.metric_alert_promql_parameters.non_null_percentage` |
| `ibm_logs_alert_invalid_condition__flow__parameters__metric_alert_promql_parameters__promql_text` | `condition.flow.parameters.metric_alert_promql_parameters.promql_text` |
| `ibm_logs_alert_invalid_condition__flow__parameters__metric_alert_promql_parameters__sample_threshold_percentage` | `condition.flow.parameters.metric_alert_promql_parameters.sample_threshold_percentage` |
| `ibm_logs_alert_invalid_condition__flow__parameters__metric_alert_promql_parameters__swap_null_values` | `condition.flow.parameters.metric_alert_promql_parameters.swap_null_values` |
| `ibm_logs_alert_invalid_condition__flow__parameters__related_extended_data__cleanup_deadman_duration` | `condition.flow.parameters.related_extended_data.cleanup_deadman_duration` |
| `ibm_logs_alert_invalid_condition__flow__parameters__related_extended_data__should_trigger_deadman` | `condition.flow.parameters.related_extended_data.should_trigger_deadman` |
| `ibm_logs_alert_invalid_condition__flow__stages__groups__next_op` | `condition.flow.stages.groups.next_op` |
| `ibm_logs_alert_invalid_condition__flow__stages__timeframe__ms` | `condition.flow.stages.timeframe.ms` |
| `ibm_logs_alert_invalid_condition__less_than__parameters__cardinality_fields` | `condition.less_than.parameters.cardinality_fields` |
| `ibm_logs_alert_invalid_condition__less_than__parameters__group_by` | `condition.less_than.parameters.group_by` |
| `ibm_logs_alert_invalid_condition__less_than__parameters__ignore_infinity` | `condition.less_than.parameters.ignore_infinity` |
| `ibm_logs_alert_invalid_condition__less_than__parameters__relative_timeframe` | `condition.less_than.parameters.relative_timeframe` |
| `ibm_logs_alert_invalid_condition__less_than__parameters__threshold` | `condition.less_than.parameters.threshold` |
| `ibm_logs_alert_invalid_condition__less_than__parameters__timeframe` | `condition.less_than.parameters.timeframe` |
| `ibm_logs_alert_invalid_condition__less_than__parameters__metric_alert_parameters__arithmetic_operator` | `condition.less_than.parameters.metric_alert_parameters.arithmetic_operator` |
| `ibm_logs_alert_invalid_condition__less_than__parameters__metric_alert_parameters__arithmetic_operator_modifier` | `condition.less_than.parameters.metric_alert_parameters.arithmetic_operator_modifier` |
| `ibm_logs_alert_invalid_condition__less_than__parameters__metric_alert_parameters__metric_field` | `condition.less_than.parameters.metric_alert_parameters.metric_field` |
| `ibm_logs_alert_invalid_condition__less_than__parameters__metric_alert_parameters__metric_source` | `condition.less_than.parameters.metric_alert_parameters.metric_source` |
| `ibm_logs_alert_invalid_condition__less_than__parameters__metric_alert_parameters__non_null_percentage` | `condition.less_than.parameters.metric_alert_parameters.non_null_percentage` |
| `ibm_logs_alert_invalid_condition__less_than__parameters__metric_alert_parameters__sample_threshold_percentage` | `condition.less_than.parameters.metric_alert_parameters.sample_threshold_percentage` |
| `ibm_logs_alert_invalid_condition__less_than__parameters__metric_alert_parameters__swap_null_values` | `condition.less_than.parameters.metric_alert_parameters.swap_null_values` |
| `ibm_logs_alert_invalid_condition__less_than__parameters__metric_alert_promql_parameters__arithmetic_operator_modifier` | `condition.less_than.parameters.metric_alert_promql_parameters.arithmetic_operator_modifier` |
| `ibm_logs_alert_invalid_condition__less_than__parameters__metric_alert_promql_parameters__non_null_percentage` | `condition.less_than.parameters.metric_alert_promql_parameters.non_null_percentage` |
| `ibm_logs_alert_invalid_condition__less_than__parameters__metric_alert_promql_parameters__promql_text` | `condition.less_than.parameters.metric_alert_promql_parameters.promql_text` |
| `ibm_logs_alert_invalid_condition__less_than__parameters__metric_alert_promql_parameters__sample_threshold_percentage` | `condition.less_than.parameters.metric_alert_promql_parameters.sample_threshold_percentage` |
| `ibm_logs_alert_invalid_condition__less_than__parameters__metric_alert_promql_parameters__swap_null_values` | `condition.less_than.parameters.metric_alert_promql_parameters.swap_null_values` |
| `ibm_logs_alert_invalid_condition__less_than__parameters__related_extended_data__cleanup_deadman_duration` | `condition.less_than.parameters.related_extended_data.cleanup_deadman_duration` |
| `ibm_logs_alert_invalid_condition__less_than__parameters__related_extended_data__should_trigger_deadman` | `condition.less_than.parameters.related_extended_data.should_trigger_deadman` |
| `ibm_logs_alert_invalid_condition__less_than_usual__parameters__cardinality_fields` | `condition.less_than_usual.parameters.cardinality_fields` |
| `ibm_logs_alert_invalid_condition__less_than_usual__parameters__group_by` | `condition.less_than_usual.parameters.group_by` |
| `ibm_logs_alert_invalid_condition__less_than_usual__parameters__ignore_infinity` | `condition.less_than_usual.parameters.ignore_infinity` |
| `ibm_logs_alert_invalid_condition__less_than_usual__parameters__relative_timeframe` | `condition.less_than_usual.parameters.relative_timeframe` |
| `ibm_logs_alert_invalid_condition__less_than_usual__parameters__threshold` | `condition.less_than_usual.parameters.threshold` |
| `ibm_logs_alert_invalid_condition__less_than_usual__parameters__timeframe` | `condition.less_than_usual.parameters.timeframe` |
| `ibm_logs_alert_invalid_condition__less_than_usual__parameters__metric_alert_parameters__arithmetic_operator` | `condition.less_than_usual.parameters.metric_alert_parameters.arithmetic_operator` |
| `ibm_logs_alert_invalid_condition__less_than_usual__parameters__metric_alert_parameters__arithmetic_operator_modifier` | `condition.less_than_usual.parameters.metric_alert_parameters.arithmetic_operator_modifier` |
| `ibm_logs_alert_invalid_condition__less_than_usual__parameters__metric_alert_parameters__metric_field` | `condition.less_than_usual.parameters.metric_alert_parameters.metric_field` |
| `ibm_logs_alert_invalid_condition__less_than_usual__parameters__metric_alert_parameters__metric_source` | `condition.less_than_usual.parameters.metric_alert_parameters.metric_source` |
| `ibm_logs_alert_invalid_condition__less_than_usual__parameters__metric_alert_parameters__non_null_percentage` | `condition.less_than_usual.parameters.metric_alert_parameters.non_null_percentage` |
| `ibm_logs_alert_invalid_condition__less_than_usual__parameters__metric_alert_parameters__sample_threshold_percentage` | `condition.less_than_usual.parameters.metric_alert_parameters.sample_threshold_percentage` |
| `ibm_logs_alert_invalid_condition__less_than_usual__parameters__metric_alert_parameters__swap_null_values` | `condition.less_than_usual.parameters.metric_alert_parameters.swap_null_values` |
| `ibm_logs_alert_invalid_condition__less_than_usual__parameters__metric_alert_promql_parameters__arithmetic_operator_modifier` | `condition.less_than_usual.parameters.metric_alert_promql_parameters.arithmetic_operator_modifier` |
| `ibm_logs_alert_invalid_condition__less_than_usual__parameters__metric_alert_promql_parameters__non_null_percentage` | `condition.less_than_usual.parameters.metric_alert_promql_parameters.non_null_percentage` |
| `ibm_logs_alert_invalid_condition__less_than_usual__parameters__metric_alert_promql_parameters__promql_text` | `condition.less_than_usual.parameters.metric_alert_promql_parameters.promql_text` |
| `ibm_logs_alert_invalid_condition__less_than_usual__parameters__metric_alert_promql_parameters__sample_threshold_percentage` | `condition.less_than_usual.parameters.metric_alert_promql_parameters.sample_threshold_percentage` |
| `ibm_logs_alert_invalid_condition__less_than_usual__parameters__metric_alert_promql_parameters__swap_null_values` | `condition.less_than_usual.parameters.metric_alert_promql_parameters.swap_null_values` |
| `ibm_logs_alert_invalid_condition__less_than_usual__parameters__related_extended_data__cleanup_deadman_duration` | `condition.less_than_usual.parameters.related_extended_data.cleanup_deadman_duration` |
| `ibm_logs_alert_invalid_condition__less_than_usual__parameters__related_extended_data__should_trigger_deadman` | `condition.less_than_usual.parameters.related_extended_data.should_trigger_deadman` |
| `ibm_logs_alert_invalid_condition__more_than__evaluation_window` | `condition.more_than.evaluation_window` |
| `ibm_logs_alert_invalid_condition__more_than__parameters__cardinality_fields` | `condition.more_than.parameters.cardinality_fields` |
| `ibm_logs_alert_invalid_condition__more_than__parameters__group_by` | `condition.more_than.parameters.group_by` |
| `ibm_logs_alert_invalid_condition__more_than__parameters__ignore_infinity` | `condition.more_than.parameters.ignore_infinity` |
| `ibm_logs_alert_invalid_condition__more_than__parameters__relative_timeframe` | `condition.more_than.parameters.relative_timeframe` |
| `ibm_logs_alert_invalid_condition__more_than__parameters__threshold` | `condition.more_than.parameters.threshold` |
| `ibm_logs_alert_invalid_condition__more_than__parameters__timeframe` | `condition.more_than.parameters.timeframe` |
| `ibm_logs_alert_invalid_condition__more_than__parameters__metric_alert_parameters__arithmetic_operator` | `condition.more_than.parameters.metric_alert_parameters.arithmetic_operator` |
| `ibm_logs_alert_invalid_condition__more_than__parameters__metric_alert_parameters__arithmetic_operator_modifier` | `condition.more_than.parameters.metric_alert_parameters.arithmetic_operator_modifier` |
| `ibm_logs_alert_invalid_condition__more_than__parameters__metric_alert_parameters__metric_field` | `condition.more_than.parameters.metric_alert_parameters.metric_field` |
| `ibm_logs_alert_invalid_condition__more_than__parameters__metric_alert_parameters__metric_source` | `condition.more_than.parameters.metric_alert_parameters.metric_source` |
| `ibm_logs_alert_invalid_condition__more_than__parameters__metric_alert_parameters__non_null_percentage` | `condition.more_than.parameters.metric_alert_parameters.non_null_percentage` |
| `ibm_logs_alert_invalid_condition__more_than__parameters__metric_alert_parameters__sample_threshold_percentage` | `condition.more_than.parameters.metric_alert_parameters.sample_threshold_percentage` |
| `ibm_logs_alert_invalid_condition__more_than__parameters__metric_alert_parameters__swap_null_values` | `condition.more_than.parameters.metric_alert_parameters.swap_null_values` |
| `ibm_logs_alert_invalid_condition__more_than__parameters__metric_alert_promql_parameters__arithmetic_operator_modifier` | `condition.more_than.parameters.metric_alert_promql_parameters.arithmetic_operator_modifier` |
| `ibm_logs_alert_invalid_condition__more_than__parameters__metric_alert_promql_parameters__non_null_percentage` | `condition.more_than.parameters.metric_alert_promql_parameters.non_null_percentage` |
| `ibm_logs_alert_invalid_condition__more_than__parameters__metric_alert_promql_parameters__promql_text` | `condition.more_than.parameters.metric_alert_promql_parameters.promql_text` |
| `ibm_logs_alert_invalid_condition__more_than__parameters__metric_alert_promql_parameters__sample_threshold_percentage` | `condition.more_than.parameters.metric_alert_promql_parameters.sample_threshold_percentage` |
| `ibm_logs_alert_invalid_condition__more_than__parameters__metric_alert_promql_parameters__swap_null_values` | `condition.more_than.parameters.metric_alert_promql_parameters.swap_null_values` |
| `ibm_logs_alert_invalid_condition__more_than__parameters__related_extended_data__cleanup_deadman_duration` | `condition.more_than.parameters.related_extended_data.cleanup_deadman_duration` |
| `ibm_logs_alert_invalid_condition__more_than__parameters__related_extended_data__should_trigger_deadman` | `condition.more_than.parameters.related_extended_data.should_trigger_deadman` |
| `ibm_logs_alert_invalid_condition__more_than_usual__parameters__cardinality_fields` | `condition.more_than_usual.parameters.cardinality_fields` |
| `ibm_logs_alert_invalid_condition__more_than_usual__parameters__group_by` | `condition.more_than_usual.parameters.group_by` |
| `ibm_logs_alert_invalid_condition__more_than_usual__parameters__ignore_infinity` | `condition.more_than_usual.parameters.ignore_infinity` |
| `ibm_logs_alert_invalid_condition__more_than_usual__parameters__relative_timeframe` | `condition.more_than_usual.parameters.relative_timeframe` |
| `ibm_logs_alert_invalid_condition__more_than_usual__parameters__threshold` | `condition.more_than_usual.parameters.threshold` |
| `ibm_logs_alert_invalid_condition__more_than_usual__parameters__timeframe` | `condition.more_than_usual.parameters.timeframe` |
| `ibm_logs_alert_invalid_condition__more_than_usual__parameters__metric_alert_parameters__arithmetic_operator` | `condition.more_than_usual.parameters.metric_alert_parameters.arithmetic_operator` |
| `ibm_logs_alert_invalid_condition__more_than_usual__parameters__metric_alert_parameters__arithmetic_operator_modifier` | `condition.more_than_usual.parameters.metric_alert_parameters.arithmetic_operator_modifier` |
| `ibm_logs_alert_invalid_condition__more_than_usual__parameters__metric_alert_parameters__metric_field` | `condition.more_than_usual.parameters.metric_alert_parameters.metric_field` |
| `ibm_logs_alert_invalid_condition__more_than_usual__parameters__metric_alert_parameters__metric_source` | `condition.more_than_usual.parameters.metric_alert_parameters.metric_source` |
| `ibm_logs_alert_invalid_condition__more_than_usual__parameters__metric_alert_parameters__non_null_percentage` | `condition.more_than_usual.parameters.metric_alert_parameters.non_null_percentage` |
| `ibm_logs_alert_invalid_condition__more_than_usual__parameters__metric_alert_parameters__sample_threshold_percentage` | `condition.more_than_usual.parameters.metric_alert_parameters.sample_threshold_percentage` |
| `ibm_logs_alert_invalid_condition__more_than_usual__parameters__metric_alert_parameters__swap_null_values` | `condition.more_than_usual.parameters.metric_alert_parameters.swap_null_values` |
| `ibm_logs_alert_invalid_condition__more_than_usual__parameters__metric_alert_promql_parameters__arithmetic_operator_modifier` | `condition.more_than_usual.parameters.metric_alert_promql_parameters.arithmetic_operator_modifier` |
| `ibm_logs_alert_invalid_condition__more_than_usual__parameters__metric_alert_promql_parameters__non_null_percentage` | `condition.more_than_usual.parameters.metric_alert_promql_parameters.non_null_percentage` |
| `ibm_logs_alert_invalid_condition__more_than_usual__parameters__metric_alert_promql_parameters__promql_text` | `condition.more_than_usual.parameters.metric_alert_promql_parameters.promql_text` |
| `ibm_logs_alert_invalid_condition__more_than_usual__parameters__metric_alert_promql_parameters__sample_threshold_percentage` | `condition.more_than_usual.parameters.metric_alert_promql_parameters.sample_threshold_percentage` |
| `ibm_logs_alert_invalid_condition__more_than_usual__parameters__metric_alert_promql_parameters__swap_null_values` | `condition.more_than_usual.parameters.metric_alert_promql_parameters.swap_null_values` |
| `ibm_logs_alert_invalid_condition__more_than_usual__parameters__related_extended_data__cleanup_deadman_duration` | `condition.more_than_usual.parameters.related_extended_data.cleanup_deadman_duration` |
| `ibm_logs_alert_invalid_condition__more_than_usual__parameters__related_extended_data__should_trigger_deadman` | `condition.more_than_usual.parameters.related_extended_data.should_trigger_deadman` |
| `ibm_logs_alert_invalid_condition__new_value__parameters__cardinality_fields` | `condition.new_value.parameters.cardinality_fields` |
| `ibm_logs_alert_invalid_condition__new_value__parameters__group_by` | `condition.new_value.parameters.group_by` |
| `ibm_logs_alert_invalid_condition__new_value__parameters__ignore_infinity` | `condition.new_value.parameters.ignore_infinity` |
| `ibm_logs_alert_invalid_condition__new_value__parameters__relative_timeframe` | `condition.new_value.parameters.relative_timeframe` |
| `ibm_logs_alert_invalid_condition__new_value__parameters__threshold` | `condition.new_value.parameters.threshold` |
| `ibm_logs_alert_invalid_condition__new_value__parameters__timeframe` | `condition.new_value.parameters.timeframe` |
| `ibm_logs_alert_invalid_condition__new_value__parameters__metric_alert_parameters__arithmetic_operator` | `condition.new_value.parameters.metric_alert_parameters.arithmetic_operator` |
| `ibm_logs_alert_invalid_condition__new_value__parameters__metric_alert_parameters__arithmetic_operator_modifier` | `condition.new_value.parameters.metric_alert_parameters.arithmetic_operator_modifier` |
| `ibm_logs_alert_invalid_condition__new_value__parameters__metric_alert_parameters__metric_field` | `condition.new_value.parameters.metric_alert_parameters.metric_field` |
| `ibm_logs_alert_invalid_condition__new_value__parameters__metric_alert_parameters__metric_source` | `condition.new_value.parameters.metric_alert_parameters.metric_source` |
| `ibm_logs_alert_invalid_condition__new_value__parameters__metric_alert_parameters__non_null_percentage` | `condition.new_value.parameters.metric_alert_parameters.non_null_percentage` |
| `ibm_logs_alert_invalid_condition__new_value__parameters__metric_alert_parameters__sample_threshold_percentage` | `condition.new_value.parameters.metric_alert_parameters.sample_threshold_percentage` |
| `ibm_logs_alert_invalid_condition__new_value__parameters__metric_alert_parameters__swap_null_values` | `condition.new_value.parameters.metric_alert_parameters.swap_null_values` |
| `ibm_logs_alert_invalid_condition__new_value__parameters__metric_alert_promql_parameters__arithmetic_operator_modifier` | `condition.new_value.parameters.metric_alert_promql_parameters.arithmetic_operator_modifier` |
| `ibm_logs_alert_invalid_condition__new_value__parameters__metric_alert_promql_parameters__non_null_percentage` | `condition.new_value.parameters.metric_alert_promql_parameters.non_null_percentage` |
| `ibm_logs_alert_invalid_condition__new_value__parameters__metric_alert_promql_parameters__promql_text` | `condition.new_value.parameters.metric_alert_promql_parameters.promql_text` |
| `ibm_logs_alert_invalid_condition__new_value__parameters__metric_alert_promql_parameters__sample_threshold_percentage` | `condition.new_value.parameters.metric_alert_promql_parameters.sample_threshold_percentage` |
| `ibm_logs_alert_invalid_condition__new_value__parameters__metric_alert_promql_parameters__swap_null_values` | `condition.new_value.parameters.metric_alert_promql_parameters.swap_null_values` |
| `ibm_logs_alert_invalid_condition__new_value__parameters__related_extended_data__cleanup_deadman_duration` | `condition.new_value.parameters.related_extended_data.cleanup_deadman_duration` |
| `ibm_logs_alert_invalid_condition__new_value__parameters__related_extended_data__should_trigger_deadman` | `condition.new_value.parameters.related_extended_data.should_trigger_deadman` |
| `ibm_logs_alert_invalid_condition__unique_count__parameters__cardinality_fields` | `condition.unique_count.parameters.cardinality_fields` |
| `ibm_logs_alert_invalid_condition__unique_count__parameters__group_by` | `condition.unique_count.parameters.group_by` |
| `ibm_logs_alert_invalid_condition__unique_count__parameters__ignore_infinity` | `condition.unique_count.parameters.ignore_infinity` |
| `ibm_logs_alert_invalid_condition__unique_count__parameters__relative_timeframe` | `condition.unique_count.parameters.relative_timeframe` |
| `ibm_logs_alert_invalid_condition__unique_count__parameters__threshold` | `condition.unique_count.parameters.threshold` |
| `ibm_logs_alert_invalid_condition__unique_count__parameters__timeframe` | `condition.unique_count.parameters.timeframe` |
| `ibm_logs_alert_invalid_condition__unique_count__parameters__metric_alert_parameters__arithmetic_operator` | `condition.unique_count.parameters.metric_alert_parameters.arithmetic_operator` |
| `ibm_logs_alert_invalid_condition__unique_count__parameters__metric_alert_parameters__arithmetic_operator_modifier` | `condition.unique_count.parameters.metric_alert_parameters.arithmetic_operator_modifier` |
| `ibm_logs_alert_invalid_condition__unique_count__parameters__metric_alert_parameters__metric_field` | `condition.unique_count.parameters.metric_alert_parameters.metric_field` |
| `ibm_logs_alert_invalid_condition__unique_count__parameters__metric_alert_parameters__metric_source` | `condition.unique_count.parameters.metric_alert_parameters.metric_source` |
| `ibm_logs_alert_invalid_condition__unique_count__parameters__metric_alert_parameters__non_null_percentage` | `condition.unique_count.parameters.metric_alert_parameters.non_null_percentage` |
| `ibm_logs_alert_invalid_condition__unique_count__parameters__metric_alert_parameters__sample_threshold_percentage` | `condition.unique_count.parameters.metric_alert_parameters.sample_threshold_percentage` |
| `ibm_logs_alert_invalid_condition__unique_count__parameters__metric_alert_parameters__swap_null_values` | `condition.unique_count.parameters.metric_alert_parameters.swap_null_values` |
| `ibm_logs_alert_invalid_condition__unique_count__parameters__metric_alert_promql_parameters__arithmetic_operator_modifier` | `condition.unique_count.parameters.metric_alert_promql_parameters.arithmetic_operator_modifier` |
| `ibm_logs_alert_invalid_condition__unique_count__parameters__metric_alert_promql_parameters__non_null_percentage` | `condition.unique_count.parameters.metric_alert_promql_parameters.non_null_percentage` |
| `ibm_logs_alert_invalid_condition__unique_count__parameters__metric_alert_promql_parameters__promql_text` | `condition.unique_count.parameters.metric_alert_promql_parameters.promql_text` |
| `ibm_logs_alert_invalid_condition__unique_count__parameters__metric_alert_promql_parameters__sample_threshold_percentage` | `condition.unique_count.parameters.metric_alert_promql_parameters.sample_threshold_percentage` |
| `ibm_logs_alert_invalid_condition__unique_count__parameters__metric_alert_promql_parameters__swap_null_values` | `condition.unique_count.parameters.metric_alert_promql_parameters.swap_null_values` |
| `ibm_logs_alert_invalid_condition__unique_count__parameters__related_extended_data__cleanup_deadman_duration` | `condition.unique_count.parameters.related_extended_data.cleanup_deadman_duration` |
| `ibm_logs_alert_invalid_condition__unique_count__parameters__related_extended_data__should_trigger_deadman` | `condition.unique_count.parameters.related_extended_data.should_trigger_deadman` |
| `ibm_logs_alert_invalid_expiration__day` | `expiration.day` |
| `ibm_logs_alert_invalid_expiration__month` | `expiration.month` |
| `ibm_logs_alert_invalid_expiration__year` | `expiration.year` |
| `ibm_logs_alert_invalid_filters__alias` | `filters.alias` |
| `ibm_logs_alert_invalid_filters__filter_type` | `filters.filter_type` |
| `ibm_logs_alert_invalid_filters__severities` | `filters.severities` |
| `ibm_logs_alert_invalid_filters__text` | `filters.text` |
| `ibm_logs_alert_invalid_filters__metadata__applications` | `filters.metadata.applications` |
| `ibm_logs_alert_invalid_filters__metadata__subsystems` | `filters.metadata.subsystems` |
| `ibm_logs_alert_invalid_filters__ratio_alerts__alias` | `filters.ratio_alerts.alias` |
| `ibm_logs_alert_invalid_filters__ratio_alerts__applications` | `filters.ratio_alerts.applications` |
| `ibm_logs_alert_invalid_filters__ratio_alerts__group_by` | `filters.ratio_alerts.group_by` |
| `ibm_logs_alert_invalid_filters__ratio_alerts__severities` | `filters.ratio_alerts.severities` |
| `ibm_logs_alert_invalid_filters__ratio_alerts__subsystems` | `filters.ratio_alerts.subsystems` |
| `ibm_logs_alert_invalid_filters__ratio_alerts__text` | `filters.ratio_alerts.text` |
| `ibm_logs_alert_invalid_incident_settings__notify_on` | `incident_settings.notify_on` |
| `ibm_logs_alert_invalid_incident_settings__retriggering_period_seconds` | `incident_settings.retriggering_period_seconds` |
| `ibm_logs_alert_invalid_incident_settings__use_as_notification_settings` | `incident_settings.use_as_notification_settings` |
| `ibm_logs_alert_invalid_meta_labels__key` | `meta_labels.key` |
| `ibm_logs_alert_invalid_meta_labels__value` | `meta_labels.value` |
| `ibm_logs_alert_invalid_notification_groups__group_by_fields` | `notification_groups.group_by_fields` |
| `ibm_logs_alert_invalid_notification_groups__notifications__integration_id` | `notification_groups.notifications.integration_id` |
| `ibm_logs_alert_invalid_notification_groups__notifications__notify_on` | `notification_groups.notifications.notify_on` |
| `ibm_logs_alert_invalid_notification_groups__notifications__retriggering_period_seconds` | `notification_groups.notifications.retriggering_period_seconds` |
| `ibm_logs_alert_invalid_notification_groups__notifications__recipients__emails` | `notification_groups.notifications.recipients.emails` |

## Example

### Valid

```hcl
resource "ibm_logs_alert" "test" {
  is_active = true
  name = "example-name"
  severity = "critical"
  condition {
  }
  filters {
  }
}
```
