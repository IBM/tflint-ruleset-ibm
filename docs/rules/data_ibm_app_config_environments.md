# ibm_app_config_environments_data

Schema validation for the ibm_app_config_environments data source.

This rule validates `ibm_app_config_environments` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Config` (`app-configuration`).

## Configuration

```hcl
rule "ibm_app_config_environments_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `environments` | no | — |
| `expand` | no | — |
| `first` | no | — |
| `guid` | yes | — |
| `last` | no | — |
| `limit` | no | — |
| `next` | no | — |
| `offset` | no | API: The number of records that are skipped in a list. |
| `previous` | no | — |
| `tags` | no | pattern `^[a-zA-Z0-9_\. ,\-]+(:[a-zA-Z0-9_\. ,\-]+)*$` |
| `total_count` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_app_config_environments_data_invalid_expand` | `expand` |
| `ibm_app_config_environments_data_invalid_guid` | `guid` |
| `ibm_app_config_environments_data_invalid_limit` | `limit` |
| `ibm_app_config_environments_data_invalid_offset` | `offset` |
| `ibm_app_config_environments_data_invalid_tags` | `tags` |

## Example

### Valid

```hcl
data "ibm_app_config_environments" "test" {
  guid = "example"
}
```
