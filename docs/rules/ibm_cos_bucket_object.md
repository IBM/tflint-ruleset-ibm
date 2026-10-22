# ibm_cos_bucket_object

Schema validation for the ibm_cos_bucket_object resource.

This rule validates `ibm_cos_bucket_object` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_cos_bucket_object" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `body` | no | — |
| `bucket_crn` | yes | crn format |
| `bucket_location` | yes | — |
| `bypass_governance_retention` | no | — |
| `content` | no | — |
| `content_base64` | no | — |
| `content_file` | no | — |
| `content_length` | no | — |
| `content_type` | no | — |
| `endpoint_type` | no | one of: `public`, `private`, `direct` |
| `etag` | no | — |
| `force_delete` | no | — |
| `key` | yes | — |
| `last_modified` | no | — |
| `object_lock_legal_hold_status` | no | — |
| `object_lock_mode` | no | — |
| `object_lock_retain_until_date` | no | — |
| `object_sql_url` | no | — |
| `version_id` | no | — |
| `website_redirect` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cos_bucket_object_invalid_bucket_crn` | `bucket_crn` |
| `ibm_cos_bucket_object_invalid_bucket_location` | `bucket_location` |
| `ibm_cos_bucket_object_invalid_bypass_governance_retention` | `bypass_governance_retention` |
| `ibm_cos_bucket_object_invalid_content` | `content` |
| `ibm_cos_bucket_object_invalid_content_base64` | `content_base64` |
| `ibm_cos_bucket_object_invalid_content_file` | `content_file` |
| `ibm_cos_bucket_object_invalid_endpoint_type` | `endpoint_type` |
| `ibm_cos_bucket_object_invalid_etag` | `etag` |
| `ibm_cos_bucket_object_invalid_force_delete` | `force_delete` |
| `ibm_cos_bucket_object_invalid_key` | `key` |
| `ibm_cos_bucket_object_invalid_object_lock_legal_hold_status` | `object_lock_legal_hold_status` |
| `ibm_cos_bucket_object_invalid_object_lock_mode` | `object_lock_mode` |
| `ibm_cos_bucket_object_invalid_object_lock_retain_until_date` | `object_lock_retain_until_date` |
| `ibm_cos_bucket_object_invalid_website_redirect` | `website_redirect` |

## Example

### Valid

```hcl
resource "ibm_cos_bucket_object" "test" {
  bucket_crn = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
  bucket_location = "example"
  key = "example"
}
```
