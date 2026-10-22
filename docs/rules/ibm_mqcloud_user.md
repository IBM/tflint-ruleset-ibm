# ibm_mqcloud_user

Manages ibm_mqcloud_user. The shortname of the user that will be used as the IBM MQ administrator in interactions with a queue manager for this service instance.

This rule validates `ibm_mqcloud_user` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `UserDetailsPrototype` (`mq-on-cloud`).

## Configuration

```hcl
rule "ibm_mqcloud_user" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `email` | yes | API: The email of the user to be created. |
| `href` | no | — |
| `iam_managed` | no | — |
| `iam_service_id` | no | — |
| `name` | yes | pattern `^[a-z][-a-z0-9]*$` |
| `roles` | no | — |
| `service_instance_guid` | yes | pattern `^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$` |
| `user_id` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_mqcloud_user_invalid_email` | `email` |
| `ibm_mqcloud_user_invalid_name` | `name` |
| `ibm_mqcloud_user_invalid_service_instance_guid` | `service_instance_guid` |

## Example

### Valid

```hcl
resource "ibm_mqcloud_user" "test" {
  email = "aaaaaaa"
  name = "example-name"
  service_instance_guid = "01234567-89ab-cdef-0123-456789abcdef"
}
```
