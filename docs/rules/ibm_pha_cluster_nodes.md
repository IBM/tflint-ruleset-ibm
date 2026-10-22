# ibm_pha_cluster_nodes

Schema validation for the ibm_pha_cluster_nodes resource.

This rule validates `ibm_pha_cluster_nodes` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pha_cluster_nodes" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `accept_language` | no | pattern `^[a-zA-Z0-9\-_,;=.*]+$` |
| `etag` | no | — |
| `id` | no | — |
| `if_none_match` | no | pattern `^[a-zA-Z0-9\-_,;=.*]+$` |
| `instance_id` | yes | pattern `^[a-zA-Z0-9-]+$` |
| `primary_cluster_nodes` | yes | pattern `^[a-zA-Z0-9._:-]+$` |
| `primary_node_details` | no | — |
| `secondary_cluster_nodes` | no | pattern `^[a-zA-Z0-9._:-]+$` |
| `secondary_node_details` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pha_cluster_nodes_invalid_accept_language` | `accept_language` |
| `ibm_pha_cluster_nodes_invalid_if_none_match` | `if_none_match` |
| `ibm_pha_cluster_nodes_invalid_instance_id` | `instance_id` |
| `ibm_pha_cluster_nodes_invalid_primary_cluster_nodes` | `primary_cluster_nodes` |
| `ibm_pha_cluster_nodes_invalid_secondary_cluster_nodes` | `secondary_cluster_nodes` |

## Example

### Valid

```hcl
resource "ibm_pha_cluster_nodes" "test" {
  instance_id = "example-name"
  primary_cluster_nodes = ["example-name"]
}
```
