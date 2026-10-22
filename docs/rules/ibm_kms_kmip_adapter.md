# ibm_kms_kmip_adapter

Manages ibm_kms_kmip_adapter. The name of the KMIP adapter

This rule validates `ibm_kms_kmip_adapter` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `CreateKMIPAdapterRequestBodyResources` (`key-protect`).

## Configuration

```hcl
rule "ibm_kms_kmip_adapter" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `adapter_id` | no | — |
| `created_at` | no | — |
| `created_by` | no | — |
| `description` | no | pattern `^.{2,240}$` |
| `endpoint_type` | no | one of: `public`, `private` |
| `instance_id` | yes | API: The v4 UUID used to uniquely identify the instance in IBM Cloud. |
| `name` | no | pattern `^[a-zA-Z0-9-_]{2,40}$` |
| `profile` | yes | one of: `native_1.0` |
| `profile_data` | yes | API: The data specific to the KMIP Adapter profile. This is a required field for profile `native_1.0` |
| `updated_at` | no | — |
| `updated_by` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_kms_kmip_adapter_invalid_description` | `description` |
| `ibm_kms_kmip_adapter_invalid_endpoint_type` | `endpoint_type` |
| `ibm_kms_kmip_adapter_invalid_instance_id` | `instance_id` |
| `ibm_kms_kmip_adapter_invalid_name` | `name` |
| `ibm_kms_kmip_adapter_invalid_profile` | `profile` |
| `ibm_kms_kmip_adapter_invalid_profile_data` | `profile_data` |

## Example

### Valid

```hcl
resource "ibm_kms_kmip_adapter" "test" {
  instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  profile = "native_1.0"
  profile_data = "example"
}
```
