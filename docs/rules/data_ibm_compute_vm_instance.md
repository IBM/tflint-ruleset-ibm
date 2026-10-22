# ibm_compute_vm_instance_data

Schema validation for the ibm_compute_vm_instance data source.

This rule validates `ibm_compute_vm_instance` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_compute_vm_instance_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cores` | no | — |
| `datacenter` | no | — |
| `domain` | yes | — |
| `hostname` | yes | — |
| `ip_address_id` | no | — |
| `ip_address_id_private` | no | — |
| `ipv4_address` | no | — |
| `ipv4_address_private` | no | — |
| `ipv6_address` | no | — |
| `ipv6_address_id` | no | — |
| `last_known_power_state` | no | — |
| `most_recent` | no | — |
| `power_state` | no | — |
| `private_interface_id` | no | — |
| `private_subnet_id` | no | — |
| `public_interface_id` | no | — |
| `public_ipv6_subnet` | no | — |
| `public_ipv6_subnet_id` | no | — |
| `public_subnet_id` | no | — |
| `secondary_ip_addresses` | no | — |
| `secondary_ip_count` | no | — |
| `status` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_compute_vm_instance_data_invalid_domain` | `domain` |
| `ibm_compute_vm_instance_data_invalid_hostname` | `hostname` |
| `ibm_compute_vm_instance_data_invalid_most_recent` | `most_recent` |

## Example

### Valid

```hcl
data "ibm_compute_vm_instance" "test" {
  domain = "example"
  hostname = "example"
}
```
