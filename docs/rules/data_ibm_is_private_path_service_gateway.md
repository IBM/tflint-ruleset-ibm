# ibm_is_private_path_service_gateway_data

Reads ibm_is_private_path_service_gateway. The name for this private path service gateway.

This rule validates `ibm_is_private_path_service_gateway` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_is_private_path_service_gateway_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `created_at` | no | — |
| `crn` | no | — |
| `default_access_policy` | no | — |
| `endpoint_gateway_binding_auto_delete` | no | — |
| `endpoint_gateway_binding_auto_delete_timeout` | no | — |
| `endpoint_gateway_count` | no | — |
| `href` | no | — |
| `lifecycle_state` | no | — |
| `load_balancer` | no | — |
| `name` | no | — |
| `private_path_service_gateway` | no | — |
| `private_path_service_gateway_name` | no | — |
| `published` | no | — |
| `region` | no | — |
| `resource_group` | no | — |
| `resource_type` | no | — |
| `service_endpoints` | no | — |
| `vpc` | no | — |
| `zonal_affinity` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_private_path_service_gateway_data_invalid_private_path_service_gateway` | `private_path_service_gateway` |
| `ibm_is_private_path_service_gateway_data_invalid_private_path_service_gateway_name` | `private_path_service_gateway_name` |

## Example

### Valid

```hcl
data "ibm_is_private_path_service_gateway" "test" {
}
```
