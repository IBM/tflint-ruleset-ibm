# ibm_is_cluster_network_interface

Manages ibm_is_cluster_network_interface. The name for this cluster network interface.

This rule validates `ibm_is_cluster_network_interface` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ClusterNetworkInterfacePrimaryIPPrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_cluster_network_interface" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `allow_ip_spoofing` | no | — |
| `auto_delete` | no | — |
| `cluster_network_id` | yes | pattern `^[-0-9a-z_]+$` |
| `cluster_network_interface_id` | no | — |
| `created_at` | no | — |
| `enable_infrastructure_nat` | no | — |
| `etag` | no | — |
| `href` | no | — |
| `lifecycle_reasons` | no | — |
| `lifecycle_state` | no | — |
| `mac_address` | no | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `resource_type` | no | — |
| `target` | no | — |
| `vpc` | no | — |
| `zone` | no | — |
| `primary_ip` (block) | no | nested block |
| `primary_ip.address` | no | pattern `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$` |
| `primary_ip.auto_delete` | no | API: Indicates whether this cluster network subnet reserved IP member will be automatically deleted when either `target` is deleted, or the cluster network subnet reserved IP is unbound. |
| `primary_ip.href` | no | pattern `^http(s)?:\/\/([^\/?#]*)([^?#]*)(\?([^#]*))?(#(.*))?$` |
| `primary_ip.id` | no | pattern `^[-0-9a-z_]+$` |
| `primary_ip.name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `subnet` (block) | no | nested block |
| `subnet.href` | no | pattern `^http(s)?:\/\/([^\/?#]*)([^?#]*)(\?([^#]*))?(#(.*))?$` |
| `subnet.id` | no | pattern `^[-0-9a-z_]+$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_cluster_network_interface_invalid_cluster_network_id` | `cluster_network_id` |
| `ibm_is_cluster_network_interface_invalid_name` | `name` |
| `ibm_is_cluster_network_interface_invalid_primary_ip__address` | `primary_ip.address` |
| `ibm_is_cluster_network_interface_invalid_primary_ip__auto_delete` | `primary_ip.auto_delete` |
| `ibm_is_cluster_network_interface_invalid_primary_ip__href` | `primary_ip.href` |
| `ibm_is_cluster_network_interface_invalid_primary_ip__id` | `primary_ip.id` |
| `ibm_is_cluster_network_interface_invalid_primary_ip__name` | `primary_ip.name` |
| `ibm_is_cluster_network_interface_invalid_subnet__href` | `subnet.href` |
| `ibm_is_cluster_network_interface_invalid_subnet__id` | `subnet.id` |

## Example

### Valid

```hcl
resource "ibm_is_cluster_network_interface" "test" {
  cluster_network_id = "example-name"
}
```
