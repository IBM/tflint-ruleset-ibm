# ibm_is_bare_metal_server_profile_data

Reads ibm_is_bare_metal_server_profile. The name for this bare metal server profile

This rule validates `ibm_is_bare_metal_server_profile` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `BareMetalServerDiskPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_bare_metal_server_profile_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `bandwidth` | no | — |
| `console_types` | no | — |
| `cpu_architecture` | no | — |
| `cpu_core_count` | no | — |
| `cpu_socket_count` | no | — |
| `disks` | no | — |
| `family` | no | — |
| `href` | no | — |
| `memory` | no | — |
| `name` | yes | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `network_attachment_count` | no | — |
| `network_interface_count` | no | — |
| `os_architecture` | no | — |
| `reservation_terms` | no | — |
| `resource_type` | no | — |
| `supported_trusted_platform_module_modes` | no | — |
| `virtual_network_interfaces_supported` | no | — |
| `zones` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_bare_metal_server_profile_data_invalid_name` | `name` |

## Example

### Valid

```hcl
data "ibm_is_bare_metal_server_profile" "test" {
  name = "example-name"
}
```
