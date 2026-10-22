# ibm_cis_firewall_rule

Schema validation for the ibm_cis_firewall_rule resource.

This rule validates `ibm_cis_firewall_rule` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `firewall_rule_update_input` (`cis`).

## Configuration

```hcl
rule "ibm_cis_firewall_rule" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `action` | yes | one of: `log`, `allow`, `challenge`, `js_challenge`, `block` |
| `cis_id` | yes | — |
| `description` | no | API: To briefly describe the firewall rule, omitted from object if empty. |
| `domain_id` | yes | — |
| `filter_id` | yes | — |
| `paused` | no | API: Indicates if the firewall rule is active. |
| `priority` | no | API: Priority |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cis_firewall_rule_invalid_action` | `action` |
| `ibm_cis_firewall_rule_invalid_cis_id` | `cis_id` |
| `ibm_cis_firewall_rule_invalid_description` | `description` |
| `ibm_cis_firewall_rule_invalid_domain_id` | `domain_id` |
| `ibm_cis_firewall_rule_invalid_filter_id` | `filter_id` |
| `ibm_cis_firewall_rule_invalid_paused` | `paused` |
| `ibm_cis_firewall_rule_invalid_priority` | `priority` |

## Example

### Valid

```hcl
resource "ibm_cis_firewall_rule" "test" {
  action = "log"
  cis_id = "01234567-89ab-cdef-0123-456789abcdef"
  domain_id = "01234567-89ab-cdef-0123-456789abcdef"
  filter_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
