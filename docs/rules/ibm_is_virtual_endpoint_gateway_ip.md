# ibm_is_virtual_endpoint_gateway_ip

Manages ibm_is_virtual_endpoint_gateway_ip. Endpoint gateway IP name

This rule validates `ibm_is_virtual_endpoint_gateway_ip` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_is_virtual_endpoint_gateway_ip" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `address` | no | — |
| `auto_delete` | no | — |
| `created_at` | no | — |
| `gateway` | yes | — |
| `name` | no | — |
| `reserved_ip` | yes | — |
| `resource_type` | no | — |
| `target` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_virtual_endpoint_gateway_ip_invalid_gateway` | `gateway` |
| `ibm_is_virtual_endpoint_gateway_ip_invalid_reserved_ip` | `reserved_ip` |

## Example

### Valid

```hcl
resource "ibm_is_virtual_endpoint_gateway_ip" "test" {
  gateway = "example"
  reserved_ip = "10.0.0.1"
}
```
