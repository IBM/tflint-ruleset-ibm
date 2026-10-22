# ibm_pi_cloud_connection_network_attach

Schema validation for the ibm_pi_cloud_connection_network_attach resource.

This rule validates `ibm_pi_cloud_connection_network_attach` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pi_cloud_connection_network_attach" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `pi_cloud_connection_id` | yes | — |
| `pi_cloud_instance_id` | yes | — |
| `pi_network_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_cloud_connection_network_attach_invalid_pi_cloud_connection_id` | `pi_cloud_connection_id` |
| `ibm_pi_cloud_connection_network_attach_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |
| `ibm_pi_cloud_connection_network_attach_invalid_pi_network_id` | `pi_network_id` |

## Example

### Valid

```hcl
resource "ibm_pi_cloud_connection_network_attach" "test" {
  pi_cloud_connection_id = "01234567-89ab-cdef-0123-456789abcdef"
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  pi_network_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
