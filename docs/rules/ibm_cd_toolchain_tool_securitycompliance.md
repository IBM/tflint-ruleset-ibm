# ibm_cd_toolchain_tool_securitycompliance

Manages ibm_cd_toolchain_tool_securitycompliance. Name of the tool.

This rule validates `ibm_cd_toolchain_tool_securitycompliance` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Toolchain` (`toolchain`).

## Configuration

```hcl
rule "ibm_cd_toolchain_tool_securitycompliance" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `crn` | no | — |
| `href` | no | — |
| `name` | no | pattern `^([^\x00-\x7F]|[a-zA-Z0-9-._ ])+$` |
| `referent` | no | — |
| `resource_group_id` | no | — |
| `state` | no | — |
| `tool_id` | no | — |
| `toolchain_crn` | no | — |
| `toolchain_id` | yes | pattern `^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$` |
| `updated_at` | no | — |
| `parameters` (block) | yes | nested block |
| `parameters.attachment_id` | no | — |
| `parameters.cos_bucket_name` | no | — |
| `parameters.evidence_locker_type` | no | — |
| `parameters.evidence_namespace` | no | — |
| `parameters.evidence_repo_url` | no | — |
| `parameters.instance_crn` | no | crn format |
| `parameters.name` | yes | pattern `^([^\x00-\x7F]|[a-zA-Z0-9-._ ])+$` |
| `parameters.profile_name` | no | — |
| `parameters.profile_version` | no | — |
| `parameters.scc_api_key` | no | — |
| `parameters.use_profile_attachment` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cd_toolchain_tool_securitycompliance_invalid_name` | `name` |
| `ibm_cd_toolchain_tool_securitycompliance_invalid_toolchain_id` | `toolchain_id` |
| `ibm_cd_toolchain_tool_securitycompliance_invalid_parameters__attachment_id` | `parameters.attachment_id` |
| `ibm_cd_toolchain_tool_securitycompliance_invalid_parameters__cos_bucket_name` | `parameters.cos_bucket_name` |
| `ibm_cd_toolchain_tool_securitycompliance_invalid_parameters__evidence_locker_type` | `parameters.evidence_locker_type` |
| `ibm_cd_toolchain_tool_securitycompliance_invalid_parameters__evidence_namespace` | `parameters.evidence_namespace` |
| `ibm_cd_toolchain_tool_securitycompliance_invalid_parameters__evidence_repo_url` | `parameters.evidence_repo_url` |
| `ibm_cd_toolchain_tool_securitycompliance_invalid_parameters__instance_crn` | `parameters.instance_crn` |
| `ibm_cd_toolchain_tool_securitycompliance_invalid_parameters__name` | `parameters.name` |
| `ibm_cd_toolchain_tool_securitycompliance_invalid_parameters__profile_name` | `parameters.profile_name` |
| `ibm_cd_toolchain_tool_securitycompliance_invalid_parameters__profile_version` | `parameters.profile_version` |
| `ibm_cd_toolchain_tool_securitycompliance_invalid_parameters__scc_api_key` | `parameters.scc_api_key` |
| `ibm_cd_toolchain_tool_securitycompliance_invalid_parameters__use_profile_attachment` | `parameters.use_profile_attachment` |

## Example

### Valid

```hcl
resource "ibm_cd_toolchain_tool_securitycompliance" "test" {
  toolchain_id = "01234567-89ab-4cde-8123-456789abcdef"
  parameters {
    name = "example-name"
  }
}
```
