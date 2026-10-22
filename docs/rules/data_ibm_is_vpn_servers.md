# ibm_is_vpn_servers_data

Schema validation for the ibm_is_vpn_servers data source.

This rule validates `ibm_is_vpn_servers` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_is_vpn_servers_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `resource_group_id` | no | — |
| `vpn_servers` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_vpn_servers_data_invalid_resource_group_id` | `resource_group_id` |

## Example

### Valid

```hcl
data "ibm_is_vpn_servers" "test" {
}
```
