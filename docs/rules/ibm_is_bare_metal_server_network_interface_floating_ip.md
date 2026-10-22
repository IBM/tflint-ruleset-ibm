# ibm_is_bare_metal_server_network_interface_floating_ip

Manages ibm_is_bare_metal_server_network_interface_floating_ip. Name of the floating IP

This rule validates `ibm_is_bare_metal_server_network_interface_floating_ip` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ImageAllowedUsePatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_bare_metal_server_network_interface_floating_ip" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `address` | no | — |
| `bare_metal_server` | yes | pattern `^([a-zA-Z_][a-zA-Z0-9_]*|[-+*/%]|&&|\|\||!|==|!=|<|<=|>|>=|~|\bin\b|\(|\)|\[|\]|,|\.|"|'|"|'|\s+|\d+)+$` |
| `crn` | no | — |
| `floating_ip` | yes | — |
| `name` | no | — |
| `network_interface` | yes | — |
| `status` | no | — |
| `target` | no | — |
| `zone` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_bare_metal_server_network_interface_floating_ip_invalid_bare_metal_server` | `bare_metal_server` |
| `ibm_is_bare_metal_server_network_interface_floating_ip_invalid_floating_ip` | `floating_ip` |
| `ibm_is_bare_metal_server_network_interface_floating_ip_invalid_network_interface` | `network_interface` |

## Example

### Valid

```hcl
resource "ibm_is_bare_metal_server_network_interface_floating_ip" "test" {
  bare_metal_server = "example-name"
  floating_ip = "10.0.0.1"
  network_interface = "example"
}
```
