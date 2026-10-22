# ibm_cd_toolchain_tool_sonarqube

Manages ibm_cd_toolchain_tool_sonarqube. Name of the tool.

This rule validates `ibm_cd_toolchain_tool_sonarqube` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Toolchain` (`toolchain`).

## Configuration

```hcl
rule "ibm_cd_toolchain_tool_sonarqube" {
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
| `parameters.blind_connection` | no | — |
| `parameters.name` | yes | pattern `^([^\x00-\x7F]|[a-zA-Z0-9-._ ])+$` |
| `parameters.server_url` | yes | — |
| `parameters.user_login` | no | — |
| `parameters.user_password` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cd_toolchain_tool_sonarqube_invalid_name` | `name` |
| `ibm_cd_toolchain_tool_sonarqube_invalid_toolchain_id` | `toolchain_id` |
| `ibm_cd_toolchain_tool_sonarqube_invalid_parameters__blind_connection` | `parameters.blind_connection` |
| `ibm_cd_toolchain_tool_sonarqube_invalid_parameters__name` | `parameters.name` |
| `ibm_cd_toolchain_tool_sonarqube_invalid_parameters__server_url` | `parameters.server_url` |
| `ibm_cd_toolchain_tool_sonarqube_invalid_parameters__user_login` | `parameters.user_login` |
| `ibm_cd_toolchain_tool_sonarqube_invalid_parameters__user_password` | `parameters.user_password` |

## Example

### Valid

```hcl
resource "ibm_cd_toolchain_tool_sonarqube" "test" {
  toolchain_id = "01234567-89ab-4cde-8123-456789abcdef"
  parameters {
    name = "example-name"
    server_url = "example"
  }
}
```
