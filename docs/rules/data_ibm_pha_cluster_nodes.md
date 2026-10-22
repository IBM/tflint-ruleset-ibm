# ibm_pha_cluster_nodes_data

Schema validation for the ibm_pha_cluster_nodes data source.

This rule validates `ibm_pha_cluster_nodes` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pha_cluster_nodes_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `if_none_match` | no | pattern `^[a-zA-Z0-9\-_,;=.*]+$` |
| `instance_id` | yes | pattern `^[a-zA-Z0-9-]+$` |
| `primary_node_details` | no | — |
| `secondary_node_details` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pha_cluster_nodes_data_invalid_if_none_match` | `if_none_match` |
| `ibm_pha_cluster_nodes_data_invalid_instance_id` | `instance_id` |

## Example

### Valid

```hcl
data "ibm_pha_cluster_nodes" "test" {
  instance_id = "example-name"
}
```
