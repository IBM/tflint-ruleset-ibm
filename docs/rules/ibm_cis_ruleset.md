# ibm_cis_ruleset

Schema validation for the ibm_cis_ruleset resource.

This rule validates `ibm_cis_ruleset` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ruleset_create` (`cis`).

## Configuration

```hcl
rule "ibm_cis_ruleset" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cis_id` | yes | — |
| `domain_id` | no | — |
| `ruleset_id` | no | — |
| `rulesets` | no | API: List of ruleset ids to apply action to. Use "current" to apply to the current ruleset. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cis_ruleset_invalid_cis_id` | `cis_id` |
| `ibm_cis_ruleset_invalid_domain_id` | `domain_id` |
| `ibm_cis_ruleset_invalid_ruleset_id` | `ruleset_id` |
| `ibm_cis_ruleset_invalid_rulesets` | `rulesets` |

## Example

### Valid

```hcl
resource "ibm_cis_ruleset" "test" {
  cis_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
