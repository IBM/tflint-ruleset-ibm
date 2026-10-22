# ibm_is_virtual_network_interface_ip_data

Reads ibm_is_virtual_network_interface_ip. The name for this reserved IP.

This rule validates `ibm_is_virtual_network_interface_ip` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ShareMountTargetByAccessControlModeSecurityGroup` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_virtual_network_interface_ip_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `address` | no | — |
| `href` | no | — |
| `id` | no | — |
| `name` | no | — |
| `reserved_ip` | yes | — |
| `resource_type` | no | — |
| `virtual_network_interface` | yes | API: The virtual network interface for this target. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_virtual_network_interface_ip_data_invalid_reserved_ip` | `reserved_ip` |
| `ibm_is_virtual_network_interface_ip_data_invalid_virtual_network_interface` | `virtual_network_interface` |

## Example

### Valid

```hcl
data "ibm_is_virtual_network_interface_ip" "test" {
  reserved_ip = "10.0.0.1"
  virtual_network_interface = "example"
}
```
