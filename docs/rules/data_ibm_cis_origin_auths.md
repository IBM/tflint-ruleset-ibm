# ibm_cis_origin_auths_data

Schema validation for the ibm_cis_origin_auths data source.

This rule validates `ibm_cis_origin_auths` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `hostname_origin_pull_settings` (`cis`).

## Configuration

```hcl
rule "ibm_cis_origin_auths_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cis_id` | yes | — |
| `domain_id` | yes | — |
| `hostname` | no | API: the hostname on the origin for which the client certificate uploaded will be used |
| `origin_pull_certs` | no | — |
| `origin_pull_settings_enabled` | no | — |
| `request_type` | no | one of: `origin-rsa`, `origin-ecc` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cis_origin_auths_data_invalid_cis_id` | `cis_id` |
| `ibm_cis_origin_auths_data_invalid_domain_id` | `domain_id` |
| `ibm_cis_origin_auths_data_invalid_hostname` | `hostname` |
| `ibm_cis_origin_auths_data_invalid_request_type` | `request_type` |

## Example

### Valid

```hcl
data "ibm_cis_origin_auths" "test" {
  cis_id = "01234567-89ab-cdef-0123-456789abcdef"
  domain_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
