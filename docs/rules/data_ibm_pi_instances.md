# ibm_pi_instances_data

Schema validation for the ibm_pi_instances data source.

This rule validates `ibm_pi_instances` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `NetworkAddressGroupUpdate` (`power-cloud`).

## Configuration

```hcl
rule "ibm_pi_instances_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `pi_cloud_instance_id` | yes | — |
| `pvm_instances` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_instances_data_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |

## Example

### Valid

```hcl
data "ibm_pi_instances" "test" {
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
