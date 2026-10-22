# ibm_is_lb_pool

Manages ibm_is_lb_pool. Load Balancer Pool name

This rule validates `ibm_is_lb_pool` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `LoadBalancerPoolPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_lb_pool" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `algorithm` | yes | one of: `round_robin`, `weighted_round_robin`, `least_connections`, `weighted_forwarding`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `health_delay` | yes | — |
| `health_monitor_port` | no | range 1–65535; port format |
| `health_monitor_url` | no | — |
| `health_retries` | yes | — |
| `health_timeout` | yes | — |
| `health_type` | yes | one of: `http`, `tcp`, `https`, `udp` |
| `lb` | yes | — |
| `name` | yes | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `pool_id` | no | — |
| `protocol` | yes | one of: `http`, `tcp`, `https`, `udp`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `provisioning_status` | no | — |
| `proxy_protocol` | no | one of: `disabled`, `v1`, `v2`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `related_crn` | no | — |
| `session_persistence_app_cookie_name` | no | pattern `^[-A-Za-z0-9!#$%&'*+.^_`~|]+$` |
| `session_persistence_http_cookie_name` | no | — |
| `session_persistence_type` | no | one of: `source_ip`, `app_cookie`, `http_cookie` |
| `client_authentication` (block) | no | nested block |
| `client_authentication.certificate_instance` | yes | API: The backend certificate instance to use for client certificate verification. Supported by load balancers with `mtls_supported` set to `true`. The pool must have a `protocol` of `https`. |
| `failsafe_policy` (block) | no | nested block |
| `failsafe_policy.action` | no | one of: `bypass`, `drop`, `fail`, `forward`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `failsafe_policy.target` (block) | no | nested block |
| `failsafe_policy.target.href` | no | API: The URL for this load balancer pool |
| `failsafe_policy.target.id` | no | pattern `^[-0-9a-z_]+$` |
| `health_monitor` (block) | no | nested block |
| `health_monitor.request` (block) | no | nested block |
| `health_monitor.request.body` | no | pattern `^[\t\x20-\x7E\r\n]+$` |
| `health_monitor.request.method` | yes | one of: `get`, `post`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `health_monitor.request.headers` (block) | no | nested block |
| `health_monitor.request.headers.field` | no | pattern `^([a-zA-Z][-a-zA-Z]*[a-zA-Z])$` |
| `health_monitor.request.headers.value` | no | pattern `^[\t\x20-\x7E]+$` |
| `health_monitor.response` (block) | no | nested block |
| `health_monitor.response.body_regex` | no | pattern `^[\t\x20-\x7E\r\n]+$` |
| `health_monitor.response.codes` | no | pattern `^[1-5](\d{2}|XX)$` |
| `server_authentication` (block) | no | nested block |
| `server_authentication.certificate_authority` | no | API: The backend server certificate authority instance to use for server certificate verification. Supported by load balancers with `mtls_supported` set to `true`. The pool must have a `protocol` of `https`. If specified, `verify_certificate` must be `true`. |
| `server_authentication.verify_certificate` | no | API: Indicates whether server certificate verification is enabled. If set to `true`, the backend server certificate is verified by: - `certificate_authority`, if specified. - the system default certificate authorities, if `certificate_authority` is not specified. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_lb_pool_invalid_algorithm` | `algorithm` |
| `ibm_is_lb_pool_invalid_health_delay` | `health_delay` |
| `ibm_is_lb_pool_invalid_health_monitor_port` | `health_monitor_port` |
| `ibm_is_lb_pool_invalid_health_monitor_url` | `health_monitor_url` |
| `ibm_is_lb_pool_invalid_health_retries` | `health_retries` |
| `ibm_is_lb_pool_invalid_health_timeout` | `health_timeout` |
| `ibm_is_lb_pool_invalid_health_type` | `health_type` |
| `ibm_is_lb_pool_invalid_lb` | `lb` |
| `ibm_is_lb_pool_invalid_name` | `name` |
| `ibm_is_lb_pool_invalid_protocol` | `protocol` |
| `ibm_is_lb_pool_invalid_proxy_protocol` | `proxy_protocol` |
| `ibm_is_lb_pool_invalid_session_persistence_app_cookie_name` | `session_persistence_app_cookie_name` |
| `ibm_is_lb_pool_invalid_session_persistence_type` | `session_persistence_type` |
| `ibm_is_lb_pool_invalid_client_authentication__certificate_instance` | `client_authentication.certificate_instance` |
| `ibm_is_lb_pool_invalid_failsafe_policy__action` | `failsafe_policy.action` |
| `ibm_is_lb_pool_invalid_failsafe_policy__target__href` | `failsafe_policy.target.href` |
| `ibm_is_lb_pool_invalid_failsafe_policy__target__id` | `failsafe_policy.target.id` |
| `ibm_is_lb_pool_invalid_health_monitor__request__body` | `health_monitor.request.body` |
| `ibm_is_lb_pool_invalid_health_monitor__request__method` | `health_monitor.request.method` |
| `ibm_is_lb_pool_invalid_health_monitor__request__headers__field` | `health_monitor.request.headers.field` |
| `ibm_is_lb_pool_invalid_health_monitor__request__headers__value` | `health_monitor.request.headers.value` |
| `ibm_is_lb_pool_invalid_health_monitor__response__body_regex` | `health_monitor.response.body_regex` |
| `ibm_is_lb_pool_invalid_health_monitor__response__codes` | `health_monitor.response.codes` |
| `ibm_is_lb_pool_invalid_server_authentication__certificate_authority` | `server_authentication.certificate_authority` |
| `ibm_is_lb_pool_invalid_server_authentication__verify_certificate` | `server_authentication.verify_certificate` |

## Example

### Valid

```hcl
resource "ibm_is_lb_pool" "test" {
  algorithm = "round_robin"
  health_delay = 1
  health_retries = 1
  health_timeout = 1
  health_type = "http"
  lb = "example"
  name = "example-name"
  protocol = "http"
}
```
