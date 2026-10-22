# ibm_app_config_collections_data

Schema validation for the ibm_app_config_collections data source.

This rule validates `ibm_app_config_collections` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Config` (`app-configuration`).

## Configuration

```hcl
rule "ibm_app_config_collections_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `collections` | no | — |
| `expand` | no | — |
| `guid` | yes | — |
| `include` | no | — |
| `limit` | no | — |
| `offset` | no | API: The number of records that are skipped in a list. |
| `total_count` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_app_config_collections_data_invalid_expand` | `expand` |
| `ibm_app_config_collections_data_invalid_guid` | `guid` |
| `ibm_app_config_collections_data_invalid_include` | `include` |
| `ibm_app_config_collections_data_invalid_limit` | `limit` |
| `ibm_app_config_collections_data_invalid_offset` | `offset` |

## Example

### Valid

```hcl
data "ibm_app_config_collections" "test" {
  guid = "example"
}
```
