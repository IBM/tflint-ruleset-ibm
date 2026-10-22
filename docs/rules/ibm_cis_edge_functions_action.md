# ibm_cis_edge_functions_action

Schema validation for the ibm_cis_edge_functions_action resource.

This rule validates `ibm_cis_edge_functions_action` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `edge_functions_action_resp` (`cis`).

## Configuration

```hcl
rule "ibm_cis_edge_functions_action" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `action_name` | yes | — |
| `cis_id` | yes | — |
| `domain_id` | yes | — |
| `script` | yes | API: Raw script content, as a string |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cis_edge_functions_action_invalid_action_name` | `action_name` |
| `ibm_cis_edge_functions_action_invalid_cis_id` | `cis_id` |
| `ibm_cis_edge_functions_action_invalid_domain_id` | `domain_id` |
| `ibm_cis_edge_functions_action_invalid_script` | `script` |

## Example

### Valid

```hcl
resource "ibm_cis_edge_functions_action" "test" {
  action_name = "example"
  cis_id = "01234567-89ab-cdef-0123-456789abcdef"
  domain_id = "01234567-89ab-cdef-0123-456789abcdef"
  script = "example"
}
```
