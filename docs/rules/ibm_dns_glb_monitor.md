# ibm_dns_glb_monitor

Manages ibm_dns_glb_monitor. The unique identifier of a service instance.

This rule validates `ibm_dns_glb_monitor` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `monitor_input` (`dns-svcs`).

## Configuration

```hcl
rule "ibm_dns_glb_monitor" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `allow_insecure` | no | API: Do not validate the certificate when monitor use HTTPS. This parameter is currently only valid for HTTPS monitors. |
| `created_on` | no | — |
| `description` | no | API: Descriptive text of the load balancer monitor. |
| `expected_body` | no | API: A case-insensitive sub-string to look for in the response body. If this string is not found, the origin will be marked as unhealthy. This parameter is only valid for HTTP and HTTPS monitors. |
| `expected_codes` | no | one of: `200`, `201`, `202`, `203`, `204`, `205`, `206`, `207`, `208`, `226`, `2xx`, `3xx`, `4xx`, `5xx` |
| `instance_id` | yes | pattern `^[-0-9a-z_]+$` |
| `interval` | no | range 5–3600 |
| `method` | no | one of: `GET`, `HEAD` |
| `modified_on` | no | — |
| `monitor_id` | no | — |
| `name` | yes | API: The name of the load balancer monitor. |
| `path` | no | API: The endpoint path to health check against. This parameter is only valid for HTTP and HTTPS monitors. |
| `port` | no | range 1–65535; port format |
| `retries` | no | range 0–3 |
| `timeout` | no | range 1–10 |
| `type` | no | one of: `HTTP`, `HTTPS`, `TCP` |
| `headers` (block) | no | nested block |
| `headers.name` | yes | API: The name of HTTP request header. |
| `headers.value` | yes | API: The value of HTTP request header. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_dns_glb_monitor_invalid_allow_insecure` | `allow_insecure` |
| `ibm_dns_glb_monitor_invalid_description` | `description` |
| `ibm_dns_glb_monitor_invalid_expected_body` | `expected_body` |
| `ibm_dns_glb_monitor_invalid_expected_codes` | `expected_codes` |
| `ibm_dns_glb_monitor_invalid_instance_id` | `instance_id` |
| `ibm_dns_glb_monitor_invalid_interval` | `interval` |
| `ibm_dns_glb_monitor_invalid_method` | `method` |
| `ibm_dns_glb_monitor_invalid_name` | `name` |
| `ibm_dns_glb_monitor_invalid_path` | `path` |
| `ibm_dns_glb_monitor_invalid_port` | `port` |
| `ibm_dns_glb_monitor_invalid_retries` | `retries` |
| `ibm_dns_glb_monitor_invalid_timeout` | `timeout` |
| `ibm_dns_glb_monitor_invalid_type` | `type` |
| `ibm_dns_glb_monitor_invalid_headers__name` | `headers.name` |
| `ibm_dns_glb_monitor_invalid_headers__value` | `headers.value` |

## Example

### Valid

```hcl
resource "ibm_dns_glb_monitor" "test" {
  instance_id = "example-namexxxxxxxxxxxxxxxxxxxxxxxx"
  name = "example"
}
```
