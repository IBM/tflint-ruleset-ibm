# ibm_firewall

Schema validation for the ibm_firewall resource.

This rule validates `ibm_firewall` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_firewall" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `firewall_type` | no | one of: `HARDWARE_FIREWALL_DEDICATED`, `FORTIGATE_SECURITY_APPLIANCE` |
| `ha_enabled` | no | — |
| `location` | no | — |
| `password` | no | — |
| `primary_ip` | no | — |
| `public_vlan_id` | yes | — |
| `tags` | no | — |
| `username` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_firewall_invalid_firewall_type` | `firewall_type` |
| `ibm_firewall_invalid_ha_enabled` | `ha_enabled` |
| `ibm_firewall_invalid_public_vlan_id` | `public_vlan_id` |
| `ibm_firewall_invalid_tags` | `tags` |

## Example

### Valid

```hcl
resource "ibm_firewall" "test" {
  public_vlan_id = 1
}
```
