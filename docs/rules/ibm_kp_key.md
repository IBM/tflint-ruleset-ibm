# ibm_kp_key

Schema validation for the ibm_kp_key resource.

This rule validates `ibm_kp_key` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Key` (`key-protect`).

## Configuration

```hcl
rule "ibm_kp_key" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `crn` | no | — |
| `encrypted_nonce` | no | — |
| `force_delete` | no | — |
| `iv_value` | no | — |
| `key_id` | no | — |
| `key_name` | yes | — |
| `key_protect_id` | yes | — |
| `payload` | no | pattern `^.{16,32}$` |
| `resource_controller_url` | no | — |
| `resource_crn` | no | — |
| `resource_group_name` | no | — |
| `resource_name` | no | — |
| `resource_status` | no | — |
| `standard_key` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_kp_key_invalid_encrypted_nonce` | `encrypted_nonce` |
| `ibm_kp_key_invalid_force_delete` | `force_delete` |
| `ibm_kp_key_invalid_iv_value` | `iv_value` |
| `ibm_kp_key_invalid_key_name` | `key_name` |
| `ibm_kp_key_invalid_key_protect_id` | `key_protect_id` |
| `ibm_kp_key_invalid_payload` | `payload` |
| `ibm_kp_key_invalid_standard_key` | `standard_key` |

## Example

### Valid

```hcl
resource "ibm_kp_key" "test" {
  key_name = "example"
  key_protect_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
