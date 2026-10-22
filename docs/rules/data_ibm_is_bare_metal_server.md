# ibm_is_bare_metal_server_data

Reads ibm_is_bare_metal_server. Bare metal server name

This rule validates `ibm_is_bare_metal_server` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `BareMetalServerDiskPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_bare_metal_server_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_tags` | no | — |
| `bandwidth` | no | — |
| `boot_target` | no | — |
| `cpu` | no | — |
| `created_at` | no | — |
| `crn` | no | — |
| `disks` | no | — |
| `enable_secure_boot` | no | — |
| `firmware_update_type_available` | no | — |
| `health_reasons` | no | — |
| `health_state` | no | — |
| `href` | no | — |
| `identifier` | no | — |
| `image` | no | — |
| `keys` | no | — |
| `memory` | no | — |
| `metadata_service` | no | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `network_attachments` | no | — |
| `network_interfaces` | no | — |
| `primary_network_attachment` | no | — |
| `primary_network_interface` | no | — |
| `profile` | no | — |
| `reservation` | no | — |
| `reservation_affinity` | no | — |
| `resource_group` | no | — |
| `resource_type` | no | — |
| `status` | no | — |
| `status_reasons` | no | — |
| `tags` | no | — |
| `trusted_platform_module` | no | — |
| `vpc` | no | — |
| `zone` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_bare_metal_server_data_invalid_identifier` | `identifier` |
| `ibm_is_bare_metal_server_data_invalid_name` | `name` |

## Example

### Valid

```hcl
data "ibm_is_bare_metal_server" "test" {
}
```
