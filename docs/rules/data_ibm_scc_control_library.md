# ibm_scc_control_library_data

Schema validation for the ibm_scc_control_library data source.

This rule validates `ibm_scc_control_library` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ProfileControlsPrototype` (`security-compliance`).

## Configuration

```hcl
rule "ibm_scc_control_library_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_id` | no | — |
| `control_library_description` | no | — |
| `control_library_id` | yes | API: The ID of the control library that contains the profile. |
| `control_library_name` | no | — |
| `control_library_type` | no | — |
| `control_library_version` | no | — |
| `control_parents_count` | no | — |
| `controls` | no | — |
| `controls_count` | no | — |
| `created_by` | no | — |
| `created_on` | no | — |
| `hierarchy_enabled` | no | — |
| `instance_id` | yes | — |
| `latest` | no | — |
| `updated_by` | no | — |
| `updated_on` | no | — |
| `version_group_label` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_scc_control_library_data_invalid_control_library_id` | `control_library_id` |
| `ibm_scc_control_library_data_invalid_instance_id` | `instance_id` |

## Example

### Valid

```hcl
data "ibm_scc_control_library" "test" {
  control_library_id = "01234567-89ab-cdef-0123-456789abcdef"
  instance_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
