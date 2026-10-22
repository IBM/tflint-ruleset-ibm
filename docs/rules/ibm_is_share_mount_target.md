# ibm_is_share_mount_target

Manages ibm_is_share_mount_target. The user-defined name for this share target.

This rule validates `ibm_is_share_mount_target` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ShareMountTargetVirtualNetworkInterfacePrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_share_mount_target" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_control_mode` | no | — |
| `access_protocol` | no | one of: `nfs4`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `created_at` | no | — |
| `href` | no | — |
| `lifecycle_state` | no | — |
| `mount_path` | no | — |
| `mount_target` | no | — |
| `name` | yes | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `resource_type` | no | — |
| `share` | yes | — |
| `transit_encryption` | no | one of: `ipsec`, `none`, `stunnel`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `vpc` | no | — |
| `virtual_network_interface` (block) | no | nested block |
| `virtual_network_interface.auto_delete` | no | API: Indicates whether this virtual network interface will be automatically deleted when `target` is deleted |
| `virtual_network_interface.id` | no | pattern `^[-0-9a-z_]+$` |
| `virtual_network_interface.name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `virtual_network_interface.protocol_state_filtering_mode` | no | one of: `auto`, `enabled`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `virtual_network_interface.resource_group` | no | API: The resource group to use. If unspecified, the account's [default resource group](https://cloud.ibm.com/apidocs/resource-manager#introduction) will be used. |
| `virtual_network_interface.security_groups` | no | API: Identifies a security group by a unique property. |
| `virtual_network_interface.subnet` | no | API: Identifies a subnet by a unique property. |
| `virtual_network_interface.primary_ip` (block) | no | nested block |
| `virtual_network_interface.primary_ip.address` | no | pattern `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$` |
| `virtual_network_interface.primary_ip.auto_delete` | no | API: Indicates whether this virtual network interface will be automatically deleted when `target` is deleted |
| `virtual_network_interface.primary_ip.name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `virtual_network_interface.primary_ip.reserved_ip` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_share_mount_target_invalid_access_protocol` | `access_protocol` |
| `ibm_is_share_mount_target_invalid_name` | `name` |
| `ibm_is_share_mount_target_invalid_share` | `share` |
| `ibm_is_share_mount_target_invalid_transit_encryption` | `transit_encryption` |
| `ibm_is_share_mount_target_invalid_vpc` | `vpc` |
| `ibm_is_share_mount_target_invalid_virtual_network_interface__auto_delete` | `virtual_network_interface.auto_delete` |
| `ibm_is_share_mount_target_invalid_virtual_network_interface__id` | `virtual_network_interface.id` |
| `ibm_is_share_mount_target_invalid_virtual_network_interface__name` | `virtual_network_interface.name` |
| `ibm_is_share_mount_target_invalid_virtual_network_interface__protocol_state_filtering_mode` | `virtual_network_interface.protocol_state_filtering_mode` |
| `ibm_is_share_mount_target_invalid_virtual_network_interface__resource_group` | `virtual_network_interface.resource_group` |
| `ibm_is_share_mount_target_invalid_virtual_network_interface__security_groups` | `virtual_network_interface.security_groups` |
| `ibm_is_share_mount_target_invalid_virtual_network_interface__subnet` | `virtual_network_interface.subnet` |
| `ibm_is_share_mount_target_invalid_virtual_network_interface__primary_ip__address` | `virtual_network_interface.primary_ip.address` |
| `ibm_is_share_mount_target_invalid_virtual_network_interface__primary_ip__auto_delete` | `virtual_network_interface.primary_ip.auto_delete` |
| `ibm_is_share_mount_target_invalid_virtual_network_interface__primary_ip__name` | `virtual_network_interface.primary_ip.name` |
| `ibm_is_share_mount_target_invalid_virtual_network_interface__primary_ip__reserved_ip` | `virtual_network_interface.primary_ip.reserved_ip` |

## Example

### Valid

```hcl
resource "ibm_is_share_mount_target" "test" {
  name = "example-name"
  share = "example"
}
```
