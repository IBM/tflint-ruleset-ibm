# ibm_cis_domain_data

Schema validation for the ibm_cis_domain data source.

This rule validates `ibm_cis_domain` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_cis_domain_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cis_id` | yes | — |
| `cname_suffix` | no | — |
| `domain` | yes | — |
| `domain_id` | no | — |
| `name_servers` | no | — |
| `original_name_servers` | no | — |
| `paused` | no | — |
| `status` | no | — |
| `type` | no | — |
| `verification_key` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cis_domain_data_invalid_cis_id` | `cis_id` |
| `ibm_cis_domain_data_invalid_cname_suffix` | `cname_suffix` |
| `ibm_cis_domain_data_invalid_domain` | `domain` |
| `ibm_cis_domain_data_invalid_verification_key` | `verification_key` |

## Example

### Valid

```hcl
data "ibm_cis_domain" "test" {
  cis_id = "01234567-89ab-cdef-0123-456789abcdef"
  domain = "example"
}
```
