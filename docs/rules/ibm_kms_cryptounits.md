# ibm_kms_cryptounits

Schema validation for the ibm_kms_cryptounits resource.

This rule validates `ibm_kms_cryptounits` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `leaf-unique` (`hs-crypto`).

## Configuration

```hcl
rule "ibm_kms_cryptounits" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cryptounits` | no | — |
| `instance_id` | yes | API: The v4 UUID used to uniquely identify the instance in IBM Cloud. |
| `region` | no | — |
| `should_zeroize` | no | — |
| `url` | no | — |
| `use_private_endpoint` | no | — |
| `master_key` (block) | yes | nested block |
| `master_key.keyname` | yes | — |
| `master_key.keysharefile` (block) | yes | nested block |
| `master_key.keysharefile.filepath` | yes | — |
| `master_key.keysharefile.passphrase` | yes | — |
| `signature_key` (block) | yes | nested block |
| `signature_key.filepath` | yes | — |
| `signature_key.owner` | no | — |
| `signature_key.passphrase` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_kms_cryptounits_invalid_instance_id` | `instance_id` |
| `ibm_kms_cryptounits_invalid_region` | `region` |
| `ibm_kms_cryptounits_invalid_should_zeroize` | `should_zeroize` |
| `ibm_kms_cryptounits_invalid_url` | `url` |
| `ibm_kms_cryptounits_invalid_use_private_endpoint` | `use_private_endpoint` |
| `ibm_kms_cryptounits_invalid_master_key__keyname` | `master_key.keyname` |
| `ibm_kms_cryptounits_invalid_master_key__keysharefile__filepath` | `master_key.keysharefile.filepath` |
| `ibm_kms_cryptounits_invalid_master_key__keysharefile__passphrase` | `master_key.keysharefile.passphrase` |
| `ibm_kms_cryptounits_invalid_signature_key__filepath` | `signature_key.filepath` |
| `ibm_kms_cryptounits_invalid_signature_key__owner` | `signature_key.owner` |
| `ibm_kms_cryptounits_invalid_signature_key__passphrase` | `signature_key.passphrase` |

## Example

### Valid

```hcl
variable "passphrase" {
  sensitive = true
}

resource "ibm_kms_cryptounits" "test" {
  instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  master_key {
    keyname = "example"
    keysharefile {
      filepath = "example"
      passphrase = var.passphrase
    }
  }
  signature_key {
    filepath = "example"
    passphrase = var.passphrase
  }
}
```
