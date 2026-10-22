# ibm_is_vpn_server_data

Reads ibm_is_vpn_server. The unique user-defined name for this VPN server

This rule validates `ibm_is_vpn_server` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `BareMetalServerDiskPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_vpn_server_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_tags` | no | — |
| `certificate` | no | — |
| `client_authentication` | no | — |
| `client_auto_delete` | no | — |
| `client_auto_delete_timeout` | no | — |
| `client_dns_server_ips` | no | — |
| `client_idle_timeout` | no | — |
| `client_ip_pool` | no | — |
| `created_at` | no | — |
| `crn` | no | — |
| `enable_split_tunneling` | no | — |
| `health_reasons` | no | — |
| `health_state` | no | — |
| `hostname` | no | — |
| `href` | no | — |
| `identifier` | no | — |
| `lifecycle_reasons` | no | — |
| `lifecycle_state` | no | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `port` | no | — |
| `private_ips` | no | — |
| `protocol` | no | — |
| `resource_group` | no | — |
| `resource_type` | no | — |
| `security_groups` | no | — |
| `subnets` | no | — |
| `vpc` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_vpn_server_data_invalid_identifier` | `identifier` |
| `ibm_is_vpn_server_data_invalid_name` | `name` |

## Example

### Valid

```hcl
data "ibm_is_vpn_server" "test" {
}
```
