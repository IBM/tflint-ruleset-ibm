# ibm_cos_bucket_data

Schema validation for the ibm_cos_bucket data source.

This rule validates `ibm_cos_bucket` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_cos_bucket_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `abort_incomplete_multipart_upload_days` | no | — |
| `activity_tracking` | no | — |
| `allowed_ip` | no | — |
| `archive_rule` | no | — |
| `bucket_name` | yes | — |
| `bucket_region` | no | — |
| `bucket_type` | no | one of: `single_site_location`, `region_location`, `cross_region_location` |
| `crn` | no | — |
| `cross_region_location` | no | — |
| `endpoint_type` | no | one of: `public`, `private`, `direct` |
| `expire_rule` | no | — |
| `hard_quota` | no | — |
| `key_protect` | no | — |
| `kms_key_crn` | no | — |
| `lifecycle_rule` | no | — |
| `metrics_monitoring` | no | — |
| `noncurrent_version_expiration` | no | — |
| `object_lock` | no | — |
| `object_lock_configuration` | no | — |
| `object_versioning` | no | — |
| `region_location` | no | — |
| `replication_rule` | no | — |
| `resource_instance_id` | yes | pattern `^crn:.+:.+:.+:.+:.+:a\/[0-9a-f]{32}:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}\:\:$` |
| `retention_rule` | no | — |
| `s3_endpoint_direct` | no | — |
| `s3_endpoint_private` | no | — |
| `s3_endpoint_public` | no | — |
| `satellite_location_id` | no | — |
| `single_site_location` | no | — |
| `storage_class` | no | — |
| `website_configuration` | no | — |
| `website_endpoint` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cos_bucket_data_invalid_bucket_name` | `bucket_name` |
| `ibm_cos_bucket_data_invalid_bucket_region` | `bucket_region` |
| `ibm_cos_bucket_data_invalid_bucket_type` | `bucket_type` |
| `ibm_cos_bucket_data_invalid_endpoint_type` | `endpoint_type` |
| `ibm_cos_bucket_data_invalid_resource_instance_id` | `resource_instance_id` |
| `ibm_cos_bucket_data_invalid_satellite_location_id` | `satellite_location_id` |

## Example

### Valid

```hcl
data "ibm_cos_bucket" "test" {
  bucket_name = "example"
  resource_instance_id = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
}
```
