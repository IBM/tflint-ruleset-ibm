# ibm_pi_instance_networks_data

Schema validation for the ibm_pi_instance_networks data source.

This rule validates `ibm_pi_instance_networks` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pi_instance_networks_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `networks` | no | — |
| `pi_cloud_instance_id` | yes | — |
| `pi_instance_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_instance_networks_data_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |
| `ibm_pi_instance_networks_data_invalid_pi_instance_id` | `pi_instance_id` |

## Example

### Valid

```hcl
data "ibm_pi_instance_networks" "test" {
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  pi_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
