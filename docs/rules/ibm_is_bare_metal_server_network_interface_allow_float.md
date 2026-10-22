# ibm_is_bare_metal_server_network_interface_allow_float

Manages ibm_is_bare_metal_server_network_interface_allow_float. The user-defined name for this network interface

This rule validates `ibm_is_bare_metal_server_network_interface_allow_float` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `VirtualNetworkInterfacePrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_bare_metal_server_network_interface_allow_float" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `allow_interface_to_float` | no | — |
| `allow_ip_spoofing` | no | API: Indicates whether source IP spoofing is allowed on this interface. If `false`, source IP spoofing is prevented on this interface. If `true`, source IP spoofing is allowed on this interface. |
| `bare_metal_server` | yes | pattern `^([a-zA-Z_][a-zA-Z0-9_]*|[-+*/%]|&&|\|\||!|==|!=|<|<=|>|>=|~|\bin\b|\(|\)|\[|\]|,|\.|"|'|"|'|\s+|\d+)+$` |
| `enable_infrastructure_nat` | no | API: If `true`: - The VPC infrastructure performs any needed NAT operations. - `floating_ips` must not have more than one floating IP. If `false`: - Packets are passed unchanged to/from the virtual network interface, allowing the workload to perform any needed NAT operations. - `allow_ip_spoofing` must be `false`. - Can only be attached to a `target` with a `resource_type` of `bare_metal_server_network |
| `floating_bare_metal_server` | no | — |
| `floating_ips` | no | — |
| `href` | no | — |
| `interface_type` | no | — |
| `mac_address` | no | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `network_interface` | no | — |
| `port_speed` | no | — |
| `resource_type` | no | — |
| `security_groups` | no | API: Identifies a security group by a unique property. |
| `status` | no | — |
| `subnet` | yes | API: Identifies a subnet by a unique property. |
| `type` | no | — |
| `vlan` | yes | range 1–4094 |
| `primary_ip` (block) | no | nested block |
| `primary_ip.address` | no | pattern `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$` |
| `primary_ip.auto_delete` | no | API: Indicates whether this reserved IP member will be automatically deleted when either `target` is deleted, or the reserved IP is unbound. |
| `primary_ip.name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `primary_ip.reserved_ip` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_bare_metal_server_network_interface_allow_float_invalid_allow_ip_spoofing` | `allow_ip_spoofing` |
| `ibm_is_bare_metal_server_network_interface_allow_float_invalid_bare_metal_server` | `bare_metal_server` |
| `ibm_is_bare_metal_server_network_interface_allow_float_invalid_enable_infrastructure_nat` | `enable_infrastructure_nat` |
| `ibm_is_bare_metal_server_network_interface_allow_float_invalid_name` | `name` |
| `ibm_is_bare_metal_server_network_interface_allow_float_invalid_security_groups` | `security_groups` |
| `ibm_is_bare_metal_server_network_interface_allow_float_invalid_subnet` | `subnet` |
| `ibm_is_bare_metal_server_network_interface_allow_float_invalid_vlan` | `vlan` |
| `ibm_is_bare_metal_server_network_interface_allow_float_invalid_primary_ip__address` | `primary_ip.address` |
| `ibm_is_bare_metal_server_network_interface_allow_float_invalid_primary_ip__auto_delete` | `primary_ip.auto_delete` |
| `ibm_is_bare_metal_server_network_interface_allow_float_invalid_primary_ip__name` | `primary_ip.name` |
| `ibm_is_bare_metal_server_network_interface_allow_float_invalid_primary_ip__reserved_ip` | `primary_ip.reserved_ip` |

## Example

### Valid

```hcl
resource "ibm_is_bare_metal_server_network_interface_allow_float" "test" {
  bare_metal_server = "example-name"
  subnet = "example"
  vlan = 1
}
```
