# ibm_cd_toolchain_tool_hostedgit_data

Reads ibm_cd_toolchain_tool_hostedgit. Name of the tool.

This rule validates `ibm_cd_toolchain_tool_hostedgit` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `params:patch:/toolchains/{toolchain_id}/tools/{tool_id}` (`toolchain`).

## Configuration

```hcl
rule "ibm_cd_toolchain_tool_hostedgit_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `crn` | no | — |
| `href` | no | — |
| `name` | no | — |
| `parameters` | no | — |
| `referent` | no | — |
| `resource_group_id` | no | — |
| `state` | no | — |
| `tool_id` | yes | pattern `^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$` |
| `toolchain_crn` | no | — |
| `toolchain_id` | yes | pattern `^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$` |
| `updated_at` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cd_toolchain_tool_hostedgit_data_invalid_tool_id` | `tool_id` |
| `ibm_cd_toolchain_tool_hostedgit_data_invalid_toolchain_id` | `toolchain_id` |

## Example

### Valid

```hcl
data "ibm_cd_toolchain_tool_hostedgit" "test" {
  tool_id = "01234567-89ab-4cde-8123-456789abcdef"
  toolchain_id = "01234567-89ab-4cde-8123-456789abcdef"
}
```
