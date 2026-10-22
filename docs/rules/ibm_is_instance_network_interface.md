# ibm_is_instance_network_interface

Manages ibm_is_instance_network_interface. The user-defined name for this network interface.

This rule validates `ibm_is_instance_network_interface` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `NetworkInterfacePrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_instance_network_interface" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `allow_ip_spoofing` | no | API: Indicates whether source IP spoofing is allowed on this instance network interface. If this instance has network attachments, this network interface is a [read-only representation](https://cloud.ibm.com/docs/vpc?topic=vpc-vni-about#vni-old-api-clients) of its corresponding network attachment and its attached virtual network interface, and source IP spoofing is managed on the attached virtual netwo |
| `created_at` | no | — |
| `floating_ip` | no | — |
| `floating_ips` | no | — |
| `href` | no | — |
| `instance` | yes | pattern `^([a-zA-Z_][a-zA-Z0-9_]*|[-+*/%]|&&|\|\||!|==|!=|<|<=|>|>=|~|\bin\b|\(|\)|\[|\]|,|\.|"|'|"|'|\s+|\d+)+$` |
| `name` | yes | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `network_interface` | no | — |
| `port_speed` | no | — |
| `primary_ipv4_address` | no | pattern `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$` |
| `resource_type` | no | — |
| `security_groups` | no | API: Identifies a security group by a unique property. |
| `status` | no | — |
| `subnet` | yes | API: Identifies a subnet by a unique property. |
| `type` | no | — |
| `primary_ip` (block) | no | nested block |
| `primary_ip.address` | no | pattern `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$` |
| `primary_ip.auto_delete` | no | API: Indicates whether this reserved IP member will be automatically deleted when either `target` is deleted, or the reserved IP is unbound. |
| `primary_ip.name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `primary_ip.reserved_ip` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_instance_network_interface_invalid_allow_ip_spoofing` | `allow_ip_spoofing` |
| `ibm_is_instance_network_interface_invalid_floating_ip` | `floating_ip` |
| `ibm_is_instance_network_interface_invalid_instance` | `instance` |
| `ibm_is_instance_network_interface_invalid_name` | `name` |
| `ibm_is_instance_network_interface_invalid_primary_ipv4_address` | `primary_ipv4_address` |
| `ibm_is_instance_network_interface_invalid_security_groups` | `security_groups` |
| `ibm_is_instance_network_interface_invalid_subnet` | `subnet` |
| `ibm_is_instance_network_interface_invalid_primary_ip__address` | `primary_ip.address` |
| `ibm_is_instance_network_interface_invalid_primary_ip__auto_delete` | `primary_ip.auto_delete` |
| `ibm_is_instance_network_interface_invalid_primary_ip__name` | `primary_ip.name` |
| `ibm_is_instance_network_interface_invalid_primary_ip__reserved_ip` | `primary_ip.reserved_ip` |

## Example

### Valid

```hcl
resource "ibm_is_instance_network_interface" "test" {
  instance = "example-name"
  name = "example-name"
  subnet = "example"
}
```
