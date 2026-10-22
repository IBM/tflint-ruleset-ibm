# ibm_pi_volume_onboarding_data

Schema validation for the ibm_pi_volume_onboarding data source.

This rule validates `ibm_pi_volume_onboarding` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pi_volume_onboarding_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `create_time` | no | — |
| `description` | no | — |
| `input_volumes` | no | — |
| `pi_cloud_instance_id` | yes | — |
| `pi_volume_onboarding_id` | yes | — |
| `progress` | no | — |
| `results_onboarded_volumes` | no | — |
| `results_volume_onboarding_failures` | no | — |
| `status` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_volume_onboarding_data_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |
| `ibm_pi_volume_onboarding_data_invalid_pi_volume_onboarding_id` | `pi_volume_onboarding_id` |

## Example

### Valid

```hcl
data "ibm_pi_volume_onboarding" "test" {
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  pi_volume_onboarding_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
