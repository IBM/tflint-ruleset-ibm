# ibm_is_vpn_gateway_advertised_cidr

Schema validation for the ibm_is_vpn_gateway_advertised_cidr resource.

This rule validates `ibm_is_vpn_gateway_advertised_cidr` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ClusterNetworkSubnetPrefixPrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_vpn_gateway_advertised_cidr" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cidr` | yes | cidr format |
| `vpn_gateway` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_vpn_gateway_advertised_cidr_invalid_cidr` | `cidr` |
| `ibm_is_vpn_gateway_advertised_cidr_invalid_vpn_gateway` | `vpn_gateway` |

## Example

### Valid

```hcl
resource "ibm_is_vpn_gateway_advertised_cidr" "test" {
  cidr = "10.0.0.0/24"
  vpn_gateway = "example"
}
```
