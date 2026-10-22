# ibm_is_cluster_network_subnet_reserved_ips_data

Reads ibm_is_cluster_network_subnet_reserved_ips. Filters the collection to resources with a `name` property matching the exact specified name.

This rule validates `ibm_is_cluster_network_subnet_reserved_ips` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `BareMetalServerDiskPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_cluster_network_subnet_reserved_ips_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cluster_network_id` | yes | — |
| `cluster_network_subnet_id` | yes | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `reserved_ips` | no | — |
| `sort` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_cluster_network_subnet_reserved_ips_data_invalid_cluster_network_id` | `cluster_network_id` |
| `ibm_is_cluster_network_subnet_reserved_ips_data_invalid_cluster_network_subnet_id` | `cluster_network_subnet_id` |
| `ibm_is_cluster_network_subnet_reserved_ips_data_invalid_name` | `name` |
| `ibm_is_cluster_network_subnet_reserved_ips_data_invalid_sort` | `sort` |

## Example

### Valid

```hcl
data "ibm_is_cluster_network_subnet_reserved_ips" "test" {
  cluster_network_id = "01234567-89ab-cdef-0123-456789abcdef"
  cluster_network_subnet_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
