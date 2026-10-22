# ibm_is_private_path_service_gateway_account_policies_data

Schema validation for the ibm_is_private_path_service_gateway_account_policies data source.

This rule validates `ibm_is_private_path_service_gateway_account_policies` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `PrivatePathServiceGatewayRevokeAccountOptions` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_private_path_service_gateway_account_policies_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account` | no | API: Identifies an account by a unique property. |
| `account_policies` | no | — |
| `private_path_service_gateway` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_private_path_service_gateway_account_policies_data_invalid_account` | `account` |
| `ibm_is_private_path_service_gateway_account_policies_data_invalid_private_path_service_gateway` | `private_path_service_gateway` |

## Example

### Valid

```hcl
data "ibm_is_private_path_service_gateway_account_policies" "test" {
  private_path_service_gateway = "example"
}
```
