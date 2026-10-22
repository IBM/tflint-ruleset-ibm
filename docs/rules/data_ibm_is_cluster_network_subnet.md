# ibm_is_cluster_network_subnet_data

Reads ibm_is_cluster_network_subnet. The name for this cluster network subnet.

This rule validates `ibm_is_cluster_network_subnet` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_is_cluster_network_subnet_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `available_ipv4_address_count` | no | — |
| `cluster_network_id` | yes | pattern `^[-0-9a-z_]+$` |
| `cluster_network_subnet_id` | yes | — |
| `created_at` | no | — |
| `href` | no | — |
| `ip_version` | no | — |
| `ipv4_cidr_block` | no | — |
| `lifecycle_reasons` | no | — |
| `lifecycle_state` | no | — |
| `name` | no | — |
| `resource_type` | no | — |
| `total_ipv4_address_count` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_cluster_network_subnet_data_invalid_cluster_network_id` | `cluster_network_id` |
| `ibm_is_cluster_network_subnet_data_invalid_cluster_network_subnet_id` | `cluster_network_subnet_id` |

## Example

### Valid

```hcl
data "ibm_is_cluster_network_subnet" "test" {
  cluster_network_id = "example-name"
  cluster_network_subnet_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
