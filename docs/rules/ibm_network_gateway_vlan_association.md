# ibm_network_gateway_vlan_association

Schema validation for the ibm_network_gateway_vlan_association resource.

This rule validates `ibm_network_gateway_vlan_association` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_network_gateway_vlan_association" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `bypass` | no | — |
| `gateway_id` | yes | — |
| `network_vlan_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_network_gateway_vlan_association_invalid_bypass` | `bypass` |
| `ibm_network_gateway_vlan_association_invalid_gateway_id` | `gateway_id` |
| `ibm_network_gateway_vlan_association_invalid_network_vlan_id` | `network_vlan_id` |

## Example

### Valid

```hcl
resource "ibm_network_gateway_vlan_association" "test" {
  gateway_id = 1
  network_vlan_id = 1
}
```
