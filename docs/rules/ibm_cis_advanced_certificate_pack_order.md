# ibm_cis_advanced_certificate_pack_order

Schema validation for the ibm_cis_advanced_certificate_pack_order resource.

This rule validates `ibm_cis_advanced_certificate_pack_order` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/v2/{crn}/zones/{zone_identifier}/ssl/certificate_packs/order:post` (`cis`).

## Configuration

```hcl
rule "ibm_cis_advanced_certificate_pack_order" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `certificate_authority` | yes | one of: `lets_encrypt`, `google`, `ssl_com` |
| `certificate_id` | no | — |
| `cis_id` | yes | — |
| `cloudflare_branding` | no | API: Indicates whether or not to add Cloudflare Branding for the order. This will add sni.cloudflaressl.com as the Common Name if set to `true`. |
| `domain_id` | yes | — |
| `hosts` | yes | API: Host name |
| `status` | no | — |
| `type` | no | one of: `advanced` |
| `validation_method` | yes | one of: `txt`, `http`, `email` |
| `validity` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cis_advanced_certificate_pack_order_invalid_certificate_authority` | `certificate_authority` |
| `ibm_cis_advanced_certificate_pack_order_invalid_cis_id` | `cis_id` |
| `ibm_cis_advanced_certificate_pack_order_invalid_cloudflare_branding` | `cloudflare_branding` |
| `ibm_cis_advanced_certificate_pack_order_invalid_domain_id` | `domain_id` |
| `ibm_cis_advanced_certificate_pack_order_invalid_hosts` | `hosts` |
| `ibm_cis_advanced_certificate_pack_order_invalid_type` | `type` |
| `ibm_cis_advanced_certificate_pack_order_invalid_validation_method` | `validation_method` |
| `ibm_cis_advanced_certificate_pack_order_invalid_validity` | `validity` |

## Example

### Valid

```hcl
resource "ibm_cis_advanced_certificate_pack_order" "test" {
  certificate_authority = "lets_encrypt"
  cis_id = "01234567-89ab-cdef-0123-456789abcdef"
  domain_id = "01234567-89ab-cdef-0123-456789abcdef"
  hosts = "example"
  validation_method = "txt"
  validity = 1
}
```
