# ibm_scc_profile_data

Schema validation for the ibm_scc_profile data source.

This rule validates `ibm_scc_profile` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ProfileAttachmentResponse` (`security-compliance`).

## Configuration

```hcl
rule "ibm_scc_profile_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `attachments_count` | no | — |
| `control_parents_count` | no | — |
| `controls` | no | — |
| `controls_count` | no | — |
| `created_by` | no | — |
| `created_on` | no | — |
| `default_parameters` | no | — |
| `hierarchy_enabled` | no | — |
| `instance_id` | yes | API: The ID of the associated Security and Compliance Center instance |
| `latest` | no | — |
| `profile_description` | no | — |
| `profile_id` | yes | API: The ID of the profile |
| `profile_name` | no | — |
| `profile_type` | no | — |
| `profile_version` | no | — |
| `updated_by` | no | — |
| `updated_on` | no | — |
| `version_group_label` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_scc_profile_data_invalid_instance_id` | `instance_id` |
| `ibm_scc_profile_data_invalid_profile_id` | `profile_id` |

## Example

### Valid

```hcl
data "ibm_scc_profile" "test" {
  instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  profile_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
