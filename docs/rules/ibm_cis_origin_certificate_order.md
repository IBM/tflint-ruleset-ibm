# ibm_cis_origin_certificate_order

Schema validation for the ibm_cis_origin_certificate_order resource.

This rule validates `ibm_cis_origin_certificate_order` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `origin_certificate_req` (`cis`).

## Configuration

```hcl
rule "ibm_cis_origin_certificate_order" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `certificate` | no | — |
| `certificate_id` | no | — |
| `cis_id` | yes | — |
| `csr` | yes | API: The Certificate Signing Request (CSR) |
| `domain_id` | yes | — |
| `expires_on` | no | — |
| `hostnames` | yes | API: Array of hostnames or wildcard names (for example, *.example.com) bound to the certificate |
| `private_key` | no | — |
| `request_type` | yes | one of: `origin-rsa`, `origin-ecc` |
| `requested_validity` | yes | one of: `7`, `30`, `90`, `365`, `730`, `1095`, `5475` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cis_origin_certificate_order_invalid_cis_id` | `cis_id` |
| `ibm_cis_origin_certificate_order_invalid_csr` | `csr` |
| `ibm_cis_origin_certificate_order_invalid_domain_id` | `domain_id` |
| `ibm_cis_origin_certificate_order_invalid_hostnames` | `hostnames` |
| `ibm_cis_origin_certificate_order_invalid_request_type` | `request_type` |
| `ibm_cis_origin_certificate_order_invalid_requested_validity` | `requested_validity` |

## Example

### Valid

```hcl
resource "ibm_cis_origin_certificate_order" "test" {
  cis_id = "01234567-89ab-cdef-0123-456789abcdef"
  csr = "example"
  domain_id = "01234567-89ab-cdef-0123-456789abcdef"
  hostnames = "example"
  request_type = "origin-rsa"
  requested_validity = "7"
}
```
