# ibm_satellite_location_data

Schema validation for the ibm_satellite_location data source.

This rule validates `ibm_satellite_location` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_satellite_location_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `capabilities` | no | — |
| `coreos_enabled` | no | — |
| `created_on` | no | — |
| `crn` | no | — |
| `description` | no | — |
| `host_attached_count` | no | — |
| `host_available_count` | no | — |
| `hosts` | no | — |
| `ingress_hostname` | no | — |
| `ingress_secret` | no | — |
| `location` | yes | — |
| `logging_account_id` | no | — |
| `managed_from` | no | — |
| `physical_address` | no | — |
| `pod_subnet` | no | — |
| `resource_group_id` | no | — |
| `resource_group_name` | no | — |
| `service_subnet` | no | — |
| `tags` | no | — |
| `zones` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_satellite_location_data_invalid_location` | `location` |

## Example

### Valid

```hcl
data "ibm_satellite_location" "test" {
  location = "example"
}
```
