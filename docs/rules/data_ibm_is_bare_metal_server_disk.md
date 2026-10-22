# ibm_is_bare_metal_server_disk_data

Reads ibm_is_bare_metal_server_disk. The user-defined name for this disk

This rule validates `ibm_is_bare_metal_server_disk` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ImageAllowedUsePatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_bare_metal_server_disk_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `allowed_use` | no | — |
| `bare_metal_server` | yes | pattern `^([a-zA-Z_][a-zA-Z0-9_]*|[-+*/%]|&&|\|\||!|==|!=|<|<=|>|>=|~|\bin\b|\(|\)|\[|\]|,|\.|"|'|"|'|\s+|\d+)+$` |
| `disk` | yes | — |
| `href` | no | — |
| `id` | no | — |
| `interface_type` | no | — |
| `name` | no | — |
| `resource_type` | no | — |
| `size` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_bare_metal_server_disk_data_invalid_bare_metal_server` | `bare_metal_server` |
| `ibm_is_bare_metal_server_disk_data_invalid_disk` | `disk` |

## Example

### Valid

```hcl
data "ibm_is_bare_metal_server_disk" "test" {
  bare_metal_server = "example-name"
  disk = "example"
}
```
