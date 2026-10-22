# ibm_pi_image_export

Schema validation for the ibm_pi_image_export resource.

This rule validates `ibm_pi_image_export` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pi_image_export" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `pi_cloud_instance_id` | yes | — |
| `pi_image_access_key` | yes | — |
| `pi_image_bucket_name` | yes | — |
| `pi_image_bucket_region` | yes | — |
| `pi_image_id` | yes | — |
| `pi_image_secret_key` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_image_export_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |
| `ibm_pi_image_export_invalid_pi_image_access_key` | `pi_image_access_key` |
| `ibm_pi_image_export_invalid_pi_image_bucket_name` | `pi_image_bucket_name` |
| `ibm_pi_image_export_invalid_pi_image_bucket_region` | `pi_image_bucket_region` |
| `ibm_pi_image_export_invalid_pi_image_id` | `pi_image_id` |
| `ibm_pi_image_export_invalid_pi_image_secret_key` | `pi_image_secret_key` |

## Example

### Valid

```hcl
variable "pi_image_access_key" {
  sensitive = true
}

variable "pi_image_secret_key" {
  sensitive = true
}

resource "ibm_pi_image_export" "test" {
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  pi_image_access_key = var.pi_image_access_key
  pi_image_bucket_name = "example"
  pi_image_bucket_region = "example"
  pi_image_id = "01234567-89ab-cdef-0123-456789abcdef"
  pi_image_secret_key = var.pi_image_secret_key
}
```
