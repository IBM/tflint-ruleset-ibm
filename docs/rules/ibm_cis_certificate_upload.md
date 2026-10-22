# ibm_cis_certificate_upload

Schema validation for the ibm_cis_certificate_upload resource.

This rule validates `ibm_cis_certificate_upload` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/v1/{crn}/zones/{zone_identifier}/custom_certificates/{custom_cert_id}:patch` (`cis`).

## Configuration

```hcl
rule "ibm_cis_certificate_upload" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `bundle_method` | no | one of: `ubiquitous`, `optimal`, `force` |
| `certificate` | yes | API: Certificates |
| `cis_id` | yes | — |
| `custom_cert_id` | no | — |
| `domain_id` | yes | — |
| `expires_on` | no | — |
| `hosts` | no | — |
| `issuer` | no | — |
| `modified_on` | no | — |
| `priority` | no | API: Priority |
| `private_key` | yes | API: Private key |
| `signature` | no | — |
| `status` | no | — |
| `uploaded_on` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cis_certificate_upload_invalid_bundle_method` | `bundle_method` |
| `ibm_cis_certificate_upload_invalid_certificate` | `certificate` |
| `ibm_cis_certificate_upload_invalid_cis_id` | `cis_id` |
| `ibm_cis_certificate_upload_invalid_domain_id` | `domain_id` |
| `ibm_cis_certificate_upload_invalid_priority` | `priority` |
| `ibm_cis_certificate_upload_invalid_private_key` | `private_key` |

## Example

### Valid

```hcl
variable "certificate" {
  sensitive = true
}

variable "private_key" {
  sensitive = true
}

resource "ibm_cis_certificate_upload" "test" {
  certificate = var.certificate
  cis_id = "01234567-89ab-cdef-0123-456789abcdef"
  domain_id = "01234567-89ab-cdef-0123-456789abcdef"
  private_key = var.private_key
}
```
