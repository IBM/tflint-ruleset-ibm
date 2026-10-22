# ibm_pi_storage_tiers_data

Schema validation for the ibm_pi_storage_tiers data source.

This rule validates `ibm_pi_storage_tiers` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pi_storage_tiers_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `pi_cloud_instance_id` | yes | — |
| `region_storage_tiers` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_storage_tiers_data_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |

## Example

### Valid

```hcl
data "ibm_pi_storage_tiers" "test" {
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
