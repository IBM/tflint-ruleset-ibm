# ibm_is_bare_metal_server

Manages ibm_is_bare_metal_server. Bare metal server name

This rule validates `ibm_is_bare_metal_server` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ShareMountTargetVirtualNetworkInterfacePrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_bare_metal_server" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_tags` | no | pattern `^([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-]):([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-])$` |
| `action` | no | one of: `start`, `restart`, `stop` |
| `bandwidth` | no | — |
| `boot_target` | no | — |
| `cpu` | no | — |
| `crn` | no | — |
| `delete_type` | no | — |
| `disks` | no | — |
| `enable_secure_boot` | no | — |
| `firmware_update_type_available` | no | — |
| `health_reasons` | no | — |
| `health_state` | no | — |
| `href` | no | — |
| `image` | yes | — |
| `keys` | yes | — |
| `memory` | no | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `profile` | yes | — |
| `reservation` | no | — |
| `resource_group` | no | API: The resource group to use. If unspecified, the account's [default resource group](https://cloud.ibm.com/apidocs/resource-manager#introduction) will be used. |
| `resource_type` | no | — |
| `status` | no | — |
| `status_reasons` | no | — |
| `tags` | no | pattern `^[A-Za-z0-9:_ .-]+$` |
| `user_data` | no | API: The [user data](https://cloud.ibm.com/docs/vpc?topic=vpc-user-data) to make available when setting up the bare metal server. |
| `vpc` | no | — |
| `zone` | yes | zone format |
| `default_trusted_profile` (block) | no | nested block |
| `default_trusted_profile.auto_link` | no | — |
| `default_trusted_profile.target` (block) | no | nested block |
| `default_trusted_profile.target.crn` | no | crn format |
| `default_trusted_profile.target.id` | no | pattern `^[-0-9a-z_]+$` |
| `metadata_service` (block) | no | nested block |
| `metadata_service.enabled` | no | — |
| `metadata_service.protocol` | no | — |
| `network_attachments` (block) | no | nested block |
| `network_attachments.allow_to_float` | no | — |
| `network_attachments.allowed_vlans` | no | range 1–4094 |
| `network_attachments.interface_type` | no | — |
| `network_attachments.name` | yes | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `network_attachments.vlan` | no | range 1–4094 |
| `network_attachments.virtual_network_interface` (block) | no | nested block |
| `network_attachments.virtual_network_interface.allow_ip_spoofing` | no | API: Indicates whether source IP spoofing is allowed on this interface. If `false`, source IP spoofing is prevented on this interface. If `true`, source IP spoofing is allowed on this interface. |
| `network_attachments.virtual_network_interface.auto_delete` | no | API: Indicates whether this virtual network interface will be automatically deleted when `target` is deleted |
| `network_attachments.virtual_network_interface.enable_infrastructure_nat` | no | API: If `true`: - The VPC infrastructure performs any needed NAT operations. - `floating_ips` must not have more than one floating IP. If `false`: - Packets are passed unchanged to/from the virtual network interface, allowing the workload to perform any needed NAT operations. - `allow_ip_spoofing` must be `false`. - Can only be attached to a `target` with a `resource_type` of `bare_metal_server_network |
| `network_attachments.virtual_network_interface.id` | no | pattern `^[-0-9a-z_]+$` |
| `network_attachments.virtual_network_interface.name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `network_attachments.virtual_network_interface.protocol_state_filtering_mode` | no | one of: `auto`, `enabled`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `network_attachments.virtual_network_interface.resource_group` | no | API: The resource group to use. If unspecified, the account's [default resource group](https://cloud.ibm.com/apidocs/resource-manager#introduction) will be used. |
| `network_attachments.virtual_network_interface.security_groups` | no | API: Identifies a security group by a unique property. |
| `network_attachments.virtual_network_interface.subnet` | no | API: Identifies a subnet by a unique property. |
| `network_attachments.virtual_network_interface.ips` (block) | no | nested block |
| `network_attachments.virtual_network_interface.ips.reserved_ip` | yes | — |
| `network_attachments.virtual_network_interface.primary_ip` (block) | no | nested block |
| `network_attachments.virtual_network_interface.primary_ip.address` | no | pattern `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$` |
| `network_attachments.virtual_network_interface.primary_ip.name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `network_attachments.virtual_network_interface.primary_ip.reserved_ip` | no | — |
| `network_interfaces` (block) | no | nested block |
| `network_interfaces.allow_interface_to_float` | no | — |
| `network_interfaces.allow_ip_spoofing` | no | API: Indicates whether source IP spoofing is allowed on this interface. If `false`, source IP spoofing is prevented on this interface. If `true`, source IP spoofing is allowed on this interface. |
| `network_interfaces.allowed_vlans` | no | range 1–4094 |
| `network_interfaces.enable_infrastructure_nat` | no | API: If `true`: - The VPC infrastructure performs any needed NAT operations. - `floating_ips` must not have more than one floating IP. If `false`: - Packets are passed unchanged to/from the virtual network interface, allowing the workload to perform any needed NAT operations. - `allow_ip_spoofing` must be `false`. - Can only be attached to a `target` with a `resource_type` of `bare_metal_server_network |
| `network_interfaces.interface_type` | no | — |
| `network_interfaces.name` | yes | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `network_interfaces.security_groups` | no | API: Identifies a security group by a unique property. |
| `network_interfaces.subnet` | yes | API: Identifies a subnet by a unique property. |
| `network_interfaces.vlan` | no | range 1–4094 |
| `network_interfaces.primary_ip` (block) | no | nested block |
| `network_interfaces.primary_ip.address` | no | pattern `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$` |
| `network_interfaces.primary_ip.auto_delete` | no | API: Indicates whether this virtual network interface will be automatically deleted when `target` is deleted |
| `network_interfaces.primary_ip.name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `network_interfaces.primary_ip.reserved_ip` | no | — |
| `primary_network_attachment` (block) | no | nested block |
| `primary_network_attachment.allow_interface_to_float` | no | — |
| `primary_network_attachment.allowed_vlans` | no | range 1–4094 |
| `primary_network_attachment.interface_type` | no | — |
| `primary_network_attachment.name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `primary_network_attachment.vlan` | no | range 1–4094 |
| `primary_network_attachment.virtual_network_interface` (block) | no | nested block |
| `primary_network_attachment.virtual_network_interface.allow_ip_spoofing` | no | API: Indicates whether source IP spoofing is allowed on this interface. If `false`, source IP spoofing is prevented on this interface. If `true`, source IP spoofing is allowed on this interface. |
| `primary_network_attachment.virtual_network_interface.auto_delete` | no | API: Indicates whether this virtual network interface will be automatically deleted when `target` is deleted |
| `primary_network_attachment.virtual_network_interface.enable_infrastructure_nat` | no | API: If `true`: - The VPC infrastructure performs any needed NAT operations. - `floating_ips` must not have more than one floating IP. If `false`: - Packets are passed unchanged to/from the virtual network interface, allowing the workload to perform any needed NAT operations. - `allow_ip_spoofing` must be `false`. - Can only be attached to a `target` with a `resource_type` of `bare_metal_server_network |
| `primary_network_attachment.virtual_network_interface.id` | no | pattern `^[-0-9a-z_]+$` |
| `primary_network_attachment.virtual_network_interface.name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `primary_network_attachment.virtual_network_interface.protocol_state_filtering_mode` | no | one of: `auto`, `enabled`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `primary_network_attachment.virtual_network_interface.resource_group` | no | API: The resource group to use. If unspecified, the account's [default resource group](https://cloud.ibm.com/apidocs/resource-manager#introduction) will be used. |
| `primary_network_attachment.virtual_network_interface.security_groups` | no | API: Identifies a security group by a unique property. |
| `primary_network_attachment.virtual_network_interface.subnet` | no | API: Identifies a subnet by a unique property. |
| `primary_network_attachment.virtual_network_interface.ips` (block) | no | nested block |
| `primary_network_attachment.virtual_network_interface.ips.reserved_ip` | yes | — |
| `primary_network_attachment.virtual_network_interface.primary_ip` (block) | no | nested block |
| `primary_network_attachment.virtual_network_interface.primary_ip.address` | no | pattern `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$` |
| `primary_network_attachment.virtual_network_interface.primary_ip.name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `primary_network_attachment.virtual_network_interface.primary_ip.reserved_ip` | no | — |
| `primary_network_interface` (block) | no | nested block |
| `primary_network_interface.allow_ip_spoofing` | no | API: Indicates whether source IP spoofing is allowed on this interface. If `false`, source IP spoofing is prevented on this interface. If `true`, source IP spoofing is allowed on this interface. |
| `primary_network_interface.allowed_vlans` | no | range 1–4094 |
| `primary_network_interface.enable_infrastructure_nat` | no | API: If `true`: - The VPC infrastructure performs any needed NAT operations. - `floating_ips` must not have more than one floating IP. If `false`: - Packets are passed unchanged to/from the virtual network interface, allowing the workload to perform any needed NAT operations. - `allow_ip_spoofing` must be `false`. - Can only be attached to a `target` with a `resource_type` of `bare_metal_server_network |
| `primary_network_interface.interface_type` | no | — |
| `primary_network_interface.name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `primary_network_interface.security_groups` | no | API: Identifies a security group by a unique property. |
| `primary_network_interface.subnet` | yes | API: Identifies a subnet by a unique property. |
| `primary_network_interface.primary_ip` (block) | no | nested block |
| `primary_network_interface.primary_ip.address` | no | pattern `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$` |
| `primary_network_interface.primary_ip.auto_delete` | no | API: Indicates whether this virtual network interface will be automatically deleted when `target` is deleted |
| `primary_network_interface.primary_ip.name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `primary_network_interface.primary_ip.reserved_ip` | no | — |
| `reservation_affinity` (block) | no | nested block |
| `reservation_affinity.policy` | no | one of: `automatic`, `disabled`, `manual`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `reservation_affinity.pool` (block) | no | nested block |
| `reservation_affinity.pool.id` | no | pattern `^[-0-9a-z_]+$` |
| `trusted_platform_module` (block) | no | nested block |
| `trusted_platform_module.mode` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_bare_metal_server_invalid_access_tags` | `access_tags` |
| `ibm_is_bare_metal_server_invalid_action` | `action` |
| `ibm_is_bare_metal_server_invalid_bandwidth` | `bandwidth` |
| `ibm_is_bare_metal_server_invalid_delete_type` | `delete_type` |
| `ibm_is_bare_metal_server_invalid_enable_secure_boot` | `enable_secure_boot` |
| `ibm_is_bare_metal_server_invalid_image` | `image` |
| `ibm_is_bare_metal_server_invalid_keys` | `keys` |
| `ibm_is_bare_metal_server_invalid_name` | `name` |
| `ibm_is_bare_metal_server_invalid_profile` | `profile` |
| `ibm_is_bare_metal_server_invalid_resource_group` | `resource_group` |
| `ibm_is_bare_metal_server_invalid_tags` | `tags` |
| `ibm_is_bare_metal_server_invalid_user_data` | `user_data` |
| `ibm_is_bare_metal_server_invalid_vpc` | `vpc` |
| `ibm_is_bare_metal_server_invalid_zone` | `zone` |
| `ibm_is_bare_metal_server_invalid_default_trusted_profile__auto_link` | `default_trusted_profile.auto_link` |
| `ibm_is_bare_metal_server_invalid_default_trusted_profile__target__crn` | `default_trusted_profile.target.crn` |
| `ibm_is_bare_metal_server_invalid_default_trusted_profile__target__id` | `default_trusted_profile.target.id` |
| `ibm_is_bare_metal_server_invalid_metadata_service__enabled` | `metadata_service.enabled` |
| `ibm_is_bare_metal_server_invalid_metadata_service__protocol` | `metadata_service.protocol` |
| `ibm_is_bare_metal_server_invalid_network_attachments__allow_to_float` | `network_attachments.allow_to_float` |
| `ibm_is_bare_metal_server_invalid_network_attachments__allowed_vlans` | `network_attachments.allowed_vlans` |
| `ibm_is_bare_metal_server_invalid_network_attachments__interface_type` | `network_attachments.interface_type` |
| `ibm_is_bare_metal_server_invalid_network_attachments__name` | `network_attachments.name` |
| `ibm_is_bare_metal_server_invalid_network_attachments__vlan` | `network_attachments.vlan` |
| `ibm_is_bare_metal_server_invalid_network_attachments__virtual_network_interface__allow_ip_spoofing` | `network_attachments.virtual_network_interface.allow_ip_spoofing` |
| `ibm_is_bare_metal_server_invalid_network_attachments__virtual_network_interface__auto_delete` | `network_attachments.virtual_network_interface.auto_delete` |
| `ibm_is_bare_metal_server_invalid_network_attachments__virtual_network_interface__enable_infrastructure_nat` | `network_attachments.virtual_network_interface.enable_infrastructure_nat` |
| `ibm_is_bare_metal_server_invalid_network_attachments__virtual_network_interface__id` | `network_attachments.virtual_network_interface.id` |
| `ibm_is_bare_metal_server_invalid_network_attachments__virtual_network_interface__name` | `network_attachments.virtual_network_interface.name` |
| `ibm_is_bare_metal_server_invalid_network_attachments__virtual_network_interface__protocol_state_filtering_mode` | `network_attachments.virtual_network_interface.protocol_state_filtering_mode` |
| `ibm_is_bare_metal_server_invalid_network_attachments__virtual_network_interface__resource_group` | `network_attachments.virtual_network_interface.resource_group` |
| `ibm_is_bare_metal_server_invalid_network_attachments__virtual_network_interface__security_groups` | `network_attachments.virtual_network_interface.security_groups` |
| `ibm_is_bare_metal_server_invalid_network_attachments__virtual_network_interface__subnet` | `network_attachments.virtual_network_interface.subnet` |
| `ibm_is_bare_metal_server_invalid_network_attachments__virtual_network_interface__ips__reserved_ip` | `network_attachments.virtual_network_interface.ips.reserved_ip` |
| `ibm_is_bare_metal_server_invalid_network_attachments__virtual_network_interface__primary_ip__address` | `network_attachments.virtual_network_interface.primary_ip.address` |
| `ibm_is_bare_metal_server_invalid_network_attachments__virtual_network_interface__primary_ip__name` | `network_attachments.virtual_network_interface.primary_ip.name` |
| `ibm_is_bare_metal_server_invalid_network_attachments__virtual_network_interface__primary_ip__reserved_ip` | `network_attachments.virtual_network_interface.primary_ip.reserved_ip` |
| `ibm_is_bare_metal_server_invalid_network_interfaces__allow_interface_to_float` | `network_interfaces.allow_interface_to_float` |
| `ibm_is_bare_metal_server_invalid_network_interfaces__allow_ip_spoofing` | `network_interfaces.allow_ip_spoofing` |
| `ibm_is_bare_metal_server_invalid_network_interfaces__allowed_vlans` | `network_interfaces.allowed_vlans` |
| `ibm_is_bare_metal_server_invalid_network_interfaces__enable_infrastructure_nat` | `network_interfaces.enable_infrastructure_nat` |
| `ibm_is_bare_metal_server_invalid_network_interfaces__interface_type` | `network_interfaces.interface_type` |
| `ibm_is_bare_metal_server_invalid_network_interfaces__name` | `network_interfaces.name` |
| `ibm_is_bare_metal_server_invalid_network_interfaces__security_groups` | `network_interfaces.security_groups` |
| `ibm_is_bare_metal_server_invalid_network_interfaces__subnet` | `network_interfaces.subnet` |
| `ibm_is_bare_metal_server_invalid_network_interfaces__vlan` | `network_interfaces.vlan` |
| `ibm_is_bare_metal_server_invalid_network_interfaces__primary_ip__address` | `network_interfaces.primary_ip.address` |
| `ibm_is_bare_metal_server_invalid_network_interfaces__primary_ip__auto_delete` | `network_interfaces.primary_ip.auto_delete` |
| `ibm_is_bare_metal_server_invalid_network_interfaces__primary_ip__name` | `network_interfaces.primary_ip.name` |
| `ibm_is_bare_metal_server_invalid_network_interfaces__primary_ip__reserved_ip` | `network_interfaces.primary_ip.reserved_ip` |
| `ibm_is_bare_metal_server_invalid_primary_network_attachment__allow_interface_to_float` | `primary_network_attachment.allow_interface_to_float` |
| `ibm_is_bare_metal_server_invalid_primary_network_attachment__allowed_vlans` | `primary_network_attachment.allowed_vlans` |
| `ibm_is_bare_metal_server_invalid_primary_network_attachment__interface_type` | `primary_network_attachment.interface_type` |
| `ibm_is_bare_metal_server_invalid_primary_network_attachment__name` | `primary_network_attachment.name` |
| `ibm_is_bare_metal_server_invalid_primary_network_attachment__vlan` | `primary_network_attachment.vlan` |
| `ibm_is_bare_metal_server_invalid_primary_network_attachment__virtual_network_interface__allow_ip_spoofing` | `primary_network_attachment.virtual_network_interface.allow_ip_spoofing` |
| `ibm_is_bare_metal_server_invalid_primary_network_attachment__virtual_network_interface__auto_delete` | `primary_network_attachment.virtual_network_interface.auto_delete` |
| `ibm_is_bare_metal_server_invalid_primary_network_attachment__virtual_network_interface__enable_infrastructure_nat` | `primary_network_attachment.virtual_network_interface.enable_infrastructure_nat` |
| `ibm_is_bare_metal_server_invalid_primary_network_attachment__virtual_network_interface__id` | `primary_network_attachment.virtual_network_interface.id` |
| `ibm_is_bare_metal_server_invalid_primary_network_attachment__virtual_network_interface__name` | `primary_network_attachment.virtual_network_interface.name` |
| `ibm_is_bare_metal_server_invalid_primary_network_attachment__virtual_network_interface__protocol_state_filtering_mode` | `primary_network_attachment.virtual_network_interface.protocol_state_filtering_mode` |
| `ibm_is_bare_metal_server_invalid_primary_network_attachment__virtual_network_interface__resource_group` | `primary_network_attachment.virtual_network_interface.resource_group` |
| `ibm_is_bare_metal_server_invalid_primary_network_attachment__virtual_network_interface__security_groups` | `primary_network_attachment.virtual_network_interface.security_groups` |
| `ibm_is_bare_metal_server_invalid_primary_network_attachment__virtual_network_interface__subnet` | `primary_network_attachment.virtual_network_interface.subnet` |
| `ibm_is_bare_metal_server_invalid_primary_network_attachment__virtual_network_interface__ips__reserved_ip` | `primary_network_attachment.virtual_network_interface.ips.reserved_ip` |
| `ibm_is_bare_metal_server_invalid_primary_network_attachment__virtual_network_interface__primary_ip__address` | `primary_network_attachment.virtual_network_interface.primary_ip.address` |
| `ibm_is_bare_metal_server_invalid_primary_network_attachment__virtual_network_interface__primary_ip__name` | `primary_network_attachment.virtual_network_interface.primary_ip.name` |
| `ibm_is_bare_metal_server_invalid_primary_network_attachment__virtual_network_interface__primary_ip__reserved_ip` | `primary_network_attachment.virtual_network_interface.primary_ip.reserved_ip` |
| `ibm_is_bare_metal_server_invalid_primary_network_interface__allow_ip_spoofing` | `primary_network_interface.allow_ip_spoofing` |
| `ibm_is_bare_metal_server_invalid_primary_network_interface__allowed_vlans` | `primary_network_interface.allowed_vlans` |
| `ibm_is_bare_metal_server_invalid_primary_network_interface__enable_infrastructure_nat` | `primary_network_interface.enable_infrastructure_nat` |
| `ibm_is_bare_metal_server_invalid_primary_network_interface__interface_type` | `primary_network_interface.interface_type` |
| `ibm_is_bare_metal_server_invalid_primary_network_interface__name` | `primary_network_interface.name` |
| `ibm_is_bare_metal_server_invalid_primary_network_interface__security_groups` | `primary_network_interface.security_groups` |
| `ibm_is_bare_metal_server_invalid_primary_network_interface__subnet` | `primary_network_interface.subnet` |
| `ibm_is_bare_metal_server_invalid_primary_network_interface__primary_ip__address` | `primary_network_interface.primary_ip.address` |
| `ibm_is_bare_metal_server_invalid_primary_network_interface__primary_ip__auto_delete` | `primary_network_interface.primary_ip.auto_delete` |
| `ibm_is_bare_metal_server_invalid_primary_network_interface__primary_ip__name` | `primary_network_interface.primary_ip.name` |
| `ibm_is_bare_metal_server_invalid_primary_network_interface__primary_ip__reserved_ip` | `primary_network_interface.primary_ip.reserved_ip` |
| `ibm_is_bare_metal_server_invalid_reservation_affinity__policy` | `reservation_affinity.policy` |
| `ibm_is_bare_metal_server_invalid_reservation_affinity__pool__id` | `reservation_affinity.pool.id` |
| `ibm_is_bare_metal_server_invalid_trusted_platform_module__mode` | `trusted_platform_module.mode` |

## Example

### Valid

```hcl
resource "ibm_is_bare_metal_server" "test" {
  image = "example"
  keys = "example"
  profile = "example"
  zone = "us-south-1"
}
```
