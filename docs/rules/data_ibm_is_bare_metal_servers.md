# ibm_is_bare_metal_servers_data

Reads ibm_is_bare_metal_servers. The name of the bare metal server

This rule validates `ibm_is_bare_metal_servers` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `NetworkACLPrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_bare_metal_servers_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `network_interfaces_subnet` | no | — |
| `network_interfaces_subnet_crn` | no | crn format |
| `network_interfaces_subnet_name` | no | — |
| `resource_group` | no | API: The resource group to use. If unspecified, the account's [default resource group](https://cloud.ibm.com/apidocs/resource-manager#introduction) will be used. |
| `servers` | no | — |
| `vpc` | no | API: Identifies a VPC by a unique property. |
| `vpc_crn` | no | crn format |
| `vpc_name` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_bare_metal_servers_data_invalid_name` | `name` |
| `ibm_is_bare_metal_servers_data_invalid_network_interfaces_subnet` | `network_interfaces_subnet` |
| `ibm_is_bare_metal_servers_data_invalid_network_interfaces_subnet_crn` | `network_interfaces_subnet_crn` |
| `ibm_is_bare_metal_servers_data_invalid_network_interfaces_subnet_name` | `network_interfaces_subnet_name` |
| `ibm_is_bare_metal_servers_data_invalid_resource_group` | `resource_group` |
| `ibm_is_bare_metal_servers_data_invalid_vpc` | `vpc` |
| `ibm_is_bare_metal_servers_data_invalid_vpc_crn` | `vpc_crn` |
| `ibm_is_bare_metal_servers_data_invalid_vpc_name` | `vpc_name` |

## Example

### Valid

```hcl
data "ibm_is_bare_metal_servers" "test" {
}
```
