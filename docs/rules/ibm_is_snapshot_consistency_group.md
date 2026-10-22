# ibm_is_snapshot_consistency_group

Manages ibm_is_snapshot_consistency_group. The name for this snapshot consistency group.

This rule validates `ibm_is_snapshot_consistency_group` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `SnapshotConsistencyGroupPrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_snapshot_consistency_group" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_tags` | no | pattern `^([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-]):([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-])$` |
| `backup_policy_job` | no | — |
| `backup_policy_plan` | no | — |
| `created_at` | no | — |
| `crn` | no | — |
| `delete_snapshots_on_delete` | no | API: Indicates whether deleting the snapshot consistency group will also delete the snapshots in the group. |
| `href` | no | — |
| `lifecycle_state` | no | — |
| `name` | yes | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `resource_group` | no | API: The resource group to use. If unspecified, the account's [default resource group](https://cloud.ibm.com/apidocs/resource-manager#introduction) will be used. |
| `resource_type` | no | — |
| `service_tags` | no | — |
| `snapshot_reference` | no | — |
| `tags` | no | pattern `^[A-Za-z0-9:_ .-]+$` |
| `snapshots` (block) | yes | nested block |
| `snapshots.name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `snapshots.source_volume` | yes | — |
| `snapshots.tags` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_snapshot_consistency_group_invalid_access_tags` | `access_tags` |
| `ibm_is_snapshot_consistency_group_invalid_delete_snapshots_on_delete` | `delete_snapshots_on_delete` |
| `ibm_is_snapshot_consistency_group_invalid_name` | `name` |
| `ibm_is_snapshot_consistency_group_invalid_resource_group` | `resource_group` |
| `ibm_is_snapshot_consistency_group_invalid_tags` | `tags` |
| `ibm_is_snapshot_consistency_group_invalid_snapshots__name` | `snapshots.name` |
| `ibm_is_snapshot_consistency_group_invalid_snapshots__source_volume` | `snapshots.source_volume` |
| `ibm_is_snapshot_consistency_group_invalid_snapshots__tags` | `snapshots.tags` |

## Example

### Valid

```hcl
resource "ibm_is_snapshot_consistency_group" "test" {
  name = "example-name"
  snapshots {
    source_volume = "example"
  }
}
```
