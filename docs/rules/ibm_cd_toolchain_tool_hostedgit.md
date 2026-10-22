# ibm_cd_toolchain_tool_hostedgit

Manages ibm_cd_toolchain_tool_hostedgit. Name of the tool.

This rule validates `ibm_cd_toolchain_tool_hostedgit` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Toolchain` (`toolchain`).

## Configuration

```hcl
rule "ibm_cd_toolchain_tool_hostedgit" {
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
| `initialization` (block) | yes | nested block |
| `initialization.git_id` | no | — |
| `initialization.owner_id` | no | — |
| `initialization.private_repo` | no | — |
| `initialization.repo_name` | no | — |
| `initialization.repo_url` | no | — |
| `initialization.source_repo_url` | no | — |
| `initialization.type` | yes | — |
| `parameters` (block) | yes | nested block |
| `parameters.api_root_url` | no | — |
| `parameters.api_token` | no | — |
| `parameters.auth_type` | no | — |
| `parameters.default_branch` | no | — |
| `parameters.enable_traceability` | no | — |
| `parameters.git_id` | no | — |
| `parameters.integration_owner` | no | — |
| `parameters.owner_id` | no | — |
| `parameters.private_repo` | no | — |
| `parameters.repo_id` | no | — |
| `parameters.repo_name` | no | — |
| `parameters.repo_url` | no | — |
| `parameters.source_repo_url` | no | — |
| `parameters.token_url` | no | — |
| `parameters.toolchain_issues_enabled` | no | — |
| `parameters.type` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cd_toolchain_tool_hostedgit_invalid_name` | `name` |
| `ibm_cd_toolchain_tool_hostedgit_invalid_toolchain_id` | `toolchain_id` |
| `ibm_cd_toolchain_tool_hostedgit_invalid_initialization__git_id` | `initialization.git_id` |
| `ibm_cd_toolchain_tool_hostedgit_invalid_initialization__owner_id` | `initialization.owner_id` |
| `ibm_cd_toolchain_tool_hostedgit_invalid_initialization__private_repo` | `initialization.private_repo` |
| `ibm_cd_toolchain_tool_hostedgit_invalid_initialization__repo_name` | `initialization.repo_name` |
| `ibm_cd_toolchain_tool_hostedgit_invalid_initialization__repo_url` | `initialization.repo_url` |
| `ibm_cd_toolchain_tool_hostedgit_invalid_initialization__source_repo_url` | `initialization.source_repo_url` |
| `ibm_cd_toolchain_tool_hostedgit_invalid_initialization__type` | `initialization.type` |
| `ibm_cd_toolchain_tool_hostedgit_invalid_parameters__api_root_url` | `parameters.api_root_url` |
| `ibm_cd_toolchain_tool_hostedgit_invalid_parameters__api_token` | `parameters.api_token` |
| `ibm_cd_toolchain_tool_hostedgit_invalid_parameters__auth_type` | `parameters.auth_type` |
| `ibm_cd_toolchain_tool_hostedgit_invalid_parameters__default_branch` | `parameters.default_branch` |
| `ibm_cd_toolchain_tool_hostedgit_invalid_parameters__enable_traceability` | `parameters.enable_traceability` |
| `ibm_cd_toolchain_tool_hostedgit_invalid_parameters__git_id` | `parameters.git_id` |
| `ibm_cd_toolchain_tool_hostedgit_invalid_parameters__integration_owner` | `parameters.integration_owner` |
| `ibm_cd_toolchain_tool_hostedgit_invalid_parameters__owner_id` | `parameters.owner_id` |
| `ibm_cd_toolchain_tool_hostedgit_invalid_parameters__private_repo` | `parameters.private_repo` |
| `ibm_cd_toolchain_tool_hostedgit_invalid_parameters__repo_id` | `parameters.repo_id` |
| `ibm_cd_toolchain_tool_hostedgit_invalid_parameters__repo_name` | `parameters.repo_name` |
| `ibm_cd_toolchain_tool_hostedgit_invalid_parameters__repo_url` | `parameters.repo_url` |
| `ibm_cd_toolchain_tool_hostedgit_invalid_parameters__source_repo_url` | `parameters.source_repo_url` |
| `ibm_cd_toolchain_tool_hostedgit_invalid_parameters__token_url` | `parameters.token_url` |
| `ibm_cd_toolchain_tool_hostedgit_invalid_parameters__toolchain_issues_enabled` | `parameters.toolchain_issues_enabled` |
| `ibm_cd_toolchain_tool_hostedgit_invalid_parameters__type` | `parameters.type` |

## Example

### Valid

```hcl
resource "ibm_cd_toolchain_tool_hostedgit" "test" {
  toolchain_id = "01234567-89ab-4cde-8123-456789abcdef"
  initialization {
    type = "example"
  }
  parameters {
  }
}
```
