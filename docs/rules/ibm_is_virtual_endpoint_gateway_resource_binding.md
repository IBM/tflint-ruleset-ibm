# ibm_is_virtual_endpoint_gateway_resource_binding

Manages ibm_is_virtual_endpoint_gateway_resource_binding. The name for this resource binding.

This rule validates `ibm_is_virtual_endpoint_gateway_resource_binding` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `CloudObjectStorageBucketIdentity` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_virtual_endpoint_gateway_resource_binding" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `created_at` | no | — |
| `endpoint_gateway_id` | yes | pattern `^[-0-9a-z_]+$` |
| `endpoint_gateway_resource_binding_id` | no | — |
| `href` | no | — |
| `lifecycle_reasons` | no | — |
| `lifecycle_state` | no | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `resource_type` | no | — |
| `service_endpoint` | no | — |
| `type` | no | — |
| `target` (block) | yes | nested block |
| `target.crn` | no | crn format |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_virtual_endpoint_gateway_resource_binding_invalid_endpoint_gateway_id` | `endpoint_gateway_id` |
| `ibm_is_virtual_endpoint_gateway_resource_binding_invalid_name` | `name` |
| `ibm_is_virtual_endpoint_gateway_resource_binding_invalid_target__crn` | `target.crn` |

## Example

### Valid

```hcl
resource "ibm_is_virtual_endpoint_gateway_resource_binding" "test" {
  endpoint_gateway_id = "example-name"
  target {
  }
}
```
