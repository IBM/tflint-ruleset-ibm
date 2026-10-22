# ibm_cis_waf_rule

Schema validation for the ibm_cis_waf_rule resource.

This rule validates `ibm_cis_waf_rule` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `waf-groups-request` (`cis`).

## Configuration

```hcl
rule "ibm_cis_waf_rule" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `allowed_modes` | no | — |
| `cis_id` | yes | — |
| `description` | no | — |
| `domain_id` | yes | — |
| `group` | no | — |
| `mode` | yes | one of: `on`, `off`, `default`, `disable`, `simulate`, `block`, `challenge` |
| `package_id` | yes | — |
| `priority` | no | — |
| `rule_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cis_waf_rule_invalid_cis_id` | `cis_id` |
| `ibm_cis_waf_rule_invalid_domain_id` | `domain_id` |
| `ibm_cis_waf_rule_invalid_mode` | `mode` |
| `ibm_cis_waf_rule_invalid_package_id` | `package_id` |
| `ibm_cis_waf_rule_invalid_rule_id` | `rule_id` |

## Example

### Valid

```hcl
resource "ibm_cis_waf_rule" "test" {
  cis_id = "01234567-89ab-cdef-0123-456789abcdef"
  domain_id = "01234567-89ab-cdef-0123-456789abcdef"
  mode = "on"
  package_id = "01234567-89ab-cdef-0123-456789abcdef"
  rule_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
