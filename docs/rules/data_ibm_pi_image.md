# ibm_pi_image_data

Reads ibm_pi_image. The name of an image.

This rule validates `ibm_pi_image` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pi_image_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `architecture` | no | — |
| `container_format` | no | — |
| `crn` | no | — |
| `disk_format` | no | — |
| `endianness` | no | — |
| `hypervisor` | no | — |
| `image_type` | no | — |
| `name` | no | — |
| `operating_system` | no | — |
| `pi_cloud_instance_id` | yes | — |
| `pi_image_id` | no | — |
| `pi_image_name` | no | — |
| `shared` | no | — |
| `size` | no | — |
| `source_checksum` | no | — |
| `state` | no | — |
| `storage_pool` | no | — |
| `storage_type` | no | — |
| `user_tags` | no | — |
| `volumes` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_image_data_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |
| `ibm_pi_image_data_invalid_pi_image_id` | `pi_image_id` |
| `ibm_pi_image_data_invalid_pi_image_name` | `pi_image_name` |

## Example

### Valid

```hcl
data "ibm_pi_image" "test" {
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
