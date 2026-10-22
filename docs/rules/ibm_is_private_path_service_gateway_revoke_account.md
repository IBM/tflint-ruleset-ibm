# ibm_is_private_path_service_gateway_revoke_account

Schema validation for the ibm_is_private_path_service_gateway_revoke_account resource.

This rule validates `ibm_is_private_path_service_gateway_revoke_account` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `PrivatePathServiceGatewayRevokeAccountOptions` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_private_path_service_gateway_revoke_account" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account` | yes | API: Identifies an account by a unique property. |
| `private_path_service_gateway` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_private_path_service_gateway_revoke_account_invalid_account` | `account` |
| `ibm_is_private_path_service_gateway_revoke_account_invalid_private_path_service_gateway` | `private_path_service_gateway` |

## Example

### Valid

```hcl
resource "ibm_is_private_path_service_gateway_revoke_account" "test" {
  account = "example"
  private_path_service_gateway = "example"
}
```
