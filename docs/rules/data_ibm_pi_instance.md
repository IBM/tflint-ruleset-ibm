# ibm_pi_instance_data

Schema validation for the ibm_pi_instance data source.

This rule validates `ibm_pi_instance` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `NetworkAddressGroupUpdate` (`power-cloud`).

## Configuration

```hcl
rule "ibm_pi_instance_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `allow_remote_restart` | no | — |
| `crn` | no | — |
| `dedicated_host_id` | no | — |
| `default_trusted_profile` | no | — |
| `deployment_type` | no | — |
| `effective_processor_compatibility_mode` | no | — |
| `fault` | no | — |
| `health_status` | no | — |
| `ibmi_css` | no | — |
| `ibmi_pha` | no | — |
| `ibmi_rds` | no | — |
| `ibmi_rds_users` | no | — |
| `license_repository_capacity` | no | — |
| `max_virtual_cores` | no | — |
| `maxmem` | no | — |
| `maxproc` | no | — |
| `memory` | no | — |
| `metadata_service` | no | — |
| `min_virtual_cores` | no | — |
| `minmem` | no | — |
| `minproc` | no | — |
| `networks` | no | — |
| `pi_cloud_instance_id` | yes | — |
| `pi_instance_id` | no | — |
| `pi_instance_name` | no | — |
| `pin_policy` | no | — |
| `placement_group_id` | no | — |
| `preferred_processor_compatibility_mode` | no | — |
| `processors` | no | — |
| `proctype` | no | — |
| `server_name` | no | — |
| `shared_processor_pool` | no | — |
| `shared_processor_pool_id` | no | — |
| `status` | no | — |
| `storage_connection` | no | — |
| `storage_pool` | no | — |
| `storage_pool_affinity` | no | — |
| `storage_type` | no | — |
| `user_tags` | no | — |
| `virtual_cores_assigned` | no | — |
| `virtual_serial_number` | no | — |
| `volumes` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_instance_data_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |
| `ibm_pi_instance_data_invalid_pi_instance_id` | `pi_instance_id` |
| `ibm_pi_instance_data_invalid_pi_instance_name` | `pi_instance_name` |

## Example

### Valid

```hcl
data "ibm_pi_instance" "test" {
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
