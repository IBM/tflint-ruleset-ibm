# ibm_cis_routing

Schema validation for the ibm_cis_routing resource.

This rule validates `ibm_cis_routing` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_cis_routing" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cis_id` | yes | — |
| `domain_id` | yes | — |
| `smart_routing` | no | one of: `on`, `off` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cis_routing_invalid_cis_id` | `cis_id` |
| `ibm_cis_routing_invalid_domain_id` | `domain_id` |
| `ibm_cis_routing_invalid_smart_routing` | `smart_routing` |

## Example

### Valid

```hcl
resource "ibm_cis_routing" "test" {
  cis_id = "01234567-89ab-cdef-0123-456789abcdef"
  domain_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
