# ibm_pha_deployment_data

Schema validation for the ibm_pha_deployment data source.

This rule validates `ibm_pha_deployment` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pha_deployment_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cloud_account_id` | no | — |
| `connectivity_type` | no | — |
| `creation_time` | no | — |
| `custom_network` | no | — |
| `deprovision_time` | no | — |
| `guid` | no | — |
| `if_none_match` | no | pattern `^[a-zA-Z0-9\-_,;=.*]+$` |
| `instance_id` | yes | pattern `^[a-zA-Z0-9-]+$` |
| `is_duplicate` | no | — |
| `plan_id` | no | — |
| `plan_name` | no | — |
| `powerha_cluster_name` | no | — |
| `powerha_cluster_type` | no | — |
| `powerha_level` | no | — |
| `primary_cluster_nodes_details` | no | — |
| `primary_location` | no | — |
| `primary_region_name` | no | — |
| `primary_workspace` | no | — |
| `primary_workspace_name` | no | — |
| `provision_end_time` | no | — |
| `provision_start_time` | no | — |
| `provision_status` | no | — |
| `region_id` | no | — |
| `resource_group` | no | — |
| `resource_group_crn` | no | — |
| `resource_instance` | no | — |
| `secondary_cluster_nodes` | no | — |
| `secondary_location` | no | — |
| `secondary_workspace` | no | — |
| `service_description` | no | — |
| `service_id` | no | — |
| `service_name` | no | — |
| `standby_region_name` | no | — |
| `standby_workspace_name` | no | — |
| `user_tags` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pha_deployment_data_invalid_if_none_match` | `if_none_match` |
| `ibm_pha_deployment_data_invalid_instance_id` | `instance_id` |

## Example

### Valid

```hcl
data "ibm_pha_deployment" "test" {
  instance_id = "example-name"
}
```
