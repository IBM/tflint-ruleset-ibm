# ibm_pi_pvm_snapshots_data

Schema validation for the ibm_pi_pvm_snapshots data source.

This rule validates `ibm_pi_pvm_snapshots` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pi_pvm_snapshots_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `pi_cloud_instance_id` | yes | — |
| `pi_instance_id` | no | — |
| `pi_instance_name` | no | — |
| `pvm_snapshots` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_pvm_snapshots_data_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |
| `ibm_pi_pvm_snapshots_data_invalid_pi_instance_id` | `pi_instance_id` |
| `ibm_pi_pvm_snapshots_data_invalid_pi_instance_name` | `pi_instance_name` |

## Example

### Valid

```hcl
data "ibm_pi_pvm_snapshots" "test" {
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
