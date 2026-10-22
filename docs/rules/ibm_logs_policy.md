# ibm_logs_policy

Manages ibm_logs_policy. Name of policy.

This rule validates `ibm_logs_policy` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Policy` (`logs-service-api`).

## Configuration

```hcl
rule "ibm_logs_policy" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `archive_retention_tag` | no | pattern `^[a-zA-Z0-9_-]+$` |
| `company_id` | no | — |
| `created_at` | no | — |
| `deleted` | no | — |
| `description` | no | pattern `^[\p{L}0-9_\-\s]+$` |
| `enabled` | no | API: Flag to enable or disable a policy. This flag is supported only while updating a policy, since the policies are always enabled during creation. |
| `name` | yes | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `order` | no | — |
| `policy_id` | no | — |
| `priority` | yes | one of: `type_block`, `type_high`, `type_low`, `type_medium`, `type_unspecified` |
| `updated_at` | no | — |
| `application_rule` (block) | no | nested block |
| `application_rule.name` | yes | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `application_rule.rule_type_id` | yes | one of: `unspecified`, `is`, `is_not`, `start_with`, `includes` |
| `archive_retention` (block) | no | nested block |
| `archive_retention.id` | yes | pattern `^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$` |
| `before` (block) | no | nested block |
| `before.id` | no | pattern `^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$` |
| `log_rules` (block) | no | nested block |
| `log_rules.severities` | no | — |
| `subsystem_rule` (block) | no | nested block |
| `subsystem_rule.name` | yes | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `subsystem_rule.rule_type_id` | yes | one of: `unspecified`, `is`, `is_not`, `start_with`, `includes` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_logs_policy_invalid_archive_retention_tag` | `archive_retention_tag` |
| `ibm_logs_policy_invalid_description` | `description` |
| `ibm_logs_policy_invalid_enabled` | `enabled` |
| `ibm_logs_policy_invalid_name` | `name` |
| `ibm_logs_policy_invalid_priority` | `priority` |
| `ibm_logs_policy_invalid_application_rule__name` | `application_rule.name` |
| `ibm_logs_policy_invalid_application_rule__rule_type_id` | `application_rule.rule_type_id` |
| `ibm_logs_policy_invalid_archive_retention__id` | `archive_retention.id` |
| `ibm_logs_policy_invalid_before__id` | `before.id` |
| `ibm_logs_policy_invalid_log_rules__severities` | `log_rules.severities` |
| `ibm_logs_policy_invalid_subsystem_rule__name` | `subsystem_rule.name` |
| `ibm_logs_policy_invalid_subsystem_rule__rule_type_id` | `subsystem_rule.rule_type_id` |

## Example

### Valid

```hcl
resource "ibm_logs_policy" "test" {
  name = "example-name"
  priority = "type_block"
}
```
