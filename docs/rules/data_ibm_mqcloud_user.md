# ibm_mqcloud_user_data

Reads ibm_mqcloud_user. The shortname of the user that will be used as the IBM MQ administrator in interactions with a queue manager for this service instance.

This rule validates `ibm_mqcloud_user` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ApplicationDetailsPrototype` (`mq-on-cloud`).

## Configuration

```hcl
rule "ibm_mqcloud_user_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `name` | no | pattern `^[a-z][-a-z0-9]*$` |
| `service_instance_guid` | yes | pattern `^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$` |
| `users` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_mqcloud_user_data_invalid_name` | `name` |
| `ibm_mqcloud_user_data_invalid_service_instance_guid` | `service_instance_guid` |

## Example

### Valid

```hcl
data "ibm_mqcloud_user" "test" {
  service_instance_guid = "01234567-89ab-cdef-0123-456789abcdef"
}
```
