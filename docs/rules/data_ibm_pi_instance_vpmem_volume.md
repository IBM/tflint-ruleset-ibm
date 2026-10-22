# ibm_pi_instance_vpmem_volume_data

Reads ibm_pi_instance_vpmem_volume. Volume name.

This rule validates `ibm_pi_instance_vpmem_volume` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pi_instance_vpmem_volume_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `creation_date` | no | — |
| `crn` | no | — |
| `error_code` | no | — |
| `href` | no | — |
| `name` | no | — |
| `pi_cloud_instance_id` | yes | — |
| `pi_pvm_instance_id` | yes | — |
| `pi_vpmem_volume_id` | yes | — |
| `pvm_instance_id` | no | — |
| `reason` | no | — |
| `size` | no | — |
| `status` | no | — |
| `updated_date` | no | — |
| `user_tags` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_instance_vpmem_volume_data_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |
| `ibm_pi_instance_vpmem_volume_data_invalid_pi_pvm_instance_id` | `pi_pvm_instance_id` |
| `ibm_pi_instance_vpmem_volume_data_invalid_pi_vpmem_volume_id` | `pi_vpmem_volume_id` |

## Example

### Valid

```hcl
data "ibm_pi_instance_vpmem_volume" "test" {
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  pi_pvm_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  pi_vpmem_volume_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
