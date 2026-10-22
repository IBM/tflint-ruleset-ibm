# ibm_is_share_snapshots_data

Reads ibm_is_share_snapshots. Filters the collection to resources with a `name` property matching the exact specified name.

This rule validates `ibm_is_share_snapshots` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `BareMetalServerDiskPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_share_snapshots_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `backup_policy_plan` | no | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `share` | no | — |
| `snapshots` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_share_snapshots_data_invalid_backup_policy_plan` | `backup_policy_plan` |
| `ibm_is_share_snapshots_data_invalid_name` | `name` |
| `ibm_is_share_snapshots_data_invalid_share` | `share` |

## Example

### Valid

```hcl
data "ibm_is_share_snapshots" "test" {
}
```
