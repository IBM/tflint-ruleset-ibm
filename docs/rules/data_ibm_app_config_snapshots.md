# ibm_app_config_snapshots_data

Schema validation for the ibm_app_config_snapshots data source.

This rule validates `ibm_app_config_snapshots` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `CreateGitConfig` (`app-configuration`).

## Configuration

```hcl
rule "ibm_app_config_snapshots_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `collection_id` | no | API: Collection Id. |
| `environment_id` | no | API: Environment Id. |
| `git_config` | no | — |
| `guid` | yes | — |
| `limit` | no | — |
| `offset` | no | API: The number of records that are skipped in a list. |
| `total_count` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_app_config_snapshots_data_invalid_collection_id` | `collection_id` |
| `ibm_app_config_snapshots_data_invalid_environment_id` | `environment_id` |
| `ibm_app_config_snapshots_data_invalid_guid` | `guid` |
| `ibm_app_config_snapshots_data_invalid_limit` | `limit` |
| `ibm_app_config_snapshots_data_invalid_offset` | `offset` |

## Example

### Valid

```hcl
data "ibm_app_config_snapshots" "test" {
  guid = "example"
}
```
