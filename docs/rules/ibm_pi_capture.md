# ibm_pi_capture

Schema validation for the ibm_pi_capture resource.

This rule validates `ibm_pi_capture` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pi_capture" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `crn` | no | — |
| `image_id` | no | — |
| `pi_capture_cloud_storage_access_key` | no | — |
| `pi_capture_cloud_storage_region` | no | — |
| `pi_capture_cloud_storage_secret_key` | no | — |
| `pi_capture_destination` | yes | one of: `image-catalog`, `cloud-storage`, `both` |
| `pi_capture_name` | yes | — |
| `pi_capture_storage_image_path` | no | — |
| `pi_capture_volume_ids` | no | — |
| `pi_cloud_instance_id` | yes | — |
| `pi_instance_name` | yes | — |
| `pi_user_tags` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_capture_invalid_pi_capture_cloud_storage_access_key` | `pi_capture_cloud_storage_access_key` |
| `ibm_pi_capture_invalid_pi_capture_cloud_storage_region` | `pi_capture_cloud_storage_region` |
| `ibm_pi_capture_invalid_pi_capture_cloud_storage_secret_key` | `pi_capture_cloud_storage_secret_key` |
| `ibm_pi_capture_invalid_pi_capture_destination` | `pi_capture_destination` |
| `ibm_pi_capture_invalid_pi_capture_name` | `pi_capture_name` |
| `ibm_pi_capture_invalid_pi_capture_storage_image_path` | `pi_capture_storage_image_path` |
| `ibm_pi_capture_invalid_pi_capture_volume_ids` | `pi_capture_volume_ids` |
| `ibm_pi_capture_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |
| `ibm_pi_capture_invalid_pi_instance_name` | `pi_instance_name` |
| `ibm_pi_capture_invalid_pi_user_tags` | `pi_user_tags` |

## Example

### Valid

```hcl
resource "ibm_pi_capture" "test" {
  pi_capture_destination = "image-catalog"
  pi_capture_name = "example"
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  pi_instance_name = "example"
}
```
