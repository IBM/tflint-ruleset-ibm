# ibm_pi_dhcp_data

Schema validation for the ibm_pi_dhcp data source.

This rule validates `ibm_pi_dhcp` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pi_dhcp_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `dhcp_id` | no | — |
| `leases` | no | — |
| `network_id` | no | — |
| `network_name` | no | — |
| `pi_cloud_instance_id` | yes | — |
| `pi_dhcp_id` | yes | — |
| `status` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_dhcp_data_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |
| `ibm_pi_dhcp_data_invalid_pi_dhcp_id` | `pi_dhcp_id` |

## Example

### Valid

```hcl
data "ibm_pi_dhcp" "test" {
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  pi_dhcp_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
