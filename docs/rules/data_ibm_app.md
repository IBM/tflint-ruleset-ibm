# ibm_app_data

Reads ibm_app. The name for the app

This rule validates `ibm_app` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_app_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `buildpack` | no | — |
| `disk_quota` | no | — |
| `environment_json` | no | — |
| `health_check_http_endpoint` | no | — |
| `health_check_timeout` | no | — |
| `health_check_type` | no | — |
| `instances` | no | — |
| `memory` | no | — |
| `name` | yes | — |
| `package_state` | no | — |
| `route_guid` | no | — |
| `service_instance_guid` | no | — |
| `space_guid` | yes | — |
| `state` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_app_data_invalid_name` | `name` |
| `ibm_app_data_invalid_space_guid` | `space_guid` |

## Example

### Valid

```hcl
data "ibm_app" "test" {
  name = "example"
  space_guid = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
}
```
