# ibm_is_cluster_network

Manages ibm_is_cluster_network. The name for this cluster network.

This rule validates `ibm_is_cluster_network` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ClusterNetworkPrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_cluster_network" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `created_at` | no | — |
| `crn` | no | — |
| `etag` | no | — |
| `href` | no | — |
| `lifecycle_reasons` | no | — |
| `lifecycle_state` | no | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `profile` | yes | API: Identifies an cluster network profile by a unique property. |
| `resource_group` | no | API: The resource group to use. If unspecified, the account's [default resource group](https://cloud.ibm.com/apidocs/resource-manager#introduction) will be used. |
| `resource_type` | no | — |
| `zone` | yes | zone format |
| `subnet_prefixes` (block) | no | nested block |
| `subnet_prefixes.cidr` | yes | cidr format |
| `vpc` (block) | yes | nested block |
| `vpc.id` | yes | pattern `^[-0-9a-z_]+$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_cluster_network_invalid_name` | `name` |
| `ibm_is_cluster_network_invalid_profile` | `profile` |
| `ibm_is_cluster_network_invalid_resource_group` | `resource_group` |
| `ibm_is_cluster_network_invalid_zone` | `zone` |
| `ibm_is_cluster_network_invalid_subnet_prefixes__cidr` | `subnet_prefixes.cidr` |
| `ibm_is_cluster_network_invalid_vpc__id` | `vpc.id` |

## Example

### Valid

```hcl
resource "ibm_is_cluster_network" "test" {
  profile = "example"
  zone = "us-south-1"
  vpc {
    id = "example-name"
  }
}
```
