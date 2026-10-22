# ibm_kms_instance_policies

Schema validation for the ibm_kms_instance_policies resource.

This rule validates `ibm_kms_instance_policies` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `InstancePolicyAllAttributes` (`hs-crypto`).

## Configuration

```hcl
rule "ibm_kms_instance_policies" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `endpoint_type` | no | one of: `public`, `private` |
| `instance_id` | yes | API: The v4 UUID used to uniquely identify the instance in IBM Cloud. |
| `dual_auth_delete` (block) | no | nested block |
| `dual_auth_delete.enabled` | yes | — |
| `key_create_import_access` (block) | no | nested block |
| `key_create_import_access.create_root_key` | no | API: If set to `false`, the service prevents you or any authorized users from using Hyper Protect Crypto Services to create root keys in the specified service instance. If set to `true`, Hyper Protect Crypto Services allows you or any authorized users to create root keys in the instance. **Note:** If omitted, `POST /instance/policies` will set this attribute to the default value (`true`). |
| `key_create_import_access.create_standard_key` | no | API: If set to `false`, the service prevents you or any authorized users from using Hyper Protect Crypto Services to create standard keys in the specified service instance. If set to `true`, Hyper Protect Crypto Services allows you or any authorized users to create standard keys in the instance. **Note:** If omitted, `POST /instance/policies` will set this attribute to the default value (`true`). |
| `key_create_import_access.enabled` | yes | — |
| `key_create_import_access.enforce_token` | no | API: If set to `true`, the service prevents you or any authorized users from importing key material into the specified service instance without using an import token. If set to `false`, Hyper Protect Crypto Services allows you or any authorized users to import key material into the instance without the use of an import token. **Note:** If omitted, `POST /instance/policies` will set this attribute to th |
| `key_create_import_access.import_root_key` | no | API: If set to `false`, the service prevents you or any authorized users from importing root keys into the specified service instance. If set to `true`, Hyper Protect Crypto Services allows you or any authorized users to import root keys into the instance. **Note:** If omitted, `POST /instance/policies` will set this attribute to the default value (`true`). |
| `key_create_import_access.import_standard_key` | no | API: If set to `false`, the service prevents you or any authorized users from importing standard keys into the specified service instance. If set to `true`, Hyper Protect Crypto Services allows you or any authorized users to import standard keys into the instance. **Note:** If omitted, `POST /instance/policies` will set this attribute to the default value (`true`). |
| `metrics` (block) | no | nested block |
| `metrics.enabled` | yes | — |
| `rotation` (block) | no | nested block |
| `rotation.enabled` | yes | — |
| `rotation.interval_month` | no | range 1–12 |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_kms_instance_policies_invalid_endpoint_type` | `endpoint_type` |
| `ibm_kms_instance_policies_invalid_instance_id` | `instance_id` |
| `ibm_kms_instance_policies_invalid_dual_auth_delete__enabled` | `dual_auth_delete.enabled` |
| `ibm_kms_instance_policies_invalid_key_create_import_access__create_root_key` | `key_create_import_access.create_root_key` |
| `ibm_kms_instance_policies_invalid_key_create_import_access__create_standard_key` | `key_create_import_access.create_standard_key` |
| `ibm_kms_instance_policies_invalid_key_create_import_access__enabled` | `key_create_import_access.enabled` |
| `ibm_kms_instance_policies_invalid_key_create_import_access__enforce_token` | `key_create_import_access.enforce_token` |
| `ibm_kms_instance_policies_invalid_key_create_import_access__import_root_key` | `key_create_import_access.import_root_key` |
| `ibm_kms_instance_policies_invalid_key_create_import_access__import_standard_key` | `key_create_import_access.import_standard_key` |
| `ibm_kms_instance_policies_invalid_metrics__enabled` | `metrics.enabled` |
| `ibm_kms_instance_policies_invalid_rotation__enabled` | `rotation.enabled` |
| `ibm_kms_instance_policies_invalid_rotation__interval_month` | `rotation.interval_month` |

## Example

### Valid

```hcl
resource "ibm_kms_instance_policies" "test" {
  instance_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
