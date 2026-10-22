# ibm_sm_public_certificate

Manages ibm_sm_public_certificate. A human-readable name to assign to your secret.

This rule validates `ibm_sm_public_certificate` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `PublicCertificatePrototype` (`secrets-manager-secrets-manager-v2`).

## Configuration

```hcl
rule "ibm_sm_public_certificate" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `alt_names` | no | API: With the Subject Alternative Name field, you can specify additional hostnames to be protected by a single SSL certificate. |
| `bundle_certs` | no | API: This field indicates whether your issued certificate is bundled with intermediate certificates. Set to `false` for the certificate file to contain only the issued certificate. |
| `ca` | yes | API: The name of the certificate authority configuration. |
| `certificate` | no | — |
| `common_name` | no | pattern `^(\*\.)?(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\-]*[A-Za-z0-9])\.?$` |
| `created_at` | no | — |
| `created_by` | no | — |
| `crn` | no | — |
| `custom_metadata` | no | API: The secret metadata that a user can customize. |
| `description` | no | API: An extended description of your secret. To protect your privacy, do not use personal data, such as your name or location, as a description for your secret group. |
| `dns` | yes | API: The name of the DNS provider configuration. |
| `downloaded` | no | — |
| `expiration_date` | no | — |
| `intermediate` | no | — |
| `issuance_info` | no | — |
| `issuer` | no | — |
| `key_algorithm` | no | pattern `^(RSA2048|RSA4096|ECDSA256|ECDSA384)$` |
| `labels` | no | API: Label that you can use to search for secrets in your instance. Label can be between 2-64 characters, including spaces. To protect your privacy, do not use personal data, such as your name or location, as a label for your secret. |
| `locks_total` | no | — |
| `name` | yes | pattern `^[A-Za-z0-9_][A-Za-z0-9_]*(?:_*-*\.*[A-Za-z0-9]*)*[A-Za-z0-9]+$` |
| `private_key` | no | — |
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
| `version_custom_metadata` | no | API: The secret version metadata that a user can customize. |
| `versions_total` | no | — |
| `akamai` (block) | no | nested block |
| `akamai.config` (block) | no | nested block |
| `akamai.config.access_token` | no | — |
| `akamai.config.client_secret` | no | — |
| `akamai.config.client_token` | no | — |
| `akamai.config.host` | no | — |
| `akamai.edgerc` (block) | no | nested block |
| `akamai.edgerc.config_section` | no | — |
| `akamai.edgerc.path_to_edgerc` | no | — |
| `rotation` (block) | no | nested block |
| `rotation.auto_rotate` | no | API: This field indicates whether Secrets Manager rotates your secret automatically. The default is `false`. If `auto_rotate` is set to `true` the service rotates your secret based on the defined interval. |
| `rotation.rotate_keys` | no | API: This field indicates whether Secrets Manager rotates the private key for your public certificate automatically. The default is `false`. If it is set to `true`, the service generates and stores a new private key for your rotated certificate. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_sm_public_certificate_invalid_alt_names` | `alt_names` |
| `ibm_sm_public_certificate_invalid_bundle_certs` | `bundle_certs` |
| `ibm_sm_public_certificate_invalid_ca` | `ca` |
| `ibm_sm_public_certificate_invalid_common_name` | `common_name` |
| `ibm_sm_public_certificate_invalid_custom_metadata` | `custom_metadata` |
| `ibm_sm_public_certificate_invalid_description` | `description` |
| `ibm_sm_public_certificate_invalid_dns` | `dns` |
| `ibm_sm_public_certificate_invalid_key_algorithm` | `key_algorithm` |
| `ibm_sm_public_certificate_invalid_labels` | `labels` |
| `ibm_sm_public_certificate_invalid_name` | `name` |
| `ibm_sm_public_certificate_invalid_secret_group_id` | `secret_group_id` |
| `ibm_sm_public_certificate_invalid_version_custom_metadata` | `version_custom_metadata` |
| `ibm_sm_public_certificate_invalid_akamai__config__access_token` | `akamai.config.access_token` |
| `ibm_sm_public_certificate_invalid_akamai__config__client_secret` | `akamai.config.client_secret` |
| `ibm_sm_public_certificate_invalid_akamai__config__client_token` | `akamai.config.client_token` |
| `ibm_sm_public_certificate_invalid_akamai__config__host` | `akamai.config.host` |
| `ibm_sm_public_certificate_invalid_akamai__edgerc__config_section` | `akamai.edgerc.config_section` |
| `ibm_sm_public_certificate_invalid_akamai__edgerc__path_to_edgerc` | `akamai.edgerc.path_to_edgerc` |
| `ibm_sm_public_certificate_invalid_rotation__auto_rotate` | `rotation.auto_rotate` |
| `ibm_sm_public_certificate_invalid_rotation__rotate_keys` | `rotation.rotate_keys` |

## Example

### Valid

```hcl
resource "ibm_sm_public_certificate" "test" {
  ca = "aaaaaaa"
  dns = "aaaaaaa"
  name = "example-name"
}
```
