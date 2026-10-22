# ibm_cis_mtls

Manages ibm_cis_mtls. Certificate name

This rule validates `ibm_cis_mtls` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/v1/{crn}/zones/{zone_id}/access/certificates:post` (`cis`).

## Configuration

```hcl
rule "ibm_cis_mtls" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `associated_hostnames` | yes | hostname format |
| `certificate` | yes | API: Access certificate |
| `cis_id` | yes | — |
| `created_at` | no | — |
| `domain_id` | yes | — |
| `expires_on` | no | — |
| `id` | no | — |
| `mtls_id` | no | — |
| `name` | yes | API: Access certificate name |
| `updated_at` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cis_mtls_invalid_associated_hostnames` | `associated_hostnames` |
| `ibm_cis_mtls_invalid_certificate` | `certificate` |
| `ibm_cis_mtls_invalid_cis_id` | `cis_id` |
| `ibm_cis_mtls_invalid_domain_id` | `domain_id` |
| `ibm_cis_mtls_invalid_name` | `name` |

## Example

### Valid

```hcl
variable "certificate" {
  sensitive = true
}

resource "ibm_cis_mtls" "test" {
  associated_hostnames = "example.com"
  certificate = var.certificate
  cis_id = "01234567-89ab-cdef-0123-456789abcdef"
  domain_id = "01234567-89ab-cdef-0123-456789abcdef"
  name = "example"
}
```
