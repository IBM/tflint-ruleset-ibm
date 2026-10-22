# ibm_app_config_integration_en

Schema validation for the ibm_app_config_integration_en resource.

This rule validates `ibm_app_config_integration_en` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Integration` (`app-configuration`).

## Configuration

```hcl
rule "ibm_app_config_integration_en" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `created_time` | no | — |
| `description` | no | API: Environment description |
| `en_endpoint` | yes | — |
| `en_instance_crn` | yes | crn format |
| `en_source_name` | yes | — |
| `guid` | yes | — |
| `href` | no | — |
| `integration_id` | yes | API: Integration id. |
| `integration_type` | no | — |
| `updated_time` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_app_config_integration_en_invalid_description` | `description` |
| `ibm_app_config_integration_en_invalid_en_endpoint` | `en_endpoint` |
| `ibm_app_config_integration_en_invalid_en_instance_crn` | `en_instance_crn` |
| `ibm_app_config_integration_en_invalid_en_source_name` | `en_source_name` |
| `ibm_app_config_integration_en_invalid_guid` | `guid` |
| `ibm_app_config_integration_en_invalid_integration_id` | `integration_id` |

## Example

### Valid

```hcl
resource "ibm_app_config_integration_en" "test" {
  en_endpoint = "example"
  en_instance_crn = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
  en_source_name = "example"
  guid = "example"
  integration_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
