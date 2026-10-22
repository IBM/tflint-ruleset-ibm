# ibm_is_cluster_network_subnets_data

Reads ibm_is_cluster_network_subnets. Filters the collection to resources with a `name` property matching the exact specified name.

This rule validates `ibm_is_cluster_network_subnets` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `BareMetalServerDiskPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_cluster_network_subnets_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cluster_network_id` | yes | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `sort` | no | — |
| `subnets` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_cluster_network_subnets_data_invalid_cluster_network_id` | `cluster_network_id` |
| `ibm_is_cluster_network_subnets_data_invalid_name` | `name` |
| `ibm_is_cluster_network_subnets_data_invalid_sort` | `sort` |

## Example

### Valid

```hcl
data "ibm_is_cluster_network_subnets" "test" {
  cluster_network_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
