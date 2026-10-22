# ibm_cis_domain

Schema validation for the ibm_cis_domain resource.

This rule validates `ibm_cis_domain` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/v1/{crn}/zones:post` (`cis`).

## Configuration

```hcl
rule "ibm_cis_domain" {
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
| `type` | no | one of: `full`, `partial` |
| `verification_key` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cis_domain_invalid_cis_id` | `cis_id` |
| `ibm_cis_domain_invalid_domain` | `domain` |
| `ibm_cis_domain_invalid_type` | `type` |

## Example

### Valid

```hcl
resource "ibm_cis_domain" "test" {
  cis_id = "01234567-89ab-cdef-0123-456789abcdef"
  domain = "example"
}
```
