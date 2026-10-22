# ibm_cis_waf_package

Manages ibm_cis_waf_package. WAF pakcage name

This rule validates `ibm_cis_waf_package` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/v1/{crn}/zones/{zone_id}/firewall/waf/packages/{package_id}:patch` (`cis`).

## Configuration

```hcl
rule "ibm_cis_waf_package" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `action_mode` | yes | one of: `simulate`, `block`, `challenge` |
| `cis_id` | yes | — |
| `description` | no | — |
| `detection_mode` | no | — |
| `domain_id` | yes | — |
| `name` | no | — |
| `package_id` | yes | — |
| `sensitivity` | yes | one of: `high`, `medium`, `low`, `off` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cis_waf_package_invalid_action_mode` | `action_mode` |
| `ibm_cis_waf_package_invalid_cis_id` | `cis_id` |
| `ibm_cis_waf_package_invalid_domain_id` | `domain_id` |
| `ibm_cis_waf_package_invalid_package_id` | `package_id` |
| `ibm_cis_waf_package_invalid_sensitivity` | `sensitivity` |

## Example

### Valid

```hcl
resource "ibm_cis_waf_package" "test" {
  action_mode = "simulate"
  cis_id = "01234567-89ab-cdef-0123-456789abcdef"
  domain_id = "01234567-89ab-cdef-0123-456789abcdef"
  package_id = "01234567-89ab-cdef-0123-456789abcdef"
  sensitivity = "high"
}
```
