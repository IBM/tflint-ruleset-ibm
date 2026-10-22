# ibm_app

Manages ibm_app. The name for the app

This rule validates `ibm_app` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_app" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `app_path` | yes | — |
| `app_version` | no | — |
| `buildpack` | no | — |
| `command` | no | — |
| `disk_quota` | no | — |
| `environment_json` | no | — |
| `health_check_http_endpoint` | no | — |
| `health_check_timeout` | no | — |
| `health_check_type` | no | one of: `port`, `process` |
| `instances` | no | — |
| `memory` | no | — |
| `name` | yes | — |
| `route_guid` | no | — |
| `service_instance_guid` | no | — |
| `space_guid` | yes | — |
| `tags` | no | — |
| `wait_time_minutes` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_app_invalid_app_path` | `app_path` |
| `ibm_app_invalid_app_version` | `app_version` |
| `ibm_app_invalid_buildpack` | `buildpack` |
| `ibm_app_invalid_command` | `command` |
| `ibm_app_invalid_disk_quota` | `disk_quota` |
| `ibm_app_invalid_environment_json` | `environment_json` |
| `ibm_app_invalid_health_check_http_endpoint` | `health_check_http_endpoint` |
| `ibm_app_invalid_health_check_timeout` | `health_check_timeout` |
| `ibm_app_invalid_health_check_type` | `health_check_type` |
| `ibm_app_invalid_instances` | `instances` |
| `ibm_app_invalid_memory` | `memory` |
| `ibm_app_invalid_name` | `name` |
| `ibm_app_invalid_route_guid` | `route_guid` |
| `ibm_app_invalid_service_instance_guid` | `service_instance_guid` |
| `ibm_app_invalid_space_guid` | `space_guid` |
| `ibm_app_invalid_tags` | `tags` |
| `ibm_app_invalid_wait_time_minutes` | `wait_time_minutes` |

## Example

### Valid

```hcl
resource "ibm_app" "test" {
  app_path = "example"
  name = "example"
  space_guid = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
}
```
