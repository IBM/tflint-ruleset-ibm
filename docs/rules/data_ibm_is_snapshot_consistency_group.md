# ibm_is_snapshot_consistency_group_data

Reads ibm_is_snapshot_consistency_group. The name for this snapshot consistency group.

This rule validates `ibm_is_snapshot_consistency_group` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `BareMetalServerDiskPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_snapshot_consistency_group_data" {
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
| `delete_snapshots_on_delete` | no | — |
| `href` | no | — |
| `identifier` | no | — |
| `lifecycle_state` | no | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `resource_group` | no | — |
| `resource_type` | no | — |
| `service_tags` | no | — |
| `snapshots` | no | — |
| `tags` | no | pattern `^[A-Za-z0-9:_ .-]+$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_snapshot_consistency_group_data_invalid_access_tags` | `access_tags` |
| `ibm_is_snapshot_consistency_group_data_invalid_identifier` | `identifier` |
| `ibm_is_snapshot_consistency_group_data_invalid_name` | `name` |
| `ibm_is_snapshot_consistency_group_data_invalid_tags` | `tags` |

## Example

### Valid

```hcl
data "ibm_is_snapshot_consistency_group" "test" {
}
```
