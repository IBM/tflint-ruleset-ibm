# ibm_is_private_path_service_gateway_operations

Schema validation for the ibm_is_private_path_service_gateway_operations resource.

This rule validates `ibm_is_private_path_service_gateway_operations` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_is_private_path_service_gateway_operations" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `private_path_service_gateway` | yes | — |
| `published` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_private_path_service_gateway_operations_invalid_private_path_service_gateway` | `private_path_service_gateway` |
| `ibm_is_private_path_service_gateway_operations_invalid_published` | `published` |

## Example

### Valid

```hcl
resource "ibm_is_private_path_service_gateway_operations" "test" {
  private_path_service_gateway = "example"
  published = true
}
```
