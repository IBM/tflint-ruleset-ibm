# ibm_satellite_location

Schema validation for the ibm_satellite_location resource.

This rule validates `ibm_satellite_location` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_satellite_location" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `capabilities` | no | one of: `on-prem` |
| `coreos_enabled` | no | — |
| `created_on` | no | — |
| `crn` | no | — |
| `description` | no | — |
| `host_attached_count` | no | — |
| `host_available_count` | no | — |
| `ingress_hostname` | no | — |
| `ingress_secret` | no | — |
| `location` | yes | — |
| `logging_account_id` | no | — |
| `managed_from` | yes | — |
| `physical_address` | no | — |
| `pod_subnet` | no | — |
| `resource_group_id` | no | — |
| `resource_group_name` | no | — |
| `service_subnet` | no | — |
| `tags` | no | pattern `^[A-Za-z0-9:_ .-]+$` |
| `zones` | no | — |
| `cos_config` (block) | no | nested block |
| `cos_config.bucket` | no | — |
| `cos_config.endpoint` | no | — |
| `cos_config.region` | no | — |
| `cos_credentials` (block) | no | nested block |
| `cos_credentials.access_key_id` | no | — |
| `cos_credentials.secret_access_key` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_satellite_location_invalid_capabilities` | `capabilities` |
| `ibm_satellite_location_invalid_coreos_enabled` | `coreos_enabled` |
| `ibm_satellite_location_invalid_description` | `description` |
| `ibm_satellite_location_invalid_location` | `location` |
| `ibm_satellite_location_invalid_logging_account_id` | `logging_account_id` |
| `ibm_satellite_location_invalid_managed_from` | `managed_from` |
| `ibm_satellite_location_invalid_physical_address` | `physical_address` |
| `ibm_satellite_location_invalid_pod_subnet` | `pod_subnet` |
| `ibm_satellite_location_invalid_resource_group_id` | `resource_group_id` |
| `ibm_satellite_location_invalid_service_subnet` | `service_subnet` |
| `ibm_satellite_location_invalid_tags` | `tags` |
| `ibm_satellite_location_invalid_zones` | `zones` |
| `ibm_satellite_location_invalid_cos_config__bucket` | `cos_config.bucket` |
| `ibm_satellite_location_invalid_cos_config__endpoint` | `cos_config.endpoint` |
| `ibm_satellite_location_invalid_cos_config__region` | `cos_config.region` |
| `ibm_satellite_location_invalid_cos_credentials__access_key_id` | `cos_credentials.access_key_id` |
| `ibm_satellite_location_invalid_cos_credentials__secret_access_key` | `cos_credentials.secret_access_key` |

## Example

### Valid

```hcl
resource "ibm_satellite_location" "test" {
  location = "example"
  managed_from = "example"
}
```
