# ibm_dl_gateway

Manages ibm_dl_gateway. The unique user-defined name for this gateway

This rule validates `ibm_dl_gateway` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Gateway` (`direct_link`).

## Configuration

```hcl
rule "ibm_dl_gateway" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `authentication_key` | no | API: A reference to a [Hyper Protect Crypto Service Standard Key](https://cloud.ibm.com/docs/hs-crypto?topic=hs-crypto-get-started). |
| `bfd_interval` | no | range 300–255000 |
| `bfd_multiplier` | no | range 1–255 |
| `bfd_status` | no | one of: `init`, `up`, `down` |
| `bfd_status_updated_at` | no | API: Date and time bfd status was updated |
| `bgp_asn` | yes | API: Customer BGP ASN |
| `bgp_base_cidr` | no | cidr format |
| `bgp_cer_cidr` | no | cidr format |
| `bgp_ibm_asn` | no | — |
| `bgp_ibm_cidr` | no | cidr format |
| `bgp_status` | no | — |
| `bgp_status_updated_at` | no | — |
| `carrier_name` | no | API: Carrier name. Only set for type=dedicated gateways. |
| `change_request` | no | — |
| `completion_notice_reject_reason` | no | — |
| `connection_mode` | no | one of: `direct`, `transit` |
| `created_at` | no | — |
| `crn` | no | — |
| `cross_connect_router` | no | API: Cross connect router. Only included on type=dedicated gateways. |
| `customer_name` | no | API: Customer name. Only set for type=dedicated gateways. |
| `default_export_route_filter` | no | one of: `permit`, `deny` |
| `default_import_route_filter` | no | one of: `permit`, `deny` |
| `global` | yes | API: Gateways with global routing (`true`) can connect to networks outside their associated region. |
| `link_status` | no | — |
| `link_status_updated_at` | no | — |
| `loa_reject_reason` | no | — |
| `location_display_name` | no | — |
| `location_name` | no | API: Gateway location |
| `metered` | yes | API: Metered billing option. When `true` gateway usage is billed per gigabyte. When `false` there is no per gigabyte usage charge, instead a flat rate is charged for the gateway. |
| `name` | yes | pattern `^([a-zA-Z]|[a-zA-Z][-_a-zA-Z0-9]*[a-zA-Z0-9])$` |
| `operational_status` | no | — |
| `port` | no | API: Port information for type=connect gateways. |
| `provider_api_managed` | no | — |
| `remove_vlan` | no | — |
| `resource_controller_url` | no | — |
| `resource_crn` | no | — |
| `resource_group` | no | API: Resource group reference |
| `resource_group_name` | no | — |
| `resource_name` | no | — |
| `resource_status` | no | — |
| `speed_mbps` | yes | API: Gateway speed in megabits per second |
| `tags` | no | pattern `^[A-Za-z0-9:_ .-]+$` |
| `type` | yes | one of: `dedicated`, `connect` |
| `vlan` | no | range 2–3967 |
| `as_prepends` (block) | no | nested block |
| `as_prepends.id` | no | pattern `^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$` |
| `as_prepends.length` | yes | range 3–10 |
| `as_prepends.policy` | yes | one of: `import`, `export` |
| `as_prepends.prefix` | no | API: Comma separated list of prefixes this AS Prepend applies to. If empty, this applies to all prefixes. |
| `as_prepends.specific_prefixes` | no | pattern `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])(\/(3[0-2]|[1-2][0-9]|[0-9]))$` |
| `export_route_filters` (block) | no | nested block |
| `export_route_filters.action` | yes | — |
| `export_route_filters.before` | no | pattern `^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$` |
| `export_route_filters.ge` | no | — |
| `export_route_filters.le` | no | — |
| `export_route_filters.prefix` | yes | API: Comma separated list of prefixes this AS Prepend applies to. If empty, this applies to all prefixes. |
| `import_route_filters` (block) | no | nested block |
| `import_route_filters.action` | yes | — |
| `import_route_filters.before` | no | pattern `^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$` |
| `import_route_filters.ge` | no | — |
| `import_route_filters.le` | no | — |
| `import_route_filters.prefix` | yes | API: Comma separated list of prefixes this AS Prepend applies to. If empty, this applies to all prefixes. |
| `macsec_config` (block) | no | nested block |
| `macsec_config.active` | yes | API: Indicates if the MACsec feature is currently active (true) or inactive (false) for a direct link. |
| `macsec_config.fallback_cak` | no | — |
| `macsec_config.primary_cak` | yes | — |
| `macsec_config.window_size` | no | range 0–65535 |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_dl_gateway_invalid_authentication_key` | `authentication_key` |
| `ibm_dl_gateway_invalid_bfd_interval` | `bfd_interval` |
| `ibm_dl_gateway_invalid_bfd_multiplier` | `bfd_multiplier` |
| `ibm_dl_gateway_invalid_bfd_status` | `bfd_status` |
| `ibm_dl_gateway_invalid_bfd_status_updated_at` | `bfd_status_updated_at` |
| `ibm_dl_gateway_invalid_bgp_asn` | `bgp_asn` |
| `ibm_dl_gateway_invalid_bgp_base_cidr` | `bgp_base_cidr` |
| `ibm_dl_gateway_invalid_bgp_cer_cidr` | `bgp_cer_cidr` |
| `ibm_dl_gateway_invalid_bgp_ibm_cidr` | `bgp_ibm_cidr` |
| `ibm_dl_gateway_invalid_carrier_name` | `carrier_name` |
| `ibm_dl_gateway_invalid_connection_mode` | `connection_mode` |
| `ibm_dl_gateway_invalid_cross_connect_router` | `cross_connect_router` |
| `ibm_dl_gateway_invalid_customer_name` | `customer_name` |
| `ibm_dl_gateway_invalid_default_export_route_filter` | `default_export_route_filter` |
| `ibm_dl_gateway_invalid_default_import_route_filter` | `default_import_route_filter` |
| `ibm_dl_gateway_invalid_global` | `global` |
| `ibm_dl_gateway_invalid_loa_reject_reason` | `loa_reject_reason` |
| `ibm_dl_gateway_invalid_location_name` | `location_name` |
| `ibm_dl_gateway_invalid_metered` | `metered` |
| `ibm_dl_gateway_invalid_name` | `name` |
| `ibm_dl_gateway_invalid_port` | `port` |
| `ibm_dl_gateway_invalid_remove_vlan` | `remove_vlan` |
| `ibm_dl_gateway_invalid_resource_group` | `resource_group` |
| `ibm_dl_gateway_invalid_speed_mbps` | `speed_mbps` |
| `ibm_dl_gateway_invalid_tags` | `tags` |
| `ibm_dl_gateway_invalid_type` | `type` |
| `ibm_dl_gateway_invalid_vlan` | `vlan` |
| `ibm_dl_gateway_invalid_as_prepends__id` | `as_prepends.id` |
| `ibm_dl_gateway_invalid_as_prepends__length` | `as_prepends.length` |
| `ibm_dl_gateway_invalid_as_prepends__policy` | `as_prepends.policy` |
| `ibm_dl_gateway_invalid_as_prepends__prefix` | `as_prepends.prefix` |
| `ibm_dl_gateway_invalid_as_prepends__specific_prefixes` | `as_prepends.specific_prefixes` |
| `ibm_dl_gateway_invalid_export_route_filters__action` | `export_route_filters.action` |
| `ibm_dl_gateway_invalid_export_route_filters__before` | `export_route_filters.before` |
| `ibm_dl_gateway_invalid_export_route_filters__ge` | `export_route_filters.ge` |
| `ibm_dl_gateway_invalid_export_route_filters__le` | `export_route_filters.le` |
| `ibm_dl_gateway_invalid_export_route_filters__prefix` | `export_route_filters.prefix` |
| `ibm_dl_gateway_invalid_import_route_filters__action` | `import_route_filters.action` |
| `ibm_dl_gateway_invalid_import_route_filters__before` | `import_route_filters.before` |
| `ibm_dl_gateway_invalid_import_route_filters__ge` | `import_route_filters.ge` |
| `ibm_dl_gateway_invalid_import_route_filters__le` | `import_route_filters.le` |
| `ibm_dl_gateway_invalid_import_route_filters__prefix` | `import_route_filters.prefix` |
| `ibm_dl_gateway_invalid_macsec_config__active` | `macsec_config.active` |
| `ibm_dl_gateway_invalid_macsec_config__fallback_cak` | `macsec_config.fallback_cak` |
| `ibm_dl_gateway_invalid_macsec_config__primary_cak` | `macsec_config.primary_cak` |
| `ibm_dl_gateway_invalid_macsec_config__window_size` | `macsec_config.window_size` |

## Example

### Valid

```hcl
resource "ibm_dl_gateway" "test" {
  bgp_asn = 1
  global = true
  metered = true
  name = "example-name"
  speed_mbps = 1
  type = "dedicated"
}
```
