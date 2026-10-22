# ibm_is_bare_metal_server_action

Schema validation for the ibm_is_bare_metal_server_action resource.

This rule validates `ibm_is_bare_metal_server_action` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `LoadBalancerPoolFailsafePolicyPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_bare_metal_server_action" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `action` | yes | one of: `start`, `restart`, `stop`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `bare_metal_server` | yes | pattern `^([a-zA-Z_][a-zA-Z0-9_]*|[-+*/%]|&&|\|\||!|==|!=|<|<=|>|>=|~|\bin\b|\(|\)|\[|\]|,|\.|"|'|"|'|\s+|\d+)+$` |
| `status` | no | — |
| `status_reasons` | no | — |
| `stop_type` | no | one of: `soft`, `hard` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_bare_metal_server_action_invalid_action` | `action` |
| `ibm_is_bare_metal_server_action_invalid_bare_metal_server` | `bare_metal_server` |
| `ibm_is_bare_metal_server_action_invalid_stop_type` | `stop_type` |

## Example

### Valid

```hcl
resource "ibm_is_bare_metal_server_action" "test" {
  action = "start"
  bare_metal_server = "example-name"
}
```
