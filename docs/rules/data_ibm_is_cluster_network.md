# ibm_is_cluster_network_data

Reads ibm_is_cluster_network. The name for this cluster network.

This rule validates `ibm_is_cluster_network` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_is_cluster_network_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cluster_network_id` | yes | — |
| `created_at` | no | — |
| `crn` | no | — |
| `href` | no | — |
| `lifecycle_reasons` | no | — |
| `lifecycle_state` | no | — |
| `name` | no | — |
| `profile` | no | — |
| `resource_group` | no | — |
| `resource_type` | no | — |
| `subnet_prefixes` | no | — |
| `vpc` | no | — |
| `zone` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_cluster_network_data_invalid_cluster_network_id` | `cluster_network_id` |

## Example

### Valid

```hcl
data "ibm_is_cluster_network" "test" {
  cluster_network_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
