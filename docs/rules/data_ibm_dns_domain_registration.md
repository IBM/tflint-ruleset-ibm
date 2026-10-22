# ibm_dns_domain_registration_data

Reads ibm_dns_domain_registration. The name of the domain registration

This rule validates `ibm_dns_domain_registration` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_dns_domain_registration_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `id` | no | — |
| `name` | yes | — |
| `name_servers` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_dns_domain_registration_data_invalid_name` | `name` |

## Example

### Valid

```hcl
data "ibm_dns_domain_registration" "test" {
  name = "example"
}
```
