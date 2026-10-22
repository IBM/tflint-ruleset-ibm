# ibm_logs_data_access_rule

Schema validation for the ibm_logs_data_access_rule resource.

This rule validates `ibm_logs_data_access_rule` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `DataAccessRulePrototype` (`logs-service-api`).

## Configuration

```hcl
rule "ibm_logs_data_access_rule" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_rule_id` | no | — |
| `default_expression` | yes | pattern `^[A-Za-z0-9_\.,\-"{}()\[\]=!:#\/$|'<> ]+$` |
| `description` | no | pattern `^[A-Za-z0-9_\-\s]+$` |
| `display_name` | yes | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `filters` (block) | yes | nested block |
| `filters.entity_type` | yes | one of: `unspecified`, `logs` |
| `filters.expression` | yes | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_logs_data_access_rule_invalid_default_expression` | `default_expression` |
| `ibm_logs_data_access_rule_invalid_description` | `description` |
| `ibm_logs_data_access_rule_invalid_display_name` | `display_name` |
| `ibm_logs_data_access_rule_invalid_filters__entity_type` | `filters.entity_type` |
| `ibm_logs_data_access_rule_invalid_filters__expression` | `filters.expression` |

## Example

### Valid

```hcl
resource "ibm_logs_data_access_rule" "test" {
  default_expression = "example-name"
  display_name = "example-name"
  filters {
    entity_type = "unspecified"
    expression = "example-name"
  }
}
```
