# ibm_cis_ip_addresses_data

Schema validation for the ibm_cis_ip_addresses data source.

This rule validates `ibm_cis_ip_addresses` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_cis_ip_addresses_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `ipv4_cidrs` | no | — |
| `ipv6_cidrs` | no | — |

## Example

### Valid

```hcl
data "ibm_cis_ip_addresses" "test" {
}
```
