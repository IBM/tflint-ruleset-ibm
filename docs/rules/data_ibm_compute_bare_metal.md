# ibm_compute_bare_metal_data

Schema validation for the ibm_compute_bare_metal data source.

This rule validates `ibm_compute_bare_metal` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_compute_bare_metal_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `block_storage_ids` | no | — |
| `datacenter` | no | — |
| `domain` | no | — |
| `file_storage_ids` | no | — |
| `global_identifier` | no | — |
| `hostname` | no | — |
| `hourly_billing` | no | — |
| `ipv6_address` | no | — |
| `ipv6_address_id` | no | — |
| `ipv6_enabled` | no | — |
| `memory` | no | — |
| `most_recent` | no | — |
| `network_speed` | no | — |
| `notes` | no | — |
| `os_reference_code` | no | — |
| `private_ipv4_address` | no | — |
| `private_ipv4_address_id` | no | — |
| `private_network_only` | no | — |
| `private_subnet` | no | — |
| `private_vlan_id` | no | — |
| `public_bandwidth` | no | — |
| `public_ipv4_address` | no | — |
| `public_ipv4_address_id` | no | — |
| `public_subnet` | no | — |
| `public_vlan_id` | no | — |
| `redundant_network` | no | — |
| `redundant_power_supply` | no | — |
| `secondary_ip_addresses` | no | — |
| `secondary_ip_count` | no | — |
| `tags` | no | — |
| `unbonded_network` | no | — |
| `user_metadata` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_compute_bare_metal_data_invalid_domain` | `domain` |
| `ibm_compute_bare_metal_data_invalid_global_identifier` | `global_identifier` |
| `ibm_compute_bare_metal_data_invalid_hostname` | `hostname` |
| `ibm_compute_bare_metal_data_invalid_most_recent` | `most_recent` |

## Example

### Valid

```hcl
data "ibm_compute_bare_metal" "test" {
}
```
