# ibm_app_config_integration_kms

Schema validation for the ibm_app_config_integration_kms resource.

This rule validates `ibm_app_config_integration_kms` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `CreateKmsIntegrationMetadata` (`app-configuration`).

## Configuration

```hcl
rule "ibm_app_config_integration_kms" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `created_time` | no | — |
| `guid` | yes | — |
| `href` | no | — |
| `integration_id` | yes | — |
| `integration_type` | no | — |
| `kms_endpoint` | yes | API: The URL endpoint of Key Protect instance. |
| `kms_instance_crn` | yes | crn format |
| `kms_schema_type` | no | — |
| `root_key_id` | yes | API: The rootkey id of Key Protect instance. |
| `updated_time` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_app_config_integration_kms_invalid_guid` | `guid` |
| `ibm_app_config_integration_kms_invalid_integration_id` | `integration_id` |
| `ibm_app_config_integration_kms_invalid_kms_endpoint` | `kms_endpoint` |
| `ibm_app_config_integration_kms_invalid_kms_instance_crn` | `kms_instance_crn` |
| `ibm_app_config_integration_kms_invalid_root_key_id` | `root_key_id` |

## Example

### Valid

```hcl
resource "ibm_app_config_integration_kms" "test" {
  guid = "example"
  integration_id = "01234567-89ab-cdef-0123-456789abcdef"
  kms_endpoint = "example"
  kms_instance_crn = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
  root_key_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
