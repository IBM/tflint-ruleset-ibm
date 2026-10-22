# ibm_container_cluster_config_data

Schema validation for the ibm_container_cluster_config data source.

This rule validates `ibm_container_cluster_config` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Cluster` (`kubernetes-containers-v1-v2`).

## Configuration

```hcl
rule "ibm_container_cluster_config_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_guid` | no | — |
| `admin` | no | — |
| `admin_certificate` | no | — |
| `admin_key` | no | — |
| `ca_certificate` | no | — |
| `calico_config_file_path` | no | — |
| `cluster_name_id` | yes | — |
| `config_dir` | no | — |
| `config_file_path` | no | — |
| `download` | no | — |
| `endpoint_type` | no | — |
| `host` | no | — |
| `network` | no | — |
| `org_guid` | no | — |
| `region` | no | API: The region value. |
| `resource_group_id` | no | — |
| `space_guid` | no | — |
| `token` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_container_cluster_config_data_invalid_account_guid` | `account_guid` |
| `ibm_container_cluster_config_data_invalid_admin` | `admin` |
| `ibm_container_cluster_config_data_invalid_cluster_name_id` | `cluster_name_id` |
| `ibm_container_cluster_config_data_invalid_config_dir` | `config_dir` |
| `ibm_container_cluster_config_data_invalid_download` | `download` |
| `ibm_container_cluster_config_data_invalid_endpoint_type` | `endpoint_type` |
| `ibm_container_cluster_config_data_invalid_network` | `network` |
| `ibm_container_cluster_config_data_invalid_org_guid` | `org_guid` |
| `ibm_container_cluster_config_data_invalid_region` | `region` |
| `ibm_container_cluster_config_data_invalid_resource_group_id` | `resource_group_id` |
| `ibm_container_cluster_config_data_invalid_space_guid` | `space_guid` |

## Example

### Valid

```hcl
data "ibm_container_cluster_config" "test" {
  cluster_name_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
