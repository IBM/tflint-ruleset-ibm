# ibm_app_config_feature

Manages ibm_app_config_feature. Feature name.

This rule validates `ibm_app_config_feature` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Feature` (`app-configuration`).

## Configuration

```hcl
rule "ibm_app_config_feature" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `created_time` | no | — |
| `description` | no | API: Feature description |
| `disabled_value` | yes | API: Value of the feature when it is disabled. The value can be Boolean, Numeric, String - TEXT, String - JSON, String - YAML value as per the `type` and `format` attributes. |
| `enabled` | no | API: The state of the feature flag. |
| `enabled_value` | yes | API: Value of the feature when it is enabled. The value can be Boolean, Numeric, String - TEXT, String - JSON, String - YAML value as per the `type` and `format` attributes. |
| `environment_id` | yes | API: Environment identifier. Only present for FEATURE, PROPERTY and ENVIRONMENT. Not present for SEGMENT and COLLECTION. |
| `feature_id` | yes | API: Feature id. Allowed special characters are dot ( . ), hyphen( - ), underscore ( _ ) only |
| `format` | no | one of: `TEXT`, `JSON`, `YAML` |
| `guid` | yes | — |
| `href` | no | — |
| `name` | yes | API: Feature name. Allowed special characters are dot ( . ), hyphen( - ), underscore ( _ ) only |
| `rollout_percentage` | no | range 0–100 |
| `segment_exists` | no | — |
| `tags` | no | pattern `^[a-zA-Z0-9_\. ,\-]+(:[a-zA-Z0-9_\. ,\-]+)*$` |
| `type` | yes | one of: `BOOLEAN`, `NUMERIC`, `STRING` |
| `updated_time` | no | — |
| `collections` (block) | no | nested block |
| `collections.collection_id` | yes | API: Collection id. |
| `collections.deleted` | no | — |
| `segment_rules` (block) | no | nested block |
| `segment_rules.order` | yes | API: Order of the rule, used during evaluation. The evaluation is performed in the order defined and the value associated with the first matching rule is used for evaluation. |
| `segment_rules.rollout_percentage` | no | range 0–100 |
| `segment_rules.value` | yes | API: Value to be used for evaluation for this rule. The value can be Boolean, SecretRef, String - TEXT , String - JSON , String - YAML or a Numeric value as per the `type` and `format` attributes. |
| `segment_rules.rules` (block) | yes | nested block |
| `segment_rules.rules.segments` | yes | API: List of segment ids that are used for targeting using the rule. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_app_config_feature_invalid_description` | `description` |
| `ibm_app_config_feature_invalid_disabled_value` | `disabled_value` |
| `ibm_app_config_feature_invalid_enabled` | `enabled` |
| `ibm_app_config_feature_invalid_enabled_value` | `enabled_value` |
| `ibm_app_config_feature_invalid_environment_id` | `environment_id` |
| `ibm_app_config_feature_invalid_feature_id` | `feature_id` |
| `ibm_app_config_feature_invalid_format` | `format` |
| `ibm_app_config_feature_invalid_guid` | `guid` |
| `ibm_app_config_feature_invalid_name` | `name` |
| `ibm_app_config_feature_invalid_rollout_percentage` | `rollout_percentage` |
| `ibm_app_config_feature_invalid_tags` | `tags` |
| `ibm_app_config_feature_invalid_type` | `type` |
| `ibm_app_config_feature_invalid_collections__collection_id` | `collections.collection_id` |
| `ibm_app_config_feature_invalid_collections__deleted` | `collections.deleted` |
| `ibm_app_config_feature_invalid_segment_rules__order` | `segment_rules.order` |
| `ibm_app_config_feature_invalid_segment_rules__rollout_percentage` | `segment_rules.rollout_percentage` |
| `ibm_app_config_feature_invalid_segment_rules__value` | `segment_rules.value` |
| `ibm_app_config_feature_invalid_segment_rules__rules__segments` | `segment_rules.rules.segments` |

## Example

### Valid

```hcl
resource "ibm_app_config_feature" "test" {
  disabled_value = "example"
  enabled_value = "example"
  environment_id = "01234567-89ab-cdef-0123-456789abcdef"
  feature_id = "01234567-89ab-cdef-0123-456789abcdef"
  guid = "example"
  name = "aaaaaaa"
  type = "BOOLEAN"
}
```
