# ibm_event_streams_quota

Schema validation for the ibm_event_streams_quota resource.

This rule validates `ibm_event_streams_quota` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/admin/quotas/{entity_name}:patch` (`event-streams-adminrest`).

## Configuration

```hcl
rule "ibm_event_streams_quota" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `consumer_byte_rate` | yes | min -1 |
| `entity` | yes | — |
| `producer_byte_rate` | yes | min -1 |
| `resource_instance_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_event_streams_quota_invalid_consumer_byte_rate` | `consumer_byte_rate` |
| `ibm_event_streams_quota_invalid_entity` | `entity` |
| `ibm_event_streams_quota_invalid_producer_byte_rate` | `producer_byte_rate` |
| `ibm_event_streams_quota_invalid_resource_instance_id` | `resource_instance_id` |

## Example

### Valid

```hcl
resource "ibm_event_streams_quota" "test" {
  consumer_byte_rate = -1
  entity = "example"
  producer_byte_rate = -1
  resource_instance_id = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
}
```
