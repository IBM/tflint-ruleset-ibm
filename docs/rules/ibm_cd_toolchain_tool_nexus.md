# ibm_cd_toolchain_tool_nexus

Manages ibm_cd_toolchain_tool_nexus. Name of the tool.

This rule validates `ibm_cd_toolchain_tool_nexus` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Toolchain` (`toolchain`).

## Configuration

```hcl
rule "ibm_cd_toolchain_tool_nexus" {
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
| `parameters.mirror_url` | no | — |
| `parameters.name` | yes | pattern `^([^\x00-\x7F]|[a-zA-Z0-9-._ ])+$` |
| `parameters.release_url` | no | — |
| `parameters.server_url` | no | — |
| `parameters.snapshot_url` | no | — |
| `parameters.token` | no | — |
| `parameters.type` | yes | — |
| `parameters.user_id` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cd_toolchain_tool_nexus_invalid_name` | `name` |
| `ibm_cd_toolchain_tool_nexus_invalid_toolchain_id` | `toolchain_id` |
| `ibm_cd_toolchain_tool_nexus_invalid_parameters__mirror_url` | `parameters.mirror_url` |
| `ibm_cd_toolchain_tool_nexus_invalid_parameters__name` | `parameters.name` |
| `ibm_cd_toolchain_tool_nexus_invalid_parameters__release_url` | `parameters.release_url` |
| `ibm_cd_toolchain_tool_nexus_invalid_parameters__server_url` | `parameters.server_url` |
| `ibm_cd_toolchain_tool_nexus_invalid_parameters__snapshot_url` | `parameters.snapshot_url` |
| `ibm_cd_toolchain_tool_nexus_invalid_parameters__token` | `parameters.token` |
| `ibm_cd_toolchain_tool_nexus_invalid_parameters__type` | `parameters.type` |
| `ibm_cd_toolchain_tool_nexus_invalid_parameters__user_id` | `parameters.user_id` |

## Example

### Valid

```hcl
resource "ibm_cd_toolchain_tool_nexus" "test" {
  toolchain_id = "01234567-89ab-4cde-8123-456789abcdef"
  parameters {
    name = "example-name"
    type = "example"
  }
}
```
