# ibm_cis_edge_functions_trigger

Schema validation for the ibm_cis_edge_functions_trigger resource.

This rule validates `ibm_cis_edge_functions_trigger` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_cis_edge_functions_trigger" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `action_name` | no | — |
| `cis_id` | yes | — |
| `domain_id` | yes | — |
| `pattern_url` | yes | — |
| `request_limit_fail_open` | no | — |
| `trigger_id` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cis_edge_functions_trigger_invalid_action_name` | `action_name` |
| `ibm_cis_edge_functions_trigger_invalid_cis_id` | `cis_id` |
| `ibm_cis_edge_functions_trigger_invalid_domain_id` | `domain_id` |
| `ibm_cis_edge_functions_trigger_invalid_pattern_url` | `pattern_url` |

## Example

### Valid

```hcl
resource "ibm_cis_edge_functions_trigger" "test" {
  cis_id = "01234567-89ab-cdef-0123-456789abcdef"
  domain_id = "01234567-89ab-cdef-0123-456789abcdef"
  pattern_url = "example"
}
```
