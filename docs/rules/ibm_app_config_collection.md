# ibm_app_config_collection

Manages ibm_app_config_collection. Collection name.

This rule validates `ibm_app_config_collection` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ImportCollectionSchema` (`app-configuration`).

## Configuration

```hcl
rule "ibm_app_config_collection" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `collection_id` | yes | API: Collection id. Allowed special characters are dot ( . ), hyphen( - ), underscore ( _ ) only |
| `created_time` | no | — |
| `description` | no | API: Description of the collection |
| `features_count` | no | — |
| `guid` | yes | — |
| `href` | no | — |
| `name` | yes | API: Collection name. Allowed special characters are dot ( . ), hyphen( - ), underscore ( _ ) only |
| `properties_count` | no | — |
| `tags` | no | pattern `^[a-zA-Z0-9_\. ,\-]+(:[a-zA-Z0-9_\. ,\-]+)*$` |
| `updated_time` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_app_config_collection_invalid_collection_id` | `collection_id` |
| `ibm_app_config_collection_invalid_description` | `description` |
| `ibm_app_config_collection_invalid_guid` | `guid` |
| `ibm_app_config_collection_invalid_name` | `name` |
| `ibm_app_config_collection_invalid_tags` | `tags` |

## Example

### Valid

```hcl
resource "ibm_app_config_collection" "test" {
  collection_id = "01234567-89ab-cdef-0123-456789abcdef"
  guid = "example"
  name = "aaaaaaa"
}
```
