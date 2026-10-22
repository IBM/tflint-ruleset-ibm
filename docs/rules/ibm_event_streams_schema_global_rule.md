# ibm_event_streams_schema_global_rule

Schema validation for the ibm_event_streams_schema_global_rule resource.

This rule validates `ibm_event_streams_schema_global_rule` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `rule` (`event-streams-schemaregistry`).

## Configuration

```hcl
rule "ibm_event_streams_schema_global_rule" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `config` | yes | one of: `BACKWARD`, `BACKWARD_TRANSITIVE`, `FORWARD`, `FORWARD_TRANSITIVE`, `FULL`, `FULL_TRANSITIVE`, `NONE` |
| `resource_instance_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_event_streams_schema_global_rule_invalid_config` | `config` |
| `ibm_event_streams_schema_global_rule_invalid_resource_instance_id` | `resource_instance_id` |

## Example

### Valid

```hcl
resource "ibm_event_streams_schema_global_rule" "test" {
  config = "BACKWARD"
  resource_instance_id = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
}
```
