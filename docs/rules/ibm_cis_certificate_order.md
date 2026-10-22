# ibm_cis_certificate_order

Schema validation for the ibm_cis_certificate_order resource.

This rule validates `ibm_cis_certificate_order` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `certificate_req` (`cis`).

## Configuration

```hcl
rule "ibm_cis_certificate_order" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `certificate_id` | no | — |
| `cis_id` | yes | — |
| `domain_id` | yes | — |
| `hosts` | yes | API: Host name |
| `status` | no | — |
| `type` | no | one of: `dedicated` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cis_certificate_order_invalid_cis_id` | `cis_id` |
| `ibm_cis_certificate_order_invalid_domain_id` | `domain_id` |
| `ibm_cis_certificate_order_invalid_hosts` | `hosts` |
| `ibm_cis_certificate_order_invalid_type` | `type` |

## Example

### Valid

```hcl
resource "ibm_cis_certificate_order" "test" {
  cis_id = "01234567-89ab-cdef-0123-456789abcdef"
  domain_id = "01234567-89ab-cdef-0123-456789abcdef"
  hosts = "example"
}
```
