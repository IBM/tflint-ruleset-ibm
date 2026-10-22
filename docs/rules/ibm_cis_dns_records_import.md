# ibm_cis_dns_records_import

Schema validation for the ibm_cis_dns_records_import resource.

This rule validates `ibm_cis_dns_records_import` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/v1/{crn}/zones/{zone_identifier}/dns_records_bulk:post` (`cis`).

## Configuration

```hcl
rule "ibm_cis_dns_records_import" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cis_id` | yes | — |
| `domain_id` | yes | — |
| `file` | yes | API: File to upload |
| `records_added` | no | — |
| `total_records_parsed` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cis_dns_records_import_invalid_cis_id` | `cis_id` |
| `ibm_cis_dns_records_import_invalid_domain_id` | `domain_id` |
| `ibm_cis_dns_records_import_invalid_file` | `file` |

## Example

### Valid

```hcl
resource "ibm_cis_dns_records_import" "test" {
  cis_id = "01234567-89ab-cdef-0123-456789abcdef"
  domain_id = "01234567-89ab-cdef-0123-456789abcdef"
  file = "example"
}
```
