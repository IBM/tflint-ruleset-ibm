# ibm_kms_key_alias

Schema validation for the ibm_kms_key_alias resource.

This rule validates `ibm_kms_key_alias` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Key` (`key-protect`).

## Configuration

```hcl
rule "ibm_kms_key_alias" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `alias` | yes | — |
| `endpoint_type` | no | one of: `public`, `private` |
| `existing_alias` | no | — |
| `instance_id` | yes | API: The v4 UUID used to uniquely identify the instance in IBM Cloud. |
| `key_id` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_kms_key_alias_invalid_alias` | `alias` |
| `ibm_kms_key_alias_invalid_endpoint_type` | `endpoint_type` |
| `ibm_kms_key_alias_invalid_existing_alias` | `existing_alias` |
| `ibm_kms_key_alias_invalid_instance_id` | `instance_id` |
| `ibm_kms_key_alias_invalid_key_id` | `key_id` |

## Example

### Valid

```hcl
resource "ibm_kms_key_alias" "test" {
  alias = "example"
  instance_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
