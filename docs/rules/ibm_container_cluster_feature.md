# ibm_container_cluster_feature

Schema validation for the ibm_container_cluster_feature resource.

This rule validates `ibm_container_cluster_feature` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Cluster` (`kubernetes-containers-v1-v2`).

## Configuration

```hcl
rule "ibm_container_cluster_feature" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cluster` | yes | API: The cluster value. |
| `private_service_endpoint` | no | — |
| `private_service_endpoint_url` | no | — |
| `public_service_endpoint` | no | — |
| `public_service_endpoint_url` | no | — |
| `refresh_api_servers` | no | — |
| `reload_workers` | no | — |
| `resource_group_id` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_container_cluster_feature_invalid_cluster` | `cluster` |
| `ibm_container_cluster_feature_invalid_private_service_endpoint` | `private_service_endpoint` |
| `ibm_container_cluster_feature_invalid_public_service_endpoint` | `public_service_endpoint` |
| `ibm_container_cluster_feature_invalid_refresh_api_servers` | `refresh_api_servers` |
| `ibm_container_cluster_feature_invalid_reload_workers` | `reload_workers` |
| `ibm_container_cluster_feature_invalid_resource_group_id` | `resource_group_id` |

## Example

### Valid

```hcl
resource "ibm_container_cluster_feature" "test" {
  cluster = "aaaaaaa"
}
```
