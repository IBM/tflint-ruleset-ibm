# ibm_is_virtual_network_interface_data

Reads ibm_is_virtual_network_interface. The name for this virtual network interface.

This rule validates `ibm_is_virtual_network_interface` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ShareMountTargetByAccessControlModeSecurityGroup` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_virtual_network_interface_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_tags` | no | — |
| `allow_ip_spoofing` | no | — |
| `auto_delete` | no | — |
| `created_at` | no | — |
| `crn` | no | — |
| `enable_infrastructure_nat` | no | — |
| `href` | no | — |
| `ips` | no | — |
| `lifecycle_state` | no | — |
| `mac_address` | no | — |
| `name` | no | — |
| `primary_ip` | no | — |
| `protocol_state_filtering_mode` | no | — |
| `resource_group` | no | — |
| `resource_type` | no | — |
| `security_groups` | no | — |
| `subnet` | no | — |
| `tags` | no | — |
| `target` | no | — |
| `virtual_network_interface` | yes | API: The virtual network interface for this target. |
| `vpc` | no | — |
| `zone` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_virtual_network_interface_data_invalid_virtual_network_interface` | `virtual_network_interface` |

## Example

### Valid

```hcl
data "ibm_is_virtual_network_interface" "test" {
  virtual_network_interface = "example"
}
```
