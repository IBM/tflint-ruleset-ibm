# ibm_cd_toolchain_tool_slack

Manages ibm_cd_toolchain_tool_slack. Name of the tool.

This rule validates `ibm_cd_toolchain_tool_slack` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Toolchain` (`toolchain`).

## Configuration

```hcl
rule "ibm_cd_toolchain_tool_slack" {
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
| `parameters.channel_name` | yes | — |
| `parameters.pipeline_fail` | no | — |
| `parameters.pipeline_start` | no | — |
| `parameters.pipeline_success` | no | — |
| `parameters.team_name` | no | — |
| `parameters.toolchain_bind` | no | — |
| `parameters.toolchain_unbind` | no | — |
| `parameters.webhook` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cd_toolchain_tool_slack_invalid_name` | `name` |
| `ibm_cd_toolchain_tool_slack_invalid_toolchain_id` | `toolchain_id` |
| `ibm_cd_toolchain_tool_slack_invalid_parameters__channel_name` | `parameters.channel_name` |
| `ibm_cd_toolchain_tool_slack_invalid_parameters__pipeline_fail` | `parameters.pipeline_fail` |
| `ibm_cd_toolchain_tool_slack_invalid_parameters__pipeline_start` | `parameters.pipeline_start` |
| `ibm_cd_toolchain_tool_slack_invalid_parameters__pipeline_success` | `parameters.pipeline_success` |
| `ibm_cd_toolchain_tool_slack_invalid_parameters__team_name` | `parameters.team_name` |
| `ibm_cd_toolchain_tool_slack_invalid_parameters__toolchain_bind` | `parameters.toolchain_bind` |
| `ibm_cd_toolchain_tool_slack_invalid_parameters__toolchain_unbind` | `parameters.toolchain_unbind` |
| `ibm_cd_toolchain_tool_slack_invalid_parameters__webhook` | `parameters.webhook` |

## Example

### Valid

```hcl
variable "webhook" {
  sensitive = true
}

resource "ibm_cd_toolchain_tool_slack" "test" {
  toolchain_id = "01234567-89ab-4cde-8123-456789abcdef"
  parameters {
    channel_name = "example"
    webhook = var.webhook
  }
}
```
