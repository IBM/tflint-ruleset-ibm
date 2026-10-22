# ibm_pi_virtual_serial_number_data

Schema validation for the ibm_pi_virtual_serial_number data source.

This rule validates `ibm_pi_virtual_serial_number` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pi_virtual_serial_number_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `description` | no | — |
| `instance_id` | no | — |
| `pi_cloud_instance_id` | yes | — |
| `pi_serial` | yes | — |
| `software_tier` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_virtual_serial_number_data_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |
| `ibm_pi_virtual_serial_number_data_invalid_pi_serial` | `pi_serial` |

## Example

### Valid

```hcl
data "ibm_pi_virtual_serial_number" "test" {
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  pi_serial = "example"
}
```
