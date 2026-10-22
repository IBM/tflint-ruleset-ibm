# ibm_scc_rule

Manages ibm_scc_rule. A human-readable alias to assign to your rule.

This rule validates `ibm_scc_rule` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Rule` (`security-compliance`).

## Configuration

```hcl
rule "ibm_scc_rule" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_id` | no | — |
| `created_by` | no | — |
| `created_on` | no | — |
| `creation_date` | no | — |
| `description` | yes | pattern `[A-Za-z0-9]+` |
| `etag` | no | — |
| `instance_id` | yes | — |
| `labels` | no | pattern `^[A-Za-z0-9]+$` |
| `modification_date` | no | — |
| `modified_by` | no | — |
| `name` | no | pattern `^[A-Za-z0-9]+$` |
| `rule_id` | no | — |
| `rule_type` | no | — |
| `type` | no | — |
| `updated_by` | no | — |
| `updated_on` | no | — |
| `version` | no | pattern `^[0-9][0-9.]*$` |
| `enforcement_actions` (block) | no | nested block |
| `enforcement_actions.action` | yes | — |
| `import` (block) | no | nested block |
| `import.parameters` (block) | no | nested block |
| `import.parameters.description` | no | pattern `^[A-Za-z0-9]+$` |
| `import.parameters.display_name` | no | pattern `^[A-Za-z0-9]+$` |
| `import.parameters.name` | no | pattern `^[A-Za-z0-9]+$` |
| `import.parameters.type` | no | one of: `string`, `numeric`, `general`, `boolean`, `string_list`, `ip_list`, `timestamp`; pattern `^[A-Za-z]+$` |
| `required_config` (block) | yes | nested block |
| `required_config.description` | no | pattern `^[A-Za-z0-9]+$` |
| `required_config.operator` | no | one of: `string_equals`, `string_not_equals`, `string_match`, `string_not_match`, `string_contains`, `string_not_contains`, `num_equals`, `num_not_equals`, `num_less_than`, `num_less_than_equals`, `num_greater_than`, `num_greater_than_equals`, `is_empty`, `is_not_empty`, `is_true`, `is_false`, `strings_in_list`, `strings_allowed`, `strings_required`, `ips_in_range`, `ips_equals`, `ips_not_equals`, `days_less_than` |
| `required_config.property` | no | pattern `^[A-Za-z0-9]+$` |
| `required_config.value` | no | — |
| `required_config.all` (block) | no | nested block |
| `required_config.all_if` (block) | no | nested block |
| `required_config.and` (block) | no | nested block |
| `required_config.any` (block) | no | nested block |
| `required_config.any_if` (block) | no | nested block |
| `required_config.or` (block) | no | nested block |
| `target` (block) | yes | nested block |
| `target.reference_name` | no | — |
| `target.resource_kind` | yes | pattern `^[A-Za-z0-9]+$` |
| `target.service_display_name` | no | API: The display name of the target service. |
| `target.service_name` | yes | API: The target service name. |
| `target.additional_target_attributes` (block) | no | nested block |
| `target.additional_target_attributes.name` | no | pattern `^[_\-A-Za-z0-9]+$` |
| `target.additional_target_attributes.operator` | no | one of: `string_equals`, `string_not_equals`, `string_match`, `string_not_match`, `string_contains`, `string_not_contains`, `num_equals`, `num_not_equals`, `num_less_than`, `num_less_than_equals`, `num_greater_than`, `num_greater_than_equals`, `is_empty`, `is_not_empty`, `is_true`, `is_false`, `strings_in_list`, `strings_allowed`, `strings_required`, `ips_in_range`, `ips_equals`, `ips_not_equals`, `days_less_than` |
| `target.additional_target_attributes.value` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_scc_rule_invalid_description` | `description` |
| `ibm_scc_rule_invalid_instance_id` | `instance_id` |
| `ibm_scc_rule_invalid_labels` | `labels` |
| `ibm_scc_rule_invalid_name` | `name` |
| `ibm_scc_rule_invalid_version` | `version` |
| `ibm_scc_rule_invalid_enforcement_actions__action` | `enforcement_actions.action` |
| `ibm_scc_rule_invalid_import__parameters__description` | `import.parameters.description` |
| `ibm_scc_rule_invalid_import__parameters__display_name` | `import.parameters.display_name` |
| `ibm_scc_rule_invalid_import__parameters__name` | `import.parameters.name` |
| `ibm_scc_rule_invalid_import__parameters__type` | `import.parameters.type` |
| `ibm_scc_rule_invalid_required_config__description` | `required_config.description` |
| `ibm_scc_rule_invalid_required_config__operator` | `required_config.operator` |
| `ibm_scc_rule_invalid_required_config__property` | `required_config.property` |
| `ibm_scc_rule_invalid_required_config__value` | `required_config.value` |
| `ibm_scc_rule_invalid_target__reference_name` | `target.reference_name` |
| `ibm_scc_rule_invalid_target__resource_kind` | `target.resource_kind` |
| `ibm_scc_rule_invalid_target__service_display_name` | `target.service_display_name` |
| `ibm_scc_rule_invalid_target__service_name` | `target.service_name` |
| `ibm_scc_rule_invalid_target__additional_target_attributes__name` | `target.additional_target_attributes.name` |
| `ibm_scc_rule_invalid_target__additional_target_attributes__operator` | `target.additional_target_attributes.operator` |
| `ibm_scc_rule_invalid_target__additional_target_attributes__value` | `target.additional_target_attributes.value` |

## Example

### Valid

```hcl
resource "ibm_scc_rule" "test" {
  description = "example-name"
  instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  required_config {
  }
  target {
    resource_kind = "example"
    service_name = "aaaaaaa"
  }
}
```
