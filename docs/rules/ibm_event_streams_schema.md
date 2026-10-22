# ibm_event_streams_schema

Schema validation for the ibm_event_streams_schema resource.

This rule validates `ibm_event_streams_schema` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/artifacts:post` (`event-streams-schemaregistry`).

## Configuration

```hcl
rule "ibm_event_streams_schema" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `kafka_http_url` | no | — |
| `resource_instance_id` | yes | — |
| `schema` | yes | API: The AVRO schema |
| `schema_id` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_event_streams_schema_invalid_resource_instance_id` | `resource_instance_id` |
| `ibm_event_streams_schema_invalid_schema` | `schema` |
| `ibm_event_streams_schema_invalid_schema_id` | `schema_id` |

## Example

### Valid

```hcl
resource "ibm_event_streams_schema" "test" {
  resource_instance_id = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
  schema = "example"
}
```
