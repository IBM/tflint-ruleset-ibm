# ibm_is_instance_template_data

Schema validation for the ibm_is_instance_template data source.

This rule validates `ibm_is_instance_template` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `BareMetalServerDiskPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_instance_template_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `availability` | no | — |
| `availability_policy` | no | — |
| `availability_policy_host_failure` | no | — |
| `boot_volume_attachment` | no | — |
| `catalog_offering` | no | — |
| `cluster_network_attachments` | no | — |
| `confidential_compute_mode` | no | — |
| `crn` | no | — |
| `default_trusted_profile_auto_link` | no | — |
| `default_trusted_profile_target` | no | — |
| `enable_secure_boot` | no | — |
| `href` | no | — |
| `identifier` | no | — |
| `image` | no | — |
| `keys` | no | — |
| `metadata_service` | no | — |
| `metadata_service_enabled` | no | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `network_attachments` | no | — |
| `network_interfaces` | no | — |
| `placement_target` | no | — |
| `primary_network_attachment` | no | — |
| `primary_network_interface` | no | — |
| `profile` | no | — |
| `reservation_affinity` | no | — |
| `resource_group` | no | — |
| `threads_per_core` | no | — |
| `total_volume_bandwidth` | no | — |
| `user_data` | no | — |
| `vcpu` | no | — |
| `volume_attachments` | no | — |
| `volume_bandwidth_qos_mode` | no | — |
| `vpc` | no | — |
| `zone` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_instance_template_data_invalid_identifier` | `identifier` |
| `ibm_is_instance_template_data_invalid_name` | `name` |

## Example

### Valid

```hcl
data "ibm_is_instance_template" "test" {
}
```
