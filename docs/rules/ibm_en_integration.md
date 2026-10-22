# ibm_en_integration

Schema validation for the ibm_en_integration resource.

This rule validates `ibm_en_integration` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `IntegrationMetadata` (`event-notifications`).

## Configuration

```hcl
rule "ibm_en_integration" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `instance_guid` | yes | — |
| `integration_id` | yes | — |
| `type` | yes | — |
| `updated_at` | no | — |
| `metadata` (block) | yes | nested block |
| `metadata.crn` | yes | crn format |
| `metadata.endpoint` | yes | API: KMS url for key management or url for COS bucket |
| `metadata.root_key_id` | yes | API: Root Key ID of KMS |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_en_integration_invalid_instance_guid` | `instance_guid` |
| `ibm_en_integration_invalid_integration_id` | `integration_id` |
| `ibm_en_integration_invalid_type` | `type` |
| `ibm_en_integration_invalid_metadata__crn` | `metadata.crn` |
| `ibm_en_integration_invalid_metadata__endpoint` | `metadata.endpoint` |
| `ibm_en_integration_invalid_metadata__root_key_id` | `metadata.root_key_id` |

## Example

### Valid

```hcl
resource "ibm_en_integration" "test" {
  instance_guid = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
  integration_id = "01234567-89ab-cdef-0123-456789abcdef"
  type = "example"
  metadata {
    crn = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
    endpoint = "aaaaaaa"
    root_key_id = "01234567-89ab-cdef-0123-456789abcdef"
  }
}
```
