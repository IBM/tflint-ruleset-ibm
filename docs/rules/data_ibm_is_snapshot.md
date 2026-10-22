# ibm_is_snapshot_data

Reads ibm_is_snapshot. Snapshot name

This rule validates `ibm_is_snapshot` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `CloudObjectStorageBucketIdentity` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_snapshot_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_tags` | no | — |
| `allowed_use` | no | — |
| `backup_policy_plan` | no | — |
| `bootable` | no | — |
| `captured_at` | no | — |
| `catalog_offering` | no | — |
| `clones` | no | — |
| `copies` | no | — |
| `crn` | no | — |
| `encryption` | no | — |
| `encryption_key` | no | — |
| `href` | no | — |
| `identifier` | no | — |
| `lifecycle_state` | no | — |
| `minimum_capacity` | no | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `operating_system` | no | — |
| `resource_group` | no | — |
| `resource_type` | no | — |
| `service_tags` | no | — |
| `size` | no | — |
| `snapshot_consistency_group` | no | — |
| `source_image` | no | — |
| `source_volume` | no | — |
| `tags` | no | — |
| `source_snapshot` (block) | no | nested block |
| `source_snapshot.crn` | no | crn format |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_snapshot_data_invalid_identifier` | `identifier` |
| `ibm_is_snapshot_data_invalid_name` | `name` |
| `ibm_is_snapshot_data_invalid_source_snapshot__crn` | `source_snapshot.crn` |

## Example

### Valid

```hcl
data "ibm_is_snapshot" "test" {
}
```
