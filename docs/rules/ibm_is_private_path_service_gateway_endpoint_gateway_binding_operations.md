# ibm_is_private_path_service_gateway_endpoint_gateway_binding_operations

Schema validation for the ibm_is_private_path_service_gateway_endpoint_gateway_binding_operations resource.

This rule validates `ibm_is_private_path_service_gateway_endpoint_gateway_binding_operations` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `PrivatePathServiceGatewayAccountPolicyPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_private_path_service_gateway_endpoint_gateway_binding_operations" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_policy` | yes | one of: `deny`, `permit`, `review`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `endpoint_gateway_binding` | yes | — |
| `private_path_service_gateway` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_private_path_service_gateway_endpoint_gateway_binding_operations_invalid_access_policy` | `access_policy` |
| `ibm_is_private_path_service_gateway_endpoint_gateway_binding_operations_invalid_endpoint_gateway_binding` | `endpoint_gateway_binding` |
| `ibm_is_private_path_service_gateway_endpoint_gateway_binding_operations_invalid_private_path_service_gateway` | `private_path_service_gateway` |

## Example

### Valid

```hcl
resource "ibm_is_private_path_service_gateway_endpoint_gateway_binding_operations" "test" {
  access_policy = "deny"
  endpoint_gateway_binding = "example"
  private_path_service_gateway = "example"
}
```
