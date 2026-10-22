# ibm_scc_control_library

Schema validation for the ibm_scc_control_library resource.

This rule validates `ibm_scc_control_library` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ControlPrototype` (`security-compliance`).

## Configuration

```hcl
rule "ibm_scc_control_library" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_id` | no | — |
| `control_library_description` | yes | pattern `[A-Za-z0-9]+` |
| `control_library_id` | no | — |
| `control_library_name` | yes | pattern `^[a-zA-Z0-9_\s\-]*$` |
| `control_library_type` | yes | one of: `custom`, `predefined` |
| `control_library_version` | no | pattern `^[a-zA-Z0-9_\-.]*$` |
| `control_parents_count` | no | — |
| `controls_count` | no | — |
| `created_by` | no | — |
| `created_on` | no | — |
| `hierarchy_enabled` | no | — |
| `instance_id` | yes | — |
| `latest` | no | — |
| `updated_by` | no | — |
| `updated_on` | no | — |
| `version_group_label` | no | — |
| `controls` (block) | yes | nested block |
| `controls.control_category` | no | API: The association of the control |
| `controls.control_description` | no | API: The control description. |
| `controls.control_id` | no | — |
| `controls.control_name` | no | API: The ID of the control library that contains the profile. |
| `controls.control_parent` | no | API: The ID of the parent control |
| `controls.control_requirement` | no | API: true if the control can be automated, false if the control cannot. |
| `controls.control_tags` | no | — |
| `controls.status` | no | API: Details if a control library can be used or not. |
| `controls.control_docs` (block) | no | nested block |
| `controls.control_docs.control_docs_id` | no | API: The ID of the control doc |
| `controls.control_docs.control_docs_type` | no | API: The type of the control doc |
| `controls.control_specifications` (block) | no | nested block |
| `controls.control_specifications.component_id` | no | API: The ID of the component. The component_id can be found from the 'service_name' using the Get Services method. |
| `controls.control_specifications.component_name` | no | — |
| `controls.control_specifications.control_specification_description` | no | API: Information about the Control Specification |
| `controls.control_specifications.control_specification_id` | no | — |
| `controls.control_specifications.environment` | no | one of: `ibm-cloud` |
| `controls.control_specifications.responsibility` | no | — |
| `controls.control_specifications.assessments` (block) | no | nested block |
| `controls.control_specifications.assessments.assessment_description` | no | API: Details on the intent of the rule for an assessment. |
| `controls.control_specifications.assessments.assessment_id` | yes | API: The ID of the rule to target. A list of rules can be obtained from the list_rules method. |
| `controls.control_specifications.assessments.assessment_method` | no | — |
| `controls.control_specifications.assessments.assessment_type` | no | API: The type of evaluation. |
| `controls.control_specifications.assessments.parameters` (block) | no | nested block |
| `controls.control_specifications.assessments.parameters.parameter_display_name` | no | — |
| `controls.control_specifications.assessments.parameters.parameter_name` | no | — |
| `controls.control_specifications.assessments.parameters.parameter_type` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_scc_control_library_invalid_control_library_description` | `control_library_description` |
| `ibm_scc_control_library_invalid_control_library_name` | `control_library_name` |
| `ibm_scc_control_library_invalid_control_library_type` | `control_library_type` |
| `ibm_scc_control_library_invalid_control_library_version` | `control_library_version` |
| `ibm_scc_control_library_invalid_instance_id` | `instance_id` |
| `ibm_scc_control_library_invalid_latest` | `latest` |
| `ibm_scc_control_library_invalid_controls__control_category` | `controls.control_category` |
| `ibm_scc_control_library_invalid_controls__control_description` | `controls.control_description` |
| `ibm_scc_control_library_invalid_controls__control_id` | `controls.control_id` |
| `ibm_scc_control_library_invalid_controls__control_name` | `controls.control_name` |
| `ibm_scc_control_library_invalid_controls__control_parent` | `controls.control_parent` |
| `ibm_scc_control_library_invalid_controls__control_requirement` | `controls.control_requirement` |
| `ibm_scc_control_library_invalid_controls__control_tags` | `controls.control_tags` |
| `ibm_scc_control_library_invalid_controls__status` | `controls.status` |
| `ibm_scc_control_library_invalid_controls__control_docs__control_docs_id` | `controls.control_docs.control_docs_id` |
| `ibm_scc_control_library_invalid_controls__control_docs__control_docs_type` | `controls.control_docs.control_docs_type` |
| `ibm_scc_control_library_invalid_controls__control_specifications__component_id` | `controls.control_specifications.component_id` |
| `ibm_scc_control_library_invalid_controls__control_specifications__component_name` | `controls.control_specifications.component_name` |
| `ibm_scc_control_library_invalid_controls__control_specifications__control_specification_description` | `controls.control_specifications.control_specification_description` |
| `ibm_scc_control_library_invalid_controls__control_specifications__control_specification_id` | `controls.control_specifications.control_specification_id` |
| `ibm_scc_control_library_invalid_controls__control_specifications__environment` | `controls.control_specifications.environment` |
| `ibm_scc_control_library_invalid_controls__control_specifications__responsibility` | `controls.control_specifications.responsibility` |
| `ibm_scc_control_library_invalid_controls__control_specifications__assessments__assessment_description` | `controls.control_specifications.assessments.assessment_description` |
| `ibm_scc_control_library_invalid_controls__control_specifications__assessments__assessment_id` | `controls.control_specifications.assessments.assessment_id` |
| `ibm_scc_control_library_invalid_controls__control_specifications__assessments__assessment_method` | `controls.control_specifications.assessments.assessment_method` |
| `ibm_scc_control_library_invalid_controls__control_specifications__assessments__assessment_type` | `controls.control_specifications.assessments.assessment_type` |
| `ibm_scc_control_library_invalid_controls__control_specifications__assessments__parameters__parameter_display_name` | `controls.control_specifications.assessments.parameters.parameter_display_name` |
| `ibm_scc_control_library_invalid_controls__control_specifications__assessments__parameters__parameter_name` | `controls.control_specifications.assessments.parameters.parameter_name` |
| `ibm_scc_control_library_invalid_controls__control_specifications__assessments__parameters__parameter_type` | `controls.control_specifications.assessments.parameters.parameter_type` |

## Example

### Valid

```hcl
resource "ibm_scc_control_library" "test" {
  control_library_description = "example-name"
  control_library_name = "example-name"
  control_library_type = "custom"
  instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  controls {
  }
}
```
