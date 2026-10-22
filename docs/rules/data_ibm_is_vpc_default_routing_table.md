# ibm_is_vpc_default_routing_table_data

Reads ibm_is_vpc_default_routing_table. Default Routing table Name

This rule validates `ibm_is_vpc_default_routing_table` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ShareMountTargetByAccessControlModeVPC` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_vpc_default_routing_table_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_tags` | no | — |
| `created_at` | no | — |
| `crn` | no | — |
| `default_routing_table` | no | — |
| `href` | no | — |
| `is_default` | no | — |
| `lifecycle_state` | no | — |
| `name` | no | — |
| `resource_group` | no | — |
| `resource_type` | no | — |
| `route_direct_link_ingress` | no | — |
| `route_internet_ingress` | no | — |
| `route_transit_gateway_ingress` | no | — |
| `route_vpc_zone_ingress` | no | — |
| `routes` | no | — |
| `subnets` | no | — |
| `tags` | no | — |
| `vpc` | yes | API: Identifies a VPC by a unique property. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_vpc_default_routing_table_data_invalid_vpc` | `vpc` |

## Example

### Valid

```hcl
data "ibm_is_vpc_default_routing_table" "test" {
  vpc = "example"
}
```
