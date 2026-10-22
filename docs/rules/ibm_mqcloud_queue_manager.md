# ibm_mqcloud_queue_manager

Manages ibm_mqcloud_queue_manager. A queue manager name conforming to MQ restrictions.

This rule validates `ibm_mqcloud_queue_manager` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `QueueManagerDetailsPrototype` (`mq-on-cloud`).

## Configuration

```hcl
rule "ibm_mqcloud_queue_manager" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `administrator_api_endpoint_url` | no | — |
| `available_upgrade_versions_uri` | no | — |
| `connection_info_uri` | no | — |
| `date_created` | no | — |
| `display_name` | no | API: A displayable name for the queue manager - limited only in length. |
| `href` | no | — |
| `location` | yes | pattern `^([^[:ascii:]]|[a-zA-Z0-9-._: ])+$` |
| `name` | yes | pattern `^[a-zA-Z0-9_.]*$` |
| `queue_manager_id` | no | — |
| `rest_api_endpoint_url` | no | — |
| `service_instance_guid` | yes | pattern `^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$` |
| `size` | yes | one of: `large`, `medium`, `small`, `xsmall` |
| `status_uri` | no | — |
| `upgrade_available` | no | — |
| `version` | no | pattern `^[0-9]+.[0-9]+.[0-9]+_[0-9]+$` |
| `web_console_url` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_mqcloud_queue_manager_invalid_display_name` | `display_name` |
| `ibm_mqcloud_queue_manager_invalid_location` | `location` |
| `ibm_mqcloud_queue_manager_invalid_name` | `name` |
| `ibm_mqcloud_queue_manager_invalid_service_instance_guid` | `service_instance_guid` |
| `ibm_mqcloud_queue_manager_invalid_size` | `size` |
| `ibm_mqcloud_queue_manager_invalid_version` | `version` |

## Example

### Valid

```hcl
resource "ibm_mqcloud_queue_manager" "test" {
  location = "example-name"
  name = "example"
  service_instance_guid = "01234567-89ab-cdef-0123-456789abcdef"
  size = "large"
}
```
