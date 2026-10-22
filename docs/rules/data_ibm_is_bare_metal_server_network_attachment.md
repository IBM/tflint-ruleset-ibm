# ibm_is_bare_metal_server_network_attachment_data

Reads ibm_is_bare_metal_server_network_attachment. The name for this bare metal server network attachment.

This rule validates `ibm_is_bare_metal_server_network_attachment` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ImageAllowedUsePatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_bare_metal_server_network_attachment_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `allow_to_float` | no | — |
| `allowed_vlans` | no | — |
| `bare_metal_server` | yes | pattern `^[-0-9a-z_]+$` |
| `bare_metal_server_network_attachment_id` | no | — |
| `created_at` | no | — |
| `href` | no | — |
| `interface_type` | no | — |
| `lifecycle_state` | no | — |
| `name` | no | — |
| `network_attachment` | yes | — |
| `port_speed` | no | — |
| `primary_ip` | no | — |
| `resource_type` | no | — |
| `subnet` | no | — |
| `type` | no | — |
| `virtual_network_interface` | no | — |
| `vlan` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_bare_metal_server_network_attachment_data_invalid_bare_metal_server` | `bare_metal_server` |
| `ibm_is_bare_metal_server_network_attachment_data_invalid_network_attachment` | `network_attachment` |

## Example

### Valid

```hcl
data "ibm_is_bare_metal_server_network_attachment" "test" {
  bare_metal_server = "example-name"
  network_attachment = "example"
}
```
