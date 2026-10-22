# ibm_is_bare_metal_server_network_attachment

Manages ibm_is_bare_metal_server_network_attachment. The name for this bare metal server network attachment.

This rule validates `ibm_is_bare_metal_server_network_attachment` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ShareMountTargetVirtualNetworkInterfacePrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_bare_metal_server_network_attachment" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `allow_to_float` | no | — |
| `allowed_vlans` | no | range 1–4094 |
| `bare_metal_server` | yes | pattern `^[-0-9a-z_]+$` |
| `created_at` | no | — |
| `floating_bare_metal_server` | no | — |
| `hard_stop` | no | — |
| `href` | no | — |
| `interface_type` | no | one of: `pci`, `vlan` |
| `lifecycle_state` | no | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `network_attachment` | no | — |
| `port_speed` | no | — |
| `resource_type` | no | — |
| `type` | no | — |
| `virtual_network_interface_id` | no | — |
| `vlan` | no | range 1–4094 |
| `virtual_network_interface` (block) | no | nested block |
| `virtual_network_interface.allow_ip_spoofing` | no | API: Indicates whether source IP spoofing is allowed on this interface. If `false`, source IP spoofing is prevented on this interface. If `true`, source IP spoofing is allowed on this interface. |
| `virtual_network_interface.auto_delete` | no | API: Indicates whether this virtual network interface will be automatically deleted when `target` is deleted |
| `virtual_network_interface.enable_infrastructure_nat` | no | API: If `true`: - The VPC infrastructure performs any needed NAT operations. - `floating_ips` must not have more than one floating IP. If `false`: - Packets are passed unchanged to/from the virtual network interface, allowing the workload to perform any needed NAT operations. - `allow_ip_spoofing` must be `false`. - Can only be attached to a `target` with a `resource_type` of `bare_metal_server_network |
| `virtual_network_interface.id` | no | pattern `^[-0-9a-z_]+$` |
| `virtual_network_interface.name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `virtual_network_interface.protocol_state_filtering_mode` | no | one of: `auto`, `enabled`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `virtual_network_interface.resource_group` | no | API: The resource group to use. If unspecified, the account's [default resource group](https://cloud.ibm.com/apidocs/resource-manager#introduction) will be used. |
| `virtual_network_interface.security_groups` | no | API: Identifies a security group by a unique property. |
| `virtual_network_interface.subnet` | no | API: Identifies a subnet by a unique property. |
| `virtual_network_interface.ips` (block) | no | nested block |
| `virtual_network_interface.ips.reserved_ip` | yes | — |
| `virtual_network_interface.primary_ip` (block) | no | nested block |
| `virtual_network_interface.primary_ip.address` | no | pattern `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$` |
| `virtual_network_interface.primary_ip.name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `virtual_network_interface.primary_ip.reserved_ip` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_bare_metal_server_network_attachment_invalid_allow_to_float` | `allow_to_float` |
| `ibm_is_bare_metal_server_network_attachment_invalid_allowed_vlans` | `allowed_vlans` |
| `ibm_is_bare_metal_server_network_attachment_invalid_bare_metal_server` | `bare_metal_server` |
| `ibm_is_bare_metal_server_network_attachment_invalid_hard_stop` | `hard_stop` |
| `ibm_is_bare_metal_server_network_attachment_invalid_interface_type` | `interface_type` |
| `ibm_is_bare_metal_server_network_attachment_invalid_name` | `name` |
| `ibm_is_bare_metal_server_network_attachment_invalid_vlan` | `vlan` |
| `ibm_is_bare_metal_server_network_attachment_invalid_virtual_network_interface__allow_ip_spoofing` | `virtual_network_interface.allow_ip_spoofing` |
| `ibm_is_bare_metal_server_network_attachment_invalid_virtual_network_interface__auto_delete` | `virtual_network_interface.auto_delete` |
| `ibm_is_bare_metal_server_network_attachment_invalid_virtual_network_interface__enable_infrastructure_nat` | `virtual_network_interface.enable_infrastructure_nat` |
| `ibm_is_bare_metal_server_network_attachment_invalid_virtual_network_interface__id` | `virtual_network_interface.id` |
| `ibm_is_bare_metal_server_network_attachment_invalid_virtual_network_interface__name` | `virtual_network_interface.name` |
| `ibm_is_bare_metal_server_network_attachment_invalid_virtual_network_interface__protocol_state_filtering_mode` | `virtual_network_interface.protocol_state_filtering_mode` |
| `ibm_is_bare_metal_server_network_attachment_invalid_virtual_network_interface__resource_group` | `virtual_network_interface.resource_group` |
| `ibm_is_bare_metal_server_network_attachment_invalid_virtual_network_interface__security_groups` | `virtual_network_interface.security_groups` |
| `ibm_is_bare_metal_server_network_attachment_invalid_virtual_network_interface__subnet` | `virtual_network_interface.subnet` |
| `ibm_is_bare_metal_server_network_attachment_invalid_virtual_network_interface__ips__reserved_ip` | `virtual_network_interface.ips.reserved_ip` |
| `ibm_is_bare_metal_server_network_attachment_invalid_virtual_network_interface__primary_ip__address` | `virtual_network_interface.primary_ip.address` |
| `ibm_is_bare_metal_server_network_attachment_invalid_virtual_network_interface__primary_ip__name` | `virtual_network_interface.primary_ip.name` |
| `ibm_is_bare_metal_server_network_attachment_invalid_virtual_network_interface__primary_ip__reserved_ip` | `virtual_network_interface.primary_ip.reserved_ip` |

## Example

### Valid

```hcl
resource "ibm_is_bare_metal_server_network_attachment" "test" {
  bare_metal_server = "example-name"
}
```
