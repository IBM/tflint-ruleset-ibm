# ibm_pi_volume_attach

Schema validation for the ibm_pi_volume_attach resource.

This rule validates `ibm_pi_volume_attach` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pi_volume_attach" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `pi_cloud_instance_id` | yes | — |
| `pi_instance_id` | yes | — |
| `pi_volume_id` | yes | — |
| `status` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_volume_attach_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |
| `ibm_pi_volume_attach_invalid_pi_instance_id` | `pi_instance_id` |
| `ibm_pi_volume_attach_invalid_pi_volume_id` | `pi_volume_id` |

## Example

### Valid

```hcl
resource "ibm_pi_volume_attach" "test" {
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  pi_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  pi_volume_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
