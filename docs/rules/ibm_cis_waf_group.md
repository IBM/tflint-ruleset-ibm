# ibm_cis_waf_group

Manages ibm_cis_waf_group. WAF Rule group name

This rule validates `ibm_cis_waf_group` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_cis_waf_group" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `check_mode` | no | — |
| `cis_id` | yes | — |
| `description` | no | — |
| `domain_id` | yes | — |
| `group_id` | yes | — |
| `mode` | yes | one of: `on`, `off` |
| `modified_rules_count` | no | — |
| `name` | no | — |
| `package_id` | yes | — |
| `rules_count` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cis_waf_group_invalid_check_mode` | `check_mode` |
| `ibm_cis_waf_group_invalid_cis_id` | `cis_id` |
| `ibm_cis_waf_group_invalid_domain_id` | `domain_id` |
| `ibm_cis_waf_group_invalid_group_id` | `group_id` |
| `ibm_cis_waf_group_invalid_mode` | `mode` |
| `ibm_cis_waf_group_invalid_package_id` | `package_id` |

## Example

### Valid

```hcl
resource "ibm_cis_waf_group" "test" {
  cis_id = "01234567-89ab-cdef-0123-456789abcdef"
  domain_id = "01234567-89ab-cdef-0123-456789abcdef"
  group_id = "01234567-89ab-cdef-0123-456789abcdef"
  mode = "on"
  package_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
