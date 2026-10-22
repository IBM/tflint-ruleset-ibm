# ibm_cd_toolchain_tool_eventnotifications

Manages ibm_cd_toolchain_tool_eventnotifications. Name of the tool.

This rule validates `ibm_cd_toolchain_tool_eventnotifications` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Toolchain` (`toolchain`).

## Configuration

```hcl
rule "ibm_cd_toolchain_tool_eventnotifications" {
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
| `parameters.instance_crn` | yes | crn format |
| `parameters.name` | yes | pattern `^([^\x00-\x7F]|[a-zA-Z0-9-._ ])+$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cd_toolchain_tool_eventnotifications_invalid_name` | `name` |
| `ibm_cd_toolchain_tool_eventnotifications_invalid_toolchain_id` | `toolchain_id` |
| `ibm_cd_toolchain_tool_eventnotifications_invalid_parameters__instance_crn` | `parameters.instance_crn` |
| `ibm_cd_toolchain_tool_eventnotifications_invalid_parameters__name` | `parameters.name` |

## Example

### Valid

```hcl
resource "ibm_cd_toolchain_tool_eventnotifications" "test" {
  toolchain_id = "01234567-89ab-4cde-8123-456789abcdef"
  parameters {
    instance_crn = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
    name = "example-name"
  }
}
```
