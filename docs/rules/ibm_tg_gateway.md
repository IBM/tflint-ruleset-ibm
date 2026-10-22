# ibm_tg_gateway

Manages ibm_tg_gateway. Name Transit Gateway Services

This rule validates `ibm_tg_gateway` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `TGCreateTemplate` (`transit-gateway`).

## Configuration

```hcl
rule "ibm_tg_gateway" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `created_at` | no | — |
| `crn` | no | — |
| `global` | no | API: Allow global routing for a Transit Gateway. If unspecified, the default value is false. |
| `gre_enhanced_route_propagation` | no | API: Allow route propagation across all GREs connected to the same transit gateway. This affects connections on the gateway of type `redundant_gre`, `unbound_gre_tunnel` and `gre_tunnel`. |
| `location` | yes | pattern `^([a-z]|[a-z][-_a-z0-9]*[a-z0-9]+)$` |
| `name` | yes | pattern `^([a-zA-Z]|[a-zA-Z][-_a-zA-Z0-9]*[a-zA-Z0-9])$` |
| `resource_controller_url` | no | — |
| `resource_crn` | no | — |
| `resource_group` | no | API: The resource group to use. If unspecified, the account's [default resource group](https://console.bluemix.net/apidocs/resource-manager#introduction) is used. |
| `resource_group_name` | no | — |
| `resource_name` | no | — |
| `resource_status` | no | — |
| `status` | no | — |
| `tags` | no | pattern `^[A-Za-z0-9:_ .-]+$` |
| `updated_at` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_tg_gateway_invalid_global` | `global` |
| `ibm_tg_gateway_invalid_gre_enhanced_route_propagation` | `gre_enhanced_route_propagation` |
| `ibm_tg_gateway_invalid_location` | `location` |
| `ibm_tg_gateway_invalid_name` | `name` |
| `ibm_tg_gateway_invalid_resource_group` | `resource_group` |
| `ibm_tg_gateway_invalid_tags` | `tags` |

## Example

### Valid

```hcl
resource "ibm_tg_gateway" "test" {
  location = "example-name"
  name = "example-name"
}
```
