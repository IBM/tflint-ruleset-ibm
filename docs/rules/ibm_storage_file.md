# ibm_storage_file

Schema validation for the ibm_storage_file resource.

This rule validates `ibm_storage_file` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_storage_file" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `allowed_hardware_ids` | no | — |
| `allowed_ip_addresses` | no | — |
| `allowed_subnets` | no | — |
| `allowed_virtual_guest_ids` | no | — |
| `capacity` | yes | — |
| `datacenter` | yes | — |
| `hostname` | no | — |
| `hourly_billing` | no | — |
| `iops` | yes | — |
| `mountpoint` | no | — |
| `notes` | no | — |
| `resource_controller_url` | no | — |
| `resource_name` | no | — |
| `resource_status` | no | — |
| `snapshot_capacity` | no | — |
| `tags` | no | — |
| `type` | yes | — |
| `volumename` | no | — |
| `snapshot_schedule` (block) | no | nested block |
| `snapshot_schedule.day_of_week` | no | — |
| `snapshot_schedule.enable` | no | — |
| `snapshot_schedule.hour` | no | — |
| `snapshot_schedule.minute` | no | — |
| `snapshot_schedule.retention_count` | yes | — |
| `snapshot_schedule.schedule_type` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_storage_file_invalid_allowed_hardware_ids` | `allowed_hardware_ids` |
| `ibm_storage_file_invalid_allowed_ip_addresses` | `allowed_ip_addresses` |
| `ibm_storage_file_invalid_allowed_subnets` | `allowed_subnets` |
| `ibm_storage_file_invalid_allowed_virtual_guest_ids` | `allowed_virtual_guest_ids` |
| `ibm_storage_file_invalid_capacity` | `capacity` |
| `ibm_storage_file_invalid_datacenter` | `datacenter` |
| `ibm_storage_file_invalid_hourly_billing` | `hourly_billing` |
| `ibm_storage_file_invalid_iops` | `iops` |
| `ibm_storage_file_invalid_notes` | `notes` |
| `ibm_storage_file_invalid_snapshot_capacity` | `snapshot_capacity` |
| `ibm_storage_file_invalid_tags` | `tags` |
| `ibm_storage_file_invalid_type` | `type` |
| `ibm_storage_file_invalid_snapshot_schedule__day_of_week` | `snapshot_schedule.day_of_week` |
| `ibm_storage_file_invalid_snapshot_schedule__enable` | `snapshot_schedule.enable` |
| `ibm_storage_file_invalid_snapshot_schedule__hour` | `snapshot_schedule.hour` |
| `ibm_storage_file_invalid_snapshot_schedule__minute` | `snapshot_schedule.minute` |
| `ibm_storage_file_invalid_snapshot_schedule__retention_count` | `snapshot_schedule.retention_count` |
| `ibm_storage_file_invalid_snapshot_schedule__schedule_type` | `snapshot_schedule.schedule_type` |

## Example

### Valid

```hcl
resource "ibm_storage_file" "test" {
  capacity = 1
  datacenter = "example"
  iops = 1
  type = "example"
}
```
