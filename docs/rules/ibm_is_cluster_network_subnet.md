# ibm_is_cluster_network_subnet

Manages ibm_is_cluster_network_subnet. The name for this cluster network subnet.

This rule validates `ibm_is_cluster_network_subnet` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Subnet` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_cluster_network_subnet" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `available_ipv4_address_count` | no | — |
| `cluster_network_id` | yes | pattern `^[-0-9a-z_]+$` |
| `cluster_network_subnet_id` | no | — |
| `created_at` | no | — |
| `etag` | no | — |
| `href` | no | — |
| `ip_version` | no | one of: `ipv4`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `ipv4_cidr_block` | no | pattern `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])(\/(3[0-2]|[1-2][0-9]|[0-9]))$` |
| `lifecycle_reasons` | no | — |
| `lifecycle_state` | no | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `resource_type` | no | — |
| `total_ipv4_address_count` | no | API: The total number of IPv4 addresses in this subnet. Note: This is calculated as 2&lt;sup&gt;(32 - prefix length)&lt;/sup&gt;. For example, the prefix length `/24` gives:&lt;br&gt; 2&lt;sup&gt;(32 - 24)&lt;/sup&gt; = 2&lt;sup&gt;8&lt;/sup&gt; = 256 addresses. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_cluster_network_subnet_invalid_cluster_network_id` | `cluster_network_id` |
| `ibm_is_cluster_network_subnet_invalid_ip_version` | `ip_version` |
| `ibm_is_cluster_network_subnet_invalid_ipv4_cidr_block` | `ipv4_cidr_block` |
| `ibm_is_cluster_network_subnet_invalid_name` | `name` |
| `ibm_is_cluster_network_subnet_invalid_total_ipv4_address_count` | `total_ipv4_address_count` |

## Example

### Valid

```hcl
resource "ibm_is_cluster_network_subnet" "test" {
  cluster_network_id = "example-name"
}
```
