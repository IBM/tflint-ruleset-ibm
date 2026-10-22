# ibm_event_streams_topic_data

Reads ibm_event_streams_topic. The name of the topic

This rule validates `ibm_event_streams_topic` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `topic_create_request` (`event-streams-adminrest`).

## Configuration

```hcl
rule "ibm_event_streams_topic_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `config` | no | — |
| `kafka_brokers_sasl` | no | — |
| `kafka_http_url` | no | — |
| `name` | yes | API: The name of topic to be created. |
| `partitions` | no | — |
| `resource_instance_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_event_streams_topic_data_invalid_name` | `name` |
| `ibm_event_streams_topic_data_invalid_resource_instance_id` | `resource_instance_id` |

## Example

### Valid

```hcl
data "ibm_event_streams_topic" "test" {
  name = "example"
  resource_instance_id = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
}
```
