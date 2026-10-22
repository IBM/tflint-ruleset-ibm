# ibm_cd_toolchain_tool_appconfig

Manages ibm_cd_toolchain_tool_appconfig. Name of the tool.

This rule validates `ibm_cd_toolchain_tool_appconfig` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Toolchain` (`toolchain`).

## Configuration

```hcl
rule "ibm_cd_toolchain_tool_appconfig" {
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
| `parameters.collection_id` | yes | — |
| `parameters.environment_id` | yes | — |
| `parameters.instance_id` | yes | — |
| `parameters.location` | yes | API: Toolchain region |
| `parameters.name` | yes | pattern `^([^\x00-\x7F]|[a-zA-Z0-9-._ ])+$` |
| `parameters.resource_group_name` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cd_toolchain_tool_appconfig_invalid_name` | `name` |
| `ibm_cd_toolchain_tool_appconfig_invalid_toolchain_id` | `toolchain_id` |
| `ibm_cd_toolchain_tool_appconfig_invalid_parameters__collection_id` | `parameters.collection_id` |
| `ibm_cd_toolchain_tool_appconfig_invalid_parameters__environment_id` | `parameters.environment_id` |
| `ibm_cd_toolchain_tool_appconfig_invalid_parameters__instance_id` | `parameters.instance_id` |
| `ibm_cd_toolchain_tool_appconfig_invalid_parameters__location` | `parameters.location` |
| `ibm_cd_toolchain_tool_appconfig_invalid_parameters__name` | `parameters.name` |
| `ibm_cd_toolchain_tool_appconfig_invalid_parameters__resource_group_name` | `parameters.resource_group_name` |

## Example

### Valid

```hcl
resource "ibm_cd_toolchain_tool_appconfig" "test" {
  toolchain_id = "01234567-89ab-4cde-8123-456789abcdef"
  parameters {
    collection_id = "01234567-89ab-cdef-0123-456789abcdef"
    environment_id = "01234567-89ab-cdef-0123-456789abcdef"
    instance_id = "01234567-89ab-cdef-0123-456789abcdef"
    location = "example"
    name = "example-name"
    resource_group_name = "example"
  }
}
```
