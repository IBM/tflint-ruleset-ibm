# ibm_logs_rule_group_data

Reads ibm_logs_rule_group. The name of the rule group.

This rule validates `ibm_logs_rule_group` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `params:put:/v1/rule_groups/{group_id}` (`logs-service-api`).

## Configuration

```hcl
rule "ibm_logs_rule_group_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `description` | no | — |
| `enabled` | no | — |
| `group_id` | yes | pattern `^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$` |
| `name` | no | — |
| `order` | no | — |
| `rule_matchers` | no | — |
| `rule_subgroups` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_logs_rule_group_data_invalid_group_id` | `group_id` |

## Example

### Valid

```hcl
data "ibm_logs_rule_group" "test" {
  group_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
