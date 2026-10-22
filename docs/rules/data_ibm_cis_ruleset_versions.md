# ibm_cis_ruleset_versions_data

Schema validation for the ibm_cis_ruleset_versions data source.

This rule validates `ibm_cis_ruleset_versions` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ruleset_create` (`cis`).

## Configuration

```hcl
rule "ibm_cis_ruleset_versions_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cis_id` | yes | — |
| `domain_id` | no | — |
| `ruleset_id` | yes | — |
| `ruleset_versions` | no | — |
| `rulesets` | no | — |
| `version` | no | API: The version of the ruleset. Use "latest" to get the latest version. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cis_ruleset_versions_data_invalid_cis_id` | `cis_id` |
| `ibm_cis_ruleset_versions_data_invalid_domain_id` | `domain_id` |
| `ibm_cis_ruleset_versions_data_invalid_ruleset_id` | `ruleset_id` |
| `ibm_cis_ruleset_versions_data_invalid_version` | `version` |

## Example

### Valid

```hcl
data "ibm_cis_ruleset_versions" "test" {
  cis_id = "01234567-89ab-cdef-0123-456789abcdef"
  ruleset_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
