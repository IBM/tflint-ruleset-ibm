# ibm_pi_volume_group_data

Schema validation for the ibm_pi_volume_group data source.

This rule validates `ibm_pi_volume_group` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pi_volume_group_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `auxiliary` | no | — |
| `consistency_group_name` | no | — |
| `pi_cloud_instance_id` | yes | — |
| `pi_volume_group_id` | yes | — |
| `replication_sites` | no | — |
| `replication_status` | no | — |
| `status` | no | — |
| `status_description_errors` | no | — |
| `storage_pool` | no | — |
| `volume_group_name` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_volume_group_data_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |
| `ibm_pi_volume_group_data_invalid_pi_volume_group_id` | `pi_volume_group_id` |

## Example

### Valid

```hcl
data "ibm_pi_volume_group" "test" {
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  pi_volume_group_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
