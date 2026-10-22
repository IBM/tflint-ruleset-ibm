# ibm_sm_public_certificate_configuration_ca_lets_encrypt

Manages ibm_sm_public_certificate_configuration_ca_lets_encrypt. A human-readable unique name to assign to your configuration.

This rule validates `ibm_sm_public_certificate_configuration_ca_lets_encrypt` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `PublicCertificateConfigurationCALetsEncryptPrototype` (`secrets-manager-secrets-manager-v2`).

## Configuration

```hcl
rule "ibm_sm_public_certificate_configuration_ca_lets_encrypt" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `config_type` | no | — |
| `lets_encrypt_environment` | yes | one of: `production`, `staging` |
| `lets_encrypt_preferred_chain` | no | API: This field supports only the chains that Let's Encrypt provides. Keep empty to use the default or supply a valid Let's Encrypt-provided value. For a list of supported chains, see: https://letsencrypt.org/certificates/. |
| `lets_encrypt_private_key` | yes | pattern `(^-----BEGIN PRIVATE KEY-----.*?)|(^-----BEGIN RSA PRIVATE KEY-----.*?)` |
| `name` | yes | pattern `^[A-Za-z0-9_][A-Za-z0-9_]*(?:_*-*\.*[A-Za-z0-9]*)*[A-Za-z0-9]+$` |
| `secret_type` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_sm_public_certificate_configuration_ca_lets_encrypt_invalid_lets_encrypt_environment` | `lets_encrypt_environment` |
| `ibm_sm_public_certificate_configuration_ca_lets_encrypt_invalid_lets_encrypt_preferred_chain` | `lets_encrypt_preferred_chain` |
| `ibm_sm_public_certificate_configuration_ca_lets_encrypt_invalid_lets_encrypt_private_key` | `lets_encrypt_private_key` |
| `ibm_sm_public_certificate_configuration_ca_lets_encrypt_invalid_name` | `name` |

## Example

### Valid

```hcl
variable "lets_encrypt_private_key" {
  sensitive = true
}

resource "ibm_sm_public_certificate_configuration_ca_lets_encrypt" "test" {
  lets_encrypt_environment = "production"
  lets_encrypt_private_key = var.lets_encrypt_private_key
  name = "example-name"
}
```
