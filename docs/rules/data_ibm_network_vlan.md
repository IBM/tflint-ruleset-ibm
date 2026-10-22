# ibm_network_vlan_data

Schema validation for the ibm_network_vlan data source.

This rule validates `ibm_network_vlan` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_network_vlan_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `id` | no | — |
| `name` | no | — |
| `number` | no | — |
| `router_hostname` | no | — |
| `subnets` | no | — |
| `virtual_guests` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_network_vlan_data_invalid_name` | `name` |
| `ibm_network_vlan_data_invalid_number` | `number` |
| `ibm_network_vlan_data_invalid_router_hostname` | `router_hostname` |

## Example

### Valid

```hcl
data "ibm_network_vlan" "test" {
}
```
