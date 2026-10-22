# ibm_dns_domain_registration_nameservers

Schema validation for the ibm_dns_domain_registration_nameservers resource.

This rule validates `ibm_dns_domain_registration_nameservers` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_dns_domain_registration_nameservers" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `dns_registration_id` | yes | — |
| `name_servers` | yes | — |
| `original_name_servers` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_dns_domain_registration_nameservers_invalid_dns_registration_id` | `dns_registration_id` |
| `ibm_dns_domain_registration_nameservers_invalid_name_servers` | `name_servers` |

## Example

### Valid

```hcl
resource "ibm_dns_domain_registration_nameservers" "test" {
  dns_registration_id = "01234567-89ab-cdef-0123-456789abcdef"
  name_servers = "example"
}
```
