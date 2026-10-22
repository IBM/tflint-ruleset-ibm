# ibm_is_subnet_public_gateway_attachment

Manages ibm_is_subnet_public_gateway_attachment. Name of the Public gateway instance

This rule validates `ibm_is_subnet_public_gateway_attachment` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ClusterNetworkInterfacePrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_subnet_public_gateway_attachment" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `crn` | no | — |
| `floating_ip` | no | — |
| `name` | no | — |
| `public_gateway` | yes | — |
| `resource_group` | no | — |
| `resource_group_name` | no | — |
| `resource_type` | no | — |
| `status` | no | — |
| `subnet` | yes | API: Identifies a cluster network subnet by a unique property. |
| `vpc` | no | — |
| `zone` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_subnet_public_gateway_attachment_invalid_public_gateway` | `public_gateway` |
| `ibm_is_subnet_public_gateway_attachment_invalid_subnet` | `subnet` |

## Example

### Valid

```hcl
resource "ibm_is_subnet_public_gateway_attachment" "test" {
  public_gateway = "example"
  subnet = "example"
}
```
