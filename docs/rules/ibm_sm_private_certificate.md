# ibm_sm_private_certificate

Manages ibm_sm_private_certificate. A human-readable name to assign to your secret.

This rule validates `ibm_sm_private_certificate` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `PrivateCertificatePrototype` (`secrets-manager-secrets-manager-v2`).

## Configuration

```hcl
rule "ibm_sm_private_certificate" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `alt_names` | no | API: With the Subject Alternative Name field, you can specify additional hostnames to be protected by a single SSL certificate. |
| `ca_chain` | no | — |
| `certificate` | no | — |
| `certificate_authority` | no | — |
| `certificate_template` | yes | pattern `^[A-Za-z0-9][A-Za-z0-9]*(?:_?-?\.?[A-Za-z0-9]+)*$` |
| `common_name` | yes | API: The Common Name (CN) represents the server name that is protected by the SSL certificate. |
| `created_at` | no | — |
| `created_by` | no | — |
| `crn` | no | — |
| `csr` | no | pattern `^(-{5}BEGIN.+?-{5}[\s\S]+-{5}END.+?-{5}[\s]*)$` |
| `custom_metadata` | no | API: The secret metadata that a user can customize. |
| `description` | no | API: An extended description of your secret. To protect your privacy, do not use personal data, such as your name or location, as a description for your secret group. |
| `downloaded` | no | — |
| `exclude_cn_from_sans` | no | API: This parameter controls whether the common name is excluded from Subject Alternative Names (SANs). If the common name is set to `true`, it is not included in DNS, or email SANs if they apply. This field can be useful if the common name is a human-readable identifier, instead of a hostname or an email address. |
| `expiration_date` | no | — |
| `format` | no | one of: `pem`, `pem_bundle` |
| `ip_sans` | no | API: The IP Subject Alternative Names to define for the CA certificate, in a comma-delimited list. |
| `issuer` | no | — |
| `issuing_ca` | no | — |
| `key_algorithm` | no | — |
| `labels` | no | API: Label that you can use to search for secrets in your instance. Label can be between 2-64 characters, including spaces. To protect your privacy, do not use personal data, such as your name or location, as a label for your secret. |
| `locks_total` | no | — |
| `name` | yes | pattern `^[A-Za-z0-9_][A-Za-z0-9_]*(?:_*-*\.*[A-Za-z0-9]*)*[A-Za-z0-9]+$` |
| `next_rotation_date` | no | — |
| `other_sans` | no | API: The custom Object Identifier (OID) or UTF8-string Subject Alternative Names to define for the CA certificate. The alternative names must match the values that are specified in the `allowed_other_sans` field in the associated certificate template. The format is the same as OpenSSL: `&lt;oid&gt;:&lt;type&gt;:&lt;value&gt;` where the current valid type is `UTF8`. |
| `private_key` | no | — |
| `private_key_format` | no | one of: `der`, `pkcs8` |
| `retrieved_at` | no | — |
| `revocation_time_rfc3339` | no | — |
| `revocation_time_seconds` | no | — |
| `secret_group_id` | no | pattern `^([0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}|default)$` |
| `secret_id` | no | — |
| `secret_type` | no | — |
| `serial_number` | no | — |
| `signing_algorithm` | no | — |
| `state` | no | — |
| `state_description` | no | — |
| `ttl` | no | pattern `^[0-9]+[s,m,h,d]{0,1}$` |
| `updated_at` | no | — |
| `uri_sans` | no | API: The URI Subject Alternative Names to define for the CA certificate, in a comma-delimited list. |
| `validity` | no | — |
| `version_custom_metadata` | no | API: The secret version metadata that a user can customize. |
| `versions_total` | no | — |
| `rotation` (block) | no | nested block |
| `rotation.auto_rotate` | no | API: This field indicates whether Secrets Manager rotates your secret automatically. The default is `false`. If `auto_rotate` is set to `true` the service rotates your secret based on the defined interval. |
| `rotation.interval` | no | range 1–1000 |
| `rotation.unit` | no | one of: `day`, `month`, `hour` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_sm_private_certificate_invalid_alt_names` | `alt_names` |
| `ibm_sm_private_certificate_invalid_certificate_template` | `certificate_template` |
| `ibm_sm_private_certificate_invalid_common_name` | `common_name` |
| `ibm_sm_private_certificate_invalid_csr` | `csr` |
| `ibm_sm_private_certificate_invalid_custom_metadata` | `custom_metadata` |
| `ibm_sm_private_certificate_invalid_description` | `description` |
| `ibm_sm_private_certificate_invalid_exclude_cn_from_sans` | `exclude_cn_from_sans` |
| `ibm_sm_private_certificate_invalid_format` | `format` |
| `ibm_sm_private_certificate_invalid_ip_sans` | `ip_sans` |
| `ibm_sm_private_certificate_invalid_labels` | `labels` |
| `ibm_sm_private_certificate_invalid_name` | `name` |
| `ibm_sm_private_certificate_invalid_other_sans` | `other_sans` |
| `ibm_sm_private_certificate_invalid_private_key_format` | `private_key_format` |
| `ibm_sm_private_certificate_invalid_secret_group_id` | `secret_group_id` |
| `ibm_sm_private_certificate_invalid_ttl` | `ttl` |
| `ibm_sm_private_certificate_invalid_uri_sans` | `uri_sans` |
| `ibm_sm_private_certificate_invalid_version_custom_metadata` | `version_custom_metadata` |
| `ibm_sm_private_certificate_invalid_rotation__auto_rotate` | `rotation.auto_rotate` |
| `ibm_sm_private_certificate_invalid_rotation__interval` | `rotation.interval` |
| `ibm_sm_private_certificate_invalid_rotation__unit` | `rotation.unit` |

## Example

### Valid

```hcl
resource "ibm_sm_private_certificate" "test" {
  certificate_template = "example-name"
  common_name = "aaaaaaa"
  name = "example-name"
}
```
