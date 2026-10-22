# ibm_cos_bucket_object_lock_configuration

Schema validation for the ibm_cos_bucket_object_lock_configuration resource.

This rule validates `ibm_cos_bucket_object_lock_configuration` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/b/{bucket}:patch` (`cos-cos-configuration`).

## Configuration

```hcl
rule "ibm_cos_bucket_object_lock_configuration" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `bucket_crn` | yes | crn format |
| `bucket_location` | yes | — |
| `endpoint_type` | no | one of: `public`, `private`, `direct` |
| `object_lock_configuration` (block) | yes | nested block |
| `object_lock_configuration.object_lock_enabled` | yes | — |
| `object_lock_configuration.object_lock_rule` (block) | no | nested block |
| `object_lock_configuration.object_lock_rule.default_retention` (block) | no | nested block |
| `object_lock_configuration.object_lock_rule.default_retention.days` | no | range 0–365243 |
| `object_lock_configuration.object_lock_rule.default_retention.mode` | yes | one of: `COMPLIANCE`, `GOVERNANCE` |
| `object_lock_configuration.object_lock_rule.default_retention.years` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cos_bucket_object_lock_configuration_invalid_bucket_crn` | `bucket_crn` |
| `ibm_cos_bucket_object_lock_configuration_invalid_bucket_location` | `bucket_location` |
| `ibm_cos_bucket_object_lock_configuration_invalid_endpoint_type` | `endpoint_type` |
| `ibm_cos_bucket_object_lock_configuration_invalid_object_lock_configuration__object_lock_enabled` | `object_lock_configuration.object_lock_enabled` |
| `ibm_cos_bucket_object_lock_configuration_invalid_object_lock_configuration__object_lock_rule__default_retention__days` | `object_lock_configuration.object_lock_rule.default_retention.days` |
| `ibm_cos_bucket_object_lock_configuration_invalid_object_lock_configuration__object_lock_rule__default_retention__mode` | `object_lock_configuration.object_lock_rule.default_retention.mode` |
| `ibm_cos_bucket_object_lock_configuration_invalid_object_lock_configuration__object_lock_rule__default_retention__years` | `object_lock_configuration.object_lock_rule.default_retention.years` |

## Example

### Valid

```hcl
resource "ibm_cos_bucket_object_lock_configuration" "test" {
  bucket_crn = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
  bucket_location = "example"
  object_lock_configuration {
    object_lock_enabled = "example"
  }
}
```
