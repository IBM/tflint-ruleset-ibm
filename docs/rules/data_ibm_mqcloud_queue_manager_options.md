# ibm_mqcloud_queue_manager_options_data

Schema validation for the ibm_mqcloud_queue_manager_options data source.

This rule validates `ibm_mqcloud_queue_manager_options` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_mqcloud_queue_manager_options_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `latest_version` | no | — |
| `locations` | no | — |
| `service_instance_guid` | yes | — |
| `sizes` | no | — |
| `versions` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_mqcloud_queue_manager_options_data_invalid_service_instance_guid` | `service_instance_guid` |

## Example

### Valid

```hcl
data "ibm_mqcloud_queue_manager_options" "test" {
  service_instance_guid = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
}
```
