# ibm_is_instance_cluster_network_attachment_data

Reads ibm_is_instance_cluster_network_attachment. The name for this instance cluster network attachment.

This rule validates `ibm_is_instance_cluster_network_attachment` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_is_instance_cluster_network_attachment_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `before` | no | — |
| `cluster_network_interface` | no | — |
| `href` | no | — |
| `instance_cluster_network_attachment_id` | yes | — |
| `instance_id` | yes | pattern `^[-0-9a-z_]+$` |
| `lifecycle_reasons` | no | — |
| `lifecycle_state` | no | — |
| `name` | no | — |
| `resource_type` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_instance_cluster_network_attachment_data_invalid_instance_cluster_network_attachment_id` | `instance_cluster_network_attachment_id` |
| `ibm_is_instance_cluster_network_attachment_data_invalid_instance_id` | `instance_id` |

## Example

### Valid

```hcl
data "ibm_is_instance_cluster_network_attachment" "test" {
  instance_cluster_network_attachment_id = "01234567-89ab-cdef-0123-456789abcdef"
  instance_id = "example-name"
}
```
