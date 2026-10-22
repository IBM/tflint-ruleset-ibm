# ibm_event_streams_schema_data

Schema validation for the ibm_event_streams_schema data source.

This rule validates `ibm_event_streams_schema` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ProduceResponseSchemaMetadata` (`event-streams-restproducer_v2`).

## Configuration

```hcl
rule "ibm_event_streams_schema_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `kafka_http_url` | no | — |
| `resource_instance_id` | yes | — |
| `schema_id` | yes | API: The schema id of the schema used in the registry. example: "schema_id": 101 |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_event_streams_schema_data_invalid_resource_instance_id` | `resource_instance_id` |
| `ibm_event_streams_schema_data_invalid_schema_id` | `schema_id` |

## Example

### Valid

```hcl
data "ibm_event_streams_schema" "test" {
  resource_instance_id = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
  schema_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
