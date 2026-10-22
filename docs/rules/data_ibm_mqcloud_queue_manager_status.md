# ibm_mqcloud_queue_manager_status_data

Schema validation for the ibm_mqcloud_queue_manager_status data source.

This rule validates `ibm_mqcloud_queue_manager_status` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `QueueManagerTaskStatus` (`mq-on-cloud`).

## Configuration

```hcl
rule "ibm_mqcloud_queue_manager_status_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `queue_manager_id` | yes | API: The queue manager id. |
| `service_instance_guid` | yes | — |
| `status` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_mqcloud_queue_manager_status_data_invalid_queue_manager_id` | `queue_manager_id` |
| `ibm_mqcloud_queue_manager_status_data_invalid_service_instance_guid` | `service_instance_guid` |

## Example

### Valid

```hcl
data "ibm_mqcloud_queue_manager_status" "test" {
  queue_manager_id = "01234567-89ab-cdef-0123-456789abcdef"
  service_instance_guid = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
}
```
