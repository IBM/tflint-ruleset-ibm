# ibm_is_snapshot_consistency_groups_data

Reads ibm_is_snapshot_consistency_groups. Filters the collection to resources with a `name` property matching the exact specified name.

This rule validates `ibm_is_snapshot_consistency_groups` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `FloatingIPPrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_snapshot_consistency_groups_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `backup_policy_job` | no | — |
| `backup_policy_plan` | no | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `resource_group` | no | API: The resource group to use. If unspecified, the account's [default resource group](https://cloud.ibm.com/apidocs/resource-manager#introduction) will be used. |
| `snapshot_consistency_groups` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_snapshot_consistency_groups_data_invalid_backup_policy_job` | `backup_policy_job` |
| `ibm_is_snapshot_consistency_groups_data_invalid_backup_policy_plan` | `backup_policy_plan` |
| `ibm_is_snapshot_consistency_groups_data_invalid_name` | `name` |
| `ibm_is_snapshot_consistency_groups_data_invalid_resource_group` | `resource_group` |

## Example

### Valid

```hcl
data "ibm_is_snapshot_consistency_groups" "test" {
}
```
