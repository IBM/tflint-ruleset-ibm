# ibm_is_cluster_network_subnet_reserved_ip

Manages ibm_is_cluster_network_subnet_reserved_ip. The name for this cluster network subnet reserved IP.

This rule validates `ibm_is_cluster_network_subnet_reserved_ip` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ClusterNetworkInterfacePrimaryIPPrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_cluster_network_subnet_reserved_ip" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `address` | no | pattern `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$` |
| `auto_delete` | no | API: Indicates whether this cluster network subnet reserved IP member will be automatically deleted when either `target` is deleted, or the cluster network subnet reserved IP is unbound. |
| `cluster_network_id` | yes | pattern `^[-0-9a-z_]+$` |
| `cluster_network_subnet_id` | yes | pattern `^[-0-9a-z_]+$` |
| `cluster_network_subnet_reserved_ip_id` | no | — |
| `created_at` | no | — |
| `etag` | no | — |
| `href` | no | — |
| `lifecycle_reasons` | no | — |
| `lifecycle_state` | no | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `owner` | no | — |
| `resource_type` | no | — |
| `target` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_cluster_network_subnet_reserved_ip_invalid_address` | `address` |
| `ibm_is_cluster_network_subnet_reserved_ip_invalid_auto_delete` | `auto_delete` |
| `ibm_is_cluster_network_subnet_reserved_ip_invalid_cluster_network_id` | `cluster_network_id` |
| `ibm_is_cluster_network_subnet_reserved_ip_invalid_cluster_network_subnet_id` | `cluster_network_subnet_id` |
| `ibm_is_cluster_network_subnet_reserved_ip_invalid_name` | `name` |

## Example

### Valid

```hcl
resource "ibm_is_cluster_network_subnet_reserved_ip" "test" {
  cluster_network_id = "example-name"
  cluster_network_subnet_id = "example-name"
}
```
