# ibm_en_ibmsource

Schema validation for the ibm_en_ibmsource resource.

This rule validates `ibm_en_ibmsource` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Source` (`event-notifications`).

## Configuration

```hcl
rule "ibm_en_ibmsource" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `enabled` | yes | API: The status of the source |
| `instance_guid` | yes | — |
| `source_id` | yes | — |
| `store_notifications` | yes | API: view the payload of incoming events for troubleshooting |
| `updated_at` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_en_ibmsource_invalid_enabled` | `enabled` |
| `ibm_en_ibmsource_invalid_instance_guid` | `instance_guid` |
| `ibm_en_ibmsource_invalid_source_id` | `source_id` |
| `ibm_en_ibmsource_invalid_store_notifications` | `store_notifications` |

## Example

### Valid

```hcl
resource "ibm_en_ibmsource" "test" {
  enabled = true
  instance_guid = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
  source_id = "01234567-89ab-cdef-0123-456789abcdef"
  store_notifications = true
}
```
