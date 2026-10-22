# ibm_is_private_path_service_gateway_endpoint_gateway_binding_data

Schema validation for the ibm_is_private_path_service_gateway_endpoint_gateway_binding data source.

This rule validates `ibm_is_private_path_service_gateway_endpoint_gateway_binding` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_is_private_path_service_gateway_endpoint_gateway_binding_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account` | no | — |
| `created_at` | no | — |
| `endpoint_gateway_binding` | yes | — |
| `expiration_at` | no | — |
| `href` | no | — |
| `lifecycle_state` | no | — |
| `private_path_service_gateway` | yes | — |
| `resource_type` | no | — |
| `status` | no | — |
| `updated_at` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_private_path_service_gateway_endpoint_gateway_binding_data_invalid_endpoint_gateway_binding` | `endpoint_gateway_binding` |
| `ibm_is_private_path_service_gateway_endpoint_gateway_binding_data_invalid_private_path_service_gateway` | `private_path_service_gateway` |

## Example

### Valid

```hcl
data "ibm_is_private_path_service_gateway_endpoint_gateway_binding" "test" {
  endpoint_gateway_binding = "example"
  private_path_service_gateway = "example"
}
```
