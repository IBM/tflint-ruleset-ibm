# ibm_logs_rule_group

Manages ibm_logs_rule_group. The name of the rule group.

This rule validates `ibm_logs_rule_group` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/v1/rule_groups/{group_id}:put` (`logs-service-api`).

## Configuration

```hcl
rule "ibm_logs_rule_group" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `description` | no | pattern `^[A-Za-z0-9_\-\s]+$` |
| `enabled` | no | API: Whether or not the rule is enabled |
| `name` | yes | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `order` | no | range 0–4294967295 |
| `rule_group_id` | no | — |
| `rule_matchers` (block) | no | nested block |
| `rule_matchers.application_name` (block) | no | nested block |
| `rule_matchers.application_name.value` | yes | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `rule_matchers.severity` (block) | no | nested block |
| `rule_matchers.severity.value` | yes | one of: `debug_or_unspecified`, `verbose`, `info`, `warning`, `error`, `critical` |
| `rule_matchers.subsystem_name` (block) | no | nested block |
| `rule_matchers.subsystem_name.value` | yes | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `rule_subgroups` (block) | yes | nested block |
| `rule_subgroups.enabled` | no | API: Whether or not the rule subgroup is enabled |
| `rule_subgroups.id` | no | — |
| `rule_subgroups.order` | yes | range 0–4294967295 |
| `rule_subgroups.rules` (block) | yes | nested block |
| `rule_subgroups.rules.description` | no | pattern `^[\p{L}0-9_\-\s]+$` |
| `rule_subgroups.rules.enabled` | yes | API: Rule parameters which define how the rule will execute |
| `rule_subgroups.rules.id` | no | — |
| `rule_subgroups.rules.name` | yes | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `rule_subgroups.rules.order` | yes | range 0–4294967295 |
| `rule_subgroups.rules.source_field` | yes | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `rule_subgroups.rules.parameters` (block) | yes | nested block |
| `rule_subgroups.rules.parameters.allow_parameters` (block) | no | nested block |
| `rule_subgroups.rules.parameters.allow_parameters.keep_blocked_logs` | yes | API: If true, matched logs will be blocked, otherwise matched logs will be kept |
| `rule_subgroups.rules.parameters.allow_parameters.rule` | yes | API: Regex which will match the source field and decide if the rule will apply |
| `rule_subgroups.rules.parameters.block_parameters` (block) | no | nested block |
| `rule_subgroups.rules.parameters.block_parameters.keep_blocked_logs` | yes | API: If true, matched logs will be kept, otherwise matched logs will be blocked |
| `rule_subgroups.rules.parameters.block_parameters.rule` | yes | API: Regex which will match the source field and decide if the rule will apply |
| `rule_subgroups.rules.parameters.extract_parameters` (block) | no | nested block |
| `rule_subgroups.rules.parameters.extract_parameters.rule` | yes | API: Regex which will parse the source field and extract the JSON keys from source field while retaining the original log |
| `rule_subgroups.rules.parameters.extract_timestamp_parameters` (block) | no | nested block |
| `rule_subgroups.rules.parameters.extract_timestamp_parameters.format` | yes | API: The time format of the source field being extracted. |
| `rule_subgroups.rules.parameters.extract_timestamp_parameters.standard` | yes | one of: `strftime_or_unspecified`, `javasdf`, `golang`, `secondsts`, `millits`, `microts`, `nanots` |
| `rule_subgroups.rules.parameters.json_extract_parameters` (block) | no | nested block |
| `rule_subgroups.rules.parameters.json_extract_parameters.destination_field` | no | one of: `category_or_unspecified`, `classname`, `methodname`, `threadid`, `severity` |
| `rule_subgroups.rules.parameters.json_parse_parameters` (block) | no | nested block |
| `rule_subgroups.rules.parameters.json_parse_parameters.delete_source` | no | API: Whether or not to delete the source field after running this rule |
| `rule_subgroups.rules.parameters.json_parse_parameters.destination_field` | yes | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `rule_subgroups.rules.parameters.json_parse_parameters.override_dest` | yes | API: The destination field where the JSON stringified content will be placed. |
| `rule_subgroups.rules.parameters.json_stringify_parameters` (block) | no | nested block |
| `rule_subgroups.rules.parameters.json_stringify_parameters.delete_source` | no | API: Whether or not to delete the source field after running this rule |
| `rule_subgroups.rules.parameters.json_stringify_parameters.destination_field` | yes | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `rule_subgroups.rules.parameters.parse_parameters` (block) | no | nested block |
| `rule_subgroups.rules.parameters.parse_parameters.destination_field` | yes | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `rule_subgroups.rules.parameters.parse_parameters.rule` | yes | API: Regex which will parse the source field and extract the JSON keys from it while removing the source field |
| `rule_subgroups.rules.parameters.remove_fields_parameters` (block) | no | nested block |
| `rule_subgroups.rules.parameters.remove_fields_parameters.fields` | yes | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `rule_subgroups.rules.parameters.replace_parameters` (block) | no | nested block |
| `rule_subgroups.rules.parameters.replace_parameters.destination_field` | yes | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `rule_subgroups.rules.parameters.replace_parameters.replace_new_val` | yes | API: The value replacing the matched text. |
| `rule_subgroups.rules.parameters.replace_parameters.rule` | yes | API: Regex which will match parts in the text to replace |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_logs_rule_group_invalid_description` | `description` |
| `ibm_logs_rule_group_invalid_enabled` | `enabled` |
| `ibm_logs_rule_group_invalid_name` | `name` |
| `ibm_logs_rule_group_invalid_order` | `order` |
| `ibm_logs_rule_group_invalid_rule_matchers__application_name__value` | `rule_matchers.application_name.value` |
| `ibm_logs_rule_group_invalid_rule_matchers__severity__value` | `rule_matchers.severity.value` |
| `ibm_logs_rule_group_invalid_rule_matchers__subsystem_name__value` | `rule_matchers.subsystem_name.value` |
| `ibm_logs_rule_group_invalid_rule_subgroups__enabled` | `rule_subgroups.enabled` |
| `ibm_logs_rule_group_invalid_rule_subgroups__id` | `rule_subgroups.id` |
| `ibm_logs_rule_group_invalid_rule_subgroups__order` | `rule_subgroups.order` |
| `ibm_logs_rule_group_invalid_rule_subgroups__rules__description` | `rule_subgroups.rules.description` |
| `ibm_logs_rule_group_invalid_rule_subgroups__rules__enabled` | `rule_subgroups.rules.enabled` |
| `ibm_logs_rule_group_invalid_rule_subgroups__rules__id` | `rule_subgroups.rules.id` |
| `ibm_logs_rule_group_invalid_rule_subgroups__rules__name` | `rule_subgroups.rules.name` |
| `ibm_logs_rule_group_invalid_rule_subgroups__rules__order` | `rule_subgroups.rules.order` |
| `ibm_logs_rule_group_invalid_rule_subgroups__rules__source_field` | `rule_subgroups.rules.source_field` |
| `ibm_logs_rule_group_invalid_rule_subgroups__rules__parameters__allow_parameters__keep_blocked_logs` | `rule_subgroups.rules.parameters.allow_parameters.keep_blocked_logs` |
| `ibm_logs_rule_group_invalid_rule_subgroups__rules__parameters__allow_parameters__rule` | `rule_subgroups.rules.parameters.allow_parameters.rule` |
| `ibm_logs_rule_group_invalid_rule_subgroups__rules__parameters__block_parameters__keep_blocked_logs` | `rule_subgroups.rules.parameters.block_parameters.keep_blocked_logs` |
| `ibm_logs_rule_group_invalid_rule_subgroups__rules__parameters__block_parameters__rule` | `rule_subgroups.rules.parameters.block_parameters.rule` |
| `ibm_logs_rule_group_invalid_rule_subgroups__rules__parameters__extract_parameters__rule` | `rule_subgroups.rules.parameters.extract_parameters.rule` |
| `ibm_logs_rule_group_invalid_rule_subgroups__rules__parameters__extract_timestamp_parameters__format` | `rule_subgroups.rules.parameters.extract_timestamp_parameters.format` |
| `ibm_logs_rule_group_invalid_rule_subgroups__rules__parameters__extract_timestamp_parameters__standard` | `rule_subgroups.rules.parameters.extract_timestamp_parameters.standard` |
| `ibm_logs_rule_group_invalid_rule_subgroups__rules__parameters__json_extract_parameters__destination_field` | `rule_subgroups.rules.parameters.json_extract_parameters.destination_field` |
| `ibm_logs_rule_group_invalid_rule_subgroups__rules__parameters__json_parse_parameters__delete_source` | `rule_subgroups.rules.parameters.json_parse_parameters.delete_source` |
| `ibm_logs_rule_group_invalid_rule_subgroups__rules__parameters__json_parse_parameters__destination_field` | `rule_subgroups.rules.parameters.json_parse_parameters.destination_field` |
| `ibm_logs_rule_group_invalid_rule_subgroups__rules__parameters__json_parse_parameters__override_dest` | `rule_subgroups.rules.parameters.json_parse_parameters.override_dest` |
| `ibm_logs_rule_group_invalid_rule_subgroups__rules__parameters__json_stringify_parameters__delete_source` | `rule_subgroups.rules.parameters.json_stringify_parameters.delete_source` |
| `ibm_logs_rule_group_invalid_rule_subgroups__rules__parameters__json_stringify_parameters__destination_field` | `rule_subgroups.rules.parameters.json_stringify_parameters.destination_field` |
| `ibm_logs_rule_group_invalid_rule_subgroups__rules__parameters__parse_parameters__destination_field` | `rule_subgroups.rules.parameters.parse_parameters.destination_field` |
| `ibm_logs_rule_group_invalid_rule_subgroups__rules__parameters__parse_parameters__rule` | `rule_subgroups.rules.parameters.parse_parameters.rule` |
| `ibm_logs_rule_group_invalid_rule_subgroups__rules__parameters__remove_fields_parameters__fields` | `rule_subgroups.rules.parameters.remove_fields_parameters.fields` |
| `ibm_logs_rule_group_invalid_rule_subgroups__rules__parameters__replace_parameters__destination_field` | `rule_subgroups.rules.parameters.replace_parameters.destination_field` |
| `ibm_logs_rule_group_invalid_rule_subgroups__rules__parameters__replace_parameters__replace_new_val` | `rule_subgroups.rules.parameters.replace_parameters.replace_new_val` |
| `ibm_logs_rule_group_invalid_rule_subgroups__rules__parameters__replace_parameters__rule` | `rule_subgroups.rules.parameters.replace_parameters.rule` |

## Example

### Valid

```hcl
resource "ibm_logs_rule_group" "test" {
  name = "example-name"
  rule_subgroups {
    order = 0
    rules {
      enabled = true
      name = "example-name"
      order = 0
      source_field = "example-name"
      parameters {
      }
    }
  }
}
```
