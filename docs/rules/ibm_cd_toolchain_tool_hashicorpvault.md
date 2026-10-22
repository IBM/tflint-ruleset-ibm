# ibm_cd_toolchain_tool_hashicorpvault

Manages ibm_cd_toolchain_tool_hashicorpvault. Name of the tool.

This rule validates `ibm_cd_toolchain_tool_hashicorpvault` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Toolchain` (`toolchain`).

## Configuration

```hcl
rule "ibm_cd_toolchain_tool_hashicorpvault" {
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
| `parameters.authentication_method` | yes | — |
| `parameters.dashboard_url` | yes | — |
| `parameters.default_secret` | no | — |
| `parameters.name` | yes | pattern `^([^\x00-\x7F]|[a-zA-Z0-9-._ ])+$` |
| `parameters.password` | no | — |
| `parameters.path` | yes | — |
| `parameters.role_id` | no | — |
| `parameters.secret_filter` | no | — |
| `parameters.secret_id` | no | — |
| `parameters.server_url` | yes | — |
| `parameters.token` | no | — |
| `parameters.username` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cd_toolchain_tool_hashicorpvault_invalid_name` | `name` |
| `ibm_cd_toolchain_tool_hashicorpvault_invalid_toolchain_id` | `toolchain_id` |
| `ibm_cd_toolchain_tool_hashicorpvault_invalid_parameters__authentication_method` | `parameters.authentication_method` |
| `ibm_cd_toolchain_tool_hashicorpvault_invalid_parameters__dashboard_url` | `parameters.dashboard_url` |
| `ibm_cd_toolchain_tool_hashicorpvault_invalid_parameters__default_secret` | `parameters.default_secret` |
| `ibm_cd_toolchain_tool_hashicorpvault_invalid_parameters__name` | `parameters.name` |
| `ibm_cd_toolchain_tool_hashicorpvault_invalid_parameters__password` | `parameters.password` |
| `ibm_cd_toolchain_tool_hashicorpvault_invalid_parameters__path` | `parameters.path` |
| `ibm_cd_toolchain_tool_hashicorpvault_invalid_parameters__role_id` | `parameters.role_id` |
| `ibm_cd_toolchain_tool_hashicorpvault_invalid_parameters__secret_filter` | `parameters.secret_filter` |
| `ibm_cd_toolchain_tool_hashicorpvault_invalid_parameters__secret_id` | `parameters.secret_id` |
| `ibm_cd_toolchain_tool_hashicorpvault_invalid_parameters__server_url` | `parameters.server_url` |
| `ibm_cd_toolchain_tool_hashicorpvault_invalid_parameters__token` | `parameters.token` |
| `ibm_cd_toolchain_tool_hashicorpvault_invalid_parameters__username` | `parameters.username` |

## Example

### Valid

```hcl
resource "ibm_cd_toolchain_tool_hashicorpvault" "test" {
  toolchain_id = "01234567-89ab-4cde-8123-456789abcdef"
  parameters {
    authentication_method = "example"
    dashboard_url = "example"
    name = "example-name"
    path = "example"
    server_url = "example"
  }
}
```
