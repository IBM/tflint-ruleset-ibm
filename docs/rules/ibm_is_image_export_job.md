# ibm_is_image_export_job

Manages ibm_is_image_export_job. The user-defined name for this image export job.

This rule validates `ibm_is_image_export_job` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ErrorBadRequest` (`vpc-identity-latest`).

## Configuration

```hcl
rule "ibm_is_image_export_job" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `completed_at` | no | — |
| `created_at` | no | — |
| `encrypted_data_key` | no | — |
| `format` | no | one of: `qcow2`, `vhd`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `href` | no | — |
| `image` | yes | — |
| `image_export_job` | no | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `resource_type` | no | — |
| `started_at` | no | — |
| `status` | no | — |
| `status_reasons` | no | — |
| `storage_href` | no | — |
| `storage_object` | no | — |
| `storage_bucket` (block) | yes | nested block |
| `storage_bucket.crn` | no | crn format |
| `storage_bucket.name` | no | API: The name of the problematic field, query parameter, or header |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_image_export_job_invalid_format` | `format` |
| `ibm_is_image_export_job_invalid_image` | `image` |
| `ibm_is_image_export_job_invalid_name` | `name` |
| `ibm_is_image_export_job_invalid_storage_bucket__crn` | `storage_bucket.crn` |
| `ibm_is_image_export_job_invalid_storage_bucket__name` | `storage_bucket.name` |

## Example

### Valid

```hcl
resource "ibm_is_image_export_job" "test" {
  image = "example"
  storage_bucket {
  }
}
```
