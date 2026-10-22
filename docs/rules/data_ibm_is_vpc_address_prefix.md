# ibm_is_vpc_address_prefix_data

Reads ibm_is_vpc_address_prefix. The user-defined name for this address prefix.

This rule validates `ibm_is_vpc_address_prefix` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `VPCDNSResolutionBindingPrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_vpc_address_prefix_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `address_prefix` | no | — |
| `address_prefix_name` | no | — |
| `cidr` | no | — |
| `created_at` | no | — |
| `has_subnets` | no | — |
| `href` | no | — |
| `is_default` | no | — |
| `name` | no | — |
| `vpc` | no | API: The VPC to bind this VPC to for DNS resolution. The VPC must be different from the VPC specified in the URL, must have `dns.enable_hub` set to `true`, and may be in a different account (subject to IAM policies). Additionally, the VPC specified in the URL (this VPC) must have `dns.enable_hub` set to `false` and a `dns.resolution_binding_count` of zero. |
| `vpc_name` | no | — |
| `zone` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_vpc_address_prefix_data_invalid_address_prefix` | `address_prefix` |
| `ibm_is_vpc_address_prefix_data_invalid_address_prefix_name` | `address_prefix_name` |
| `ibm_is_vpc_address_prefix_data_invalid_vpc` | `vpc` |
| `ibm_is_vpc_address_prefix_data_invalid_vpc_name` | `vpc_name` |

## Example

### Valid

```hcl
data "ibm_is_vpc_address_prefix" "test" {
}
```
