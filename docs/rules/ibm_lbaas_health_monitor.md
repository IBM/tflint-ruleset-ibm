# ibm_lbaas_health_monitor

Schema validation for the ibm_lbaas_health_monitor resource.

This rule validates `ibm_lbaas_health_monitor` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_lbaas_health_monitor" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `interval` | no | — |
| `lbaas_id` | yes | — |
| `max_retries` | no | — |
| `monitor_id` | yes | — |
| `port` | yes | range 1–65535; port format |
| `protocol` | yes | one of: `HTTP`, `HTTPS`, `TCP` |
| `timeout` | no | — |
| `url_path` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_lbaas_health_monitor_invalid_interval` | `interval` |
| `ibm_lbaas_health_monitor_invalid_lbaas_id` | `lbaas_id` |
| `ibm_lbaas_health_monitor_invalid_max_retries` | `max_retries` |
| `ibm_lbaas_health_monitor_invalid_monitor_id` | `monitor_id` |
| `ibm_lbaas_health_monitor_invalid_port` | `port` |
| `ibm_lbaas_health_monitor_invalid_protocol` | `protocol` |
| `ibm_lbaas_health_monitor_invalid_timeout` | `timeout` |
| `ibm_lbaas_health_monitor_invalid_url_path` | `url_path` |

## Example

### Valid

```hcl
resource "ibm_lbaas_health_monitor" "test" {
  lbaas_id = "01234567-89ab-cdef-0123-456789abcdef"
  monitor_id = "01234567-89ab-cdef-0123-456789abcdef"
  port = 80
  protocol = "HTTP"
}
```
