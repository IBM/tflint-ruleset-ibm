# ibm_kms_kmip_client_cert

Manages ibm_kms_kmip_client_cert. The name of the KMIP client certificate

This rule validates `ibm_kms_kmip_client_cert` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `CreateKMIPClientCertificateObject` (`key-protect`).

## Configuration

```hcl
rule "ibm_kms_kmip_client_cert" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `adapter_id` | no | — |
| `cert_id` | no | — |
| `certificate` | yes | pattern `^\s*-----BEGIN CERTIFICATE-----[A-Za-z0-9+/\=\r\\n]+-----END CERTIFICATE-----\s*$` |
| `created_at` | no | — |
| `created_by` | no | — |
| `endpoint_type` | no | one of: `public`, `private` |
| `instance_id` | yes | API: The v4 UUID used to uniquely identify the instance in IBM Cloud. |
| `name` | no | pattern `^[a-zA-Z0-9-_]{2,40}$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_kms_kmip_client_cert_invalid_adapter_id` | `adapter_id` |
| `ibm_kms_kmip_client_cert_invalid_certificate` | `certificate` |
| `ibm_kms_kmip_client_cert_invalid_endpoint_type` | `endpoint_type` |
| `ibm_kms_kmip_client_cert_invalid_instance_id` | `instance_id` |
| `ibm_kms_kmip_client_cert_invalid_name` | `name` |

## Example

### Valid

```hcl
variable "certificate" {
  sensitive = true
}

resource "ibm_kms_kmip_client_cert" "test" {
  certificate = var.certificate
  instance_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
