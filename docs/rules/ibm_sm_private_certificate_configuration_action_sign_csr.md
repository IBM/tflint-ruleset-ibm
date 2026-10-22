# ibm_sm_private_certificate_configuration_action_sign_csr

Manages ibm_sm_private_certificate_configuration_action_sign_csr. The name that uniquely identifies a configuration

This rule validates `ibm_sm_private_certificate_configuration_action_sign_csr` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `PrivateCertificateConfigurationActionSignCSRPrototype` (`secrets-manager-secrets-manager-v2`).

## Configuration

```hcl
rule "ibm_sm_private_certificate_configuration_action_sign_csr" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `alt_names` | no | API: With the Subject Alternative Name field, you can specify additional hostnames to be protected by a single SSL certificate. |
| `common_name` | no | API: The Common Name (CN) represents the server name that is protected by the SSL certificate. |
| `country` | no | API: The Country (C) values to define in the subject field of the resulting certificate. |
| `csr` | yes | pattern `^(-{5}BEGIN.+?-{5}[\s\S]+-{5}END.+?-{5}[\s]*)$` |
| `data` | no | — |
| `exclude_cn_from_sans` | no | API: This parameter controls whether the common name is excluded from Subject Alternative Names (SANs). If the common name is set to `true`, it is not included in DNS, or email SANs if they apply. This field can be useful if the common name is a human-readable identifier, instead of a hostname or an email address. |
| `format` | no | one of: `pem`, `pem_bundle` |
| `ip_sans` | no | API: The IP Subject Alternative Names to define for the CA certificate, in a comma-delimited list. |
| `locality` | no | API: The Locality (L) values to define in the subject field of the resulting certificate. |
| `max_path_length` | no | range -1–10 |
| `name` | yes | — |
| `organization` | no | API: The Organization (O) values to define in the subject field of the resulting certificate. |
| `other_sans` | no | API: The custom Object Identifier (OID) or UTF8-string Subject Alternative Names to define for the CA certificate. The alternative names must match the values that are specified in the `allowed_other_sans` field in the associated certificate template. The format is the same as OpenSSL: `&lt;oid&gt;:&lt;type&gt;:&lt;value&gt;` where the current valid type is `UTF8`. |
| `ou` | no | API: The Organizational Unit (OU) values to define in the subject field of the resulting certificate. |
| `permitted_dns_domains` | no | API: The allowed DNS domains or subdomains for the certificates that are to be signed and issued by this CA certificate. |
| `postal_code` | no | API: The postal code values to define in the subject field of the resulting certificate. |
| `province` | no | API: The Province (ST) values to define in the subject field of the resulting certificate. |
| `serial_number` | no | pattern `[^a-fA-F0-9]` |
| `street_address` | no | API: The street address values to define in the subject field of the resulting certificate. |
| `ttl` | no | pattern `^[0-9]+[s,m,h,d]{0,1}$` |
| `uri_sans` | no | API: The URI Subject Alternative Names to define for the CA certificate, in a comma-delimited list. |
| `use_csr_values` | no | API: This field indicates whether to use values from a certificate signing request (CSR) to complete a `private_cert_configuration_action_sign_csr` action. If it is set to `true`, then: 1) Subject information, including names and alternate names, are preserved from the CSR rather than by using the values that are provided in the other parameters to this operation. 2) Any key usage, for example, non-rep |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_sm_private_certificate_configuration_action_sign_csr_invalid_alt_names` | `alt_names` |
| `ibm_sm_private_certificate_configuration_action_sign_csr_invalid_common_name` | `common_name` |
| `ibm_sm_private_certificate_configuration_action_sign_csr_invalid_country` | `country` |
| `ibm_sm_private_certificate_configuration_action_sign_csr_invalid_csr` | `csr` |
| `ibm_sm_private_certificate_configuration_action_sign_csr_invalid_exclude_cn_from_sans` | `exclude_cn_from_sans` |
| `ibm_sm_private_certificate_configuration_action_sign_csr_invalid_format` | `format` |
| `ibm_sm_private_certificate_configuration_action_sign_csr_invalid_ip_sans` | `ip_sans` |
| `ibm_sm_private_certificate_configuration_action_sign_csr_invalid_locality` | `locality` |
| `ibm_sm_private_certificate_configuration_action_sign_csr_invalid_max_path_length` | `max_path_length` |
| `ibm_sm_private_certificate_configuration_action_sign_csr_invalid_name` | `name` |
| `ibm_sm_private_certificate_configuration_action_sign_csr_invalid_organization` | `organization` |
| `ibm_sm_private_certificate_configuration_action_sign_csr_invalid_other_sans` | `other_sans` |
| `ibm_sm_private_certificate_configuration_action_sign_csr_invalid_ou` | `ou` |
| `ibm_sm_private_certificate_configuration_action_sign_csr_invalid_permitted_dns_domains` | `permitted_dns_domains` |
| `ibm_sm_private_certificate_configuration_action_sign_csr_invalid_postal_code` | `postal_code` |
| `ibm_sm_private_certificate_configuration_action_sign_csr_invalid_province` | `province` |
| `ibm_sm_private_certificate_configuration_action_sign_csr_invalid_serial_number` | `serial_number` |
| `ibm_sm_private_certificate_configuration_action_sign_csr_invalid_street_address` | `street_address` |
| `ibm_sm_private_certificate_configuration_action_sign_csr_invalid_ttl` | `ttl` |
| `ibm_sm_private_certificate_configuration_action_sign_csr_invalid_uri_sans` | `uri_sans` |
| `ibm_sm_private_certificate_configuration_action_sign_csr_invalid_use_csr_values` | `use_csr_values` |

## Example

### Valid

```hcl
variable "csr" {
  sensitive = true
}

resource "ibm_sm_private_certificate_configuration_action_sign_csr" "test" {
  csr = var.csr
  name = "example"
}
```
