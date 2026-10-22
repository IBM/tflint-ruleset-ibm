# ibm_is_vpn_server

Manages ibm_is_vpn_server. The user-defined name for this VPN server.

This rule validates `ibm_is_vpn_server` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `VPNServerPrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_vpn_server" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_tags` | no | pattern `^([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-]):([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-])$` |
| `certificate_crn` | yes | crn format |
| `client_auto_delete` | no | — |
| `client_auto_delete_timeout` | no | — |
| `client_dns_server_ips` | no | API: The DNS server addresses that will be provided to VPN clients connected to this VPN server. |
| `client_idle_timeout` | no | range 0–28800 |
| `client_ip_pool` | yes | pattern `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])(\/(3[0-2]|[1-2][0-9]|[0-9]))$` |
| `created_at` | no | — |
| `crn` | no | — |
| `enable_split_tunneling` | no | API: Indicates whether the split tunneling is enabled on this VPN server. |
| `health_reasons` | no | — |
| `health_state` | no | — |
| `hostname` | no | — |
| `href` | no | — |
| `lifecycle_reasons` | no | — |
| `lifecycle_state` | no | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `port` | no | range 1–65535 |
| `private_ips` | no | — |
| `protocol` | no | one of: `tcp`, `udp`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `resource_group` | no | API: The resource group to use. If unspecified, the account's [default resource group](https://cloud.ibm.com/apidocs/resource-manager#introduction) will be used. |
| `resource_type` | no | one of: `vpn_server` |
| `security_groups` | no | API: Identifies a security group by a unique property. |
| `subnets` | yes | API: Identifies a subnet by a unique property. |
| `tags` | no | pattern `^[A-Za-z0-9:_ .-]+$` |
| `vpc` | no | — |
| `vpn_server` | no | — |
| `client_authentication` (block) | yes | nested block |
| `client_authentication.client_ca_crn` | no | crn format |
| `client_authentication.identity_provider` | no | one of: `iam` |
| `client_authentication.method` | yes | one of: `certificate`, `username`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_vpn_server_invalid_access_tags` | `access_tags` |
| `ibm_is_vpn_server_invalid_certificate_crn` | `certificate_crn` |
| `ibm_is_vpn_server_invalid_client_dns_server_ips` | `client_dns_server_ips` |
| `ibm_is_vpn_server_invalid_client_idle_timeout` | `client_idle_timeout` |
| `ibm_is_vpn_server_invalid_client_ip_pool` | `client_ip_pool` |
| `ibm_is_vpn_server_invalid_enable_split_tunneling` | `enable_split_tunneling` |
| `ibm_is_vpn_server_invalid_name` | `name` |
| `ibm_is_vpn_server_invalid_port` | `port` |
| `ibm_is_vpn_server_invalid_protocol` | `protocol` |
| `ibm_is_vpn_server_invalid_resource_group` | `resource_group` |
| `ibm_is_vpn_server_invalid_resource_type` | `resource_type` |
| `ibm_is_vpn_server_invalid_security_groups` | `security_groups` |
| `ibm_is_vpn_server_invalid_subnets` | `subnets` |
| `ibm_is_vpn_server_invalid_tags` | `tags` |
| `ibm_is_vpn_server_invalid_client_authentication__client_ca_crn` | `client_authentication.client_ca_crn` |
| `ibm_is_vpn_server_invalid_client_authentication__identity_provider` | `client_authentication.identity_provider` |
| `ibm_is_vpn_server_invalid_client_authentication__method` | `client_authentication.method` |

## Example

### Valid

```hcl
resource "ibm_is_vpn_server" "test" {
  certificate_crn = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
  client_ip_pool = "10.0.0.0/24"
  subnets = "example"
  client_authentication {
    method = "certificate"
  }
}
```
