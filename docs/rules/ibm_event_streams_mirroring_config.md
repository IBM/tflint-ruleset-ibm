# ibm_event_streams_mirroring_config

Schema validation for the ibm_event_streams_mirroring_config resource.

This rule validates `ibm_event_streams_mirroring_config` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_event_streams_mirroring_config" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `mirroring_topic_patterns` | yes | — |
| `resource_instance_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_event_streams_mirroring_config_invalid_mirroring_topic_patterns` | `mirroring_topic_patterns` |
| `ibm_event_streams_mirroring_config_invalid_resource_instance_id` | `resource_instance_id` |

## Example

### Valid

```hcl
resource "ibm_event_streams_mirroring_config" "test" {
  mirroring_topic_patterns = "example"
  resource_instance_id = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
}
```
