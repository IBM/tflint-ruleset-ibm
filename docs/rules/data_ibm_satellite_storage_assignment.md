# ibm_satellite_storage_assignment_data

Schema validation for the ibm_satellite_storage_assignment data source.

This rule validates `ibm_satellite_storage_assignment` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_satellite_storage_assignment_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `assignment_name` | no | — |
| `assignment_type` | no | — |
| `cluster` | no | — |
| `config` | no | — |
| `config_uuid` | no | — |
| `config_version` | no | — |
| `config_version_uuid` | no | — |
| `created` | no | — |
| `groups` | no | — |
| `is_assignment_upgrade_available` | no | — |
| `owner` | no | — |
| `rollout_error_count` | no | — |
| `rollout_success_count` | no | — |
| `sat_cluster` | no | — |
| `svc_cluster` | no | — |
| `uuid` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_satellite_storage_assignment_data_invalid_uuid` | `uuid` |

## Example

### Valid

```hcl
data "ibm_satellite_storage_assignment" "test" {
  uuid = "01234567-89ab-cdef-0123-456789abcdef"
}
```
