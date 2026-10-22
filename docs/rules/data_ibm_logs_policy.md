# ibm_logs_policy_data

Reads ibm_logs_policy. Name of policy.

This rule validates `ibm_logs_policy` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_logs_policy_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `application_rule` | no | — |
| `archive_retention` | no | — |
| `archive_retention_tag` | no | — |
| `before` | no | — |
| `company_id` | no | — |
| `created_at` | no | — |
| `deleted` | no | — |
| `description` | no | — |
| `enabled` | no | — |
| `log_rules` | no | — |
| `logs_policy_id` | yes | — |
| `name` | no | — |
| `order` | no | — |
| `priority` | no | — |
| `subsystem_rule` | no | — |
| `updated_at` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_logs_policy_data_invalid_logs_policy_id` | `logs_policy_id` |

## Example

### Valid

```hcl
data "ibm_logs_policy" "test" {
  logs_policy_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
