# ibm_is_private_path_service_gateway

Manages ibm_is_private_path_service_gateway. The name of this PPSG

This rule validates `ibm_is_private_path_service_gateway` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `PrivatePathServiceGatewayPrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_private_path_service_gateway" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_tags` | no | pattern `^([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-]):([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-])$` |
| `created_at` | no | — |
| `crn` | no | — |
| `default_access_policy` | no | one of: `deny`, `permit`, `review`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `endpoint_gateway_binding_auto_delete` | no | — |
| `endpoint_gateway_binding_auto_delete_timeout` | no | — |
| `endpoint_gateway_count` | no | — |
| `href` | no | — |
| `lifecycle_state` | no | — |
| `load_balancer` | yes | API: Identifies a load balancer by a unique property. |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `private_path_service_gateway` | no | — |
| `published` | no | — |
| `resource_group` | no | API: The resource group to use. If unspecified, the account's [default resource group](https://cloud.ibm.com/apidocs/resource-manager#introduction) will be used. |
| `resource_type` | no | — |
| `service_endpoints` | yes | API: A fully qualified domain name for this resource that permits an optional wildcard prefix |
| `tags` | no | pattern `^[A-Za-z0-9:_ .-]+$` |
| `updated_at` | no | — |
| `vpc` | no | — |
| `zonal_affinity` | no | API: Indicates whether this private path service gateway has zonal affinity. - `true`: Traffic to the service from a zone the service resides in will remain in that zone. - `false`: Traffic to the service from a zone will be load balanced across all zones in the region the service resides in. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_private_path_service_gateway_invalid_access_tags` | `access_tags` |
| `ibm_is_private_path_service_gateway_invalid_default_access_policy` | `default_access_policy` |
| `ibm_is_private_path_service_gateway_invalid_load_balancer` | `load_balancer` |
| `ibm_is_private_path_service_gateway_invalid_name` | `name` |
| `ibm_is_private_path_service_gateway_invalid_resource_group` | `resource_group` |
| `ibm_is_private_path_service_gateway_invalid_service_endpoints` | `service_endpoints` |
| `ibm_is_private_path_service_gateway_invalid_tags` | `tags` |
| `ibm_is_private_path_service_gateway_invalid_zonal_affinity` | `zonal_affinity` |

## Example

### Valid

```hcl
resource "ibm_is_private_path_service_gateway" "test" {
  load_balancer = "example"
  service_endpoints = "aaaaaaa"
}
```
