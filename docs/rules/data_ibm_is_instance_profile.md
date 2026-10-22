# ibm_is_instance_profile_data

Schema validation for the ibm_is_instance_profile data source.

This rule validates `ibm_is_instance_profile` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `BareMetalServerDiskPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_instance_profile_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `architecture` | no | — |
| `architecture_type` | no | — |
| `architecture_values` | no | — |
| `availability_class` | no | — |
| `bandwidth` | no | — |
| `cluster_network_attachment_count` | no | — |
| `confidential_compute_modes` | no | — |
| `disks` | no | — |
| `family` | no | — |
| `gpu_count` | no | — |
| `gpu_manufacturer` | no | — |
| `gpu_memory` | no | — |
| `gpu_model` | no | — |
| `href` | no | — |
| `memory` | no | — |
| `name` | yes | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `network_attachment_count` | no | — |
| `network_bandwidth_mode` | no | — |
| `network_interface_count` | no | — |
| `numa_count` | no | — |
| `port_speed` | no | — |
| `reservation_terms` | no | — |
| `secure_boot_modes` | no | — |
| `status` | no | — |
| `supported_cluster_network_profiles` | no | — |
| `supported_vcpu_count` | no | — |
| `threads_per_core` | no | — |
| `total_volume_bandwidth` | no | — |
| `vcpu_architecture` | no | — |
| `vcpu_burst_limit` | no | — |
| `vcpu_count` | no | — |
| `vcpu_manufacturer` | no | — |
| `vcpu_percentage` | no | — |
| `volume_bandwidth_qos_modes` | no | — |
| `zones` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_instance_profile_data_invalid_name` | `name` |

## Example

### Valid

```hcl
data "ibm_is_instance_profile" "test" {
  name = "example-name"
}
```
