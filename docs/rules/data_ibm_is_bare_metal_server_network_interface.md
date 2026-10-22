# ibm_is_bare_metal_server_network_interface_data

Reads ibm_is_bare_metal_server_network_interface. The user-defined name for this network interface

This rule validates `ibm_is_bare_metal_server_network_interface` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ImageAllowedUsePatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_bare_metal_server_network_interface_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `allow_interface_to_float` | no | — |
| `allow_ip_spoofing` | no | — |
| `allowed_vlans` | no | — |
| `bare_metal_server` | yes | pattern `^([a-zA-Z_][a-zA-Z0-9_]*|[-+*/%]|&&|\|\||!|==|!=|<|<=|>|>=|~|\bin\b|\(|\)|\[|\]|,|\.|"|'|"|'|\s+|\d+)+$` |
| `enable_infrastructure_nat` | no | — |
| `floating_ips` | no | — |
| `href` | no | — |
| `interface_type` | no | — |
| `mac_address` | no | — |
| `name` | no | — |
| `network_interface` | yes | — |
| `port_speed` | no | — |
| `primary_ip` | no | — |
| `resource_type` | no | — |
| `security_groups` | no | — |
| `status` | no | — |
| `subnet` | no | — |
| `type` | no | — |
| `vlan` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_bare_metal_server_network_interface_data_invalid_bare_metal_server` | `bare_metal_server` |
| `ibm_is_bare_metal_server_network_interface_data_invalid_network_interface` | `network_interface` |

## Example

### Valid

```hcl
data "ibm_is_bare_metal_server_network_interface" "test" {
  bare_metal_server = "example-name"
  network_interface = "example"
}
```
