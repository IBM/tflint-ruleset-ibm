# ibm_cis_ruleset_version_detach

Schema validation for the ibm_cis_ruleset_version_detach resource.

This rule validates `ibm_cis_ruleset_version_detach` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ruleset_create` (`cis`).

## Configuration

```hcl
rule "ibm_cis_ruleset_version_detach" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cis_id` | yes | — |
| `domain_id` | no | — |
| `ruleset_id` | yes | — |
| `version` | yes | API: The version of the ruleset. Use "latest" to get the latest version. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cis_ruleset_version_detach_invalid_cis_id` | `cis_id` |
| `ibm_cis_ruleset_version_detach_invalid_domain_id` | `domain_id` |
| `ibm_cis_ruleset_version_detach_invalid_ruleset_id` | `ruleset_id` |
| `ibm_cis_ruleset_version_detach_invalid_version` | `version` |

## Example

### Valid

```hcl
resource "ibm_cis_ruleset_version_detach" "test" {
  cis_id = "01234567-89ab-cdef-0123-456789abcdef"
  ruleset_id = "01234567-89ab-cdef-0123-456789abcdef"
  version = "example"
}
```
