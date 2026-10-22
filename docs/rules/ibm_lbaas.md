# ibm_lbaas

Manages ibm_lbaas. The load balancer's name.

This rule validates `ibm_lbaas` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_lbaas" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `datacenter` | no | — |
| `description` | no | — |
| `health_monitors` | no | — |
| `name` | yes | — |
| `resource_controller_url` | no | — |
| `resource_name` | no | — |
| `resource_status` | no | — |
| `ssl_ciphers` | no | — |
| `status` | no | — |
| `subnets` | yes | — |
| `type` | no | one of: `PUBLIC`, `PRIVATE` |
| `use_system_public_ip_pool` | no | — |
| `vip` | no | — |
| `wait_time_minutes` | no | — |
| `protocols` (block) | no | nested block |
| `protocols.backend_port` | yes | range 1–65535; port format |
| `protocols.backend_protocol` | yes | one of: `HTTP`, `HTTPS`, `TCP` |
| `protocols.frontend_port` | yes | range 1–65535; port format |
| `protocols.frontend_protocol` | yes | one of: `HTTP`, `HTTPS`, `TCP` |
| `protocols.load_balancing_method` | no | one of: `round_robin`, `weighted_round_robin`, `least_connection` |
| `protocols.max_conn` | no | — |
| `protocols.session_stickiness` | no | one of: `SOURCE_IP`, `HTTP_COOKIE` |
| `protocols.tls_certificate_id` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_lbaas_invalid_description` | `description` |
| `ibm_lbaas_invalid_name` | `name` |
| `ibm_lbaas_invalid_ssl_ciphers` | `ssl_ciphers` |
| `ibm_lbaas_invalid_subnets` | `subnets` |
| `ibm_lbaas_invalid_type` | `type` |
| `ibm_lbaas_invalid_use_system_public_ip_pool` | `use_system_public_ip_pool` |
| `ibm_lbaas_invalid_wait_time_minutes` | `wait_time_minutes` |
| `ibm_lbaas_invalid_protocols__backend_port` | `protocols.backend_port` |
| `ibm_lbaas_invalid_protocols__backend_protocol` | `protocols.backend_protocol` |
| `ibm_lbaas_invalid_protocols__frontend_port` | `protocols.frontend_port` |
| `ibm_lbaas_invalid_protocols__frontend_protocol` | `protocols.frontend_protocol` |
| `ibm_lbaas_invalid_protocols__load_balancing_method` | `protocols.load_balancing_method` |
| `ibm_lbaas_invalid_protocols__max_conn` | `protocols.max_conn` |
| `ibm_lbaas_invalid_protocols__session_stickiness` | `protocols.session_stickiness` |
| `ibm_lbaas_invalid_protocols__tls_certificate_id` | `protocols.tls_certificate_id` |

## Example

### Valid

```hcl
resource "ibm_lbaas" "test" {
  name = "example"
  subnets = "example"
}
```
