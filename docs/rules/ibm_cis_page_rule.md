# ibm_cis_page_rule

Schema validation for the ibm_cis_page_rule resource.

This rule validates `ibm_cis_page_rule` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `rule_create` (`cis`).

## Configuration

```hcl
rule "ibm_cis_page_rule" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cis_id` | yes | — |
| `domain_id` | yes | — |
| `priority` | no | API: Priority |
| `rule_id` | no | — |
| `status` | no | one of: `active`, `disabled` |
| `actions` (block) | yes | nested block |
| `actions.css` | no | one of: `on`, `off` |
| `actions.html` | no | one of: `on`, `off` |
| `actions.id` | yes | API: Unique ID of the ruleset |
| `actions.js` | no | one of: `on`, `off` |
| `actions.status_code` | no | range 400–499 |
| `actions.url` | no | API: The URL pattern to match comprised of the host and path, i.e. example.org/path. Wildcard are expanded to match applicable traffic, query strings are not matched. Use * for all traffic to your zone. |
| `actions.value` | no | — |
| `targets` (block) | yes | nested block |
| `targets.target` | yes | — |
| `targets.constraint` (block) | yes | nested block |
| `targets.constraint.operator` | yes | — |
| `targets.constraint.value` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cis_page_rule_invalid_cis_id` | `cis_id` |
| `ibm_cis_page_rule_invalid_domain_id` | `domain_id` |
| `ibm_cis_page_rule_invalid_priority` | `priority` |
| `ibm_cis_page_rule_invalid_status` | `status` |
| `ibm_cis_page_rule_invalid_actions__css` | `actions.css` |
| `ibm_cis_page_rule_invalid_actions__html` | `actions.html` |
| `ibm_cis_page_rule_invalid_actions__id` | `actions.id` |
| `ibm_cis_page_rule_invalid_actions__js` | `actions.js` |
| `ibm_cis_page_rule_invalid_actions__status_code` | `actions.status_code` |
| `ibm_cis_page_rule_invalid_actions__url` | `actions.url` |
| `ibm_cis_page_rule_invalid_actions__value` | `actions.value` |
| `ibm_cis_page_rule_invalid_targets__target` | `targets.target` |
| `ibm_cis_page_rule_invalid_targets__constraint__operator` | `targets.constraint.operator` |
| `ibm_cis_page_rule_invalid_targets__constraint__value` | `targets.constraint.value` |

## Example

### Valid

```hcl
resource "ibm_cis_page_rule" "test" {
  cis_id = "01234567-89ab-cdef-0123-456789abcdef"
  domain_id = "01234567-89ab-cdef-0123-456789abcdef"
  actions {
    id = "example"
  }
  targets {
    target = "example"
    constraint {
      operator = "example"
      value = "example"
    }
  }
}
```
