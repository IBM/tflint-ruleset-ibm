# ibm_dns_record

Schema validation for the ibm_dns_record resource.

This rule validates `ibm_dns_record` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_dns_record" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `data` | yes | — |
| `domain_id` | yes | — |
| `expire` | no | — |
| `host` | yes | — |
| `minimum_ttl` | no | — |
| `mx_priority` | no | — |
| `port` | no | range 1–65535; port format |
| `priority` | no | — |
| `protocol` | no | — |
| `refresh` | no | — |
| `responsible_person` | no | — |
| `retry` | no | — |
| `service` | no | — |
| `tags` | no | — |
| `ttl` | yes | — |
| `type` | yes | — |
| `weight` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_dns_record_invalid_data` | `data` |
| `ibm_dns_record_invalid_domain_id` | `domain_id` |
| `ibm_dns_record_invalid_expire` | `expire` |
| `ibm_dns_record_invalid_host` | `host` |
| `ibm_dns_record_invalid_minimum_ttl` | `minimum_ttl` |
| `ibm_dns_record_invalid_mx_priority` | `mx_priority` |
| `ibm_dns_record_invalid_port` | `port` |
| `ibm_dns_record_invalid_priority` | `priority` |
| `ibm_dns_record_invalid_protocol` | `protocol` |
| `ibm_dns_record_invalid_refresh` | `refresh` |
| `ibm_dns_record_invalid_responsible_person` | `responsible_person` |
| `ibm_dns_record_invalid_retry` | `retry` |
| `ibm_dns_record_invalid_service` | `service` |
| `ibm_dns_record_invalid_tags` | `tags` |
| `ibm_dns_record_invalid_ttl` | `ttl` |
| `ibm_dns_record_invalid_type` | `type` |
| `ibm_dns_record_invalid_weight` | `weight` |

## Example

### Valid

```hcl
resource "ibm_dns_record" "test" {
  data = "example"
  domain_id = 1
  host = "example"
  ttl = 1
  type = "example"
}
```
