# ibm_is_image_export_job_data

Reads ibm_is_image_export_job. The user-defined name for this image export job.

This rule validates `ibm_is_image_export_job` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `InstanceByImageAndCatalogOffering` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_image_export_job_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `completed_at` | no | — |
| `created_at` | no | — |
| `encrypted_data_key` | no | — |
| `format` | no | — |
| `href` | no | — |
| `image` | yes | API: Identifies an image by a unique property. |
| `image_export_job` | yes | — |
| `name` | no | — |
| `resource_type` | no | — |
| `started_at` | no | — |
| `status` | no | — |
| `status_reasons` | no | — |
| `storage_bucket` | no | — |
| `storage_href` | no | — |
| `storage_object` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_image_export_job_data_invalid_image` | `image` |
| `ibm_is_image_export_job_data_invalid_image_export_job` | `image_export_job` |

## Example

### Valid

```hcl
data "ibm_is_image_export_job" "test" {
  image = "example"
  image_export_job = "example"
}
```
