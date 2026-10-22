# ibm_is_virtual_endpoint_gateway_resource_binding_data

Reads ibm_is_virtual_endpoint_gateway_resource_binding. The name for this resource binding.

This rule validates `ibm_is_virtual_endpoint_gateway_resource_binding` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_is_virtual_endpoint_gateway_resource_binding_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `created_at` | no | — |
| `endpoint_gateway_id` | yes | pattern `^[-0-9a-z_]+$` |
| `endpoint_gateway_resource_binding_id` | yes | — |
| `href` | no | — |
| `lifecycle_reasons` | no | — |
| `lifecycle_state` | no | — |
| `name` | no | — |
| `resource_type` | no | — |
| `service_endpoint` | no | — |
| `target` | no | — |
| `type` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_virtual_endpoint_gateway_resource_binding_data_invalid_endpoint_gateway_id` | `endpoint_gateway_id` |
| `ibm_is_virtual_endpoint_gateway_resource_binding_data_invalid_endpoint_gateway_resource_binding_id` | `endpoint_gateway_resource_binding_id` |

## Example

### Valid

```hcl
data "ibm_is_virtual_endpoint_gateway_resource_binding" "test" {
  endpoint_gateway_id = "example-name"
  endpoint_gateway_resource_binding_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
