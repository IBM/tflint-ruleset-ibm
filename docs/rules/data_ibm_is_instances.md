# ibm_is_instances_data

Schema validation for the ibm_is_instances data source.

This rule validates `ibm_is_instances` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `NetworkACLPrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_instances_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cluster_network_crn` | no | crn format |
| `cluster_network_id` | no | — |
| `cluster_network_name` | no | — |
| `dedicated_host` | no | — |
| `dedicated_host_name` | no | — |
| `instance_group` | no | — |
| `instance_group_name` | no | — |
| `instances` | no | — |
| `placement_group` | no | — |
| `placement_group_name` | no | — |
| `resource_group` | no | API: The resource group to use. If unspecified, the account's [default resource group](https://cloud.ibm.com/apidocs/resource-manager#introduction) will be used. |
| `vpc` | no | API: Identifies a VPC by a unique property. |
| `vpc_crn` | no | crn format |
| `vpc_name` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_instances_data_invalid_cluster_network_crn` | `cluster_network_crn` |
| `ibm_is_instances_data_invalid_cluster_network_id` | `cluster_network_id` |
| `ibm_is_instances_data_invalid_cluster_network_name` | `cluster_network_name` |
| `ibm_is_instances_data_invalid_dedicated_host` | `dedicated_host` |
| `ibm_is_instances_data_invalid_dedicated_host_name` | `dedicated_host_name` |
| `ibm_is_instances_data_invalid_instance_group` | `instance_group` |
| `ibm_is_instances_data_invalid_instance_group_name` | `instance_group_name` |
| `ibm_is_instances_data_invalid_placement_group` | `placement_group` |
| `ibm_is_instances_data_invalid_placement_group_name` | `placement_group_name` |
| `ibm_is_instances_data_invalid_resource_group` | `resource_group` |
| `ibm_is_instances_data_invalid_vpc` | `vpc` |
| `ibm_is_instances_data_invalid_vpc_crn` | `vpc_crn` |
| `ibm_is_instances_data_invalid_vpc_name` | `vpc_name` |

## Example

### Valid

```hcl
data "ibm_is_instances" "test" {
}
```
