# ibm_app_config_segment

Manages ibm_app_config_segment. Segment name.

This rule validates `ibm_app_config_segment` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Segment` (`app-configuration`).

## Configuration

```hcl
rule "ibm_app_config_segment" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `created_time` | no | — |
| `description` | no | API: Segment description |
| `guid` | yes | — |
| `href` | no | — |
| `name` | yes | API: Segment name. Allowed special characters are dot ( . ), hyphen( - ), underscore ( _ ) only |
| `segment_id` | yes | API: Segment id. Allowed special characters are dot ( . ), hyphen( - ), underscore ( _ ) only |
| `tags` | no | pattern `^[a-zA-Z0-9_\. ,\-]+(:[a-zA-Z0-9_\. ,\-]+)*$` |
| `updated_time` | no | — |
| `rules` (block) | yes | nested block |
| `rules.attribute_name` | yes | API: Attribute name. |
| `rules.operator` | yes | one of: `is`, `isNot`, `contains`, `notContains`, `startsWith`, `notStartsWith`, `endsWith`, `notEndsWith`, `greaterThan`, `lesserThan`, `greaterThanEquals`, `lesserThanEquals` |
| `rules.values` | yes | API: Value to be used to evaluate if the entity belongs to the segment. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_app_config_segment_invalid_description` | `description` |
| `ibm_app_config_segment_invalid_guid` | `guid` |
| `ibm_app_config_segment_invalid_name` | `name` |
| `ibm_app_config_segment_invalid_segment_id` | `segment_id` |
| `ibm_app_config_segment_invalid_tags` | `tags` |
| `ibm_app_config_segment_invalid_rules__attribute_name` | `rules.attribute_name` |
| `ibm_app_config_segment_invalid_rules__operator` | `rules.operator` |
| `ibm_app_config_segment_invalid_rules__values` | `rules.values` |

## Example

### Valid

```hcl
resource "ibm_app_config_segment" "test" {
  guid = "example"
  name = "aaaaaaa"
  segment_id = "01234567-89ab-cdef-0123-456789abcdef"
  rules {
    attribute_name = "example"
    operator = "is"
    values = "example"
  }
}
```
