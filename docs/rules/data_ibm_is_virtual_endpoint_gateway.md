# ibm_is_virtual_endpoint_gateway_data

Reads ibm_is_virtual_endpoint_gateway. Endpoint gateway name

This rule validates `ibm_is_virtual_endpoint_gateway` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ClusterNetworkInterfacePrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_virtual_endpoint_gateway_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_tags` | no | — |
| `allow_dns_resolution_binding` | no | — |
| `created_at` | no | — |
| `crn` | no | — |
| `dns_resolution_binding_mode` | no | — |
| `health_state` | no | — |
| `ips` | no | — |
| `lifecycle_reasons` | no | — |
| `lifecycle_state` | no | — |
| `name` | yes | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `resource_group` | no | — |
| `resource_type` | no | — |
| `security_groups` | no | — |
| `service_endpoints` | no | — |
| `tags` | no | — |
| `target` | no | — |
| `vpc` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_virtual_endpoint_gateway_data_invalid_name` | `name` |

## Example

### Valid

```hcl
data "ibm_is_virtual_endpoint_gateway" "test" {
  name = "example-name"
}
```
