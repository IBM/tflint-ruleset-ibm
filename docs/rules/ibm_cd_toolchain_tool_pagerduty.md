# ibm_cd_toolchain_tool_pagerduty

Manages ibm_cd_toolchain_tool_pagerduty. Name of the tool.

This rule validates `ibm_cd_toolchain_tool_pagerduty` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Toolchain` (`toolchain`).

## Configuration

```hcl
rule "ibm_cd_toolchain_tool_pagerduty" {
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
| `parameters.service_id` | no | — |
| `parameters.service_key` | yes | — |
| `parameters.service_url` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cd_toolchain_tool_pagerduty_invalid_name` | `name` |
| `ibm_cd_toolchain_tool_pagerduty_invalid_toolchain_id` | `toolchain_id` |
| `ibm_cd_toolchain_tool_pagerduty_invalid_parameters__service_id` | `parameters.service_id` |
| `ibm_cd_toolchain_tool_pagerduty_invalid_parameters__service_key` | `parameters.service_key` |
| `ibm_cd_toolchain_tool_pagerduty_invalid_parameters__service_url` | `parameters.service_url` |

## Example

### Valid

```hcl
variable "service_key" {
  sensitive = true
}

resource "ibm_cd_toolchain_tool_pagerduty" "test" {
  toolchain_id = "01234567-89ab-4cde-8123-456789abcdef"
  parameters {
    service_key = var.service_key
    service_url = "example"
  }
}
```
