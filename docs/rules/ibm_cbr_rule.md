# ibm_cbr_rule

Schema validation for the ibm_cbr_rule resource.

This rule validates `ibm_cbr_rule` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Rule` (`context-based-restrictions`).

## Configuration

```hcl
rule "ibm_cbr_rule" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `created_at` | no | — |
| `created_by_id` | no | — |
| `crn` | no | — |
| `description` | no | API: The description of the rule. |
| `enforcement_mode` | no | one of: `disabled`, `enabled`, `report` |
| `etag` | no | — |
| `href` | no | — |
| `last_modified_at` | no | — |
| `last_modified_by_id` | no | — |
| `transaction_id` | no | pattern `^[a-zA-Z0-9 ,\-_]+$` |
| `x_correlation_id` | no | pattern `^[a-zA-Z0-9 ,\-_]+$` |
| `contexts` (block) | no | nested block |
| `contexts.attributes` (block) | yes | nested block |
| `contexts.attributes.name` | yes | pattern `^[a-zA-Z0-9]+$` |
| `contexts.attributes.value` | yes | pattern `^[\S\s]+$` |
| `operations` (block) | no | nested block |
| `operations.api_types` (block) | yes | nested block |
| `operations.api_types.api_type_id` | yes | pattern `^[a-zA-Z0-9_.\-:]+$` |
| `operations.api_types.description` | no | API: The description of the rule. |
| `operations.api_types.display_name` | no | — |
| `resources` (block) | yes | nested block |
| `resources.attributes` (block) | yes | nested block |
| `resources.attributes.name` | yes | pattern `^[a-zA-Z0-9_]+$` |
| `resources.attributes.operator` | no | pattern `^[a-zA-Z0-9]+$` |
| `resources.attributes.value` | yes | pattern `^[\S\s]+$` |
| `resources.tags` (block) | no | nested block |
| `resources.tags.name` | yes | pattern `^[a-zA-Z0-9 _.-]+$` |
| `resources.tags.operator` | no | pattern `^[a-zA-Z0-9]+$` |
| `resources.tags.value` | yes | pattern `^[a-zA-Z0-9 _*?.-]+$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cbr_rule_invalid_description` | `description` |
| `ibm_cbr_rule_invalid_enforcement_mode` | `enforcement_mode` |
| `ibm_cbr_rule_invalid_transaction_id` | `transaction_id` |
| `ibm_cbr_rule_invalid_x_correlation_id` | `x_correlation_id` |
| `ibm_cbr_rule_invalid_contexts__attributes__name` | `contexts.attributes.name` |
| `ibm_cbr_rule_invalid_contexts__attributes__value` | `contexts.attributes.value` |
| `ibm_cbr_rule_invalid_operations__api_types__api_type_id` | `operations.api_types.api_type_id` |
| `ibm_cbr_rule_invalid_operations__api_types__description` | `operations.api_types.description` |
| `ibm_cbr_rule_invalid_operations__api_types__display_name` | `operations.api_types.display_name` |
| `ibm_cbr_rule_invalid_resources__attributes__name` | `resources.attributes.name` |
| `ibm_cbr_rule_invalid_resources__attributes__operator` | `resources.attributes.operator` |
| `ibm_cbr_rule_invalid_resources__attributes__value` | `resources.attributes.value` |
| `ibm_cbr_rule_invalid_resources__tags__name` | `resources.tags.name` |
| `ibm_cbr_rule_invalid_resources__tags__operator` | `resources.tags.operator` |
| `ibm_cbr_rule_invalid_resources__tags__value` | `resources.tags.value` |

## Example

### Valid

```hcl
resource "ibm_cbr_rule" "test" {
  resources {
    attributes {
      name = "example"
      value = "example-name"
    }
  }
}
```
