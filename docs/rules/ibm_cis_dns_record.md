# ibm_cis_dns_record

Manages ibm_cis_dns_record. DNS record name

This rule validates `ibm_cis_dns_record` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/v1/{crn}/zones/{zone_identifier}/dns_records:post` (`cis`).

## Configuration

```hcl
rule "ibm_cis_dns_record" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cis_id` | yes | — |
| `content` | no | API: DNS record content |
| `created_on` | no | — |
| `data` | no | API: For LOC, SRV, CAA, DS records only |
| `domain_id` | yes | — |
| `modified_on` | no | — |
| `name` | no | API: Required for all record types except SRV. |
| `priority` | no | API: For MX records only |
| `proxiable` | no | — |
| `proxied` | no | API: proxied |
| `record_id` | no | — |
| `ttl` | no | API: DNS record TTL value |
| `type` | yes | one of: `A`, `AAAA`, `CNAME`, `NS`, `MX`, `TXT`, `LOC`, `SRV`, `PTR`, `CAA`, `DS` |
| `zone_name` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cis_dns_record_invalid_cis_id` | `cis_id` |
| `ibm_cis_dns_record_invalid_content` | `content` |
| `ibm_cis_dns_record_invalid_data` | `data` |
| `ibm_cis_dns_record_invalid_domain_id` | `domain_id` |
| `ibm_cis_dns_record_invalid_name` | `name` |
| `ibm_cis_dns_record_invalid_priority` | `priority` |
| `ibm_cis_dns_record_invalid_proxied` | `proxied` |
| `ibm_cis_dns_record_invalid_ttl` | `ttl` |
| `ibm_cis_dns_record_invalid_type` | `type` |

## Example

### Valid

```hcl
resource "ibm_cis_dns_record" "test" {
  cis_id = "01234567-89ab-cdef-0123-456789abcdef"
  domain_id = "01234567-89ab-cdef-0123-456789abcdef"
  type = "A"
}
```
