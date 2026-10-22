# ibm_is_cluster_network_profiles_data

Schema validation for the ibm_is_cluster_network_profiles data source.

This rule validates `ibm_is_cluster_network_profiles` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_is_cluster_network_profiles_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `profiles` | no | — |

## Example

### Valid

```hcl
data "ibm_is_cluster_network_profiles" "test" {
}
```
