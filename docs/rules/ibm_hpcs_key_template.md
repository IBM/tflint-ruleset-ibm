# ibm_hpcs_key_template

Manages ibm_hpcs_key_template. Name of the template, it will be referenced when creating managed keys.

This rule validates `ibm_hpcs_key_template` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `keyPropertiesUpdate` (`uko`).

## Configuration

```hcl
rule "ibm_hpcs_key_template" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `created_at` | no | — |
| `created_by` | no | — |
| `description` | no | — |
| `href` | no | — |
| `instance_id` | yes | API: The v4 UUID used to uniquely identify the instance in IBM Cloud. |
| `name` | yes | pattern `^[A-Za-z][A-Za-z0-9-]*$` |
| `region` | yes | one of: `au-syd`, `in-che`, `jp-osa`, `jp-tok`, `kr-seo`, `eu-de`, `eu-gb`, `ca-tor`, `us-south`, `us-south-test`, `us-east`, `br-sao` |
| `uko_vault` | yes | — |
| `updated_at` | no | — |
| `updated_by` | no | — |
| `version` | no | — |
| `key` (block) | yes | nested block |
| `key.activation_date` | yes | API: Key activation date can be provided as a period definition (e.g. P1Y means 1 year). |
| `key.algorithm` | yes | one of: `aes`, `rsa`, `hmac`, `ec`, `des`, `dilithium` |
| `key.expiration_date` | yes | API: Key expiration date can be provided as a period definition (e.g. P1Y means 1 year). |
| `key.size` | yes | pattern `^[A-Za-z0-9]+$` |
| `key.state` | yes | one of: `pre_activation`, `active` |
| `keystores` (block) | yes | nested block |
| `keystores.google_key_protection_level` | no | one of: `software`, `hsm` |
| `keystores.google_key_purpose` | no | one of: `encrypt_decrypt`, `asymmetric_decrypt`, `asymmetric_sign`, `mac` |
| `keystores.google_kms_algorithm` | no | one of: `google_symmetric_encryption`, `ec_sign_p256_sha256`, `ec_sign_p384_sha384`, `ec_sign_secp256k1_sha256`, `rsa_sign_pss_2048_sha256`, `rsa_sign_pss_3072_sha256`, `rsa_sign_pss_4096_sha256`, `rsa_sign_pss_4096_sha512`, `rsa_sign_pkcs1_2048_sha256`, `rsa_sign_pkcs1_3072_sha256`, `rsa_sign_pkcs1_4096_sha256`, `rsa_sign_pkcs1_4096_sha512`, `rsa_sign_raw_pkcs1_2048`, `rsa_sign_raw_pkcs1_3072`, `rsa_sign_raw_pkcs1_4096`, `rsa_decrypt_oaep_2048_sha1`, `rsa_decrypt_oaep_2048_sha256`, `rsa_decrypt_oaep_3072_sha1`, `rsa_decrypt_oaep_3072_sha256`, `rsa_decrypt_oaep_4096_sha1`, `rsa_decrypt_oaep_4096_sha256`, `rsa_decrypt_oaep_4096_sha512`, `hmac_sha256` |
| `keystores.group` | yes | — |
| `keystores.type` | yes | — |
| `vault` (block) | yes | nested block |
| `vault.id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_hpcs_key_template_invalid_description` | `description` |
| `ibm_hpcs_key_template_invalid_instance_id` | `instance_id` |
| `ibm_hpcs_key_template_invalid_name` | `name` |
| `ibm_hpcs_key_template_invalid_region` | `region` |
| `ibm_hpcs_key_template_invalid_uko_vault` | `uko_vault` |
| `ibm_hpcs_key_template_invalid_key__activation_date` | `key.activation_date` |
| `ibm_hpcs_key_template_invalid_key__algorithm` | `key.algorithm` |
| `ibm_hpcs_key_template_invalid_key__expiration_date` | `key.expiration_date` |
| `ibm_hpcs_key_template_invalid_key__size` | `key.size` |
| `ibm_hpcs_key_template_invalid_key__state` | `key.state` |
| `ibm_hpcs_key_template_invalid_keystores__google_key_protection_level` | `keystores.google_key_protection_level` |
| `ibm_hpcs_key_template_invalid_keystores__google_key_purpose` | `keystores.google_key_purpose` |
| `ibm_hpcs_key_template_invalid_keystores__google_kms_algorithm` | `keystores.google_kms_algorithm` |
| `ibm_hpcs_key_template_invalid_keystores__group` | `keystores.group` |
| `ibm_hpcs_key_template_invalid_keystores__type` | `keystores.type` |
| `ibm_hpcs_key_template_invalid_vault__id` | `vault.id` |

## Example

### Valid

```hcl
resource "ibm_hpcs_key_template" "test" {
  instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  name = "example-name"
  region = "au-syd"
  uko_vault = "example"
  key {
    activation_date = "aaaaaaa"
    algorithm = "aes"
    expiration_date = "aaaaaaa"
    size = "example"
    state = "pre_activation"
  }
  keystores {
    group = "example"
    type = "example"
  }
  vault {
    id = "example"
  }
}
```
