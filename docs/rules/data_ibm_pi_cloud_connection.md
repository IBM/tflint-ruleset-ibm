# ibm_pi_cloud_connection_data

Schema validation for the ibm_pi_cloud_connection data source.

This rule validates `ibm_pi_cloud_connection` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pi_cloud_connection_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `classic_enabled` | no | — |
| `connection_mode` | no | — |
| `global_routing` | no | — |
| `gre_destination_address` | no | — |
| `gre_source_address` | no | — |
| `ibm_ip_address` | no | — |
| `metered` | no | — |
| `networks` | no | — |
| `pi_cloud_connection_name` | yes | — |
| `pi_cloud_instance_id` | yes | — |
| `port` | no | — |
| `speed` | no | — |
| `status` | no | — |
| `user_ip_address` | no | — |
| `vpc_crns` | no | — |
| `vpc_enabled` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_cloud_connection_data_invalid_pi_cloud_connection_name` | `pi_cloud_connection_name` |
| `ibm_pi_cloud_connection_data_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |

## Example

### Valid

```hcl
data "ibm_pi_cloud_connection" "test" {
  pi_cloud_connection_name = "example"
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
