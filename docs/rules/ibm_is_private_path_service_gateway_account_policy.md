# ibm_is_private_path_service_gateway_account_policy

Schema validation for the ibm_is_private_path_service_gateway_account_policy resource.

This rule validates `ibm_is_private_path_service_gateway_account_policy` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `PrivatePathServiceGatewayAccountPolicyPrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_private_path_service_gateway_account_policy" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_policy` | yes | one of: `deny`, `permit`, `review`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `account` | yes | API: Identifies an account by a unique property. |
| `account_policy` | no | — |
| `created_at` | no | — |
| `href` | no | — |
| `private_path_service_gateway` | yes | — |
| `resource_type` | no | — |
| `updated_at` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_private_path_service_gateway_account_policy_invalid_access_policy` | `access_policy` |
| `ibm_is_private_path_service_gateway_account_policy_invalid_account` | `account` |
| `ibm_is_private_path_service_gateway_account_policy_invalid_private_path_service_gateway` | `private_path_service_gateway` |

## Example

### Valid

```hcl
resource "ibm_is_private_path_service_gateway_account_policy" "test" {
  access_policy = "deny"
  account = "example"
  private_path_service_gateway = "example"
}
```
