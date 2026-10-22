# ibm_dns_domain

Manages ibm_dns_domain. DNS name

This rule validates `ibm_dns_domain` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_dns_domain" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `name` | yes | — |
| `serial` | no | — |
| `tags` | no | — |
| `target` | no | — |
| `update_date` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_dns_domain_invalid_name` | `name` |
| `ibm_dns_domain_invalid_tags` | `tags` |
| `ibm_dns_domain_invalid_target` | `target` |

## Example

### Valid

```hcl
resource "ibm_dns_domain" "test" {
  name = "example"
}
```
