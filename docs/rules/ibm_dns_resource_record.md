# ibm_dns_resource_record

Manages ibm_dns_resource_record. DNS record name

This rule validates `ibm_dns_resource_record` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `resource_record_input` (`dns-svcs`).

## Configuration

```hcl
rule "ibm_dns_resource_record" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `created_on` | no | — |
| `instance_id` | yes | pattern `^[-0-9a-z_]+$` |
| `modified_on` | no | — |
| `name` | yes | API: Name of the resource record. |
| `port` | no | range 1–65535; port format |
| `preference` | no | range 0–65535 |
| `priority` | no | range 0–65535 |
| `protocol` | no | API: Only used for SRV record. |
| `rdata` | yes | API: Content of the resource record. |
| `resource_record_id` | no | — |
| `service` | no | API: Only used for SRV record. |
| `ttl` | no | API: Time to live in second. |
| `type` | yes | one of: `A`, `AAAA`, `CNAME`, `MX`, `SRV`, `TXT`, `PTR` |
| `weight` | no | range 0–65535 |
| `zone_id` | yes | pattern `^[-0-9a-z_]+$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_dns_resource_record_invalid_instance_id` | `instance_id` |
| `ibm_dns_resource_record_invalid_name` | `name` |
| `ibm_dns_resource_record_invalid_port` | `port` |
| `ibm_dns_resource_record_invalid_preference` | `preference` |
| `ibm_dns_resource_record_invalid_priority` | `priority` |
| `ibm_dns_resource_record_invalid_protocol` | `protocol` |
| `ibm_dns_resource_record_invalid_rdata` | `rdata` |
| `ibm_dns_resource_record_invalid_service` | `service` |
| `ibm_dns_resource_record_invalid_ttl` | `ttl` |
| `ibm_dns_resource_record_invalid_type` | `type` |
| `ibm_dns_resource_record_invalid_weight` | `weight` |
| `ibm_dns_resource_record_invalid_zone_id` | `zone_id` |

## Example

### Valid

```hcl
resource "ibm_dns_resource_record" "test" {
  instance_id = "example-namexxxxxxxxxxxxxxxxxxxxxxxx"
  name = "example"
  rdata = "example"
  type = "A"
  zone_id = "example-namexxxxxxxxxxxxxxxxxxxxxxxx"
}
```
