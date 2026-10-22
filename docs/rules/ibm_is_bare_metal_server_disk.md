# ibm_is_bare_metal_server_disk

Manages ibm_is_bare_metal_server_disk. Bare metal server disk name

This rule validates `ibm_is_bare_metal_server_disk` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `BareMetalServerDiskPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_bare_metal_server_disk" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `allowed_use` | no | — |
| `bare_metal_server` | yes | pattern `^([a-zA-Z_][a-zA-Z0-9_]*|[-+*/%]|&&|\|\||!|==|!=|<|<=|>|>=|~|\bin\b|\(|\)|\[|\]|,|\.|"|'|"|'|\s+|\d+)+$` |
| `created_at` | no | — |
| `disk` | yes | — |
| `href` | no | — |
| `interface_type` | no | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `resource_type` | no | — |
| `size` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_bare_metal_server_disk_invalid_bare_metal_server` | `bare_metal_server` |
| `ibm_is_bare_metal_server_disk_invalid_disk` | `disk` |
| `ibm_is_bare_metal_server_disk_invalid_name` | `name` |

## Example

### Valid

```hcl
resource "ibm_is_bare_metal_server_disk" "test" {
  bare_metal_server = "example-name"
  disk = "example"
}
```
