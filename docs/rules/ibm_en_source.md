# ibm_en_source

Manages ibm_en_source. The Source name.

This rule validates `ibm_en_source` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/v1/instances/{instance_id}/sources/{id}:patch` (`event-notifications`).

## Configuration

```hcl
rule "ibm_en_source" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `description` | no | API: Description of the source |
| `enabled` | yes | API: Whether the source is enabled or not |
| `instance_guid` | yes | — |
| `name` | yes | API: Name of the source |
| `source_id` | no | — |
| `store_notifications` | yes | API: enable to view the payload of incoming events for troubleshooting |
| `updated_at` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_en_source_invalid_description` | `description` |
| `ibm_en_source_invalid_enabled` | `enabled` |
| `ibm_en_source_invalid_instance_guid` | `instance_guid` |
| `ibm_en_source_invalid_name` | `name` |
| `ibm_en_source_invalid_store_notifications` | `store_notifications` |

## Example

### Valid

```hcl
resource "ibm_en_source" "test" {
  enabled = true
  instance_guid = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
  name = "aaaaaaa"
  store_notifications = true
}
```
