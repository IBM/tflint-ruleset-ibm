# ibm_cis_origin_auth

Schema validation for the ibm_cis_origin_auth resource.

This rule validates `ibm_cis_origin_auth` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/v1/{crn}/zones/{zone_identifier}/origin_tls_client_auth:post` (`cis`).

## Configuration

```hcl
rule "ibm_cis_origin_auth" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `auth_id` | no | — |
| `cert_id` | no | — |
| `certificate` | yes | API: the zone's leaf certificate |
| `cis_id` | yes | — |
| `domain_id` | yes | — |
| `enabled` | no | — |
| `expires_on` | no | — |
| `hostname` | no | hostname format |
| `level` | yes | — |
| `private_key` | yes | API: the zone's private key |
| `status` | no | — |
| `uploaded_on` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cis_origin_auth_invalid_certificate` | `certificate` |
| `ibm_cis_origin_auth_invalid_cis_id` | `cis_id` |
| `ibm_cis_origin_auth_invalid_domain_id` | `domain_id` |
| `ibm_cis_origin_auth_invalid_enabled` | `enabled` |
| `ibm_cis_origin_auth_invalid_hostname` | `hostname` |
| `ibm_cis_origin_auth_invalid_level` | `level` |
| `ibm_cis_origin_auth_invalid_private_key` | `private_key` |

## Example

### Valid

```hcl
variable "certificate" {
  sensitive = true
}

variable "private_key" {
  sensitive = true
}

resource "ibm_cis_origin_auth" "test" {
  certificate = var.certificate
  cis_id = "01234567-89ab-cdef-0123-456789abcdef"
  domain_id = "01234567-89ab-cdef-0123-456789abcdef"
  level = "example"
  private_key = var.private_key
}
```
