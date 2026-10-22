# ibm_cd_toolchains_data

Reads ibm_cd_toolchains. Exact name of toolchain to look up.

This rule validates `ibm_cd_toolchains` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/toolchains:post` (`toolchain`).

## Configuration

```hcl
rule "ibm_cd_toolchains_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `name` | no | pattern `^([^\x00-\x7F]|[a-zA-Z0-9-._ ])+$` |
| `resource_group_id` | yes | pattern `^[0-9a-f]{32}$` |
| `toolchains` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cd_toolchains_data_invalid_name` | `name` |
| `ibm_cd_toolchains_data_invalid_resource_group_id` | `resource_group_id` |

## Example

### Valid

```hcl
data "ibm_cd_toolchains" "test" {
  resource_group_id = "0123456789abcdef0123456789abcdef"
}
```
