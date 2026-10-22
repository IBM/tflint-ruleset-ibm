# ibm_cos_bucket_lifecycle_configuration

Schema validation for the ibm_cos_bucket_lifecycle_configuration resource.

This rule validates `ibm_cos_bucket_lifecycle_configuration` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/b/{bucket}:patch` (`cos-cos-configuration`).

## Configuration

```hcl
rule "ibm_cos_bucket_lifecycle_configuration" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `bucket_crn` | yes | crn format |
| `bucket_location` | yes | — |
| `endpoint_type` | no | one of: `public`, `private`, `direct` |
| `lifecycle_rule` (block) | yes | nested block |
| `lifecycle_rule.rule_id` | yes | — |
| `lifecycle_rule.status` | yes | one of: `enable`, `disable` |
| `lifecycle_rule.abort_incomplete_multipart_upload` (block) | no | nested block |
| `lifecycle_rule.abort_incomplete_multipart_upload.days_after_initiation` | no | — |
| `lifecycle_rule.expiration` (block) | no | nested block |
| `lifecycle_rule.expiration.date` | no | — |
| `lifecycle_rule.expiration.days` | no | range 0–365243 |
| `lifecycle_rule.expiration.expired_object_delete_marker` | no | — |
| `lifecycle_rule.filter` (block) | yes | nested block |
| `lifecycle_rule.filter.object_size_greater_than` | no | min 0 |
| `lifecycle_rule.filter.object_size_less_than` | no | min 1 |
| `lifecycle_rule.filter.prefix` | no | — |
| `lifecycle_rule.filter.and` (block) | no | nested block |
| `lifecycle_rule.filter.and.object_size_greater_than` | no | min 0 |
| `lifecycle_rule.filter.and.object_size_less_than` | no | min 1 |
| `lifecycle_rule.filter.and.prefix` | no | — |
| `lifecycle_rule.filter.and.tags` (block) | no | nested block |
| `lifecycle_rule.filter.and.tags.key` | yes | — |
| `lifecycle_rule.filter.and.tags.value` | yes | — |
| `lifecycle_rule.filter.tag` (block) | no | nested block |
| `lifecycle_rule.filter.tag.key` | yes | — |
| `lifecycle_rule.filter.tag.value` | yes | — |
| `lifecycle_rule.noncurrent_version_expiration` (block) | no | nested block |
| `lifecycle_rule.noncurrent_version_expiration.noncurrent_days` | no | min 1 |
| `lifecycle_rule.transition` (block) | no | nested block |
| `lifecycle_rule.transition.date` | no | — |
| `lifecycle_rule.transition.days` | no | range 0–365243 |
| `lifecycle_rule.transition.storage_class` | yes | one of: `GLACIER`, `ACCELERATED` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cos_bucket_lifecycle_configuration_invalid_bucket_crn` | `bucket_crn` |
| `ibm_cos_bucket_lifecycle_configuration_invalid_bucket_location` | `bucket_location` |
| `ibm_cos_bucket_lifecycle_configuration_invalid_endpoint_type` | `endpoint_type` |
| `ibm_cos_bucket_lifecycle_configuration_invalid_lifecycle_rule__rule_id` | `lifecycle_rule.rule_id` |
| `ibm_cos_bucket_lifecycle_configuration_invalid_lifecycle_rule__status` | `lifecycle_rule.status` |
| `ibm_cos_bucket_lifecycle_configuration_invalid_lifecycle_rule__abort_incomplete_multipart_upload__days_after_initiation` | `lifecycle_rule.abort_incomplete_multipart_upload.days_after_initiation` |
| `ibm_cos_bucket_lifecycle_configuration_invalid_lifecycle_rule__expiration__date` | `lifecycle_rule.expiration.date` |
| `ibm_cos_bucket_lifecycle_configuration_invalid_lifecycle_rule__expiration__days` | `lifecycle_rule.expiration.days` |
| `ibm_cos_bucket_lifecycle_configuration_invalid_lifecycle_rule__expiration__expired_object_delete_marker` | `lifecycle_rule.expiration.expired_object_delete_marker` |
| `ibm_cos_bucket_lifecycle_configuration_invalid_lifecycle_rule__filter__object_size_greater_than` | `lifecycle_rule.filter.object_size_greater_than` |
| `ibm_cos_bucket_lifecycle_configuration_invalid_lifecycle_rule__filter__object_size_less_than` | `lifecycle_rule.filter.object_size_less_than` |
| `ibm_cos_bucket_lifecycle_configuration_invalid_lifecycle_rule__filter__prefix` | `lifecycle_rule.filter.prefix` |
| `ibm_cos_bucket_lifecycle_configuration_invalid_lifecycle_rule__filter__and__object_size_greater_than` | `lifecycle_rule.filter.and.object_size_greater_than` |
| `ibm_cos_bucket_lifecycle_configuration_invalid_lifecycle_rule__filter__and__object_size_less_than` | `lifecycle_rule.filter.and.object_size_less_than` |
| `ibm_cos_bucket_lifecycle_configuration_invalid_lifecycle_rule__filter__and__prefix` | `lifecycle_rule.filter.and.prefix` |
| `ibm_cos_bucket_lifecycle_configuration_invalid_lifecycle_rule__filter__and__tags__key` | `lifecycle_rule.filter.and.tags.key` |
| `ibm_cos_bucket_lifecycle_configuration_invalid_lifecycle_rule__filter__and__tags__value` | `lifecycle_rule.filter.and.tags.value` |
| `ibm_cos_bucket_lifecycle_configuration_invalid_lifecycle_rule__filter__tag__key` | `lifecycle_rule.filter.tag.key` |
| `ibm_cos_bucket_lifecycle_configuration_invalid_lifecycle_rule__filter__tag__value` | `lifecycle_rule.filter.tag.value` |
| `ibm_cos_bucket_lifecycle_configuration_invalid_lifecycle_rule__noncurrent_version_expiration__noncurrent_days` | `lifecycle_rule.noncurrent_version_expiration.noncurrent_days` |
| `ibm_cos_bucket_lifecycle_configuration_invalid_lifecycle_rule__transition__date` | `lifecycle_rule.transition.date` |
| `ibm_cos_bucket_lifecycle_configuration_invalid_lifecycle_rule__transition__days` | `lifecycle_rule.transition.days` |
| `ibm_cos_bucket_lifecycle_configuration_invalid_lifecycle_rule__transition__storage_class` | `lifecycle_rule.transition.storage_class` |

## Example

### Valid

```hcl
resource "ibm_cos_bucket_lifecycle_configuration" "test" {
  bucket_crn = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
  bucket_location = "example"
  lifecycle_rule {
    rule_id = "01234567-89ab-cdef-0123-456789abcdef"
    status = "enable"
    filter {
    }
  }
}
```
