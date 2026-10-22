# ibm_is_share_mount_target_data

Reads ibm_is_share_mount_target. The user-defined name for this share target.

This rule validates `ibm_is_share_mount_target` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_is_share_mount_target_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_control_mode` | no | — |
| `access_protocol` | no | — |
| `created_at` | no | — |
| `href` | no | — |
| `lifecycle_state` | no | — |
| `mount_path` | no | — |
| `mount_target` | no | — |
| `mount_target_name` | no | — |
| `name` | no | — |
| `primary_ip` | no | — |
| `resource_type` | no | — |
| `share` | no | — |
| `share_name` | no | — |
| `subnet` | no | — |
| `transit_encryption` | no | — |
| `virtual_network_interface` | no | — |
| `vpc` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_share_mount_target_data_invalid_mount_target` | `mount_target` |
| `ibm_is_share_mount_target_data_invalid_mount_target_name` | `mount_target_name` |
| `ibm_is_share_mount_target_data_invalid_share` | `share` |
| `ibm_is_share_mount_target_data_invalid_share_name` | `share_name` |

## Example

### Valid

```hcl
data "ibm_is_share_mount_target" "test" {
}
```
