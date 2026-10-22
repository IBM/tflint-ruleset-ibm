# ibm_container_cluster_versions_data

Schema validation for the ibm_container_cluster_versions data source.

This rule validates `ibm_container_cluster_versions` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Cluster` (`kubernetes-containers-v1-v2`).

## Configuration

```hcl
rule "ibm_container_cluster_versions_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_guid` | no | — |
| `default_kube_version` | no | — |
| `default_openshift_version` | no | — |
| `org_guid` | no | — |
| `region` | no | API: The region value. |
| `resource_group_id` | no | — |
| `space_guid` | no | — |
| `valid_kube_versions` | no | — |
| `valid_openshift_versions` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_container_cluster_versions_data_invalid_account_guid` | `account_guid` |
| `ibm_container_cluster_versions_data_invalid_org_guid` | `org_guid` |
| `ibm_container_cluster_versions_data_invalid_region` | `region` |
| `ibm_container_cluster_versions_data_invalid_resource_group_id` | `resource_group_id` |
| `ibm_container_cluster_versions_data_invalid_space_guid` | `space_guid` |

## Example

### Valid

```hcl
data "ibm_container_cluster_versions" "test" {
}
```
