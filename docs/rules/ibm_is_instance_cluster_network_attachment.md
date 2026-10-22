# ibm_is_instance_cluster_network_attachment

Manages ibm_is_instance_cluster_network_attachment. The name for this instance cluster network attachment.

This rule validates `ibm_is_instance_cluster_network_attachment` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ClusterNetworkInterfacePrimaryIPPrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_instance_cluster_network_attachment" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `href` | no | — |
| `instance_cluster_network_attachment_id` | no | — |
| `instance_id` | yes | pattern `^[-0-9a-z_]+$` |
| `lifecycle_reasons` | no | — |
| `lifecycle_state` | no | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `resource_type` | no | — |
| `before` (block) | no | nested block |
| `before.id` | yes | pattern `^[-0-9a-z_]+$` |
| `cluster_network_interface` (block) | yes | nested block |
| `cluster_network_interface.name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `cluster_network_interface.primary_ip` (block) | no | nested block |
| `cluster_network_interface.primary_ip.address` | yes | pattern `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$` |
| `cluster_network_interface.primary_ip.href` | yes | pattern `^http(s)?:\/\/([^\/?#]*)([^?#]*)(\?([^#]*))?(#(.*))?$` |
| `cluster_network_interface.primary_ip.id` | yes | pattern `^[-0-9a-z_]+$` |
| `cluster_network_interface.primary_ip.name` | yes | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `cluster_network_interface.subnet` (block) | yes | nested block |
| `cluster_network_interface.subnet.id` | yes | pattern `^[-0-9a-z_]+$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_instance_cluster_network_attachment_invalid_instance_id` | `instance_id` |
| `ibm_is_instance_cluster_network_attachment_invalid_name` | `name` |
| `ibm_is_instance_cluster_network_attachment_invalid_before__id` | `before.id` |
| `ibm_is_instance_cluster_network_attachment_invalid_cluster_network_interface__name` | `cluster_network_interface.name` |
| `ibm_is_instance_cluster_network_attachment_invalid_cluster_network_interface__primary_ip__address` | `cluster_network_interface.primary_ip.address` |
| `ibm_is_instance_cluster_network_attachment_invalid_cluster_network_interface__primary_ip__href` | `cluster_network_interface.primary_ip.href` |
| `ibm_is_instance_cluster_network_attachment_invalid_cluster_network_interface__primary_ip__id` | `cluster_network_interface.primary_ip.id` |
| `ibm_is_instance_cluster_network_attachment_invalid_cluster_network_interface__primary_ip__name` | `cluster_network_interface.primary_ip.name` |
| `ibm_is_instance_cluster_network_attachment_invalid_cluster_network_interface__subnet__id` | `cluster_network_interface.subnet.id` |

## Example

### Valid

```hcl
resource "ibm_is_instance_cluster_network_attachment" "test" {
  instance_id = "example-name"
  cluster_network_interface {
    subnet {
      id = "example-name"
    }
  }
}
```
