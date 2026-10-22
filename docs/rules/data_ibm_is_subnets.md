# ibm_is_subnets_data

Schema validation for the ibm_is_subnets data source.

This rule validates `ibm_is_subnets` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `PublicGatewayPrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_subnets_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `resource_group` | no | API: The resource group to use. If unspecified, the account's [default resource group](https://cloud.ibm.com/apidocs/resource-manager#introduction) will be used. |
| `routing_table` | no | — |
| `routing_table_name` | no | — |
| `subnets` | no | — |
| `vpc` | no | API: Identifies a VPC by a unique property. |
| `vpc_crn` | no | crn format |
| `vpc_name` | no | — |
| `zone` | no | zone format |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_subnets_data_invalid_resource_group` | `resource_group` |
| `ibm_is_subnets_data_invalid_routing_table` | `routing_table` |
| `ibm_is_subnets_data_invalid_routing_table_name` | `routing_table_name` |
| `ibm_is_subnets_data_invalid_vpc` | `vpc` |
| `ibm_is_subnets_data_invalid_vpc_crn` | `vpc_crn` |
| `ibm_is_subnets_data_invalid_vpc_name` | `vpc_name` |
| `ibm_is_subnets_data_invalid_zone` | `zone` |

## Example

### Valid

```hcl
data "ibm_is_subnets" "test" {
}
```
