# ibm_app_config_property

Manages ibm_app_config_property. Property name.

This rule validates `ibm_app_config_property` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Property` (`app-configuration`).

## Configuration

```hcl
rule "ibm_app_config_property" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `created_time` | no | — |
| `description` | no | API: Property description |
| `environment_id` | yes | API: Environment identifier. Only present for FEATURE, PROPERTY and ENVIRONMENT. Not present for SEGMENT and COLLECTION. |
| `evaluation_time` | no | — |
| `format` | no | one of: `TEXT`, `JSON`, `YAML` |
| `guid` | yes | — |
| `href` | no | — |
| `name` | yes | API: Property name. Allowed special characters are dot ( . ), hyphen( - ), underscore ( _ ) only |
| `property_id` | yes | API: Property id. Allowed special characters are dot ( . ), hyphen( - ), underscore ( _ ) only |
| `segment_exists` | no | — |
| `tags` | no | pattern `^[a-zA-Z0-9_\. ,\-]+(:[a-zA-Z0-9_\. ,\-]+)*$` |
| `type` | yes | one of: `BOOLEAN`, `STRING`, `NUMERIC`, `SECRETREF` |
| `updated_time` | no | — |
| `value` | yes | API: Value of the Property. The value can be Boolean, Numeric, SecretRef, String - TEXT, String - JSON, String - YAML as per the `type` and `format` attributes. |
| `collections` (block) | no | nested block |
| `collections.collection_id` | yes | API: Collection id. |
| `collections.deleted` | no | — |
| `segment_rules` (block) | no | nested block |
| `segment_rules.order` | yes | API: Order of the rule, used during evaluation. The evaluation is performed in the order defined and the value associated with the first matching rule is used for evaluation. |
| `segment_rules.value` | yes | API: Value to be used for evaluation for this rule. The value can be Boolean, SecretRef, String - TEXT , String - JSON , String - YAML or a Numeric value as per the `type` and `format` attributes. |
| `segment_rules.rules` (block) | yes | nested block |
| `segment_rules.rules.segments` | yes | API: List of segment ids that are used for targeting using the rule. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_app_config_property_invalid_description` | `description` |
| `ibm_app_config_property_invalid_environment_id` | `environment_id` |
| `ibm_app_config_property_invalid_format` | `format` |
| `ibm_app_config_property_invalid_guid` | `guid` |
| `ibm_app_config_property_invalid_name` | `name` |
| `ibm_app_config_property_invalid_property_id` | `property_id` |
| `ibm_app_config_property_invalid_tags` | `tags` |
| `ibm_app_config_property_invalid_type` | `type` |
| `ibm_app_config_property_invalid_value` | `value` |
| `ibm_app_config_property_invalid_collections__collection_id` | `collections.collection_id` |
| `ibm_app_config_property_invalid_collections__deleted` | `collections.deleted` |
| `ibm_app_config_property_invalid_segment_rules__order` | `segment_rules.order` |
| `ibm_app_config_property_invalid_segment_rules__value` | `segment_rules.value` |
| `ibm_app_config_property_invalid_segment_rules__rules__segments` | `segment_rules.rules.segments` |

## Example

### Valid

```hcl
resource "ibm_app_config_property" "test" {
  environment_id = "01234567-89ab-cdef-0123-456789abcdef"
  guid = "example"
  name = "aaaaaaa"
  property_id = "01234567-89ab-cdef-0123-456789abcdef"
  type = "BOOLEAN"
  value = "example"
}
```
