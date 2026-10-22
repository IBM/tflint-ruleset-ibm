# ibm_is_cluster_networks_data

Reads ibm_is_cluster_networks. Filters the collection to resources with a `name` property matching the exact specified name.

This rule validates `ibm_is_cluster_networks` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `BareMetalServerDiskPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_cluster_networks_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cluster_networks` | no | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `resource_group_id` | no | — |
| `sort` | no | — |
| `vpc_crn` | no | crn format |
| `vpc_id` | no | — |
| `vpc_name` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_cluster_networks_data_invalid_name` | `name` |
| `ibm_is_cluster_networks_data_invalid_resource_group_id` | `resource_group_id` |
| `ibm_is_cluster_networks_data_invalid_sort` | `sort` |
| `ibm_is_cluster_networks_data_invalid_vpc_crn` | `vpc_crn` |
| `ibm_is_cluster_networks_data_invalid_vpc_id` | `vpc_id` |
| `ibm_is_cluster_networks_data_invalid_vpc_name` | `vpc_name` |

## Example

### Valid

```hcl
data "ibm_is_cluster_networks" "test" {
}
```
