# ibm_cd_toolchain

Manages ibm_cd_toolchain. Toolchain name.

This rule validates `ibm_cd_toolchain` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ToolchainPrototype` (`toolchain`).

## Configuration

```hcl
rule "ibm_cd_toolchain" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_id` | no | — |
| `created_at` | no | — |
| `created_by` | no | — |
| `crn` | no | — |
| `description` | no | API: Describes the toolchain |
| `href` | no | — |
| `location` | no | — |
| `name` | yes | pattern `^([^\x00-\x7F]|[a-zA-Z0-9-._ ])+$` |
| `resource_group_id` | yes | pattern `^[0-9a-f]{32}$` |
| `tags` | no | — |
| `ui_href` | no | — |
| `updated_at` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cd_toolchain_invalid_description` | `description` |
| `ibm_cd_toolchain_invalid_name` | `name` |
| `ibm_cd_toolchain_invalid_resource_group_id` | `resource_group_id` |
| `ibm_cd_toolchain_invalid_tags` | `tags` |

## Example

### Valid

```hcl
resource "ibm_cd_toolchain" "test" {
  name = "example-name"
  resource_group_id = "0123456789abcdef0123456789abcdef"
}
```
