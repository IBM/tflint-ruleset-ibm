# ibm_is_instance_data

Reads ibm_is_instance. Instance name

This rule validates `ibm_is_instance` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `BareMetalServerDiskPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_instance_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_tags` | no | — |
| `availability` | no | — |
| `availability_policy` | no | — |
| `availability_policy_host_failure` | no | — |
| `bandwidth` | no | — |
| `boot_volume` | no | — |
| `catalog_offering` | no | — |
| `cluster_network` | no | — |
| `cluster_network_attachments` | no | — |
| `confidential_compute_mode` | no | — |
| `crn` | no | — |
| `disks` | no | — |
| `enable_secure_boot` | no | — |
| `gpu` | no | — |
| `health_reasons` | no | — |
| `health_state` | no | — |
| `image` | no | — |
| `keys` | no | — |
| `lifecycle_reasons` | no | — |
| `lifecycle_state` | no | — |
| `memory` | no | — |
| `metadata_service` | no | — |
| `metadata_service_enabled` | no | — |
| `name` | yes | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `network_attachments` | no | — |
| `network_interfaces` | no | — |
| `numa_count` | no | — |
| `passphrase` | no | — |
| `password` | no | — |
| `placement_target` | no | — |
| `primary_network_attachment` | no | — |
| `primary_network_interface` | no | — |
| `private_key` | no | — |
| `profile` | no | — |
| `reservation` | no | — |
| `reservation_affinity` | no | — |
| `resource_controller_url` | no | — |
| `resource_crn` | no | — |
| `resource_group` | no | — |
| `resource_group_name` | no | — |
| `resource_name` | no | — |
| `resource_status` | no | — |
| `software_attachments` | no | — |
| `status` | no | — |
| `status_reasons` | no | — |
| `tags` | no | — |
| `threads_per_core` | no | — |
| `total_network_bandwidth` | no | — |
| `total_volume_bandwidth` | no | — |
| `vcpu` | no | — |
| `volume_attachments` | no | — |
| `volume_bandwidth_qos_mode` | no | — |
| `volumes` | no | — |
| `vpc` | no | — |
| `zone` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_instance_data_invalid_name` | `name` |
| `ibm_is_instance_data_invalid_passphrase` | `passphrase` |
| `ibm_is_instance_data_invalid_private_key` | `private_key` |

## Example

### Valid

```hcl
data "ibm_is_instance" "test" {
  name = "example-name"
}
```
