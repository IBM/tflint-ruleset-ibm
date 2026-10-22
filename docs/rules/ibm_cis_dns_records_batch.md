# ibm_cis_dns_records_batch

Schema validation for the ibm_cis_dns_records_batch resource.

This rule validates `ibm_cis_dns_records_batch` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_cis_dns_records_batch" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cis_id` | yes | — |
| `domain_id` | yes | — |
| `result_deletes` | no | — |
| `result_patches` | no | — |
| `result_posts` | no | — |
| `result_puts` | no | — |
| `deletes` (block) | no | nested block |
| `patches` (block) | no | nested block |
| `posts` (block) | no | nested block |
| `puts` (block) | no | nested block |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cis_dns_records_batch_invalid_cis_id` | `cis_id` |
| `ibm_cis_dns_records_batch_invalid_domain_id` | `domain_id` |

## Example

### Valid

```hcl
resource "ibm_cis_dns_records_batch" "test" {
  cis_id = "01234567-89ab-cdef-0123-456789abcdef"
  domain_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
