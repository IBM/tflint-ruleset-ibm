# ibm_mqcloud_queue_manager_data

Reads ibm_mqcloud_queue_manager. A queue manager name conforming to MQ restrictions.

This rule validates `ibm_mqcloud_queue_manager` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ApplicationDetailsPrototype` (`mq-on-cloud`).

## Configuration

```hcl
rule "ibm_mqcloud_queue_manager_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `name` | no | pattern `^[a-zA-Z0-9_.]*$` |
| `queue_managers` | no | — |
| `service_instance_guid` | yes | pattern `^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_mqcloud_queue_manager_data_invalid_name` | `name` |
| `ibm_mqcloud_queue_manager_data_invalid_service_instance_guid` | `service_instance_guid` |

## Example

### Valid

```hcl
data "ibm_mqcloud_queue_manager" "test" {
  service_instance_guid = "01234567-89ab-cdef-0123-456789abcdef"
}
```
