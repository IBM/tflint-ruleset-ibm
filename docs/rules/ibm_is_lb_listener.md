# ibm_is_lb_listener

Schema validation for the ibm_is_lb_listener resource.

This rule validates `ibm_is_lb_listener` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `LoadBalancerListenerPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_lb_listener" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `accept_proxy_protocol` | no | API: If set to `true`, this listener will accept and forward PROXY protocol information. Supported by load balancers in the `application` family (otherwise always `false`). Additional restrictions: - If this listener has `https_redirect` specified, its `accept_proxy_protocol` value must match the `accept_proxy_protocol` value of the `https_redirect` listener. - If this listener is the target of another |
| `certificate_instance` | no | API: Identifies a certificate instance by a unique property. |
| `connection_limit` | no | range 1–15000 |
| `default_pool` | no | API: The default pool for this listener. If `https_redirect` is specified, the default pool will not be used. The specified pool must: - Belong to this load balancer - Have the same `protocol` as this listener, or have a compatible protocol. At present, the compatible protocols are `http` and `https`. - Not already be the `default_pool` for another listener. - Not already be the `failsafe_policy.target |
| `https_redirect_listener` | no | — |
| `https_redirect_status_code` | no | — |
| `https_redirect_uri` | no | — |
| `idle_connection_timeout` | no | range 50–7200 |
| `lb` | yes | — |
| `listener_id` | no | — |
| `port` | no | range 1–65535; port format |
| `port_max` | no | range 1–65535; port format |
| `port_min` | no | range 1–65535; port format |
| `protocol` | yes | one of: `https`, `http`, `tcp`, `udp`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `related_crn` | no | — |
| `status` | no | — |
| `client_authentication` (block) | no | nested block |
| `client_authentication.certificate_authority` | yes | API: The certificate instance to use for the listener client certificate authority. Specify `null` to remove an existing certificate authority. |
| `client_authentication.certificate_revocation_list` | no | API: A [PEM-encoded](https://www.rfc-editor.org/rfc/rfc7468) (with the label `X509 CRL`) certificate revocation list (CRL) to use for the listener. The CRL must be formatted using the X.509 standard as described in [RFC 5280](https://www.rfc-editor.org/rfc/rfc5280). Specify `null` to remove an existing certificate revocation list. |
| `https_redirect` (block) | no | nested block |
| `https_redirect.http_status_code` | yes | one of: `301`, `302`, `303`, `307`, `308` |
| `https_redirect.uri` | no | pattern `^\/(([a-zA-Z0-9-._~!$&'()*+,;=:@]|%[a-fA-F0-9]{2})+(\/([a-zA-Z0-9-._~!$&'()*+,;=:@]|%[a-fA-F0-9]{2})*)*)?(\?([a-zA-Z0-9-._~!$&'()*+,;=:@\/?]|%[a-fA-F0-9]{2})*)?$` |
| `https_redirect.listener` (block) | yes | nested block |
| `https_redirect.listener.id` | yes | pattern `^[-0-9a-z_]+$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_lb_listener_invalid_accept_proxy_protocol` | `accept_proxy_protocol` |
| `ibm_is_lb_listener_invalid_certificate_instance` | `certificate_instance` |
| `ibm_is_lb_listener_invalid_connection_limit` | `connection_limit` |
| `ibm_is_lb_listener_invalid_default_pool` | `default_pool` |
| `ibm_is_lb_listener_invalid_https_redirect_listener` | `https_redirect_listener` |
| `ibm_is_lb_listener_invalid_https_redirect_status_code` | `https_redirect_status_code` |
| `ibm_is_lb_listener_invalid_https_redirect_uri` | `https_redirect_uri` |
| `ibm_is_lb_listener_invalid_idle_connection_timeout` | `idle_connection_timeout` |
| `ibm_is_lb_listener_invalid_lb` | `lb` |
| `ibm_is_lb_listener_invalid_port` | `port` |
| `ibm_is_lb_listener_invalid_port_max` | `port_max` |
| `ibm_is_lb_listener_invalid_port_min` | `port_min` |
| `ibm_is_lb_listener_invalid_protocol` | `protocol` |
| `ibm_is_lb_listener_invalid_client_authentication__certificate_authority` | `client_authentication.certificate_authority` |
| `ibm_is_lb_listener_invalid_client_authentication__certificate_revocation_list` | `client_authentication.certificate_revocation_list` |
| `ibm_is_lb_listener_invalid_https_redirect__http_status_code` | `https_redirect.http_status_code` |
| `ibm_is_lb_listener_invalid_https_redirect__uri` | `https_redirect.uri` |
| `ibm_is_lb_listener_invalid_https_redirect__listener__id` | `https_redirect.listener.id` |

## Example

### Valid

```hcl
resource "ibm_is_lb_listener" "test" {
  lb = "example"
  protocol = "https"
}
```
