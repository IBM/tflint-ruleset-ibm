# ibm_is_virtual_network_interface_floating_ips_data

Schema validation for the ibm_is_virtual_network_interface_floating_ips data source.

This rule validates `ibm_is_virtual_network_interface_floating_ips` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ShareMountTargetByAccessControlModeSecurityGroup` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_virtual_network_interface_floating_ips_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `floating_ips` | no | — |
| `virtual_network_interface` | yes | API: The virtual network interface for this target. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_virtual_network_interface_floating_ips_data_invalid_virtual_network_interface` | `virtual_network_interface` |

## Example

### Valid

```hcl
data "ibm_is_virtual_network_interface_floating_ips" "test" {
  virtual_network_interface = "example"
}
```
