# ibm_is_vpc_address_prefixes_data

Reads ibm_is_vpc_address_prefixes. The user-defined name for this address prefix.

This rule validates `ibm_is_vpc_address_prefixes` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `VPCDNSResolutionBindingPrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_vpc_address_prefixes_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `address_prefixes` | no | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `vpc` | yes | API: The VPC to bind this VPC to for DNS resolution. The VPC must be different from the VPC specified in the URL, must have `dns.enable_hub` set to `true`, and may be in a different account (subject to IAM policies). Additionally, the VPC specified in the URL (this VPC) must have `dns.enable_hub` set to `false` and a `dns.resolution_binding_count` of zero. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_vpc_address_prefixes_data_invalid_name` | `name` |
| `ibm_is_vpc_address_prefixes_data_invalid_vpc` | `vpc` |

## Example

### Valid

```hcl
data "ibm_is_vpc_address_prefixes" "test" {
  vpc = "example"
}
```
