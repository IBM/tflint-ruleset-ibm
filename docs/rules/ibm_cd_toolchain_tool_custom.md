# ibm_cd_toolchain_tool_custom

Manages ibm_cd_toolchain_tool_custom. Name of the tool.

This rule validates `ibm_cd_toolchain_tool_custom` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Toolchain` (`toolchain`).

## Configuration

```hcl
rule "ibm_cd_toolchain_tool_custom" {
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
| `parameters.additional_properties` | no | — |
| `parameters.dashboard_url` | yes | — |
| `parameters.description` | no | API: Describes the toolchain |
| `parameters.documentation_url` | no | — |
| `parameters.image_url` | no | — |
| `parameters.lifecycle_phase` | yes | — |
| `parameters.name` | yes | pattern `^([^\x00-\x7F]|[a-zA-Z0-9-._ ])+$` |
| `parameters.type` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cd_toolchain_tool_custom_invalid_name` | `name` |
| `ibm_cd_toolchain_tool_custom_invalid_toolchain_id` | `toolchain_id` |
| `ibm_cd_toolchain_tool_custom_invalid_parameters__additional_properties` | `parameters.additional_properties` |
| `ibm_cd_toolchain_tool_custom_invalid_parameters__dashboard_url` | `parameters.dashboard_url` |
| `ibm_cd_toolchain_tool_custom_invalid_parameters__description` | `parameters.description` |
| `ibm_cd_toolchain_tool_custom_invalid_parameters__documentation_url` | `parameters.documentation_url` |
| `ibm_cd_toolchain_tool_custom_invalid_parameters__image_url` | `parameters.image_url` |
| `ibm_cd_toolchain_tool_custom_invalid_parameters__lifecycle_phase` | `parameters.lifecycle_phase` |
| `ibm_cd_toolchain_tool_custom_invalid_parameters__name` | `parameters.name` |
| `ibm_cd_toolchain_tool_custom_invalid_parameters__type` | `parameters.type` |

## Example

### Valid

```hcl
resource "ibm_cd_toolchain_tool_custom" "test" {
  toolchain_id = "01234567-89ab-4cde-8123-456789abcdef"
  parameters {
    dashboard_url = "example"
    lifecycle_phase = "example"
    name = "example-name"
    type = "example"
  }
}
```
