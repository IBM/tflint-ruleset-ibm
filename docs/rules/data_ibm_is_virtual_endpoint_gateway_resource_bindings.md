# ibm_is_virtual_endpoint_gateway_resource_bindings_data

Schema validation for the ibm_is_virtual_endpoint_gateway_resource_bindings data source.

This rule validates `ibm_is_virtual_endpoint_gateway_resource_bindings` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_is_virtual_endpoint_gateway_resource_bindings_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `endpoint_gateway_id` | yes | — |
| `resource_bindings` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_virtual_endpoint_gateway_resource_bindings_data_invalid_endpoint_gateway_id` | `endpoint_gateway_id` |

## Example

### Valid

```hcl
data "ibm_is_virtual_endpoint_gateway_resource_bindings" "test" {
  endpoint_gateway_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
