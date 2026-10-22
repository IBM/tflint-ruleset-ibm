# ibm_lb

Schema validation for the ibm_lb resource.

This rule validates `ibm_lb` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_lb" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `connections` | yes | — |
| `datacenter` | yes | — |
| `dedicated` | no | — |
| `ha_enabled` | no | — |
| `hostname` | no | — |
| `ip_address` | no | — |
| `security_certificate_id` | no | — |
| `ssl_enabled` | no | — |
| `ssl_offload` | no | — |
| `subnet_id` | no | — |
| `tags` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_lb_invalid_connections` | `connections` |
| `ibm_lb_invalid_datacenter` | `datacenter` |
| `ibm_lb_invalid_dedicated` | `dedicated` |
| `ibm_lb_invalid_ha_enabled` | `ha_enabled` |
| `ibm_lb_invalid_security_certificate_id` | `security_certificate_id` |
| `ibm_lb_invalid_ssl_offload` | `ssl_offload` |
| `ibm_lb_invalid_tags` | `tags` |

## Example

### Valid

```hcl
resource "ibm_lb" "test" {
  connections = 1
  datacenter = "example"
}
```
