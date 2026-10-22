# ibm_logs_e2m

Manages ibm_logs_e2m. Name of the E2M.

This rule validates `ibm_logs_e2m` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ApisLogs2metricsV2LogsQuery` (`logs-service-api`).

## Configuration

```hcl
rule "ibm_logs_e2m" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `create_time` | no | — |
| `description` | no | pattern `^[A-Za-z0-9_\-\s]+$` |
| `e2m_id` | no | — |
| `is_internal` | no | — |
| `name` | yes | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `permutations` | no | — |
| `type` | no | one of: `logs2metrics`, `unspecified` |
| `update_time` | no | — |
| `logs_query` (block) | no | nested block |
| `logs_query.alias` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `logs_query.applicationname_filters` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `logs_query.lucene` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `logs_query.severity_filters` | no | one of: `unspecified`, `debug`, `verbose`, `info`, `warning`, `error`, `critical` |
| `logs_query.subsystemname_filters` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `metric_fields` (block) | no | nested block |
| `metric_fields.source_field` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `metric_fields.target_base_metric_name` | no | pattern `^[\w/-]+$` |
| `metric_fields.aggregations` (block) | no | nested block |
| `metric_fields.aggregations.agg_type` | no | one of: `unspecified`, `min`, `max`, `count`, `avg`, `sum`, `histogram`, `samples` |
| `metric_fields.aggregations.enabled` | no | — |
| `metric_fields.aggregations.target_metric_name` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `metric_fields.aggregations.histogram` (block) | no | nested block |
| `metric_fields.aggregations.histogram.buckets` | no | API: Histogram buckets of the Events to Metrics |
| `metric_fields.aggregations.samples` (block) | no | nested block |
| `metric_fields.aggregations.samples.sample_type` | no | one of: `unspecified`, `min`, `max` |
| `metric_labels` (block) | no | nested block |
| `metric_labels.source_field` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `metric_labels.target_label` | no | pattern `^[\w/-]+$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_logs_e2m_invalid_description` | `description` |
| `ibm_logs_e2m_invalid_name` | `name` |
| `ibm_logs_e2m_invalid_type` | `type` |
| `ibm_logs_e2m_invalid_logs_query__alias` | `logs_query.alias` |
| `ibm_logs_e2m_invalid_logs_query__applicationname_filters` | `logs_query.applicationname_filters` |
| `ibm_logs_e2m_invalid_logs_query__lucene` | `logs_query.lucene` |
| `ibm_logs_e2m_invalid_logs_query__severity_filters` | `logs_query.severity_filters` |
| `ibm_logs_e2m_invalid_logs_query__subsystemname_filters` | `logs_query.subsystemname_filters` |
| `ibm_logs_e2m_invalid_metric_fields__source_field` | `metric_fields.source_field` |
| `ibm_logs_e2m_invalid_metric_fields__target_base_metric_name` | `metric_fields.target_base_metric_name` |
| `ibm_logs_e2m_invalid_metric_fields__aggregations__agg_type` | `metric_fields.aggregations.agg_type` |
| `ibm_logs_e2m_invalid_metric_fields__aggregations__enabled` | `metric_fields.aggregations.enabled` |
| `ibm_logs_e2m_invalid_metric_fields__aggregations__target_metric_name` | `metric_fields.aggregations.target_metric_name` |
| `ibm_logs_e2m_invalid_metric_fields__aggregations__histogram__buckets` | `metric_fields.aggregations.histogram.buckets` |
| `ibm_logs_e2m_invalid_metric_fields__aggregations__samples__sample_type` | `metric_fields.aggregations.samples.sample_type` |
| `ibm_logs_e2m_invalid_metric_labels__source_field` | `metric_labels.source_field` |
| `ibm_logs_e2m_invalid_metric_labels__target_label` | `metric_labels.target_label` |

## Example

### Valid

```hcl
resource "ibm_logs_e2m" "test" {
  name = "example-name"
}
```
