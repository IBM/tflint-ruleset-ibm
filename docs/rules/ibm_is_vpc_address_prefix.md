# ibm_is_vpc_address_prefix

Manages ibm_is_vpc_address_prefix. Name

This rule validates `ibm_is_vpc_address_prefix` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `AddressPrefixPrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_vpc_address_prefix" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `address_prefix` | no | — |
| `cidr` | yes | pattern `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])(\/(3[0-2]|[1-2][0-9]|[0-9]))$` |
| `has_subnets` | no | — |
| `is_default` | no | API: Indicates whether this will be the default address prefix for this zone in this VPC. If `true`, the VPC must not have a default address prefix for this zone. |
| `name` | yes | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `related_crn` | no | — |
| `vpc` | yes | — |
| `zone` | yes | zone format |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_vpc_address_prefix_invalid_cidr` | `cidr` |
| `ibm_is_vpc_address_prefix_invalid_is_default` | `is_default` |
| `ibm_is_vpc_address_prefix_invalid_name` | `name` |
| `ibm_is_vpc_address_prefix_invalid_vpc` | `vpc` |
| `ibm_is_vpc_address_prefix_invalid_zone` | `zone` |

## Example

### Valid

```hcl
resource "ibm_is_vpc_address_prefix" "test" {
  cidr = "10.0.0.0/24"
  name = "example-name"
  vpc = "example"
  zone = "us-south-1"
}
```
