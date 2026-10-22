# ibm_event_streams_quota_data

Schema validation for the ibm_event_streams_quota data source.

This rule validates `ibm_event_streams_quota` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_event_streams_quota_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `consumer_byte_rate` | no | — |
| `entity` | yes | — |
| `producer_byte_rate` | no | — |
| `resource_instance_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_event_streams_quota_data_invalid_entity` | `entity` |
| `ibm_event_streams_quota_data_invalid_resource_instance_id` | `resource_instance_id` |

## Example

### Valid

```hcl
data "ibm_event_streams_quota" "test" {
  entity = "example"
  resource_instance_id = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
}
```
