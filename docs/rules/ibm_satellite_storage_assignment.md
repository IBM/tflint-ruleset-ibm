# ibm_satellite_storage_assignment

Schema validation for the ibm_satellite_storage_assignment resource.

This rule validates `ibm_satellite_storage_assignment` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_satellite_storage_assignment" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `assignment_name` | yes | — |
| `assignment_type` | no | — |
| `cluster` | no | — |
| `config` | yes | — |
| `config_uuid` | no | — |
| `config_version` | no | — |
| `config_version_uuid` | no | — |
| `controller` | no | — |
| `created` | no | — |
| `groups` | no | — |
| `is_assignment_upgrade_available` | no | — |
| `owner` | no | — |
| `rollout_error_count` | no | — |
| `rollout_success_count` | no | — |
| `sat_cluster` | no | — |
| `svc_cluster` | no | — |
| `update_config_revision` | no | — |
| `uuid` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_satellite_storage_assignment_invalid_assignment_name` | `assignment_name` |
| `ibm_satellite_storage_assignment_invalid_cluster` | `cluster` |
| `ibm_satellite_storage_assignment_invalid_config` | `config` |
| `ibm_satellite_storage_assignment_invalid_controller` | `controller` |
| `ibm_satellite_storage_assignment_invalid_groups` | `groups` |
| `ibm_satellite_storage_assignment_invalid_update_config_revision` | `update_config_revision` |

## Example

### Valid

```hcl
resource "ibm_satellite_storage_assignment" "test" {
  assignment_name = "example"
  config = "example"
}
```
