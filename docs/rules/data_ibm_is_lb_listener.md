# ibm_is_lb_listener_data

Schema validation for the ibm_is_lb_listener data source.

This rule validates `ibm_is_lb_listener` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_is_lb_listener_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `accept_proxy_protocol` | no | — |
| `certificate_instance` | no | — |
| `client_authentication` | no | — |
| `connection_limit` | no | — |
| `created_at` | no | — |
| `default_pool` | no | — |
| `href` | no | — |
| `https_redirect` | no | — |
| `idle_connection_timeout` | no | — |
| `lb` | yes | — |
| `listener_id` | yes | — |
| `policies` | no | — |
| `port` | no | — |
| `port_max` | no | — |
| `port_min` | no | — |
| `protocol` | no | — |
| `provisioning_status` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_lb_listener_data_invalid_lb` | `lb` |
| `ibm_is_lb_listener_data_invalid_listener_id` | `listener_id` |

## Example

### Valid

```hcl
data "ibm_is_lb_listener" "test" {
  lb = "example"
  listener_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
