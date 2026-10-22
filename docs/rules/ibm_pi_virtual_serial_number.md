# ibm_pi_virtual_serial_number

Schema validation for the ibm_pi_virtual_serial_number resource.

This rule validates `ibm_pi_virtual_serial_number` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pi_virtual_serial_number" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `pi_cloud_instance_id` | yes | — |
| `pi_description` | no | — |
| `pi_instance_id` | no | — |
| `pi_retain_virtual_serial_number` | no | — |
| `pi_serial` | yes | — |
| `pi_software_tier` | no | one of: `P05`, `P10`, `P20`, `P30` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_virtual_serial_number_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |
| `ibm_pi_virtual_serial_number_invalid_pi_description` | `pi_description` |
| `ibm_pi_virtual_serial_number_invalid_pi_instance_id` | `pi_instance_id` |
| `ibm_pi_virtual_serial_number_invalid_pi_retain_virtual_serial_number` | `pi_retain_virtual_serial_number` |
| `ibm_pi_virtual_serial_number_invalid_pi_serial` | `pi_serial` |
| `ibm_pi_virtual_serial_number_invalid_pi_software_tier` | `pi_software_tier` |

## Example

### Valid

```hcl
resource "ibm_pi_virtual_serial_number" "test" {
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  pi_serial = "example"
}
```
