# ibm_cis_origin_certificates_data

Schema validation for the ibm_cis_origin_certificates data source.

This rule validates `ibm_cis_origin_certificates` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_cis_origin_certificates_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `certificate_id` | no | — |
| `cis_id` | yes | — |
| `domain_id` | yes | — |
| `origin_certificate_list` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cis_origin_certificates_data_invalid_certificate_id` | `certificate_id` |
| `ibm_cis_origin_certificates_data_invalid_cis_id` | `cis_id` |
| `ibm_cis_origin_certificates_data_invalid_domain_id` | `domain_id` |

## Example

### Valid

```hcl
data "ibm_cis_origin_certificates" "test" {
  cis_id = "01234567-89ab-cdef-0123-456789abcdef"
  domain_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
