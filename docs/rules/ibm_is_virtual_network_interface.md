# ibm_is_virtual_network_interface

Manages ibm_is_virtual_network_interface. The name for this virtual network interface.

This rule validates `ibm_is_virtual_network_interface` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `VirtualNetworkInterfacePrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_virtual_network_interface" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_tags` | no | — |
| `allow_ip_spoofing` | no | API: Indicates whether source IP spoofing is allowed on this interface. If `false`, source IP spoofing is prevented on this interface. If `true`, source IP spoofing is allowed on this interface. |
| `auto_delete` | no | API: Indicates whether this virtual network interface will be automatically deleted when `target` is deleted. Must be `false` if the virtual network interface is unbound. |
| `created_at` | no | — |
| `crn` | no | — |
| `enable_infrastructure_nat` | no | API: If `true`: - The VPC infrastructure performs any needed NAT operations. - `floating_ips` must not have more than one floating IP. If `false`: - Packets are passed unchanged to/from the virtual network interface, allowing the workload to perform any needed NAT operations. - `allow_ip_spoofing` must be `false`. - Can only be attached to a `target` with a `resource_type` of `bare_metal_server_network |
| `href` | no | — |
| `lifecycle_state` | no | — |
| `mac_address` | no | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `protocol_state_filtering_mode` | no | one of: `auto`, `enabled`, `disabled`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `resource_group` | no | API: The resource group to use. If unspecified, the account's [default resource group](https://cloud.ibm.com/apidocs/resource-manager#introduction) will be used. |
| `resource_type` | no | — |
| `security_groups` | no | API: Identifies a security group by a unique property. |
| `subnet` | no | API: Identifies a subnet by a unique property. |
| `tags` | no | — |
| `target` | no | — |
| `vpc` | no | — |
| `zone` | no | — |
| `ips` (block) | no | nested block |
| `ips.reserved_ip` | yes | — |
| `primary_ip` (block) | no | nested block |
| `primary_ip.address` | no | pattern `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$` |
| `primary_ip.auto_delete` | no | API: Indicates whether this reserved IP member will be automatically deleted when either `target` is deleted, or the reserved IP is unbound. |
| `primary_ip.name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `primary_ip.reserved_ip` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_virtual_network_interface_invalid_access_tags` | `access_tags` |
| `ibm_is_virtual_network_interface_invalid_allow_ip_spoofing` | `allow_ip_spoofing` |
| `ibm_is_virtual_network_interface_invalid_auto_delete` | `auto_delete` |
| `ibm_is_virtual_network_interface_invalid_enable_infrastructure_nat` | `enable_infrastructure_nat` |
| `ibm_is_virtual_network_interface_invalid_name` | `name` |
| `ibm_is_virtual_network_interface_invalid_protocol_state_filtering_mode` | `protocol_state_filtering_mode` |
| `ibm_is_virtual_network_interface_invalid_resource_group` | `resource_group` |
| `ibm_is_virtual_network_interface_invalid_security_groups` | `security_groups` |
| `ibm_is_virtual_network_interface_invalid_subnet` | `subnet` |
| `ibm_is_virtual_network_interface_invalid_tags` | `tags` |
| `ibm_is_virtual_network_interface_invalid_ips__reserved_ip` | `ips.reserved_ip` |
| `ibm_is_virtual_network_interface_invalid_primary_ip__address` | `primary_ip.address` |
| `ibm_is_virtual_network_interface_invalid_primary_ip__auto_delete` | `primary_ip.auto_delete` |
| `ibm_is_virtual_network_interface_invalid_primary_ip__name` | `primary_ip.name` |
| `ibm_is_virtual_network_interface_invalid_primary_ip__reserved_ip` | `primary_ip.reserved_ip` |

## Example

### Valid

```hcl
resource "ibm_is_virtual_network_interface" "test" {
}
```
