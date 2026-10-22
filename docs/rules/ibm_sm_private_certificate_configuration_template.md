# ibm_sm_private_certificate_configuration_template

Manages ibm_sm_private_certificate_configuration_template. A human-readable unique name to assign to your configuration.

This rule validates `ibm_sm_private_certificate_configuration_template` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `PrivateCertificateConfigurationTemplatePrototype` (`secrets-manager-secrets-manager-v2`).

## Configuration

```hcl
rule "ibm_sm_private_certificate_configuration_template" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `allow_any_name` | no | API: This field indicates whether to allow clients to request a private certificate that matches any common name. |
| `allow_bare_domains` | no | API: This field indicates whether to allow clients to request private certificates that match the value of the actual domains on the final certificate. For example, if you specify `example.com` in the `allowed_domains` field, you grant clients the ability to request a certificate that contains the name `example.com` as one of the DNS values on the final certificate. **Important:** In some scenarios, al |
| `allow_glob_domains` | no | API: This field indicates whether to allow glob patterns, for example, `ftp*.example.com`, in the names that are specified in the `allowed_domains` field. If set to `true`, clients are allowed to request private certificates with names that match the glob patterns. |
| `allow_ip_sans` | no | API: This field indicates whether to allow clients to request a private certificate with IP Subject Alternative Names. |
| `allow_localhost` | no | API: This field indicates whether to allow `localhost` to be included as one of the requested common names. |
| `allow_subdomains` | no | API: This field indicates whether to allow clients to request private certificates with common names (CN) that are subdomains of the CNs that are allowed by the other certificate template options. This includes wildcard subdomains. For example, if `allowed_domains` has a value of `example.com` and `allow_subdomains`is set to `true`, then the following subdomains are allowed: `foo.example.com`, `bar.exa |
| `allow_wildcard_certificates` | no | API: This field indicates whether the issuance of certificates with RFC 6125 wildcards in the CN field. When set to false, this field prevents wildcards from being issued even if they can be allowed by an option `allow_glob_domains`. |
| `allowed_domains` | no | API: The domains to define for the certificate template. This property is used along with the `allow_bare_domains` and `allow_subdomains` options. |
| `allowed_domains_template` | no | API: This field indicates whether to allow the domains that are supplied in the `allowed_domains` field to contain access control list (ACL) templates. |
| `allowed_other_sans` | no | API: The custom Object Identifier (OID) or UTF8-string Subject Alternative Names (SANs) to allow for private certificates. The format for each element in the list is the same as OpenSSL: `&lt;oid&gt;:&lt;type&gt;:&lt;value&gt;` where the current valid type is `UTF8`. To allow any value for an OID, use `*` as its value. Alternatively, specify a single `*` to allow any `other_sans` input. |
| `allowed_secret_groups` | no | API: This field scopes the creation of private certificates to only the secret groups that you specify. This field can be supplied as a comma-delimited list of secret group IDs. |
| `allowed_uri_sans` | no | API: The URI Subject Alternative Names to allow for private certificates. Values can contain glob patterns, for example `spiffe://hostname/*`. |
| `basic_constraints_valid_for_non_ca` | no | API: This field indicates whether to mark the Basic Constraints extension of an issued private certificate as valid for non-CA certificates. |
| `certificate_authority` | yes | pattern `^[A-Za-z0-9][A-Za-z0-9]*(?:_?-?\.?[A-Za-z0-9]+)*$` |
| `client_flag` | no | API: This field indicates whether private certificates are flagged for client use. |
| `code_signing_flag` | no | API: This field indicates whether private certificates are flagged for code signing use. |
| `config_type` | no | — |
| `country` | no | API: The Country (C) values to define in the subject field of the resulting certificate. |
| `created_at` | no | — |
| `created_by` | no | — |
| `email_protection_flag` | no | API: This field indicates whether private certificates are flagged for email protection use. |
| `enforce_hostnames` | no | API: This field indicates whether to enforce only valid hostnames for common names, DNS Subject Alternative Names, and the host section of email addresses. |
| `ext_key_usage` | no | pattern `^[a-zA-Z]+$` |
| `ext_key_usage_oids` | no | API: A list of extended key usage Object Identifiers (OIDs). |
| `key_bits` | no | range 224–4096 |
| `key_type` | no | one of: `rsa`, `ec` |
| `key_usage` | no | pattern `^[a-zA-Z]+$` |
| `locality` | no | API: The Locality (L) values to define in the subject field of the resulting certificate. |
| `max_ttl` | no | pattern `^[0-9]+[s,m,h,d]{0,1}$` |
| `max_ttl_seconds` | no | — |
| `name` | yes | pattern `^[A-Za-z0-9_][A-Za-z0-9_]*(?:_*-*\.*[A-Za-z0-9]*)*[A-Za-z0-9]+$` |
| `not_before_duration` | no | pattern `^[0-9]+[s,m,h,d]{0,1}$` |
| `not_before_duration_seconds` | no | — |
| `organization` | no | API: The Organization (O) values to define in the subject field of the resulting certificate. |
| `ou` | no | API: The Organizational Unit (OU) values to define in the subject field of the resulting certificate. |
| `policy_identifiers` | no | API: A list of policy Object Identifiers (OIDs). |
| `postal_code` | no | API: The postal code values to define in the subject field of the resulting certificate. |
| `province` | no | API: The Province (ST) values to define in the subject field of the resulting certificate. |
| `require_cn` | no | API: This field indicates whether to require a common name to create a private certificate. By default, a common name is required to generate a certificate. To make the `common_name` field optional, set the `require_cn` option to `false`. |
| `secret_type` | no | — |
| `serial_number` | no | pattern `[^a-fA-F0-9]` |
| `server_flag` | no | API: This field indicates whether private certificates are flagged for server use. |
| `street_address` | no | API: The street address values to define in the subject field of the resulting certificate. |
| `ttl` | no | pattern `^[0-9]+[s,m,h,d]{0,1}$` |
| `ttl_seconds` | no | — |
| `updated_at` | no | — |
| `use_csr_common_name` | no | API: When used with the `private_cert_configuration_action_sign_csr` action, this field determines whether to use the common name (CN) from a certificate signing request (CSR) instead of the CN that is included in the data of the certificate. Does not include any requested Subject Alternative Names (SANs) in the CSR. To use the alternative names, include the `use_csr_sans` property. |
| `use_csr_sans` | no | API: When used with the `private_cert_configuration_action_sign_csr` action, this field determines whether to use the Subject Alternative Names (SANs) from a certificate signing request (CSR) instead of the SANs that are included in the data of the certificate. This field does not include the common name in the CSR. To use the common name, include the `use_csr_common_name` property. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_sm_private_certificate_configuration_template_invalid_allow_any_name` | `allow_any_name` |
| `ibm_sm_private_certificate_configuration_template_invalid_allow_bare_domains` | `allow_bare_domains` |
| `ibm_sm_private_certificate_configuration_template_invalid_allow_glob_domains` | `allow_glob_domains` |
| `ibm_sm_private_certificate_configuration_template_invalid_allow_ip_sans` | `allow_ip_sans` |
| `ibm_sm_private_certificate_configuration_template_invalid_allow_localhost` | `allow_localhost` |
| `ibm_sm_private_certificate_configuration_template_invalid_allow_subdomains` | `allow_subdomains` |
| `ibm_sm_private_certificate_configuration_template_invalid_allow_wildcard_certificates` | `allow_wildcard_certificates` |
| `ibm_sm_private_certificate_configuration_template_invalid_allowed_domains` | `allowed_domains` |
| `ibm_sm_private_certificate_configuration_template_invalid_allowed_domains_template` | `allowed_domains_template` |
| `ibm_sm_private_certificate_configuration_template_invalid_allowed_other_sans` | `allowed_other_sans` |
| `ibm_sm_private_certificate_configuration_template_invalid_allowed_secret_groups` | `allowed_secret_groups` |
| `ibm_sm_private_certificate_configuration_template_invalid_allowed_uri_sans` | `allowed_uri_sans` |
| `ibm_sm_private_certificate_configuration_template_invalid_basic_constraints_valid_for_non_ca` | `basic_constraints_valid_for_non_ca` |
| `ibm_sm_private_certificate_configuration_template_invalid_certificate_authority` | `certificate_authority` |
| `ibm_sm_private_certificate_configuration_template_invalid_client_flag` | `client_flag` |
| `ibm_sm_private_certificate_configuration_template_invalid_code_signing_flag` | `code_signing_flag` |
| `ibm_sm_private_certificate_configuration_template_invalid_country` | `country` |
| `ibm_sm_private_certificate_configuration_template_invalid_email_protection_flag` | `email_protection_flag` |
| `ibm_sm_private_certificate_configuration_template_invalid_enforce_hostnames` | `enforce_hostnames` |
| `ibm_sm_private_certificate_configuration_template_invalid_ext_key_usage` | `ext_key_usage` |
| `ibm_sm_private_certificate_configuration_template_invalid_ext_key_usage_oids` | `ext_key_usage_oids` |
| `ibm_sm_private_certificate_configuration_template_invalid_key_bits` | `key_bits` |
| `ibm_sm_private_certificate_configuration_template_invalid_key_type` | `key_type` |
| `ibm_sm_private_certificate_configuration_template_invalid_key_usage` | `key_usage` |
| `ibm_sm_private_certificate_configuration_template_invalid_locality` | `locality` |
| `ibm_sm_private_certificate_configuration_template_invalid_max_ttl` | `max_ttl` |
| `ibm_sm_private_certificate_configuration_template_invalid_name` | `name` |
| `ibm_sm_private_certificate_configuration_template_invalid_not_before_duration` | `not_before_duration` |
| `ibm_sm_private_certificate_configuration_template_invalid_organization` | `organization` |
| `ibm_sm_private_certificate_configuration_template_invalid_ou` | `ou` |
| `ibm_sm_private_certificate_configuration_template_invalid_policy_identifiers` | `policy_identifiers` |
| `ibm_sm_private_certificate_configuration_template_invalid_postal_code` | `postal_code` |
| `ibm_sm_private_certificate_configuration_template_invalid_province` | `province` |
| `ibm_sm_private_certificate_configuration_template_invalid_require_cn` | `require_cn` |
| `ibm_sm_private_certificate_configuration_template_invalid_serial_number` | `serial_number` |
| `ibm_sm_private_certificate_configuration_template_invalid_server_flag` | `server_flag` |
| `ibm_sm_private_certificate_configuration_template_invalid_street_address` | `street_address` |
| `ibm_sm_private_certificate_configuration_template_invalid_ttl` | `ttl` |
| `ibm_sm_private_certificate_configuration_template_invalid_use_csr_common_name` | `use_csr_common_name` |
| `ibm_sm_private_certificate_configuration_template_invalid_use_csr_sans` | `use_csr_sans` |

## Example

### Valid

```hcl
resource "ibm_sm_private_certificate_configuration_template" "test" {
  certificate_authority = "example-name"
  name = "example-name"
}
```
