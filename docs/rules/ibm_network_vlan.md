# ibm_network_vlan

Manages ibm_network_vlan. VLAN name

This rule validates `ibm_network_vlan` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_network_vlan" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `child_resource_count` | no | — |
| `datacenter` | yes | — |
| `name` | no | — |
| `resource_controller_url` | no | — |
| `resource_name` | no | — |
| `router_hostname` | no | — |
| `softlayer_managed` | no | — |
| `subnets` | no | — |
| `tags` | no | — |
| `type` | yes | — |
| `vlan_number` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_network_vlan_invalid_datacenter` | `datacenter` |
| `ibm_network_vlan_invalid_name` | `name` |
| `ibm_network_vlan_invalid_router_hostname` | `router_hostname` |
| `ibm_network_vlan_invalid_tags` | `tags` |
| `ibm_network_vlan_invalid_type` | `type` |

## Example

### Valid

```hcl
resource "ibm_network_vlan" "test" {
  datacenter = "example"
  type = "example"
}
```
