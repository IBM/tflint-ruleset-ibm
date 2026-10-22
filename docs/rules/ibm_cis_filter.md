# ibm_cis_filter

Schema validation for the ibm_cis_filter resource.

This rule validates `ibm_cis_filter` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `filter_input` (`cis`).

## Configuration

```hcl
rule "ibm_cis_filter" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cis_id` | yes | — |
| `description` | no | one of: `Filter-creation` |
| `domain_id` | yes | — |
| `expression` | yes | API: A filter expression |
| `filter_id` | no | — |
| `paused` | no | API: Indicates if the filter is active. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cis_filter_invalid_cis_id` | `cis_id` |
| `ibm_cis_filter_invalid_description` | `description` |
| `ibm_cis_filter_invalid_domain_id` | `domain_id` |
| `ibm_cis_filter_invalid_expression` | `expression` |
| `ibm_cis_filter_invalid_paused` | `paused` |

## Example

### Valid

```hcl
resource "ibm_cis_filter" "test" {
  cis_id = "01234567-89ab-cdef-0123-456789abcdef"
  domain_id = "01234567-89ab-cdef-0123-456789abcdef"
  expression = "example"
}
```
