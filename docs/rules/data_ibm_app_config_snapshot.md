# ibm_app_config_snapshot_data

Schema validation for the ibm_app_config_snapshot data source.

This rule validates `ibm_app_config_snapshot` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `SnapshotOutput` (`app-configuration`).

## Configuration

```hcl
rule "ibm_app_config_snapshot_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `collection` | no | — |
| `created_time` | no | — |
| `environment` | no | — |
| `git_branch` | no | — |
| `git_config_id` | yes | API: Git Config id. |
| `git_config_name` | no | — |
| `git_file_path` | no | — |
| `git_url` | no | — |
| `guid` | yes | — |
| `href` | no | — |
| `last_sync_time` | no | — |
| `updated_time` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_app_config_snapshot_data_invalid_git_config_id` | `git_config_id` |
| `ibm_app_config_snapshot_data_invalid_guid` | `guid` |

## Example

### Valid

```hcl
data "ibm_app_config_snapshot" "test" {
  git_config_id = "01234567-89ab-cdef-0123-456789abcdef"
  guid = "example"
}
```
