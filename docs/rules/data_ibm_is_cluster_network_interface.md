# ibm_is_cluster_network_interface_data

Reads ibm_is_cluster_network_interface. The name for this cluster network interface.

This rule validates `ibm_is_cluster_network_interface` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_is_cluster_network_interface_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `allow_ip_spoofing` | no | — |
| `auto_delete` | no | — |
| `cluster_network_id` | yes | pattern `^[-0-9a-z_]+$` |
| `cluster_network_interface_id` | yes | — |
| `created_at` | no | — |
| `enable_infrastructure_nat` | no | — |
| `href` | no | — |
| `lifecycle_reasons` | no | — |
| `lifecycle_state` | no | — |
| `mac_address` | no | — |
| `name` | no | — |
| `primary_ip` | no | — |
| `resource_type` | no | — |
| `subnet` | no | — |
| `target` | no | — |
| `vpc` | no | — |
| `zone` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_cluster_network_interface_data_invalid_cluster_network_id` | `cluster_network_id` |
| `ibm_is_cluster_network_interface_data_invalid_cluster_network_interface_id` | `cluster_network_interface_id` |

## Example

### Valid

```hcl
data "ibm_is_cluster_network_interface" "test" {
  cluster_network_id = "example-name"
  cluster_network_interface_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
