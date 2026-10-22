# ibm_scc_profile

Schema validation for the ibm_scc_profile resource.

This rule validates `ibm_scc_profile` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Profile` (`security-compliance`).

## Configuration

```hcl
rule "ibm_scc_profile" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `attachments_count` | no | — |
| `control_parents_count` | no | — |
| `controls_count` | no | — |
| `created_by` | no | — |
| `created_on` | no | — |
| `hierarchy_enabled` | no | — |
| `instance_id` | yes | — |
| `latest` | no | — |
| `profile_description` | yes | pattern `^[a-zA-Z0-9_,'"\s\-\[\]]+$` |
| `profile_id` | no | — |
| `profile_name` | yes | pattern `^[a-zA-Z0-9_\s\-]*$` |
| `profile_type` | yes | one of: `custom`, `predefined` |
| `profile_version` | no | pattern `^\d.\d.\d$` |
| `updated_by` | no | — |
| `updated_on` | no | — |
| `version_group_label` | no | — |
| `controls` (block) | yes | nested block |
| `controls.control_id` | yes | API: The control ID. |
| `controls.control_library_id` | yes | API: The ID of the control library that contains a profile. |
| `default_parameters` (block) | no | nested block |
| `default_parameters.assessment_id` | no | API: The implementation ID of the parameter. |
| `default_parameters.assessment_type` | no | API: The type of the implementation. |
| `default_parameters.parameter_default_value` | no | API: The default value of the parameter. |
| `default_parameters.parameter_display_name` | no | API: The parameter display name. |
| `default_parameters.parameter_name` | no | API: The parameter's name. |
| `default_parameters.parameter_type` | no | API: The parameter type. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_scc_profile_invalid_instance_id` | `instance_id` |
| `ibm_scc_profile_invalid_profile_description` | `profile_description` |
| `ibm_scc_profile_invalid_profile_name` | `profile_name` |
| `ibm_scc_profile_invalid_profile_type` | `profile_type` |
| `ibm_scc_profile_invalid_profile_version` | `profile_version` |
| `ibm_scc_profile_invalid_controls__control_id` | `controls.control_id` |
| `ibm_scc_profile_invalid_controls__control_library_id` | `controls.control_library_id` |
| `ibm_scc_profile_invalid_default_parameters__assessment_id` | `default_parameters.assessment_id` |
| `ibm_scc_profile_invalid_default_parameters__assessment_type` | `default_parameters.assessment_type` |
| `ibm_scc_profile_invalid_default_parameters__parameter_default_value` | `default_parameters.parameter_default_value` |
| `ibm_scc_profile_invalid_default_parameters__parameter_display_name` | `default_parameters.parameter_display_name` |
| `ibm_scc_profile_invalid_default_parameters__parameter_name` | `default_parameters.parameter_name` |
| `ibm_scc_profile_invalid_default_parameters__parameter_type` | `default_parameters.parameter_type` |

## Example

### Valid

```hcl
resource "ibm_scc_profile" "test" {
  instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  profile_description = "example-name"
  profile_name = "example-name"
  profile_type = "custom"
  controls {
    control_id = "01234567-89ab-cdef-0123-456789abcdef"
    control_library_id = "01234567-89ab-cdef-0123-456789abcdef"
  }
}
```
