# ibm_cis_ruleset_rule

Schema validation for the ibm_cis_ruleset_rule resource.

This rule validates `ibm_cis_ruleset_rule` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_cis_ruleset_rule" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cis_id` | yes | — |
| `domain_id` | no | — |
| `rule` | no | — |
| `ruleset_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cis_ruleset_rule_invalid_cis_id` | `cis_id` |
| `ibm_cis_ruleset_rule_invalid_domain_id` | `domain_id` |
| `ibm_cis_ruleset_rule_invalid_rule` | `rule` |
| `ibm_cis_ruleset_rule_invalid_ruleset_id` | `ruleset_id` |

## Example

### Valid

```hcl
resource "ibm_cis_ruleset_rule" "test" {
  cis_id = "01234567-89ab-cdef-0123-456789abcdef"
  ruleset_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
