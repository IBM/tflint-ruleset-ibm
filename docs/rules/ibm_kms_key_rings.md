# ibm_kms_key_rings

Schema validation for the ibm_kms_key_rings resource.

This rule validates `ibm_kms_key_rings` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Key` (`key-protect`).

## Configuration

```hcl
rule "ibm_kms_key_rings" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `endpoint_type` | no | one of: `public`, `private` |
| `force_delete` | no | — |
| `instance_id` | yes | API: The v4 UUID used to uniquely identify the instance in IBM Cloud. |
| `key_ring_id` | yes | pattern `^[a-zA-Z0-9-]*$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_kms_key_rings_invalid_endpoint_type` | `endpoint_type` |
| `ibm_kms_key_rings_invalid_force_delete` | `force_delete` |
| `ibm_kms_key_rings_invalid_instance_id` | `instance_id` |
| `ibm_kms_key_rings_invalid_key_ring_id` | `key_ring_id` |

## Example

### Valid

```hcl
resource "ibm_kms_key_rings" "test" {
  instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  key_ring_id = "example-name"
}
```
