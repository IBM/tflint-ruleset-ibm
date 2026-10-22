# ibm_is_virtual_endpoint_gateway

Manages ibm_is_virtual_endpoint_gateway. Endpoint gateway name

This rule validates `ibm_is_virtual_endpoint_gateway` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `EndpointGatewayPrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_virtual_endpoint_gateway" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_tags` | no | pattern `^([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-]):([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-])$` |
| `allow_dns_resolution_binding` | no | — |
| `created_at` | no | — |
| `crn` | no | — |
| `dns_resolution_binding_mode` | no | one of: `disabled`, `per_resource_binding`, `primary`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `health_state` | no | — |
| `lifecycle_reasons` | no | — |
| `lifecycle_state` | no | — |
| `name` | yes | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `resource_group` | no | API: The resource group to use. If unspecified, the account's [default resource group](https://cloud.ibm.com/apidocs/resource-manager#introduction) will be used. |
| `resource_type` | no | — |
| `security_groups` | no | API: Identifies a security group by a unique property. |
| `service_endpoints` | no | — |
| `tags` | no | pattern `^[A-Za-z0-9:_ .-]+$` |
| `vpc` | yes | API: Identifies a VPC by a unique property. |
| `ips` (block) | no | nested block |
| `ips.id` | no | pattern `^[-0-9a-z_]+$` |
| `ips.name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `ips.subnet` | no | API: Identifies a subnet by a unique property. |
| `target` (block) | yes | nested block |
| `target.crn` | no | crn format |
| `target.name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `target.resource_type` | yes | one of: `private_path_service_gateway`, `provider_cloud_service`, `provider_infrastructure_service`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_virtual_endpoint_gateway_invalid_access_tags` | `access_tags` |
| `ibm_is_virtual_endpoint_gateway_invalid_allow_dns_resolution_binding` | `allow_dns_resolution_binding` |
| `ibm_is_virtual_endpoint_gateway_invalid_dns_resolution_binding_mode` | `dns_resolution_binding_mode` |
| `ibm_is_virtual_endpoint_gateway_invalid_name` | `name` |
| `ibm_is_virtual_endpoint_gateway_invalid_resource_group` | `resource_group` |
| `ibm_is_virtual_endpoint_gateway_invalid_security_groups` | `security_groups` |
| `ibm_is_virtual_endpoint_gateway_invalid_tags` | `tags` |
| `ibm_is_virtual_endpoint_gateway_invalid_vpc` | `vpc` |
| `ibm_is_virtual_endpoint_gateway_invalid_ips__id` | `ips.id` |
| `ibm_is_virtual_endpoint_gateway_invalid_ips__name` | `ips.name` |
| `ibm_is_virtual_endpoint_gateway_invalid_ips__subnet` | `ips.subnet` |
| `ibm_is_virtual_endpoint_gateway_invalid_target__crn` | `target.crn` |
| `ibm_is_virtual_endpoint_gateway_invalid_target__name` | `target.name` |
| `ibm_is_virtual_endpoint_gateway_invalid_target__resource_type` | `target.resource_type` |

## Example

### Valid

```hcl
resource "ibm_is_virtual_endpoint_gateway" "test" {
  name = "example-name"
  vpc = "example"
  target {
    resource_type = "private_path_service_gateway"
  }
}
```
