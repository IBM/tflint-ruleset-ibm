# ibm_dns_reverse_record

Schema validation for the ibm_dns_reverse_record resource.

This rule validates `ibm_dns_reverse_record` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_dns_reverse_record" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `hostname` | yes | — |
| `ipaddress` | yes | — |
| `ttl` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_dns_reverse_record_invalid_hostname` | `hostname` |
| `ibm_dns_reverse_record_invalid_ipaddress` | `ipaddress` |
| `ibm_dns_reverse_record_invalid_ttl` | `ttl` |

## Example

### Valid

```hcl
resource "ibm_dns_reverse_record" "test" {
  hostname = "example"
  ipaddress = "example"
}
```
