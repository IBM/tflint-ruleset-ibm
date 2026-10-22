# ibm_cd_toolchain_tool_saucelabs

Manages ibm_cd_toolchain_tool_saucelabs. Name of the tool.

This rule validates `ibm_cd_toolchain_tool_saucelabs` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Toolchain` (`toolchain`).

## Configuration

```hcl
rule "ibm_cd_toolchain_tool_saucelabs" {
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
| `parameters.access_key` | yes | — |
| `parameters.username` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cd_toolchain_tool_saucelabs_invalid_name` | `name` |
| `ibm_cd_toolchain_tool_saucelabs_invalid_toolchain_id` | `toolchain_id` |
| `ibm_cd_toolchain_tool_saucelabs_invalid_parameters__access_key` | `parameters.access_key` |
| `ibm_cd_toolchain_tool_saucelabs_invalid_parameters__username` | `parameters.username` |

## Example

### Valid

```hcl
variable "access_key" {
  sensitive = true
}

resource "ibm_cd_toolchain_tool_saucelabs" "test" {
  toolchain_id = "01234567-89ab-4cde-8123-456789abcdef"
  parameters {
    access_key = var.access_key
    username = "example"
  }
}
```
