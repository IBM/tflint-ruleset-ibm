# ibm_pi_image

Schema validation for the ibm_pi_image resource.

This rule validates `ibm_pi_image` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Image` (`power-cloud`).

## Configuration

```hcl
rule "ibm_pi_image" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `crn` | no | — |
| `image_id` | no | — |
| `pi_affinity_instance` | no | — |
| `pi_affinity_policy` | no | one of: `affinity`, `anti-affinity` |
| `pi_affinity_volume` | no | — |
| `pi_anti_affinity_instances` | no | — |
| `pi_anti_affinity_volumes` | no | — |
| `pi_cloud_instance_id` | yes | — |
| `pi_image_access_key` | no | — |
| `pi_image_bucket_access` | no | one of: `public`, `private` |
| `pi_image_bucket_file_name` | no | — |
| `pi_image_bucket_name` | no | — |
| `pi_image_bucket_region` | no | — |
| `pi_image_id` | no | — |
| `pi_image_name` | no | — |
| `pi_image_secret_key` | no | — |
| `pi_image_storage_pool` | no | — |
| `pi_image_storage_type` | no | — |
| `pi_source_checksum` | no | — |
| `pi_user_tags` | no | — |
| `pi_image_import_details` (block) | no | nested block |
| `pi_image_import_details.license_type` | yes | one of: `byol` |
| `pi_image_import_details.product` | yes | one of: `Hana`, `Netweaver` |
| `pi_image_import_details.vendor` | yes | one of: `SAP` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_image_invalid_pi_affinity_instance` | `pi_affinity_instance` |
| `ibm_pi_image_invalid_pi_affinity_policy` | `pi_affinity_policy` |
| `ibm_pi_image_invalid_pi_affinity_volume` | `pi_affinity_volume` |
| `ibm_pi_image_invalid_pi_anti_affinity_instances` | `pi_anti_affinity_instances` |
| `ibm_pi_image_invalid_pi_anti_affinity_volumes` | `pi_anti_affinity_volumes` |
| `ibm_pi_image_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |
| `ibm_pi_image_invalid_pi_image_access_key` | `pi_image_access_key` |
| `ibm_pi_image_invalid_pi_image_bucket_access` | `pi_image_bucket_access` |
| `ibm_pi_image_invalid_pi_image_bucket_file_name` | `pi_image_bucket_file_name` |
| `ibm_pi_image_invalid_pi_image_bucket_name` | `pi_image_bucket_name` |
| `ibm_pi_image_invalid_pi_image_bucket_region` | `pi_image_bucket_region` |
| `ibm_pi_image_invalid_pi_image_id` | `pi_image_id` |
| `ibm_pi_image_invalid_pi_image_name` | `pi_image_name` |
| `ibm_pi_image_invalid_pi_image_secret_key` | `pi_image_secret_key` |
| `ibm_pi_image_invalid_pi_image_storage_pool` | `pi_image_storage_pool` |
| `ibm_pi_image_invalid_pi_image_storage_type` | `pi_image_storage_type` |
| `ibm_pi_image_invalid_pi_source_checksum` | `pi_source_checksum` |
| `ibm_pi_image_invalid_pi_user_tags` | `pi_user_tags` |
| `ibm_pi_image_invalid_pi_image_import_details__license_type` | `pi_image_import_details.license_type` |
| `ibm_pi_image_invalid_pi_image_import_details__product` | `pi_image_import_details.product` |
| `ibm_pi_image_invalid_pi_image_import_details__vendor` | `pi_image_import_details.vendor` |

## Example

### Valid

```hcl
resource "ibm_pi_image" "test" {
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
