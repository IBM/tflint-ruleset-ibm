# ibm_multi_vlan_firewall

Manages ibm_multi_vlan_firewall. name

This rule validates `ibm_multi_vlan_firewall` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_multi_vlan_firewall" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `addon_configuration` | no | — |
| `datacenter` | yes | — |
| `firewall_type` | yes | one of: `FortiGate Firewall Appliance HA Option`, `FortiGate Security Appliance` |
| `name` | yes | — |
| `password` | no | — |
| `pod` | yes | — |
| `private_ip` | no | — |
| `private_vlan_id` | no | — |
| `public_ip` | no | — |
| `public_ipv6` | no | — |
| `public_vlan_id` | no | — |
| `username` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_multi_vlan_firewall_invalid_addon_configuration` | `addon_configuration` |
| `ibm_multi_vlan_firewall_invalid_datacenter` | `datacenter` |
| `ibm_multi_vlan_firewall_invalid_firewall_type` | `firewall_type` |
| `ibm_multi_vlan_firewall_invalid_name` | `name` |
| `ibm_multi_vlan_firewall_invalid_pod` | `pod` |

## Example

### Valid

```hcl
resource "ibm_multi_vlan_firewall" "test" {
  datacenter = "example"
  firewall_type = "FortiGate Firewall Appliance HA Option"
  name = "example"
  pod = "example"
}
```
