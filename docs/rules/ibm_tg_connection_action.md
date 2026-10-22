# ibm_tg_connection_action

Schema validation for the ibm_tg_connection_action resource.

This rule validates `ibm_tg_connection_action` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `TransitGatewayConnectionActions` (`transit-gateway`).

## Configuration

```hcl
rule "ibm_tg_connection_action" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `action` | yes | one of: `approve`, `reject` |
| `connection_id` | yes | pattern `^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$` |
| `gateway` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_tg_connection_action_invalid_action` | `action` |
| `ibm_tg_connection_action_invalid_connection_id` | `connection_id` |
| `ibm_tg_connection_action_invalid_gateway` | `gateway` |

## Example

### Valid

```hcl
resource "ibm_tg_connection_action" "test" {
  action = "approve"
  connection_id = "01234567-89ab-cdef-0123-456789abcdef"
  gateway = "example"
}
```
