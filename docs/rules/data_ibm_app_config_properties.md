# ibm_app_config_properties_data

Schema validation for the ibm_app_config_properties data source.

This rule validates `ibm_app_config_properties` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Config` (`app-configuration`).

## Configuration

```hcl
rule "ibm_app_config_properties_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `collections` | no | — |
| `environment_id` | yes | API: Environment id. Allowed special characters are dot ( . ), hyphen( - ), underscore ( _ ) only |
| `expand` | no | — |
| `guid` | yes | — |
| `include` | no | — |
| `limit` | no | — |
| `offset` | no | API: The number of records that are skipped in a list. |
| `properties` | no | — |
| `segments` | no | — |
| `sort` | no | — |
| `tags` | no | pattern `^[a-zA-Z0-9_\. ,\-]+(:[a-zA-Z0-9_\. ,\-]+)*$` |
| `total_count` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_app_config_properties_data_invalid_collections` | `collections` |
| `ibm_app_config_properties_data_invalid_environment_id` | `environment_id` |
| `ibm_app_config_properties_data_invalid_expand` | `expand` |
| `ibm_app_config_properties_data_invalid_guid` | `guid` |
| `ibm_app_config_properties_data_invalid_include` | `include` |
| `ibm_app_config_properties_data_invalid_limit` | `limit` |
| `ibm_app_config_properties_data_invalid_offset` | `offset` |
| `ibm_app_config_properties_data_invalid_segments` | `segments` |
| `ibm_app_config_properties_data_invalid_sort` | `sort` |
| `ibm_app_config_properties_data_invalid_tags` | `tags` |

## Example

### Valid

```hcl
data "ibm_app_config_properties" "test" {
  environment_id = "01234567-89ab-cdef-0123-456789abcdef"
  guid = "example"
}
```
