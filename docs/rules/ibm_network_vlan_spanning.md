# ibm_network_vlan_spanning

Schema validation for the ibm_network_vlan_spanning resource.

This rule validates `ibm_network_vlan_spanning` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_network_vlan_spanning" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `vlan_spanning` | yes | one of: `off`, `on` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_network_vlan_spanning_invalid_vlan_spanning` | `vlan_spanning` |

## Example

### Valid

```hcl
resource "ibm_network_vlan_spanning" "test" {
  vlan_spanning = "off"
}
```
