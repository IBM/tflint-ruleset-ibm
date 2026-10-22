# ibm_atracker_target

Manages ibm_atracker_target. The name of the target.

This rule validates `ibm_atracker_target` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `CosEndpointPrototype` (`atracker`).

## Configuration

```hcl
rule "ibm_atracker_target" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `api_version` | no | — |
| `created_at` | no | — |
| `crn` | no | — |
| `managed_by` | no | one of: `account`, `enterprise` |
| `message` | no | — |
| `name` | yes | pattern `^[a-zA-Z0-9 -._:]+$` |
| `region` | no | pattern `^[a-zA-Z0-9 -._:]+$` |
| `target_type` | yes | one of: `app_config`, `cloud_logs`, `cloud_object_storage`, `event_streams` |
| `updated_at` | no | — |
| `write_status` | no | — |
| `appconfig_endpoint` (block) | no | nested block |
| `appconfig_endpoint.target_crn` | yes | crn format |
| `cloudlogs_endpoint` (block) | no | nested block |
| `cloudlogs_endpoint.target_crn` | yes | crn format |
| `cos_endpoint` (block) | no | nested block |
| `cos_endpoint.api_key` | no | pattern `^[a-zA-Z0-9 -._:]+$` |
| `cos_endpoint.bucket` | yes | pattern `^[a-zA-Z0-9 -._:/]+$` |
| `cos_endpoint.endpoint` | yes | pattern `^[a-zA-Z0-9 -._:]+$` |
| `cos_endpoint.service_to_service_enabled` | no | API: Determines if IBM Cloud Activity Tracker Event Routing has service to service authentication enabled. Set this flag to true if service to service is enabled and do not supply an apikey. |
| `cos_endpoint.target_crn` | yes | crn format |
| `eventstreams_endpoint` (block) | no | nested block |
| `eventstreams_endpoint.api_key` | no | pattern `^[a-zA-Z0-9 -._:]+$` |
| `eventstreams_endpoint.brokers` | yes | pattern `^[a-zA-Z0-9 -._:]+$` |
| `eventstreams_endpoint.service_to_service_enabled` | no | API: Determines if IBM Cloud Activity Tracker Event Routing has service to service authentication enabled. Set this flag to true if service to service is enabled and do not supply an apikey. |
| `eventstreams_endpoint.target_crn` | yes | crn format |
| `eventstreams_endpoint.topic` | yes | pattern `^[a-zA-Z0-9 -._:/]+$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_atracker_target_invalid_managed_by` | `managed_by` |
| `ibm_atracker_target_invalid_name` | `name` |
| `ibm_atracker_target_invalid_region` | `region` |
| `ibm_atracker_target_invalid_target_type` | `target_type` |
| `ibm_atracker_target_invalid_appconfig_endpoint__target_crn` | `appconfig_endpoint.target_crn` |
| `ibm_atracker_target_invalid_cloudlogs_endpoint__target_crn` | `cloudlogs_endpoint.target_crn` |
| `ibm_atracker_target_invalid_cos_endpoint__api_key` | `cos_endpoint.api_key` |
| `ibm_atracker_target_invalid_cos_endpoint__bucket` | `cos_endpoint.bucket` |
| `ibm_atracker_target_invalid_cos_endpoint__endpoint` | `cos_endpoint.endpoint` |
| `ibm_atracker_target_invalid_cos_endpoint__service_to_service_enabled` | `cos_endpoint.service_to_service_enabled` |
| `ibm_atracker_target_invalid_cos_endpoint__target_crn` | `cos_endpoint.target_crn` |
| `ibm_atracker_target_invalid_eventstreams_endpoint__api_key` | `eventstreams_endpoint.api_key` |
| `ibm_atracker_target_invalid_eventstreams_endpoint__brokers` | `eventstreams_endpoint.brokers` |
| `ibm_atracker_target_invalid_eventstreams_endpoint__service_to_service_enabled` | `eventstreams_endpoint.service_to_service_enabled` |
| `ibm_atracker_target_invalid_eventstreams_endpoint__target_crn` | `eventstreams_endpoint.target_crn` |
| `ibm_atracker_target_invalid_eventstreams_endpoint__topic` | `eventstreams_endpoint.topic` |

## Example

### Valid

```hcl
resource "ibm_atracker_target" "test" {
  name = "example-name"
  target_type = "app_config"
}
```
