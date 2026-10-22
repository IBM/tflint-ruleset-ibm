# ibm_storage_block

Schema validation for the ibm_storage_block resource.

This rule validates `ibm_storage_block` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_storage_block" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `allowed_hardware_ids` | no | — |
| `allowed_hardware_info` | no | — |
| `allowed_host_info` | no | — |
| `allowed_ip_addresses` | no | — |
| `allowed_virtual_guest_ids` | no | — |
| `allowed_virtual_guest_info` | no | — |
| `capacity` | yes | — |
| `datacenter` | yes | — |
| `hostname` | no | — |
| `hourly_billing` | no | — |
| `iops` | yes | — |
| `lunid` | no | — |
| `notes` | no | — |
| `os_format_type` | yes | — |
| `resource_controller_url` | no | — |
| `resource_name` | no | — |
| `snapshot_capacity` | no | — |
| `tags` | no | — |
| `target_address` | no | — |
| `type` | yes | — |
| `volumename` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_storage_block_invalid_allowed_hardware_ids` | `allowed_hardware_ids` |
| `ibm_storage_block_invalid_allowed_ip_addresses` | `allowed_ip_addresses` |
| `ibm_storage_block_invalid_allowed_virtual_guest_ids` | `allowed_virtual_guest_ids` |
| `ibm_storage_block_invalid_capacity` | `capacity` |
| `ibm_storage_block_invalid_datacenter` | `datacenter` |
| `ibm_storage_block_invalid_hourly_billing` | `hourly_billing` |
| `ibm_storage_block_invalid_iops` | `iops` |
| `ibm_storage_block_invalid_notes` | `notes` |
| `ibm_storage_block_invalid_os_format_type` | `os_format_type` |
| `ibm_storage_block_invalid_snapshot_capacity` | `snapshot_capacity` |
| `ibm_storage_block_invalid_tags` | `tags` |
| `ibm_storage_block_invalid_type` | `type` |

## Example

### Valid

```hcl
resource "ibm_storage_block" "test" {
  capacity = 1
  datacenter = "example"
  iops = 1
  os_format_type = "example"
  type = "example"
}
```
