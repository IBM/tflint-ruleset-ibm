# ibm_is_share_mount_targets_data

Reads ibm_is_share_mount_targets. The user-defined name for this share target.

This rule validates `ibm_is_share_mount_targets` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `BareMetalServerDiskPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_share_mount_targets_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `mount_targets` | no | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `share` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_share_mount_targets_data_invalid_name` | `name` |
| `ibm_is_share_mount_targets_data_invalid_share` | `share` |

## Example

### Valid

```hcl
data "ibm_is_share_mount_targets" "test" {
  share = "example"
}
```
