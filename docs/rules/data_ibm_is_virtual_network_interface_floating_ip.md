# ibm_is_virtual_network_interface_floating_ip_data

Reads ibm_is_virtual_network_interface_floating_ip. The name for this floating IP.

This rule validates `ibm_is_virtual_network_interface_floating_ip` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `InstanceNetworkAttachmentPrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_virtual_network_interface_floating_ip_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `address` | no | — |
| `crn` | no | — |
| `floating_ip` | yes | — |
| `href` | no | — |
| `name` | no | — |
| `virtual_network_interface` | yes | API: A virtual network interface for the instance network attachment. This can be specified using an existing virtual network interface, or a prototype object for a new virtual network interface. If an existing virtual network interface is specified, `enable_infrastructure_nat` must be `true`. |
| `deleted` (block) | no | nested block |
| `deleted.more_info` | yes | pattern `^http(s)?:\/\/([^\/?#]*)([^?#]*)(\?([^#]*))?(#(.*))?$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_virtual_network_interface_floating_ip_data_invalid_floating_ip` | `floating_ip` |
| `ibm_is_virtual_network_interface_floating_ip_data_invalid_virtual_network_interface` | `virtual_network_interface` |
| `ibm_is_virtual_network_interface_floating_ip_data_invalid_deleted__more_info` | `deleted.more_info` |

## Example

### Valid

```hcl
data "ibm_is_virtual_network_interface_floating_ip" "test" {
  floating_ip = "10.0.0.1"
  virtual_network_interface = "example"
}
```
