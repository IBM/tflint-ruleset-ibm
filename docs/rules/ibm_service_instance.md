# ibm_service_instance

Manages ibm_service_instance. A name for the service instance

This rule validates `ibm_service_instance` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_service_instance" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `credentials` | no | — |
| `dashboard_url` | no | — |
| `name` | yes | — |
| `parameters` | no | — |
| `plan` | yes | — |
| `service` | yes | — |
| `service_keys` | no | — |
| `service_plan_guid` | no | — |
| `space_guid` | yes | — |
| `tags` | no | — |
| `wait_time_minutes` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_service_instance_invalid_name` | `name` |
| `ibm_service_instance_invalid_parameters` | `parameters` |
| `ibm_service_instance_invalid_plan` | `plan` |
| `ibm_service_instance_invalid_service` | `service` |
| `ibm_service_instance_invalid_space_guid` | `space_guid` |
| `ibm_service_instance_invalid_tags` | `tags` |
| `ibm_service_instance_invalid_wait_time_minutes` | `wait_time_minutes` |

## Example

### Valid

```hcl
resource "ibm_service_instance" "test" {
  name = "example"
  plan = "example"
  service = "example"
  space_guid = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
}
```
