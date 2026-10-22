# ibm_kms_key_policies

Schema validation for the ibm_kms_key_policies resource.

This rule validates `ibm_kms_key_policies` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `RotationKeyMetadata` (`key-protect`).

## Configuration

```hcl
rule "ibm_kms_key_policies" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `alias` | no | — |
| `endpoint_type` | no | one of: `public`, `private` |
| `instance_id` | yes | API: The v4 UUID used to uniquely identify the instance in IBM Cloud. |
| `key_id` | no | — |
| `resource_controller_url` | no | — |
| `resource_crn` | no | — |
| `resource_name` | no | — |
| `resource_status` | no | — |
| `dual_auth_delete` (block) | no | nested block |
| `dual_auth_delete.enabled` | yes | API: If set to `true`, Key Protect enables a rotation policy on a single key. |
| `rotation` (block) | no | nested block |
| `rotation.enabled` | no | API: If set to `true`, Key Protect enables a rotation policy on a single key. |
| `rotation.interval_month` | no | range 1–12 |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_kms_key_policies_invalid_alias` | `alias` |
| `ibm_kms_key_policies_invalid_endpoint_type` | `endpoint_type` |
| `ibm_kms_key_policies_invalid_instance_id` | `instance_id` |
| `ibm_kms_key_policies_invalid_key_id` | `key_id` |
| `ibm_kms_key_policies_invalid_dual_auth_delete__enabled` | `dual_auth_delete.enabled` |
| `ibm_kms_key_policies_invalid_rotation__enabled` | `rotation.enabled` |
| `ibm_kms_key_policies_invalid_rotation__interval_month` | `rotation.interval_month` |

## Example

### Valid

```hcl
resource "ibm_kms_key_policies" "test" {
  instance_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
