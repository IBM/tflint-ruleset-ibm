# ibm_pi_instance_vpmem_volumes

Schema validation for the ibm_pi_instance_vpmem_volumes resource.

This rule validates `ibm_pi_instance_vpmem_volumes` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `VPMemVolumeCreate` (`power-cloud`).

## Configuration

```hcl
rule "ibm_pi_instance_vpmem_volumes" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `pi_cloud_instance_id` | yes | — |
| `pi_pvm_instance_id` | yes | — |
| `pi_user_tags` | no | — |
| `pi_vpmem_volumes` (block) | yes | nested block |
| `pi_vpmem_volumes.name` | yes | pattern `^[a-zA-Z0-9][a-zA-Z0-9-_]*$` |
| `pi_vpmem_volumes.size` | yes | min 1 |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_instance_vpmem_volumes_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |
| `ibm_pi_instance_vpmem_volumes_invalid_pi_pvm_instance_id` | `pi_pvm_instance_id` |
| `ibm_pi_instance_vpmem_volumes_invalid_pi_user_tags` | `pi_user_tags` |
| `ibm_pi_instance_vpmem_volumes_invalid_pi_vpmem_volumes__name` | `pi_vpmem_volumes.name` |
| `ibm_pi_instance_vpmem_volumes_invalid_pi_vpmem_volumes__size` | `pi_vpmem_volumes.size` |

## Example

### Valid

```hcl
resource "ibm_pi_instance_vpmem_volumes" "test" {
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  pi_pvm_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  pi_vpmem_volumes {
    name = "example-name"
    size = 1
  }
}
```
