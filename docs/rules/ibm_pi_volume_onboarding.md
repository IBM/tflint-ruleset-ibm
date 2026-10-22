# ibm_pi_volume_onboarding

Schema validation for the ibm_pi_volume_onboarding resource.

This rule validates `ibm_pi_volume_onboarding` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pi_volume_onboarding" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `create_time` | no | — |
| `input_volumes` | no | — |
| `onboarding_id` | no | — |
| `pi_cloud_instance_id` | yes | — |
| `pi_description` | no | — |
| `progress` | no | — |
| `results_onboarded_volumes` | no | — |
| `results_volume_onboarding_failures` | no | — |
| `status` | no | — |
| `pi_onboarding_volumes` (block) | yes | nested block |
| `pi_onboarding_volumes.pi_source_crn` | yes | crn format |
| `pi_onboarding_volumes.pi_auxiliary_volumes` (block) | no | nested block |
| `pi_onboarding_volumes.pi_auxiliary_volumes.pi_auxiliary_volume_name` | yes | — |
| `pi_onboarding_volumes.pi_auxiliary_volumes.pi_display_name` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_volume_onboarding_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |
| `ibm_pi_volume_onboarding_invalid_pi_description` | `pi_description` |
| `ibm_pi_volume_onboarding_invalid_pi_onboarding_volumes__pi_source_crn` | `pi_onboarding_volumes.pi_source_crn` |
| `ibm_pi_volume_onboarding_invalid_pi_onboarding_volumes__pi_auxiliary_volumes__pi_auxiliary_volume_name` | `pi_onboarding_volumes.pi_auxiliary_volumes.pi_auxiliary_volume_name` |
| `ibm_pi_volume_onboarding_invalid_pi_onboarding_volumes__pi_auxiliary_volumes__pi_display_name` | `pi_onboarding_volumes.pi_auxiliary_volumes.pi_display_name` |

## Example

### Valid

```hcl
resource "ibm_pi_volume_onboarding" "test" {
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  pi_onboarding_volumes {
    pi_source_crn = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
  }
}
```
