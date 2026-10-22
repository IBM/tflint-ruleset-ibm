# ibm_cis_range_app

Schema validation for the ibm_cis_range_app resource.

This rule validates `ibm_cis_range_app` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `range_app_req` (`cis`).

## Configuration

```hcl
rule "ibm_cis_range_app" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `app_id` | no | — |
| `cis_id` | yes | — |
| `created_on` | no | — |
| `dns` | yes | API: Name and type of the DNS record for this application |
| `dns_type` | yes | — |
| `domain_id` | yes | — |
| `edge_ips_connectivity` | no | one of: `ipv4`, `ipv6`, `all` |
| `edge_ips_type` | no | one of: `dynamic` |
| `ip_firewall` | no | API: Enables the IP Firewall for this application. Only available for TCP applications. |
| `modified_on` | no | — |
| `origin_direct` | no | API: IP address and port of the origin for this Range application. If configuring a load balancer, use\ \ 'origin_dns' and 'origin_port'. This cannot be combined with 'origin_dns' and 'origin_port'. |
| `origin_dns` | no | API: DNS record pointing to the origin for this Range application. This is used for configuring a load\ \ balancer. When specifying an individual IP address, use 'origin_direct'. This requires 'origin_port' and\ \ cannot be combined with 'origin_direct'. |
| `origin_port` | no | range 1–65535; port format |
| `protocol` | yes | API: Defines the protocol (and optionally port) for this application. The currently supported protocols are `tcp`, `udp`, `ssh`, `rdp`, and `minecraft`. |
| `proxy_protocol` | no | one of: `off`, `v1`, `v2`, `simple` |
| `tls` | no | one of: `off`, `flexible`, `full`, `strict` |
| `traffic_type` | no | one of: `direct`, `http`, `https` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cis_range_app_invalid_cis_id` | `cis_id` |
| `ibm_cis_range_app_invalid_dns` | `dns` |
| `ibm_cis_range_app_invalid_dns_type` | `dns_type` |
| `ibm_cis_range_app_invalid_domain_id` | `domain_id` |
| `ibm_cis_range_app_invalid_edge_ips_connectivity` | `edge_ips_connectivity` |
| `ibm_cis_range_app_invalid_edge_ips_type` | `edge_ips_type` |
| `ibm_cis_range_app_invalid_ip_firewall` | `ip_firewall` |
| `ibm_cis_range_app_invalid_origin_direct` | `origin_direct` |
| `ibm_cis_range_app_invalid_origin_dns` | `origin_dns` |
| `ibm_cis_range_app_invalid_origin_port` | `origin_port` |
| `ibm_cis_range_app_invalid_protocol` | `protocol` |
| `ibm_cis_range_app_invalid_proxy_protocol` | `proxy_protocol` |
| `ibm_cis_range_app_invalid_tls` | `tls` |
| `ibm_cis_range_app_invalid_traffic_type` | `traffic_type` |

## Example

### Valid

```hcl
resource "ibm_cis_range_app" "test" {
  cis_id = "01234567-89ab-cdef-0123-456789abcdef"
  dns = "example"
  dns_type = "example"
  domain_id = "01234567-89ab-cdef-0123-456789abcdef"
  protocol = "example"
}
```
