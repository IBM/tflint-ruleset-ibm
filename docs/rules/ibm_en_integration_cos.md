# ibm_en_integration_cos

Schema validation for the ibm_en_integration_cos resource.

This rule validates `ibm_en_integration_cos` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `IntegrationCreateMetadata` (`event-notifications`).

## Configuration

```hcl
rule "ibm_en_integration_cos" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `instance_guid` | yes | — |
| `integration_id` | no | — |
| `type` | yes | — |
| `updated_at` | no | — |
| `metadata` (block) | yes | nested block |
| `metadata.bucket_name` | yes | API: Cloud Object Storage bucket name |
| `metadata.crn` | yes | crn format |
| `metadata.endpoint` | yes | API: URL for Cloud Object storage |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_en_integration_cos_invalid_instance_guid` | `instance_guid` |
| `ibm_en_integration_cos_invalid_type` | `type` |
| `ibm_en_integration_cos_invalid_metadata__bucket_name` | `metadata.bucket_name` |
| `ibm_en_integration_cos_invalid_metadata__crn` | `metadata.crn` |
| `ibm_en_integration_cos_invalid_metadata__endpoint` | `metadata.endpoint` |

## Example

### Valid

```hcl
resource "ibm_en_integration_cos" "test" {
  instance_guid = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
  type = "example"
  metadata {
    bucket_name = "aaaaaaa"
    crn = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
    endpoint = "aaaaaaa"
  }
}
```
