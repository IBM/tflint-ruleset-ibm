# ibm_dl_virtual_connection

Manages ibm_dl_virtual_connection. The user-defined name for this virtual connection.

This rule validates `ibm_dl_virtual_connection` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `GatewayVirtualConnectionTemplate` (`direct_link`).

## Configuration

```hcl
rule "ibm_dl_virtual_connection" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `created_at` | no | — |
| `gateway` | yes | — |
| `name` | yes | pattern `^([a-zA-Z]|[a-zA-Z][-_a-zA-Z0-9]*[a-zA-Z0-9])$` |
| `network_account` | no | — |
| `network_id` | no | API: Unique identifier of the target network. For type=vpc virtual connections this is the CRN of the target VPC. This field does not apply to type=classic connections. |
| `related_crn` | no | — |
| `status` | no | — |
| `type` | yes | one of: `classic`, `vpc` |
| `virtual_connection_id` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_dl_virtual_connection_invalid_gateway` | `gateway` |
| `ibm_dl_virtual_connection_invalid_name` | `name` |
| `ibm_dl_virtual_connection_invalid_network_id` | `network_id` |
| `ibm_dl_virtual_connection_invalid_type` | `type` |

## Example

### Valid

```hcl
resource "ibm_dl_virtual_connection" "test" {
  gateway = "example"
  name = "example-name"
  type = "classic"
}
```
