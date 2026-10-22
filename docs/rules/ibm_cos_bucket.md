# ibm_cos_bucket

Schema validation for the ibm_cos_bucket resource.

This rule validates `ibm_cos_bucket` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/b/{bucket}:patch` (`cos-cos-configuration`).

## Configuration

```hcl
rule "ibm_cos_bucket" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `allowed_ip` | no | API: Individual IP address or range of IP addresses. The entire array will be overwritten in a `PATCH` operation. |
| `bucket_name` | yes | — |
| `crn` | no | — |
| `cross_region_location` | no | one of: `us`, `eu`, `ap` |
| `endpoint_type` | no | one of: `public`, `private`, `direct` |
| `force_delete` | no | — |
| `hard_quota` | no | API: Maximum bytes for this bucket. |
| `key_protect` | no | — |
| `kms_key_crn` | no | crn format |
| `object_lock` | no | — |
| `region_location` | no | — |
| `resource_instance_id` | yes | pattern `^crn:.+:.+:.+:.+:.+:a\/[0-9a-f]{32}:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}\:\:$` |
| `s3_endpoint_direct` | no | — |
| `s3_endpoint_private` | no | — |
| `s3_endpoint_public` | no | — |
| `satellite_location_id` | no | — |
| `single_site_location` | no | one of: `ams03`, `che01`, `hkg02`, `mel01`, `mex01`, `mil01`, `mon01`, `osl01`, `par01`, `sjc04`, `sao01`, `seo01`, `sng01`, `tor01` |
| `storage_class` | no | one of: `standard`, `vault`, `cold`, `smart`, `flex`, `onerate_active` |
| `abort_incomplete_multipart_upload_days` (block) | no | nested block |
| `abort_incomplete_multipart_upload_days.days_after_initiation` | no | — |
| `abort_incomplete_multipart_upload_days.enable` | yes | — |
| `abort_incomplete_multipart_upload_days.prefix` | no | — |
| `abort_incomplete_multipart_upload_days.rule_id` | no | — |
| `activity_tracking` (block) | no | nested block |
| `activity_tracking.activity_tracker_crn` | no | crn format |
| `activity_tracking.management_events` | no | API: This field only applies if `activity_tracker_crn` is not populated. If set to `true`, all bucket management events will be sent to Activity Tracker Event Routing. |
| `activity_tracking.read_data_events` | no | API: If set to `true`, all object read events (i.e. downloads) will be sent to Activity Tracker Event Routing. |
| `activity_tracking.write_data_events` | no | API: If set to `true`, all object write events (i.e. uploads) will be sent to Activity Tracker Event Routing. |
| `archive_rule` (block) | no | nested block |
| `archive_rule.days` | yes | range 0–365243 |
| `archive_rule.enable` | yes | — |
| `archive_rule.rule_id` | no | — |
| `archive_rule.type` | yes | — |
| `expire_rule` (block) | no | nested block |
| `expire_rule.date` | no | — |
| `expire_rule.days` | no | range 0–365243 |
| `expire_rule.enable` | yes | — |
| `expire_rule.expired_object_delete_marker` | no | — |
| `expire_rule.prefix` | no | — |
| `expire_rule.rule_id` | no | — |
| `metrics_monitoring` (block) | no | nested block |
| `metrics_monitoring.metrics_monitoring_crn` | no | crn format |
| `metrics_monitoring.request_metrics_enabled` | no | API: If set to `true`, all request metrics (i.e. `rest.object.head`) will be sent to the monitoring service. |
| `metrics_monitoring.usage_metrics_enabled` | no | API: If set to `true`, all usage metrics (i.e. `bytes_used`) will be sent to the monitoring service. |
| `noncurrent_version_expiration` (block) | no | nested block |
| `noncurrent_version_expiration.enable` | yes | — |
| `noncurrent_version_expiration.noncurrent_days` | no | — |
| `noncurrent_version_expiration.prefix` | no | — |
| `noncurrent_version_expiration.rule_id` | no | — |
| `object_versioning` (block) | no | nested block |
| `object_versioning.enable` | no | — |
| `retention_rule` (block) | no | nested block |
| `retention_rule.default` | yes | — |
| `retention_rule.maximum` | yes | — |
| `retention_rule.minimum` | yes | — |
| `retention_rule.permanent` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cos_bucket_invalid_allowed_ip` | `allowed_ip` |
| `ibm_cos_bucket_invalid_bucket_name` | `bucket_name` |
| `ibm_cos_bucket_invalid_cross_region_location` | `cross_region_location` |
| `ibm_cos_bucket_invalid_endpoint_type` | `endpoint_type` |
| `ibm_cos_bucket_invalid_force_delete` | `force_delete` |
| `ibm_cos_bucket_invalid_hard_quota` | `hard_quota` |
| `ibm_cos_bucket_invalid_key_protect` | `key_protect` |
| `ibm_cos_bucket_invalid_kms_key_crn` | `kms_key_crn` |
| `ibm_cos_bucket_invalid_object_lock` | `object_lock` |
| `ibm_cos_bucket_invalid_region_location` | `region_location` |
| `ibm_cos_bucket_invalid_resource_instance_id` | `resource_instance_id` |
| `ibm_cos_bucket_invalid_satellite_location_id` | `satellite_location_id` |
| `ibm_cos_bucket_invalid_single_site_location` | `single_site_location` |
| `ibm_cos_bucket_invalid_storage_class` | `storage_class` |
| `ibm_cos_bucket_invalid_abort_incomplete_multipart_upload_days__days_after_initiation` | `abort_incomplete_multipart_upload_days.days_after_initiation` |
| `ibm_cos_bucket_invalid_abort_incomplete_multipart_upload_days__enable` | `abort_incomplete_multipart_upload_days.enable` |
| `ibm_cos_bucket_invalid_abort_incomplete_multipart_upload_days__prefix` | `abort_incomplete_multipart_upload_days.prefix` |
| `ibm_cos_bucket_invalid_abort_incomplete_multipart_upload_days__rule_id` | `abort_incomplete_multipart_upload_days.rule_id` |
| `ibm_cos_bucket_invalid_activity_tracking__activity_tracker_crn` | `activity_tracking.activity_tracker_crn` |
| `ibm_cos_bucket_invalid_activity_tracking__management_events` | `activity_tracking.management_events` |
| `ibm_cos_bucket_invalid_activity_tracking__read_data_events` | `activity_tracking.read_data_events` |
| `ibm_cos_bucket_invalid_activity_tracking__write_data_events` | `activity_tracking.write_data_events` |
| `ibm_cos_bucket_invalid_archive_rule__days` | `archive_rule.days` |
| `ibm_cos_bucket_invalid_archive_rule__enable` | `archive_rule.enable` |
| `ibm_cos_bucket_invalid_archive_rule__rule_id` | `archive_rule.rule_id` |
| `ibm_cos_bucket_invalid_archive_rule__type` | `archive_rule.type` |
| `ibm_cos_bucket_invalid_expire_rule__date` | `expire_rule.date` |
| `ibm_cos_bucket_invalid_expire_rule__days` | `expire_rule.days` |
| `ibm_cos_bucket_invalid_expire_rule__enable` | `expire_rule.enable` |
| `ibm_cos_bucket_invalid_expire_rule__expired_object_delete_marker` | `expire_rule.expired_object_delete_marker` |
| `ibm_cos_bucket_invalid_expire_rule__prefix` | `expire_rule.prefix` |
| `ibm_cos_bucket_invalid_expire_rule__rule_id` | `expire_rule.rule_id` |
| `ibm_cos_bucket_invalid_metrics_monitoring__metrics_monitoring_crn` | `metrics_monitoring.metrics_monitoring_crn` |
| `ibm_cos_bucket_invalid_metrics_monitoring__request_metrics_enabled` | `metrics_monitoring.request_metrics_enabled` |
| `ibm_cos_bucket_invalid_metrics_monitoring__usage_metrics_enabled` | `metrics_monitoring.usage_metrics_enabled` |
| `ibm_cos_bucket_invalid_noncurrent_version_expiration__enable` | `noncurrent_version_expiration.enable` |
| `ibm_cos_bucket_invalid_noncurrent_version_expiration__noncurrent_days` | `noncurrent_version_expiration.noncurrent_days` |
| `ibm_cos_bucket_invalid_noncurrent_version_expiration__prefix` | `noncurrent_version_expiration.prefix` |
| `ibm_cos_bucket_invalid_noncurrent_version_expiration__rule_id` | `noncurrent_version_expiration.rule_id` |
| `ibm_cos_bucket_invalid_object_versioning__enable` | `object_versioning.enable` |
| `ibm_cos_bucket_invalid_retention_rule__default` | `retention_rule.default` |
| `ibm_cos_bucket_invalid_retention_rule__maximum` | `retention_rule.maximum` |
| `ibm_cos_bucket_invalid_retention_rule__minimum` | `retention_rule.minimum` |
| `ibm_cos_bucket_invalid_retention_rule__permanent` | `retention_rule.permanent` |

## Example

### Valid

```hcl
resource "ibm_cos_bucket" "test" {
  bucket_name = "example"
  resource_instance_id = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
}
```
