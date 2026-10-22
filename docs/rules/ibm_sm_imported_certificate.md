# ibm_sm_imported_certificate

Manages ibm_sm_imported_certificate. A human-readable name to assign to your secret.

This rule validates `ibm_sm_imported_certificate` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ImportedCertificateManagedCsrResponse` (`secrets-manager-secrets-manager-v2`).

## Configuration

```hcl
rule "ibm_sm_imported_certificate" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `certificate` | no | — |
| `common_name` | no | — |
| `created_at` | no | — |
| `created_by` | no | — |
| `crn` | no | — |
| `csr` | no | — |
| `custom_metadata` | no | — |
| `description` | no | — |
| `downloaded` | no | — |
| `expiration_date` | no | — |
| `intermediate` | no | — |
| `intermediate_included` | no | — |
| `issuer` | no | — |
| `key_algorithm` | no | — |
| `labels` | no | API: Label that you can use to search for secrets in your instance. Label can be between 2-64 characters, including spaces. To protect your privacy, do not use personal data, such as your name or location, as a label for your secret. |
| `locks_total` | no | — |
| `name` | yes | — |
| `private_key` | no | API: The PEM-encoded private key that is associated with the certificate. The data must be formatted on a single line with embedded newline characters. |
| `private_key_included` | no | — |
| `retrieved_at` | no | — |
| `secret_group_id` | no | pattern `^([0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}|default)$` |
| `secret_id` | no | — |
| `secret_type` | no | — |
| `serial_number` | no | — |
| `signing_algorithm` | no | — |
| `state` | no | — |
| `state_description` | no | — |
| `updated_at` | no | — |
| `validity` | no | — |
| `version_custom_metadata` | no | — |
| `versions_total` | no | — |
| `managed_csr` (block) | no | nested block |
| `managed_csr.alt_names` | no | API: With the Subject Alternative Name field, you can specify additional hostnames to be protected by a single SSL certificate. |
| `managed_csr.client_flag` | no | API: This field indicates whether certificate is flagged for client use. |
| `managed_csr.code_signing_flag` | no | API: This field indicates whether certificate is flagged for code signing use. |
| `managed_csr.common_name` | no | pattern `^$|^(\*\.)?(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\-]*[A-Za-z0-9])\.?$` |
| `managed_csr.country` | no | API: The Country (C) values to define in the subject field of the resulting certificate. |
| `managed_csr.email_protection_flag` | no | API: This field indicates whether certificate is flagged for email protection use. |
| `managed_csr.exclude_cn_from_sans` | no | API: This parameter controls whether the common name is excluded from Subject Alternative Names (SANs). If the common name is set to `true`, it is not included in DNS, or email SANs if they apply. This field can be useful if the common name is a human-readable identifier, instead of a hostname or an email address. |
| `managed_csr.ext_key_usage` | no | API: The allowed extended key usage constraint on certificate, in a comma-delimited list. You can find valid values in the [Go x509 package documentation](https://golang.org/pkg/crypto/x509/#ExtKeyUsage). Omit the `ExtKeyUsage` part of the value. Values are not case-sensitive. To specify no key usage constraints, set this field to an empty string. |
| `managed_csr.ext_key_usage_oids` | no | API: A comma-delimited list of extended key usage Object Identifiers (OIDs). |
| `managed_csr.ip_sans` | no | API: The IP Subject Alternative Names to define for the certificate, in a comma-delimited list. |
| `managed_csr.key_bits` | no | range 0–8192 |
| `managed_csr.key_type` | no | one of: `rsa`, `ec`, `ed25519` |
| `managed_csr.key_usage` | no | API: The allowed key usage constraint to define for certificate, in a comma-delimited list. You can find valid values in the [Go x509 package documentation](https://pkg.go.dev/crypto/x509#KeyUsage). Omit the `KeyUsage` part of the value. Values are not case-sensitive. To specify no key usage constraints, set this field to an empty string. |
| `managed_csr.locality` | no | API: The Locality (L) values to define in the subject field of the resulting certificate. |
| `managed_csr.organization` | no | API: The Organization (O) values to define in the subject field of the resulting certificate. |
| `managed_csr.other_sans` | no | API: The custom Object Identifier (OID) or UTF8-string Subject Alternative Names to define for the certificate, in a comma-delimited list. The format is the same as OpenSSL: `&lt;oid&gt;:&lt;type&gt;:&lt;value&gt;` where the current valid type is `UTF8`. |
| `managed_csr.ou` | no | API: The Organizational Unit (OU) values to define in the subject field of the resulting certificate. |
| `managed_csr.policy_identifiers` | no | API: A comma-delimited list of policy Object Identifiers (OIDs). |
| `managed_csr.postal_code` | no | API: The postal code values to define in the subject field of the resulting certificate. |
| `managed_csr.province` | no | API: The Province (ST) values to define in the subject field of the resulting certificate. |
| `managed_csr.require_cn` | no | API: If set to false, makes the common_name field optional while generating a certificate. |
| `managed_csr.rotate_keys` | no | API: This field indicates whether the private key will be rotated. |
| `managed_csr.server_flag` | no | API: This field indicates whether certificate is flagged for server use. |
| `managed_csr.street_address` | no | API: The street address values to define in the subject field of the resulting certificate. |
| `managed_csr.uri_sans` | no | API: The URI Subject Alternative Names to define for the certificate, in a comma-delimited list. |
| `managed_csr.user_ids` | no | API: Specifies the list of requested User ID (OID 0.9.2342.19200300.100.1.1) Subject values to be placed on the signed certificate. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_sm_imported_certificate_invalid_certificate` | `certificate` |
| `ibm_sm_imported_certificate_invalid_custom_metadata` | `custom_metadata` |
| `ibm_sm_imported_certificate_invalid_description` | `description` |
| `ibm_sm_imported_certificate_invalid_intermediate` | `intermediate` |
| `ibm_sm_imported_certificate_invalid_labels` | `labels` |
| `ibm_sm_imported_certificate_invalid_name` | `name` |
| `ibm_sm_imported_certificate_invalid_private_key` | `private_key` |
| `ibm_sm_imported_certificate_invalid_secret_group_id` | `secret_group_id` |
| `ibm_sm_imported_certificate_invalid_version_custom_metadata` | `version_custom_metadata` |
| `ibm_sm_imported_certificate_invalid_managed_csr__alt_names` | `managed_csr.alt_names` |
| `ibm_sm_imported_certificate_invalid_managed_csr__client_flag` | `managed_csr.client_flag` |
| `ibm_sm_imported_certificate_invalid_managed_csr__code_signing_flag` | `managed_csr.code_signing_flag` |
| `ibm_sm_imported_certificate_invalid_managed_csr__common_name` | `managed_csr.common_name` |
| `ibm_sm_imported_certificate_invalid_managed_csr__country` | `managed_csr.country` |
| `ibm_sm_imported_certificate_invalid_managed_csr__email_protection_flag` | `managed_csr.email_protection_flag` |
| `ibm_sm_imported_certificate_invalid_managed_csr__exclude_cn_from_sans` | `managed_csr.exclude_cn_from_sans` |
| `ibm_sm_imported_certificate_invalid_managed_csr__ext_key_usage` | `managed_csr.ext_key_usage` |
| `ibm_sm_imported_certificate_invalid_managed_csr__ext_key_usage_oids` | `managed_csr.ext_key_usage_oids` |
| `ibm_sm_imported_certificate_invalid_managed_csr__ip_sans` | `managed_csr.ip_sans` |
| `ibm_sm_imported_certificate_invalid_managed_csr__key_bits` | `managed_csr.key_bits` |
| `ibm_sm_imported_certificate_invalid_managed_csr__key_type` | `managed_csr.key_type` |
| `ibm_sm_imported_certificate_invalid_managed_csr__key_usage` | `managed_csr.key_usage` |
| `ibm_sm_imported_certificate_invalid_managed_csr__locality` | `managed_csr.locality` |
| `ibm_sm_imported_certificate_invalid_managed_csr__organization` | `managed_csr.organization` |
| `ibm_sm_imported_certificate_invalid_managed_csr__other_sans` | `managed_csr.other_sans` |
| `ibm_sm_imported_certificate_invalid_managed_csr__ou` | `managed_csr.ou` |
| `ibm_sm_imported_certificate_invalid_managed_csr__policy_identifiers` | `managed_csr.policy_identifiers` |
| `ibm_sm_imported_certificate_invalid_managed_csr__postal_code` | `managed_csr.postal_code` |
| `ibm_sm_imported_certificate_invalid_managed_csr__province` | `managed_csr.province` |
| `ibm_sm_imported_certificate_invalid_managed_csr__require_cn` | `managed_csr.require_cn` |
| `ibm_sm_imported_certificate_invalid_managed_csr__rotate_keys` | `managed_csr.rotate_keys` |
| `ibm_sm_imported_certificate_invalid_managed_csr__server_flag` | `managed_csr.server_flag` |
| `ibm_sm_imported_certificate_invalid_managed_csr__street_address` | `managed_csr.street_address` |
| `ibm_sm_imported_certificate_invalid_managed_csr__uri_sans` | `managed_csr.uri_sans` |
| `ibm_sm_imported_certificate_invalid_managed_csr__user_ids` | `managed_csr.user_ids` |

## Example

### Valid

```hcl
resource "ibm_sm_imported_certificate" "test" {
  name = "example"
}
```
