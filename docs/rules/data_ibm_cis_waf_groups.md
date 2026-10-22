# ibm_cis_waf_groups_data

Schema validation for the ibm_cis_waf_groups data source.

This rule validates `ibm_cis_waf_groups` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_cis_waf_groups_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cis_id` | yes | — |
| `domain_id` | yes | — |
| `package_id` | yes | — |
| `waf_groups` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cis_waf_groups_data_invalid_cis_id` | `cis_id` |
| `ibm_cis_waf_groups_data_invalid_domain_id` | `domain_id` |
| `ibm_cis_waf_groups_data_invalid_package_id` | `package_id` |

## Example

### Valid

```hcl
data "ibm_cis_waf_groups" "test" {
  cis_id = "01234567-89ab-cdef-0123-456789abcdef"
  domain_id = "01234567-89ab-cdef-0123-456789abcdef"
  package_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
