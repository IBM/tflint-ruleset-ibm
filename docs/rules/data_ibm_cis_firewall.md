# ibm_cis_firewall_data

Schema validation for the ibm_cis_firewall data source.

This rule validates `ibm_cis_firewall` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_cis_firewall_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_rule` | no | — |
| `cis_id` | yes | — |
| `domain_id` | yes | — |
| `firewall_type` | yes | — |
| `lockdown` | no | — |
| `ua_rule` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cis_firewall_data_invalid_cis_id` | `cis_id` |
| `ibm_cis_firewall_data_invalid_domain_id` | `domain_id` |
| `ibm_cis_firewall_data_invalid_firewall_type` | `firewall_type` |

## Example

### Valid

```hcl
data "ibm_cis_firewall" "test" {
  cis_id = "01234567-89ab-cdef-0123-456789abcdef"
  domain_id = "01234567-89ab-cdef-0123-456789abcdef"
  firewall_type = "example"
}
```
