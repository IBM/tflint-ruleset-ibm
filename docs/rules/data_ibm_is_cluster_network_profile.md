# ibm_is_cluster_network_profile_data

Reads ibm_is_cluster_network_profile. The cluster network profile name.

This rule validates `ibm_is_cluster_network_profile` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `BareMetalServerDiskPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_cluster_network_profile_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `family` | no | — |
| `href` | no | — |
| `name` | yes | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `resource_type` | no | — |
| `supported_instance_profiles` | no | — |
| `zones` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_cluster_network_profile_data_invalid_name` | `name` |

## Example

### Valid

```hcl
data "ibm_is_cluster_network_profile" "test" {
  name = "example-name"
}
```
