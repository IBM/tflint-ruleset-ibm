# ibm_is_private_path_service_gateway_account_policy_data

Schema validation for the ibm_is_private_path_service_gateway_account_policy data source.

This rule validates `ibm_is_private_path_service_gateway_account_policy` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_is_private_path_service_gateway_account_policy_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_policy` | no | — |
| `account` | no | — |
| `account_policy` | yes | — |
| `created_at` | no | — |
| `href` | no | — |
| `private_path_service_gateway` | yes | — |
| `resource_type` | no | — |
| `updated_at` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_private_path_service_gateway_account_policy_data_invalid_account_policy` | `account_policy` |
| `ibm_is_private_path_service_gateway_account_policy_data_invalid_private_path_service_gateway` | `private_path_service_gateway` |

## Example

### Valid

```hcl
data "ibm_is_private_path_service_gateway_account_policy" "test" {
  account_policy = "example"
  private_path_service_gateway = "example"
}
```
