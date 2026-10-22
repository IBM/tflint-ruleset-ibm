# ibm_app_config_collection_data

Reads ibm_app_config_collection. Collection name.

This rule validates `ibm_app_config_collection` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Config` (`app-configuration`).

## Configuration

```hcl
rule "ibm_app_config_collection_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `collection_id` | yes | API: Collection Id. Allowed special characters are dot ( . ), hyphen( - ), underscore ( _ ) only |
| `created_time` | no | — |
| `description` | no | — |
| `expand` | no | — |
| `features` | no | — |
| `features_count` | no | — |
| `guid` | yes | — |
| `href` | no | — |
| `include` | no | — |
| `name` | no | — |
| `properties` | no | — |
| `properties_count` | no | — |
| `tags` | no | — |
| `updated_time` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_app_config_collection_data_invalid_collection_id` | `collection_id` |
| `ibm_app_config_collection_data_invalid_expand` | `expand` |
| `ibm_app_config_collection_data_invalid_guid` | `guid` |
| `ibm_app_config_collection_data_invalid_include` | `include` |

## Example

### Valid

```hcl
data "ibm_app_config_collection" "test" {
  collection_id = "01234567-89ab-cdef-0123-456789abcdef"
  guid = "example"
}
```
