# ibm_is_bare_metal_server_network_interface_floating_ips_data

Schema validation for the ibm_is_bare_metal_server_network_interface_floating_ips data source.

This rule validates `ibm_is_bare_metal_server_network_interface_floating_ips` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ClusterNetworkInterfacePrimaryIPPrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_bare_metal_server_network_interface_floating_ips_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `bare_metal_server` | yes | pattern `^([a-zA-Z_][a-zA-Z0-9_]*|[-+*/%]|&&|\|\||!|==|!=|<|<=|>|>=|~|\bin\b|\(|\)|\[|\]|,|\.|"|'|"|'|\s+|\d+)+$` |
| `floating_ips` | no | — |
| `network_interface` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_bare_metal_server_network_interface_floating_ips_data_invalid_bare_metal_server` | `bare_metal_server` |
| `ibm_is_bare_metal_server_network_interface_floating_ips_data_invalid_network_interface` | `network_interface` |

## Example

### Valid

```hcl
data "ibm_is_bare_metal_server_network_interface_floating_ips" "test" {
  bare_metal_server = "example-name"
  network_interface = "example"
}
```
