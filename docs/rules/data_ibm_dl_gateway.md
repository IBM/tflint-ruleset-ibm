# ibm_dl_gateway_data

Reads ibm_dl_gateway. The unique user-defined name for this gateway

This rule validates `ibm_dl_gateway` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Gateway` (`direct_link`).

## Configuration

```hcl
rule "ibm_dl_gateway_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `as_prepends` | no | — |
| `authentication_key` | no | — |
| `bfd_interval` | no | — |
| `bfd_multiplier` | no | — |
| `bfd_status` | no | — |
| `bfd_status_updated_at` | no | — |
| `bgp_asn` | no | — |
| `bgp_base_cidr` | no | — |
| `bgp_cer_cidr` | no | — |
| `bgp_ibm_asn` | no | — |
| `bgp_ibm_cidr` | no | — |
| `bgp_status` | no | — |
| `bgp_status_updated_at` | no | API: Date and time bgp status was updated |
| `change_request` | no | — |
| `completion_notice_reject_reason` | no | — |
| `connection_mode` | no | — |
| `created_at` | no | — |
| `crn` | no | — |
| `cross_connect_router` | no | — |
| `default_export_route_filter` | no | — |
| `default_import_route_filter` | no | — |
| `gateway_vcs` | no | — |
| `global` | no | — |
| `link_status` | no | — |
| `link_status_updated_at` | no | API: Date and time link status was updated |
| `location_display_name` | no | — |
| `location_name` | no | — |
| `macsec_config` | no | — |
| `metered` | no | — |
| `name` | yes | pattern `^([a-zA-Z]|[a-zA-Z][-_a-zA-Z0-9]*[a-zA-Z0-9])$` |
| `operational_status` | no | — |
| `port` | no | — |
| `provider_api_managed` | no | — |
| `resource_group` | no | — |
| `speed_mbps` | no | — |
| `type` | no | — |
| `vlan` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_dl_gateway_data_invalid_bgp_status_updated_at` | `bgp_status_updated_at` |
| `ibm_dl_gateway_data_invalid_link_status_updated_at` | `link_status_updated_at` |
| `ibm_dl_gateway_data_invalid_name` | `name` |

## Example

### Valid

```hcl
data "ibm_dl_gateway" "test" {
  name = "example-name"
}
```
