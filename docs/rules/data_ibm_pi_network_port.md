# ibm_pi_network_port_data

Schema validation for the ibm_pi_network_port data source.

This rule validates `ibm_pi_network_port` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pi_network_port_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `network_ports` | no | — |
| `pi_cloud_instance_id` | yes | — |
| `pi_network_id` | no | — |
| `pi_network_name` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_network_port_data_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |
| `ibm_pi_network_port_data_invalid_pi_network_id` | `pi_network_id` |
| `ibm_pi_network_port_data_invalid_pi_network_name` | `pi_network_name` |

## Example

### Valid

```hcl
data "ibm_pi_network_port" "test" {
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
