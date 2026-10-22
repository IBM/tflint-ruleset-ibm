# ibm_mqcloud_application

Manages ibm_mqcloud_application. The name of the application - conforming to MQ rules.

This rule validates `ibm_mqcloud_application` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ApplicationDetailsPrototype` (`mq-on-cloud`).

## Configuration

```hcl
rule "ibm_mqcloud_application" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `application_id` | no | — |
| `create_api_key_uri` | no | — |
| `href` | no | — |
| `iam_service_id` | no | — |
| `name` | yes | pattern `^[a-z][-a-z0-9]*$` |
| `service_instance_guid` | yes | pattern `^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_mqcloud_application_invalid_name` | `name` |
| `ibm_mqcloud_application_invalid_service_instance_guid` | `service_instance_guid` |

## Example

### Valid

```hcl
resource "ibm_mqcloud_application" "test" {
  name = "example-name"
  service_instance_guid = "01234567-89ab-cdef-0123-456789abcdef"
}
```
