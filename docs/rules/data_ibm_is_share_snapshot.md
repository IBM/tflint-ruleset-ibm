# ibm_is_share_snapshot_data

Reads ibm_is_share_snapshot. The name for this share snapshot.

This rule validates `ibm_is_share_snapshot` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_is_share_snapshot_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `backup_policy_plan` | no | — |
| `captured_at` | no | — |
| `created_at` | no | — |
| `crn` | no | — |
| `fingerprint` | no | — |
| `href` | no | — |
| `lifecycle_state` | no | — |
| `minimum_size` | no | — |
| `name` | no | — |
| `resource_group` | no | — |
| `resource_type` | no | — |
| `share` | yes | pattern `^[-0-9a-z_]+$` |
| `share_snapshot` | yes | — |
| `status` | no | — |
| `status_reasons` | no | — |
| `tags` | no | — |
| `user_tags` | no | — |
| `zone` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_share_snapshot_data_invalid_share` | `share` |
| `ibm_is_share_snapshot_data_invalid_share_snapshot` | `share_snapshot` |

## Example

### Valid

```hcl
data "ibm_is_share_snapshot" "test" {
  share = "example-name"
  share_snapshot = "example"
}
```
