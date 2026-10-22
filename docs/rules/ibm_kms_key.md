# ibm_kms_key

Schema validation for the ibm_kms_key resource.

This rule validates `ibm_kms_key` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Key` (`key-protect`).

## Configuration

```hcl
rule "ibm_kms_key" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `crn` | no | — |
| `description` | no | pattern `^.{2,240}$` |
| `encrypted_nonce` | no | — |
| `endpoint_type` | no | one of: `public`, `private` |
| `expiration_date` | no | — |
| `force_delete` | no | — |
| `instance_crn` | no | — |
| `instance_id` | yes | API: The v4 UUID used to uniquely identify the instance in IBM Cloud. |
| `iv_value` | no | — |
| `key_id` | no | — |
| `key_name` | yes | — |
| `key_ring_id` | no | — |
| `payload` | no | pattern `^.{16,32}$` |
| `registrations` | no | — |
| `resource_controller_url` | no | — |
| `resource_crn` | no | — |
| `resource_group_name` | no | — |
| `resource_name` | no | — |
| `resource_status` | no | — |
| `standard_key` | no | — |
| `type` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_kms_key_invalid_description` | `description` |
| `ibm_kms_key_invalid_encrypted_nonce` | `encrypted_nonce` |
| `ibm_kms_key_invalid_endpoint_type` | `endpoint_type` |
| `ibm_kms_key_invalid_expiration_date` | `expiration_date` |
| `ibm_kms_key_invalid_force_delete` | `force_delete` |
| `ibm_kms_key_invalid_instance_id` | `instance_id` |
| `ibm_kms_key_invalid_iv_value` | `iv_value` |
| `ibm_kms_key_invalid_key_name` | `key_name` |
| `ibm_kms_key_invalid_key_ring_id` | `key_ring_id` |
| `ibm_kms_key_invalid_payload` | `payload` |
| `ibm_kms_key_invalid_standard_key` | `standard_key` |

## Example

### Valid

```hcl
resource "ibm_kms_key" "test" {
  instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  key_name = "example"
}
```
