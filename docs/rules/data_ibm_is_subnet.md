# ibm_is_subnet_data

Schema validation for the ibm_is_subnet data source.

This rule validates `ibm_is_subnet` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `VPCDNSResolutionBindingPrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_subnet_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_tags` | no | — |
| `available_ipv4_address_count` | no | — |
| `crn` | no | — |
| `identifier` | no | — |
| `ipv4_cidr_block` | no | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `network_acl` | no | — |
| `public_gateway` | no | — |
| `resource_controller_url` | no | — |
| `resource_crn` | no | — |
| `resource_group` | no | — |
| `resource_group_name` | no | — |
| `resource_name` | no | — |
| `resource_status` | no | — |
| `routing_table` | no | — |
| `status` | no | — |
| `tags` | no | — |
| `total_ipv4_address_count` | no | — |
| `vpc` | no | API: The VPC to bind this VPC to for DNS resolution. The VPC must be different from the VPC specified in the URL, must have `dns.enable_hub` set to `true`, and may be in a different account (subject to IAM policies). Additionally, the VPC specified in the URL (this VPC) must have `dns.enable_hub` set to `false` and a `dns.resolution_binding_count` of zero. |
| `vpc_name` | no | — |
| `zone` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_subnet_data_invalid_identifier` | `identifier` |
| `ibm_is_subnet_data_invalid_name` | `name` |
| `ibm_is_subnet_data_invalid_vpc` | `vpc` |

## Example

### Valid

```hcl
data "ibm_is_subnet" "test" {
}
```
