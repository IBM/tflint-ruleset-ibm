# ibm_is_snapshots_data

Reads ibm_is_snapshots. Filters the snapshot collection by snapshot name

This rule validates `ibm_is_snapshots` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `PublicGatewayFloatingIPPrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_snapshots_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `backup_policy_plan_id` | no | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `resource_group` | no | API: The resource group to use. If unspecified, the account's [default resource group](https://cloud.ibm.com/apidocs/resource-manager#introduction) will be used. |
| `snapshot_consistency_group_crn` | no | crn format |
| `snapshot_consistency_group_id` | no | — |
| `snapshot_copies_crn` | no | crn format |
| `snapshot_copies_id` | no | — |
| `snapshot_copies_name` | no | — |
| `snapshot_copies_remote_region_name` | no | — |
| `snapshot_source_volume_remote_region_name` | no | — |
| `snapshots` | no | — |
| `source_image` | no | — |
| `source_snapshot_id` | no | — |
| `source_snapshot_remote_region_name` | no | — |
| `source_volume` | no | — |
| `tag` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_snapshots_data_invalid_backup_policy_plan_id` | `backup_policy_plan_id` |
| `ibm_is_snapshots_data_invalid_name` | `name` |
| `ibm_is_snapshots_data_invalid_resource_group` | `resource_group` |
| `ibm_is_snapshots_data_invalid_snapshot_consistency_group_crn` | `snapshot_consistency_group_crn` |
| `ibm_is_snapshots_data_invalid_snapshot_consistency_group_id` | `snapshot_consistency_group_id` |
| `ibm_is_snapshots_data_invalid_snapshot_copies_crn` | `snapshot_copies_crn` |
| `ibm_is_snapshots_data_invalid_snapshot_copies_id` | `snapshot_copies_id` |
| `ibm_is_snapshots_data_invalid_snapshot_copies_name` | `snapshot_copies_name` |
| `ibm_is_snapshots_data_invalid_snapshot_copies_remote_region_name` | `snapshot_copies_remote_region_name` |
| `ibm_is_snapshots_data_invalid_snapshot_source_volume_remote_region_name` | `snapshot_source_volume_remote_region_name` |
| `ibm_is_snapshots_data_invalid_source_image` | `source_image` |
| `ibm_is_snapshots_data_invalid_source_snapshot_id` | `source_snapshot_id` |
| `ibm_is_snapshots_data_invalid_source_snapshot_remote_region_name` | `source_snapshot_remote_region_name` |
| `ibm_is_snapshots_data_invalid_source_volume` | `source_volume` |
| `ibm_is_snapshots_data_invalid_tag` | `tag` |

## Example

### Valid

```hcl
data "ibm_is_snapshots" "test" {
}
```
