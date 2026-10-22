# ibm_is_public_gateway

Manages ibm_is_public_gateway. Name of the Public gateway instance

This rule validates `ibm_is_public_gateway` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `PublicGatewayPrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_public_gateway" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_tags` | no | pattern `^([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-]):([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-])$` |
| `crn` | no | — |
| `floating_ip` | no | API: Identifies a floating IP by a unique property. |
| `name` | yes | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `resource_controller_url` | no | — |
| `resource_crn` | no | — |
| `resource_group` | no | API: The resource group to use. If unspecified, the account's [default resource group](https://cloud.ibm.com/apidocs/resource-manager#introduction) will be used. |
| `resource_group_name` | no | — |
| `resource_name` | no | — |
| `resource_status` | no | — |
| `status` | no | — |
| `tags` | no | pattern `^[A-Za-z0-9:_ .-]+$` |
| `vpc` | yes | API: Identifies a VPC by a unique property. |
| `zone` | yes | zone format |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_public_gateway_invalid_access_tags` | `access_tags` |
| `ibm_is_public_gateway_invalid_floating_ip` | `floating_ip` |
| `ibm_is_public_gateway_invalid_name` | `name` |
| `ibm_is_public_gateway_invalid_resource_group` | `resource_group` |
| `ibm_is_public_gateway_invalid_tags` | `tags` |
| `ibm_is_public_gateway_invalid_vpc` | `vpc` |
| `ibm_is_public_gateway_invalid_zone` | `zone` |

## Example

### Valid

```hcl
resource "ibm_is_public_gateway" "test" {
  name = "example-name"
  vpc = "example"
  zone = "us-south-1"
}
```
